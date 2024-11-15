package models
import (
	"time"
	"errors"
	"regexp"
	"strconv"
	"math"
)

func ValidateRetailer(retailer string) error {
	retailerRegex := `^[\w\s\-&]+$`
	re := regexp.MustCompile(retailerRegex)
	if !re.MatchString(retailer) {
		return errors.New("retailer name " + retailer + "is invalid")
	}
	return nil
}

func ValidatePurchaseDate(purchaseDate string) error {
	_, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return errors.New("purchase date " + purchaseDate + " is invalid")
	}
	return nil
}

func ValidatePurchaseTime(purchaseTime string) error {
	timeRegex := `^(?:[01]\d|2[0-3]):[0-5]\d$`
	re := regexp.MustCompile(timeRegex)
	if !re.MatchString(purchaseTime) {
		return errors.New("purchase time " + purchaseTime + " does not meet date format HH:MM")
	}
	return nil
}

func ValidateTotal(total string, items []Item) error {
	totalRegex := `^\d+\.\d{2}$`
	re := regexp.MustCompile(totalRegex)
	if !re.MatchString(total) {
		return errors.New("total " + total + " is invalid. Note: Must include 2 decimal places.")
	}
	var itemTotal = 0.0
	// we want to check that the sum total is correct for the price of all items
	for _, item := range items {
		parsedPrice, _ := strconv.ParseFloat(item.Price, 64)
		itemTotal += parsedPrice
	}
	floatTotal, _ := strconv.ParseFloat(total, 64)
	// I noticed some weirdness with floats not rounding so I do it manually here
	floatTotal = RoundTo(floatTotal, 2)
	itemTotal = RoundTo(itemTotal, 2)
	if itemTotal != floatTotal {
		return errors.New("item total " + strconv.FormatFloat(itemTotal, 'f', 2, 64) + " does not match given total " + strconv.FormatFloat(floatTotal, 'f', -1, 64))
	}
	return nil
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
  }