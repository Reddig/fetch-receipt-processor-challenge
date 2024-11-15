package models
import (
	"errors"
	"regexp"
)



func ValidateShortDescription(description string) error {
	descriptionRegex := `^[\w\s\-]+$`
	re := regexp.MustCompile(descriptionRegex)
	if !re.MatchString(description) {
		return errors.New("item short description is invalid")
	}
	return nil
}

func ValidatePrice(price string) error {
	priceRegex := `^\d+\.\d{2}$`
	re := regexp.MustCompile(priceRegex)
	if !re.MatchString(price) {
		return errors.New("item price is invalid")
	}
	return nil
}
