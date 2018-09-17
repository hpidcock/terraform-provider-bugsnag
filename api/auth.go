package api

import (
	"fmt"
	"net/http"
)

func appendAuth(req *http.Request, authToken string) {
	req.Header.Add("X-Version", "2")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", authToken))
}
