package sdkv2betalib

import (
	"fmt"
	"net/http"
)

func edgeRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(AccessControlAllowOrigin, r.Header.Get(Origin))
		w.Header().Set(AccessControlAllowMethods, "HEAD,OPTIONS,GET,PUT,PATCH,POST,DELETE")
		w.Header().Set(AccessControlAllowHeaders, "*")
		w.Header().Set(AccessControlMaxAge, "86400")
		w.Header().Set(CacheControl, "public, max-age=86400")

		// Handle Cors Preflight
		if r.Method == http.MethodOptions &&
			r.Header[Origin] != nil &&
			r.Header[AccessControlRequestHeaders] != nil &&
			r.Header[AccessControlRequestMethod] != nil {
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

		// TODO: Sanitize Query Strings
		// TODO: Normalize Request

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

var allowedRequestHeaders = map[string]struct{}{
	Host:                            {},
	Origin:                          {},
	Accept:                          {},
	AcceptEncoding:                  {},
	AcceptCharset:                   {},
	AcceptLanguage:                  {},
	AcceptRanges:                    {},
	Authorization:                   {},
	ContentType:                     {},
	ContentLength:                   {},
	Referer:                         {},
	UserAgent:                       {},
	Connection:                      {},
	Upgrade:                         {},
	ConnectProtocolVersion:          {},
	ConnectTimeoutMs:                {},
	GrpcAcceptEncoding:              {},
	GrpcTimeout:                     {},
	Te:                              {},
	Http2Settings:                   {},
	SecChPrefersColorScheme:         {},
	SecChPrefersReducedMotion:       {},
	SecChPrefersReducedTransparency: {},
	SecChUa:                         {},
	SecChUaArch:                     {},
	SecChUaBitness:                  {},
	SecChUaFullVersionList:          {},
	SecChUaMobile:                   {},
	SecChUaModel:                    {},
	SecChUaPlatform:                 {},
	SecChUaPlatformVersion:          {},
	SecFetchDest:                    {},
	SecFetchMode:                    {},
	SecFetchSite:                    {},
	SecGpc:                          {},
	SecPurpose:                      {},
	SecWebsocketAccept:              {},
	SecWebsocketExtensions:          {},
	SecWebsocketKey:                 {},
	SecWebsocketProtocol:            {},
	SecWebsocketVersion:             {},
	AccessControlExposeHeaders:      {},
	AccessControlMaxAge:             {},
	AccessControlRequestHeaders:     {},
	AccessControlRequestMethod:      {},
	XCorrelationId:                  {},
	XClientTraceId:                  {},
	XB3Traceid:                      {},
	XB3Spanid:                       {},
	XB3Parentspanid:                 {},
	XB3Sampled:                      {},
	XB3Flags:                        {},
	XSpecApiKey:                     {},
	XSpecRouterKey:                  {},
	XSpecEnvironment:                {},
	XSpecPlatformHost:               {},
	XSpecFieldmask:                  {},
	XSpecSentAt:                     {},
	XSpecLocale:                     {},
	XSpecTimezone:                   {},
	XSpecValidateOnly:               {},
	XSpecDeviceId:                   {},
	XSpecDeviceAdvId:                {},
	XSpecDeviceManufacturer:         {},
	XSpecDeviceModel:                {},
	XSpecDeviceName:                 {},
	XSpecDeviceType:                 {},
	XSpecDeviceToken:                {},
	XSpecDeviceBluetooth:            {},
	XSpecDeviceCellular:             {},
	XSpecDeviceWifi:                 {},
	XSpecDeviceCarrier:              {},
	XSpecOsName:                     {},
	XSpecOsVersion:                  {},
	XSpecEcosystem:                  {},
	// XSpecOrganization:               {},
	// XSpecWorkspace:                  {},
	// XSpecWorkspaceJan:               {},
}

func sanitizeRequestHeaders(req *http.Request) {
	for header := range req.Header {
		// If the header is not in the allowed list, delete it
		if _, allowed := allowedRequestHeaders[header]; !allowed {
			req.Header.Del(header)
		}
	}
}
