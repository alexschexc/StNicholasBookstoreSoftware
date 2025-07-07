package restapi

import (
	"context"
	"net/http"

	"github.com/alexschexc/bookstore/api/dbmodels"
	"github.com/alexschexc/bookstore/api/inject"
	"github.com/alexschexc/bookstore/api/restapi/operations"
	"github.com/go-openapi/loads"
)

func getAPI() (*operations.BookstoreAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewBookstoreAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler(connectionString string) (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}

	sess, err := dbmodels.NewDBSession(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	if err := inject.InjectApi(api, sess); err != nil {
		return nil, err
	}

	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
