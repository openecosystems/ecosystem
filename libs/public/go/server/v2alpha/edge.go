package serverv2alphalib

import (
	"fmt"
	"net/http"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

func edgeRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(sdkv2alphalib.AccessControlAllowOrigin, r.Header.Get(sdkv2alphalib.Origin))
		w.Header().Set(sdkv2alphalib.AccessControlAllowMethods, "HEAD,OPTIONS,GET,PUT,PATCH,POST,DELETE")
		w.Header().Set(sdkv2alphalib.AccessControlAllowHeaders, "*")
		w.Header().Set(sdkv2alphalib.AccessControlMaxAge, "86400")
		w.Header().Set(sdkv2alphalib.CacheControl, "public, max-age=86400")

		// Handle Cors Preflight
		if r.Method == http.MethodOptions &&
			r.Header[sdkv2alphalib.Origin] != nil &&
			r.Header[sdkv2alphalib.AccessControlRequestHeaders] != nil &&
			r.Header[sdkv2alphalib.AccessControlRequestMethod] != nil {
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent) // HTTP 204 No Content
				return
			}
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

		// Sanitize request headers
		// sanitizeRequestHeaders(r)

		// Before sanitization
		//println("Before sanitization:")
		//for k, v := range r.Header {
		//	println(k, ":", v[0])
		//}

		// Sanitize request headers
		sanitizeRequestHeaders(r)

		// After sanitization
		//println("\nAfter sanitization:")
		//for k, v := range r.Header {
		//	println(k, ":", v[0])
		//}

		// sanitizeRequestHeaders(r)
		// Sanitize Query Strings
		// Normalize Request

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

var allowedRequestHeaders = map[string]struct{}{
	sdkv2alphalib.Host:                            {},
	sdkv2alphalib.Origin:                          {},
	sdkv2alphalib.Accept:                          {},
	sdkv2alphalib.AcceptEncoding:                  {},
	sdkv2alphalib.AcceptCharset:                   {},
	sdkv2alphalib.AcceptLanguage:                  {},
	sdkv2alphalib.AcceptRanges:                    {},
	sdkv2alphalib.Authorization:                   {},
	sdkv2alphalib.ContentType:                     {},
	sdkv2alphalib.ContentLength:                   {},
	sdkv2alphalib.Referer:                         {},
	sdkv2alphalib.UserAgent:                       {},
	sdkv2alphalib.Connection:                      {},
	sdkv2alphalib.Upgrade:                         {},
	sdkv2alphalib.ConnectProtocolVersion:          {},
	sdkv2alphalib.ConnectTimeoutMs:                {},
	sdkv2alphalib.GrpcAcceptEncoding:              {},
	sdkv2alphalib.GrpcTimeout:                     {},
	sdkv2alphalib.Te:                              {},
	sdkv2alphalib.Http2Settings:                   {},
	sdkv2alphalib.SecChPrefersColorScheme:         {},
	sdkv2alphalib.SecChPrefersReducedMotion:       {},
	sdkv2alphalib.SecChPrefersReducedTransparency: {},
	sdkv2alphalib.SecChUa:                         {},
	sdkv2alphalib.SecChUaArch:                     {},
	sdkv2alphalib.SecChUaBitness:                  {},
	sdkv2alphalib.SecChUaFullVersionList:          {},
	sdkv2alphalib.SecChUaMobile:                   {},
	sdkv2alphalib.SecChUaModel:                    {},
	sdkv2alphalib.SecChUaPlatform:                 {},
	sdkv2alphalib.SecChUaPlatformVersion:          {},
	sdkv2alphalib.SecFetchDest:                    {},
	sdkv2alphalib.SecFetchMode:                    {},
	sdkv2alphalib.SecFetchSite:                    {},
	sdkv2alphalib.SecGpc:                          {},
	sdkv2alphalib.SecPurpose:                      {},
	sdkv2alphalib.SecWebsocketAccept:              {},
	sdkv2alphalib.SecWebsocketExtensions:          {},
	sdkv2alphalib.SecWebsocketKey:                 {},
	sdkv2alphalib.SecWebsocketProtocol:            {},
	sdkv2alphalib.SecWebsocketVersion:             {},
	sdkv2alphalib.AccessControlExposeHeaders:      {},
	sdkv2alphalib.AccessControlMaxAge:             {},
	sdkv2alphalib.AccessControlRequestHeaders:     {},
	sdkv2alphalib.AccessControlRequestMethod:      {},
	sdkv2alphalib.XCorrelationId:                  {},
	sdkv2alphalib.XClientTraceId:                  {},
	sdkv2alphalib.XB3Traceid:                      {},
	sdkv2alphalib.XB3Spanid:                       {},
	sdkv2alphalib.XB3Parentspanid:                 {},
	sdkv2alphalib.XB3Sampled:                      {},
	sdkv2alphalib.XB3Flags:                        {},
	sdkv2alphalib.XSpecApiKey:                     {},
	sdkv2alphalib.XSpecRouterKey:                  {},
	sdkv2alphalib.XSpecEnvironment:                {},
	sdkv2alphalib.XSpecPlatformHost:               {},
	sdkv2alphalib.XSpecFieldmask:                  {},
	sdkv2alphalib.XSpecSentAt:                     {},
	sdkv2alphalib.XSpecLocale:                     {},
	sdkv2alphalib.XSpecTimezone:                   {},
	sdkv2alphalib.XSpecValidateOnly:               {},
	sdkv2alphalib.XSpecDeviceId:                   {},
	sdkv2alphalib.XSpecDeviceAdvId:                {},
	sdkv2alphalib.XSpecDeviceManufacturer:         {},
	sdkv2alphalib.XSpecDeviceModel:                {},
	sdkv2alphalib.XSpecDeviceName:                 {},
	sdkv2alphalib.XSpecDeviceType:                 {},
	sdkv2alphalib.XSpecDeviceToken:                {},
	sdkv2alphalib.XSpecDeviceBluetooth:            {},
	sdkv2alphalib.XSpecDeviceCellular:             {},
	sdkv2alphalib.XSpecDeviceWifi:                 {},
	sdkv2alphalib.XSpecDeviceCarrier:              {},
	sdkv2alphalib.XSpecOsName:                     {},
	sdkv2alphalib.XSpecOsVersion:                  {},
	sdkv2alphalib.XSpecOrganization:               {},
	sdkv2alphalib.XSpecWorkspace:                  {},
	sdkv2alphalib.XSpecWorkspaceJan:               {},
}

func sanitizeRequestHeaders(req *http.Request) {
	for header := range req.Header {
		// If the header is not in the allowed list, delete it
		if _, allowed := allowedRequestHeaders[header]; !allowed {
			req.Header.Del(header)
		}
	}
}
