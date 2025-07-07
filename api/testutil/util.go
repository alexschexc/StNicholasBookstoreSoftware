package testutil

import (
	"os"
	"testing"

	"github.com/onsi/gomega"
)

func TempFile(t *testing.T, g *gomega.WithT) string {
	tmp, err := os.CreateTemp("", "sqlite-bookstore.db")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	// t.Cleanup(func() { os.Remove(tmp.Name()) })
	tmp.Close()
	return tmp.Name()
}
