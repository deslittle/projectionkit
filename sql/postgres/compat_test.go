package postgres_test

import (
	"database/sql"

	"github.com/dogmatiq/projectionkit/sql/internal/drivertest"
	. "github.com/dogmatiq/projectionkit/sql/postgres"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func IsCompatibleWith()", func() {
	It("returns true if the driver is recognized", func() {
		db, err := sql.Open("postgres", "host=localhost")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(IsCompatibleWith(db)).To(BeTrue())
	})

	It("returns false if the driver is unrecognized", func() {
		Expect(IsCompatibleWith(drivertest.MockDB())).To(BeFalse())
	})
})