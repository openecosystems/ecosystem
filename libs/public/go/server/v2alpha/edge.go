package serverv2alphalib

import (
	"fmt"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"net/http"
)

func edgeRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Handle Cors Preflight
		if r.Method == http.MethodOptions &&
			r.Header[sdkv2alphalib.Origin] != nil &&
			r.Header[sdkv2alphalib.AccessControlRequestHeaders] != nil &&
			r.Header[sdkv2alphalib.AccessControlRequestMethod] != nil {

			w.Header().Set(sdkv2alphalib.AccessControlAllowOrigin, r.Header.Get(sdkv2alphalib.Origin))
			w.Header().Set(sdkv2alphalib.AccessControlAllowMethods, "HEAD,OPTIONS,GET,PUT,PATCH,POST,DELETE")
			w.Header().Set(sdkv2alphalib.AccessControlAllowHeaders, "*")
			w.Header().Set(sdkv2alphalib.AccessControlMaxAge, "86400")
			w.Header().Set(sdkv2alphalib.CacheControl, "public, max-age=86400")
		}

		// Handle robots
		if r.URL.Path == "/robots.txt" {
			w.WriteHeader(http.StatusOK)
			_, err := fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
			if err != nil {
				return
			}
			return
		}

		// Sanitize Methods

		// Sanitize Headers
		// Sanitize Query Strings
		// Normalize Request

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
