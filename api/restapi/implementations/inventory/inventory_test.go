package inventory_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexschexc/bookstore/api/database/migrations"
	"github.com/alexschexc/bookstore/api/dbmodels"
	"github.com/alexschexc/bookstore/api/restapi"
	"github.com/alexschexc/bookstore/api/testutil"
	. "github.com/onsi/gomega"
)

func TestMe(t *testing.T) {
	g := NewWithT(t)
	dbFile := testutil.TempFile(t, g)
	connectionString := fmt.Sprintf("sqlite://%s", dbFile)

	pending, err := migrations.Pending(connectionString)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).To(BeTrue())

	err = migrations.Up(connectionString)
	g.Expect(err).ToNot(HaveOccurred())

	pending, err = migrations.Pending(connectionString)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).ToNot(BeTrue())

	s, _ := dbmodels.NewDBSession(context.Background(), dbFile)
	t.Cleanup(func() { s.Close() })

	handler, err := restapi.GetAPIHandler(dbFile)
	if err != nil {
		t.Fatal("get api handler", err)
	}
	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/v1/inventory-items", nil)
	req.SetBasicAuth("wronguser", "wrongpassword")

	res, err := ts.Client().Do(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.StatusCode).To(Equal(200))
}
