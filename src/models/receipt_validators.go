package models
import (
	"time"
	"errors"
	"regexp"
	"strconv"
	// "fmt"
	"math"
)
func ValidatePurchaseDate(purchaseDate string) error {
	_, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return errors.New("purchase date does not meet date format YYYY-MM-dd")
	}
	return nil
}

func ValidatePurchaseTime(purchaseTime string) error {
	timeRegex := `^(?:[01]\d|2[0-3]):[0-5]\d$`
	re := regexp.MustCompile(timeRegex)
	if !re.MatchString(purchaseTime) {
		return errors.New("purchase time does not meet date format HH:MM")
	}
	return nil
}

func ValidateTotal(total string, items []Item) error {
	totalRegex := `^\d+\.\d{2}$`
	re := regexp.MustCompile(totalRegex)
	if !re.MatchString(total) {
		return errors.New("total is invalid")
	}
	var itemTotal = 0.0
	for _, item := range items {
		parsedPrice, _ := strconv.ParseFloat(item.Price, 64)
		itemTotal += parsedPrice
	}
	floatTotal, _ := strconv.ParseFloat(total, 64)
	floatTotal = RoundTo(floatTotal, 2)
	itemTotal = RoundTo(itemTotal, 2)
	if itemTotal != floatTotal {
		return errors.New("item total does not match given total")
	}
	return nil
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
  }