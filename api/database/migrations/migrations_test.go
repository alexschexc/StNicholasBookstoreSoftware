package migrations_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/alexschexc/bookstore/api/database/migrations"
	"github.com/alexschexc/bookstore/api/testutil"
	. "github.com/onsi/gomega"

	// load SQLite3 driver explicitly
	_ "github.com/mattn/go-sqlite3"
)

func TestMigrations(t *testing.T) {
	g := NewWithT(t)
	tmp := testutil.TempFile(t, g)
	dbURI := fmt.Sprintf("sqlite3://%s", tmp)

	pending, err := migrations.Pending(dbURI)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).To(BeTrue())

	err = migrations.Up(dbURI)
	g.Expect(err).ToNot(HaveOccurred())

	pending, err = migrations.Pending(dbURI)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).To(BeFalse())

	db, err := sql.Open("sqlite3", tmp)
	g.Expect(err).ToNot(HaveOccurred())
	t.Cleanup(func() { db.Close() })
}
