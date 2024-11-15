package models

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	// "fmt"
)

// Receipt represents a receipt with an ID and a name
type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required"`
	Total        string `json:"total" validate:"required"`
}

var validate = validator.New()

// ReceiptStorage will store our receipts in memory
var ReceiptStorage = map[string]Receipt{}

// AddReceipt adds a receipt to the storage
func AddReceipt(receipt Receipt) (string, error) {
	receipt.ID = uuid.New().String()
	ReceiptStorage[receipt.ID] = receipt
	return receipt.ID, nil
}

// GetReceipt retrieves a receipt by ID
func GetReceipt(id string) (Receipt, bool) {
	receipt, exists := ReceiptStorage[id]
	return receipt, exists
}

// GetAllReceipts returns all receipts
func GetAllReceipts() []Receipt {
	receipts := make([]Receipt, 0, len(ReceiptStorage))
	for _, receipt := range ReceiptStorage {
		receipts = append(receipts, receipt)
	}
	return receipts
}

func ValidateReceipt(receipt Receipt) error {
	if err := validate.Struct(receipt); err != nil {
		return err
	}
	if err := ValidatePurchaseDate(receipt.PurchaseDate); err != nil {
		return err
	}
	if err := ValidatePurchaseTime(receipt.PurchaseTime); err != nil {
		return err
	}
	for _, item := range receipt.Items {
		if err := ValidateItem(item); err != nil {
			return err
		}
	}
	if err := ValidateTotal(receipt.Total, receipt.Items); err != nil {
		return err
	}

	return nil
}

// calculates the points for a receipt
func CalculatePoints(receipt Receipt) int {
	var points = 0
	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)

	points += CalculateRetailerNamePoints(receipt.Retailer)
	points += CalculateWholeDollarPoints(totalFloat)
	points += CalculateQuarterMultiplesPoints(totalFloat)
	points += CalculateItemPairPoints(len(receipt.Items))
	points += CalculateDescriptionLengthPoints(receipt.Items)
	points += CalculatePurchaseDatePoints(receipt.PurchaseDate)
	points += CalculatePurchaseTimePoints(receipt.PurchaseTime)
	return points
}

func CalculateRetailerNamePoints(name string) int {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	var modifiedString = nonAlphanumericRegex.ReplaceAllString(name, "")
	return len(modifiedString)
}

func CalculateWholeDollarPoints(total float64) int {
	// if this is converted to int and back to float, it will round off the decimal places
	if FloatIsWholeNumber(total) {
		return 50
	} else {
		return 0
	}
}

func CalculateQuarterMultiplesPoints(total float64) int {
	var wholeDollar = float64(int64(total))
	// multiply by 100 to guarantee integer, assuming good input, else we round off the 1000th of cents
	var cents = int64((total - wholeDollar) * 100)
	if cents%25 == 0 { // 0, 25, 50, 75 are the multiples of 25 we are looking for
		return 25
	} else {
		return 0
	}
}

func CalculateItemPairPoints(itemCount int) int {
	return 5 * (itemCount / 2)
}

func CalculateDescriptionLengthPoints(items []Item) int {
	var points = 0
	for _, item := range items {
		points += CalculateIndividualItemDescriptionPoints(item)
	}
	return points
}

func CalculateIndividualItemDescriptionPoints(item Item) int {
	trimmedDescription := strings.TrimSpace(item.ShortDescription)
	if len(trimmedDescription)%3 == 0 {
		price, _ := strconv.ParseFloat(item.Price, 64)
		var points = 0.2 * price
		// if we've got a whole number, return it
		if FloatIsWholeNumber(points) {
			return int(points)
			// if number is not whole, round down to nearest whole number and add 1
		} else {
			return int(points) + 1
		}
	} else { // if not a multiple of 3, it gets no points
		return 0
	}
}

func CalculatePurchaseDatePoints(date string) int {
	dateValue, _ := time.Parse("2006-01-02", date)
	if dateValue.Day()%2 == 1 {
		return 6
	} else {
		return 0
	}
}

func CalculatePurchaseTimePoints(time string) int {
	splitTimes := strings.Split(time, ":")
	if splitTimes[0] == "15" {
		return 10
	} else if (splitTimes[0] == "14" || splitTimes[0] == "16") && (splitTimes[1] != "00") {
		return 10
	}
	return 0
}

func FloatIsWholeNumber(num float64) bool {
	return float64(int64(num)) == num
}
