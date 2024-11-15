package models

import (
	"github.com/google/uuid"
)

// Item represents a item with an ID and a name
type Item struct {
	ID   string `json:"id"`
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

// ItemStorage will store our items in memory
var ItemStorage = map[string]Item{}

// AddItem adds a item to the storage
func AddItem(item Item) {
	item.ID = uuid.New().String()
	ItemStorage[item.ID] = item
}

func ValidateItem(item Item) error {
	if err := validate.Struct(item); err != nil {
		return err
	}
	if err := ValidateShortDescription(item.ShortDescription); err != nil {
		return err
	}
	if err := ValidatePrice(item.Price); err != nil {
		return err
	}
	return nil
}
// GetItem retrieves a item by ID
func GetItem(id string) (Item, bool) {
	item, exists := ItemStorage[id]
	return item, exists
}

// GetAllItems returns all items
func GetAllItems() []Item {
	items := make([]Item, 0, len(ItemStorage))
	for _, item := range ItemStorage {
		items = append(items, item)
	}
	return items
}
