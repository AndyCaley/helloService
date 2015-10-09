// @APIVersion 1.0.0
// @APITitle My Cool API
// @APIDescription My API usually works as expected.
// @Contact api@contact.me
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

// @SubApi hello API [/v1/hello]
package main

import (
	"github.com/go-zoo/bone"
	"net/http"
)

func main() {
	router := bone.New()

	router.GetFunc("/v1/hello", HelloHandler)

	http.ListenAndServe(":8080", router)
}
