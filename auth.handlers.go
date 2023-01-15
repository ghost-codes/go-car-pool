package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type SignUpData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type JwtAuthClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func createJWT(id int) (string, error) {
	jwtSecrete := os.Getenv("jwtSecret")

	claims := JwtAuthClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	}
	secrete := []byte(jwtSecrete)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secrete)
}

func valuidateJwt(tokenString string) (*jwt.Token, error) {
	jwtSecrete := os.Getenv("jwtSecret")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecrete), nil
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hasedString string, password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hasedString), []byte(password))
	return result == nil
}

func (api *ApiServer) InitAuthHandlers(router *mux.Router) {
	router.HandleFunc("/login", MakeHttpHanlder(api.userLogin))
	router.HandleFunc("/sign_up", MakeHttpHanlder(api.userSignUp))
}

func (api *ApiServer) userLogin(w http.ResponseWriter, r *http.Request) ErrorI {
	if r.Method != Methods.POST {
		return ApiError{
			code: http.StatusMethodNotAllowed,
			Err:  "Method not allowed",
		}
	}

	err := SendMail(SendMailData{
		To:      "compounddork@gmail.com",
		Subject: "Test",
		Body:    "Testing my smtp server on go",
	})

	loginData := new(LoginData)

	if err := json.NewDecoder(r.Body).Decode(loginData); err != nil {
		return ApiError{
			code: http.StatusBadRequest,
			Err:  err.Error(),
		}
	}

	user, err := api.store.userStore.GetUserByIdentifier(loginData.Identifier)
	if err != nil {
		log.Fatal(err)
		return ApiError{code: http.StatusInternalServerError, Err: "Something went wrong"}
	}
	if user == nil {
		return ApiError{
			Err:  "User not found",
			code: http.StatusNotFound,
		}
	}

	if !ComparePassword(user.HashedPassword, loginData.Password) {
		return ApiError{
			code: http.StatusUnauthorized,
			Err:  "Invalid user credentials",
		}
	}
	token, err := createJWT(user.ID)
	if err != nil {
		return ApiError{Err: err.Error(),
			code: http.StatusInternalServerError,
		}
	}
	return WriteJson(w, http.StatusOK, map[string]interface{}{"user": user, "token": token})
}

func (api *ApiServer) userSignUp(w http.ResponseWriter, r *http.Request) ErrorI {
	if r.Method != Methods.POST {
		return ApiError{
			code: http.StatusMethodNotAllowed,
			Err:  "Method not allowed",
		}
	}
	signUpData := new(SignUpData)

	if err := json.NewDecoder(r.Body).Decode(signUpData); err != nil {
		return ApiError{
			code: http.StatusBadRequest,
			Err:  err.Error(),
		}
	}

	hashedPassword, err := HashPassword(signUpData.Password)
	if err != nil {
		return ApiError{
			Err:  err.Error(),
			code: http.StatusInternalServerError,
		}
	}

	signUpData.Password = hashedPassword

	user, err := api.store.userStore.CreateUser(signUpData)
	if err != nil {
		return ApiError{
			Err:  err.Error(),
			code: http.StatusInternalServerError,
		}
	}

	token, err := createJWT(user.ID)
	if err != nil {
		return ApiError{Err: err.Error(),
			code: http.StatusInternalServerError,
		}
	}
	return WriteJson(w, http.StatusOK, map[string]interface{}{"user": user, "token": token})

}
