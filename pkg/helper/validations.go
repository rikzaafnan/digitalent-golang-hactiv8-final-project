package helper

import (
	"log"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) error {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
