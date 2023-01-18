package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/jayanthkrishna/golang-csrf-project/server/middleware/myjwt"
	"github.com/justinas/alice"
)

func NewHandler() http.Handler {
	return alice.New(recoverHandler, authhandler).ThenFunc(logicHandler)
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panic("recovered! Panic : %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}

func authhandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/restricted", "/logout", "deleteUser":
		default:
		}
	}

	return http.HandlerFunc(fn)
}

func logicHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/restricted":
	case "/login":
		switch r.Method {
		case "GET":
		case "POST":
		default:

		}
	case "/register":
		switch r.Method {
		case "GET":
		case "POST":
		default:
		}
	case "logout":
	case "/deleteUser":
	default:
	}
}

func nullifyToken(w http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshCookie)
}

func setAuthAndRefreshCookies(w http.ResponseWriter, r *http.Request, authTokenString, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}
	http.SetCookie(w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshCookie)

	RefreshCookie, refreshErr := r.Cookie("RefreshToken")
	if refreshErr == http.ErrNoCookie {
		return
	} else if refreshErr != nil {
		log.Panic("Panic : %+v", refreshErr)
		http.Error(w, http.StatusText(500), 500)
	}

	myjwt.RevokeRefreshToken(RefreshCookie.Value)
}

func grabCsrfFromReq(r *http.Request) string {
	csrfFromForm := r.FormValue("X-CSRF-Token")

	if csrfFromForm != "" {
		return csrfFromForm
	} else {
		return r.Header.Get("X-CSRF-Token")
	}

}
