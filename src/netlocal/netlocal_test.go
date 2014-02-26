package netlocal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"netlocal"
)

var _ = Describe("Netlocal", func() {
	Describe("StubGet", func() {
		BeforeEach(func() {
			s := netlocal.Start()
			s.StubGet(9999, "/index.html", 202, "<body>Hello World</body>")
		})

		AfterEach(func() {
			netlocal.Clear()
		})

		It("serves HTTP on the desired port", func() {
			var _, err = http.Get("http://localhost:9999/index.html")
			Expect(err).To(BeNil())
		})

		It("responds with the desired body for the stubbed path", func() {
			var resp, _ = http.Get("http://localhost:9999/index.html")
			var body, err = ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())
			Expect(string(body)).To(Equal("<body>Hello World</body>"))
		})

		It("responds with the given status code for the stubbed path", func() {
			var resp, _ = http.Get("http://localhost:9999/index.html")
			Expect(resp.StatusCode).To(Equal(202))
		})
	})
})
