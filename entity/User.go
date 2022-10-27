package entity

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

var secret_key = "RAHASIA"

type User struct {
	ID              int64     `db:"id"`
	Username        string    `db:"username"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	Age             int64     `db:"age"`
	ProfileImageUrl string    `db:"profile_image_url"`
	CereatedAt      null.Time `db:"created_at"`
	UpdatedAt       null.Time `db:"updated_at"`
}

// hashPassword
func (u *User) HashPass() error {
	salt := 8
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Println("gagal generate password")
		return errors.New("something went error")
	}
	u.Password = string(hash)
	return nil
}

// compare password
func (u *User) ComparePassword(uerPassword string) bool {

	fmt.Println("hash password : ", u.Password)
	fmt.Println("user password : ", uerPassword)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(uerPassword))
	if err != nil {
		log.Println(err)
		return false
	}

	return err == nil

}

func (u *User) claimsForAccessToken() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secret_key))

	return signedToken
}

func (u *User) GenerateToken() string {
	tokenClaims := u.claimsForAccessToken()
	return u.signToken(tokenClaims)
}

func (u *User) ParseToken(stringToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secret_key), nil
	})

	if err != nil {
		var vErr *jwt.ValidationError
		if errors.As(err, &vErr) {
			if vErr.Errors == jwt.ValidationErrorExpired {
				return nil, errors.New("token expired")
			}
		}

		return nil, errors.New("invalid token")
	}

	return token, nil
}

func (u *User) bindTokenDataToUserEntity(mapClaims jwt.MapClaims) error {
	if v, ok := mapClaims["id"].(float64); !ok {
		return errors.New("invalid token")
	} else {
		u.ID = int64(v)
	}

	if v, ok := mapClaims["email"].(string); !ok {
		return errors.New("invalid token")
	} else {
		u.Email = v
	}

	return nil
}

func (u *User) VerifyToken(tokenStr string) error {

	if bearer := strings.HasPrefix(tokenStr, "Bearer"); !bearer {
		return errors.New("login to proceed")
	}

	stringToken := strings.Split(tokenStr, " ")[1]

	token, err := u.ParseToken(stringToken)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		mapClaims = v
	}

	err = u.bindTokenDataToUserEntity(mapClaims)

	if err != nil {
		return err
	}

	return nil

}
