package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticatePartnerHandler(w http.ResponseWriter, r *http.Request) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	mySecretKey := os.Getenv("MY_SECRET_KEY")

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var partner models.Partner
	var attendant models.Attendant
	db.DB.Where("username = ?", creds.Username).First(&partner)
	db.DB.Where("username = ?", creds.Username).First(&attendant)
	if partner.ID == 0 || !checkPasswordHash(creds.Password, partner.Password) {
		if attendant.ID == 0 || checkPasswordHash(creds.Password, partner.Password) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Credentials"))
			return
		}
	}

	var nitValue int
	if partner.Nit != 0 {
		nitValue = 1
	} 
	if attendant.ID != 0 {
		nitValue = 0
	}
	

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	if nitValue == 1 {
		claims["partner_id"] = partner.ID
		claims["vet_id"] = partner.Vets[0].ID
	} else {
		claims["user_id"] = attendant.ID
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(mySecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token":     tokenString,
		"type_user": nitValue,
	}

	json.NewEncoder(w).Encode(response)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
