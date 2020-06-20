package firebase

import (
	"log"
	"net/http"
)

// BasicAuth performs basic authentication
func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, pw, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			http.Error(w, "auth required", http.StatusUnauthorized)
			return
		}

		if ok, err := VerifyEmailCredential(usr, pw); err != nil {
			log.Printf("failed to verify email cred: %v", err)
			http.Error(w, "Internal Server Error. please try again later.", http.StatusInternalServerError)
			return
		} else if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			http.Error(w, "Access is denied due to invalid credentials.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
