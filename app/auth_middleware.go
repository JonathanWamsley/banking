package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.repo.IsAuthorized(token, currentRouteVars)
				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					writeResponse(w, http.StatusForbidden, "Unauthorized")
				}

			} else {
				writeResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
