package kinben

import (
	"net/http"
)

func setCustomMiddleware(fn func(), next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn()
		next.ServeHTTP(w, r)
	})
}
