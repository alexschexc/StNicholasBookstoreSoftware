package migrations_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/alexschexc/bookstore/api/database/migrations"
	. "github.com/onsi/gomega"

	// load SQLite3 driver explicitly
	_ "github.com/mattn/go-sqlite3"
)

func TestMigrations(t *testing.T) {
	g := NewWithT(t)
	tmp := tempFile(t, g)
	db, err := sql.Open("sqlite3", tmp)
	g.Expect(err).ToNot(HaveOccurred())
	t.Cleanup(func() { db.Close() })

	dbURI := fmt.Sprintf("sqlite3://%s", tmp)

	pending, err := migrations.Pending(dbURI)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).To(BeTrue())

	err = migrations.Up(dbURI)
	g.Expect(err).ToNot(HaveOccurred())

	pending, err = migrations.Pending(dbURI)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(pending).To(BeFalse())
}

func tempFile(t *testing.T, g *WithT) string {
	tmp, err := os.CreateTemp("", "sqlite-bookstore.db")
	g.Expect(err).ToNot(HaveOccurred())
	t.Cleanup(func() { os.Remove(tmp.Name()) })
	tmp.Close()
	return tmp.Name()
}
