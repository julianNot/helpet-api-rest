package auth

import (
	"fmt"

	"github.com/julianNot/helpet-api-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username string, password string) (*models.Partner, error) {
	// Busca el usuario en la base de datos por su nombre de usuario
	var partner models.Partner
	result := db.Where("username = ?", username).First(&partner)
	if result.Error != nil {
		return nil, result.Error
	}

	// Comprueba si la contraseña es correcta
	err := bcrypt.CompareHashAndPassword([]byte(partner.Password), []byte(password))
	if err != nil {
		return nil, err
		fmt.Println()
	}

	// Devuelve el usuario autenticado
	return &partner, nil
}
