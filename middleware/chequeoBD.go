package middleware

import (
	"net/http"

	"github.com/squok/twittor/bd"
)

/*ChequeoBD chuequea la base de datos como handler*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
