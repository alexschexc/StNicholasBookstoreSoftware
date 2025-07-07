package main

import (
	"context"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/alexschexc/bookstore/api/dbmodels"
	"github.com/alexschexc/bookstore/api/inject"
	"github.com/alexschexc/bookstore/api/restapi"
	"github.com/alexschexc/bookstore/api/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewBookstoreAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Bookstore API"
	parser.LongDescription = "TODO: do better than Bookstore API description"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	cs := os.Getenv("DB_CONNECTION_STRING")
	sess, err := dbmodels.NewDBSession(context.Background(), cs)
	if err != nil {
		log.Fatalf("failed to connect to db %q: %v", cs, err)
	}
	defer sess.Close()

	if err := inject.InjectApi(api, sess); err != nil {
		log.Fatalln(err)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
