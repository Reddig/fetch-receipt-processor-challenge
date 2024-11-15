package models

import (
	"testing"
	"fmt"
)

var validItem = &Item{
	ID: "1234",
	ShortDescription: "A description",
	Price: "10.00",
}

var validReceipt = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var validReceiptInvalidItems = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		{
			ID: "1",
			ShortDescription: "TEST",
			Price: "0",
		},
	},
	Total: "10.00",
}

var invalidReceiptRetailerName = &Receipt{
	Retailer: "CVS\n",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var invalidReceiptReceiptTotal = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10",
}


var invalidReceiptItemTotal = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.01",
}

var invalidReceiptDate = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-02-31",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var invalidReceiptTime = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:69",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

func TestValidateReceipt(t *testing.T) {
	tests := []struct {
		name string
		receipt Receipt
		wantErr bool
		expectedErr string
	}{
		{
			name: "Valid Receipt",
			receipt: *validReceipt,
			wantErr: false,
			expectedErr: "",
		},
		{
			name: "Valid Receipt Invalid Items",
			receipt: *validReceiptInvalidItems,
			wantErr: true,
			expectedErr: "item price 0 is invalid. Note: Must include 2 decimal places.",
		},
		{
			name: "Invalid Receipt Retailer Name",
			receipt: *invalidReceiptRetailerName,
			wantErr: true,
			expectedErr: "retailer name CVS\n is invalid",
		},
		{
			name: "Invalid Receipt Total",
			receipt: *invalidReceiptReceiptTotal,
			wantErr: true,
			expectedErr: "total 10 is invalid. Note: Must include 2 decimal places.",
		},
		{
			name: "Invalid Receipt Mismatched Total",
			receipt: *invalidReceiptItemTotal,
			wantErr: true,
			expectedErr: "item total 10.00 does not match given total 10.01",
		},
		{
			name: "Invalid Receipt Date",
			receipt: *invalidReceiptDate,
			wantErr: true,
			expectedErr: "purchase date 2020-02-31 is invalid",
		},
		{
			name: "Invalid Receipt Time",
			receipt: *invalidReceiptTime,
			wantErr: true,
			expectedErr: "purchase time 12:69 does not meet date format HH:MM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateReceipt(tt.receipt)
			if (err != nil)  {
				if (err != nil ) != tt.wantErr {
					t.Errorf("ValidateProduct() error = %v, wantErr %v", err, tt.wantErr)
				} else if fmt.Sprint(err) != tt.expectedErr {
					t.Errorf("ValidateProduct() incorrect error = %v, expectedErr %v", err, tt.expectedErr)
				}
			}
		})
	}
}



var receiptNonAlphaNumInRetailer = &Receipt{
	Retailer: "CVS&^123**;'",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var receiptExtraPointTime = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "14:45",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var receiptPurchaseTimeEdgeLower = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "14:00",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var receiptPurchaseTimeEdgeUpper = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "16:00",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var receiptTotalNotDivisibleByQuarters = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:00",
	Items: []Item {
		{
			ID: "1",
			Price: "1.28",
			ShortDescription: "asdf",
		},
	},
	Total: "1.28",
}

var receiptTotalDivisibleByQuartersNotWhole = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:00",
	Items: []Item {
		{
			ID: "1",
			Price: "1.25",
			ShortDescription: "asdf",
		},
	},
	Total: "1.25",
}

var receiptItemsDivisbleByTwo = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:00",
	Items: []Item {
		{
			ID: "1",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "2",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "3",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "4",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "5",
			Price: "1.25",
			ShortDescription: "asdf",
		},
	},
	Total: "6.25",
}

var receiptEvenPurchaseDay = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-02",
	PurchaseTime: "12:34",
	Items: []Item {
		*validItem,
	},
	Total: "10.00",
}

var receiptItemsDescriptionDivisibleByThree = &Receipt{
	Retailer: "CVS",
	PurchaseDate: "2020-01-01",
	PurchaseTime: "12:00",
	Items: []Item {
		{
			ID: "1",
			Price: "1",
			ShortDescription: "asd",
		},
		{
			ID: "2",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "3",
			Price: "1",
			ShortDescription: "asdfas",
		},
		{
			ID: "4",
			Price: "1.25",
			ShortDescription: "asdf",
		},
		{
			ID: "5",
			Price: "1.25",
			ShortDescription: "asdf",
		},
	},
	Total: "5.75",
}

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name string
		receipt Receipt
		expectedPoints int
	}{
		{
			name: "Eighty-Four Points",
			receipt: *validReceipt,
			//3 points retailer name, 50 points for whole dollar, 25 points for 0.25 multiple, 6 points purchase date
			expectedPoints: 84,
		},
		{
			name: "6 Point Retailer Ignoring Non-Alphanumeral Characters",
			receipt: *receiptNonAlphaNumInRetailer,
			//6 points retailer name (IGNORING NON-ALPHANUMERAL)
			expectedPoints: 87,
		},
		{
			name: "PurchaseTime Extra Point Window",
			receipt: *receiptExtraPointTime,
			// extra 10 points for time in 2:01-3:59pm window
			expectedPoints: 94,
		},
		{
			name: "PurchaseTime 2pm",
			receipt: *receiptPurchaseTimeEdgeLower,
			// no extra points at 2pm
			expectedPoints: 84,
		},
		{
			name: "PurchaseTime 4pm",
			receipt: *receiptPurchaseTimeEdgeUpper,
			// no extra points 4pm
			expectedPoints: 84,
		},
		{
			name: "Total Not Divsible By Quarters ",
			receipt: *receiptTotalNotDivisibleByQuarters,
			// no 25 point bonus because not divisble by quarters, no 50 point bonus
			expectedPoints: 9,
		},
		{
			name: "Total Divisble By Quarters, Not Whole Number",
			receipt: *receiptTotalDivisibleByQuartersNotWhole,
			// 25 point bonus because divisble by quarters, no 50 point bonus
			expectedPoints: 34,
		},
		{
			name: "Item Count Divisible By Two",
			receipt: *receiptItemsDivisbleByTwo,
			// Bonus 10 points for 5 items (5 points per 2)
			expectedPoints: 44,
		},
		{
			name: "Even Purchase Date",
			receipt: *receiptEvenPurchaseDay,
			// 84 points - 6 from starting point
			expectedPoints: 78,
		},
		{
			name: "Item Descriptions Divisible By Three",
			receipt: *receiptItemsDescriptionDivisibleByThree,
			// 44 points (receiptItemsDivisbleByTwo) + 1 for ceiling(0.2*1) + ceiling(0.2*1)
			expectedPoints: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pts := CalculatePoints(tt.receipt)
			if (pts != tt.expectedPoints)  {
				t.Errorf("CalculatePoints() points %v != expectedPoints %v", pts, tt.expectedPoints)
			}
		})
	}
}
