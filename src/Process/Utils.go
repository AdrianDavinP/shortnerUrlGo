package Process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"mraladin/shortnerUrl/src/Struct"
	"net/http"
	"strings"
	"time"
)

func ConnectionPostgres() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=Password2019 dbname=mrAladin sslmode=disable")
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	db.LogMode(true)
	return db
}

func GetToken(w http.ResponseWriter, req *http.Request) {
	var resp = map[string]interface{}{"status": false}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := Struct.Token{Company: "mrAladin", Date: time.Now().String(),
		StandardClaims: &jwt.StandardClaims{ExpiresAt: expiresAt}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
		return
	}
	resp["status"] = true
	resp["token"] = tokenString
	//resp["token"] = tokenString
	json.NewEncoder(w).Encode(resp)
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		tk := &Struct.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Invalid Token")
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
