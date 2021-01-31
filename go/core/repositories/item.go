package repositories

import (
	"github.com/francesc2509/glided-rose/entities"
)

type itemRepository struct{}

var items = []*entities.Item{
	nil,
	{Id: 1, Name: "Ordinary item 1", SellIn: 1, Quality: 50},
	{Id: 2, Name: "Ordinary item 2", SellIn: 2, Quality: 1},
	{Id: 3, Name: "Ordinary item 3", SellIn: 3, Quality: 0},
	{Id: 4, Name: "Ordinary item 4", SellIn: -3, Quality: 40},
	{Id: 5, Name: "Ordinary item 5", SellIn: -3, Quality: -4},
	{Id: 6, Name: "Sulfuras item 1", SellIn: 34, Quality: 80},
	{Id: 7, Name: "Sulfuras item 2", SellIn: 1, Quality: 50},
	{Id: 8, Name: "Sulfuras item 3", SellIn: 4, Quality: 100},
	{Id: 9, Name: "Aged Brie item 1", SellIn: 0, Quality: 49},
	{Id: 10, Name: "Aged Brie item 2", SellIn: 3, Quality: 48},
	{Id: 11, Name: "Aged Brie item 3", SellIn: -1, Quality: 47},
	{Id: 12, Name: "Aged Brie item 4", SellIn: 0, Quality: 54},
	{Id: 13, Name: "Backstage passes item 1", SellIn: 13, Quality: 49},
	{Id: 14, Name: "Backstage passes item 2", SellIn: 10, Quality: 49},
	{Id: 15, Name: "Backstage passes item 3", SellIn: 5, Quality: 48},
	{Id: 16, Name: "Backstage passes item 4", SellIn: 2, Quality: 50},
	{Id: 17, Name: "Backstage passes item 5", SellIn: 10, Quality: 46},
	{Id: 18, Name: "Backstage passes item 6", SellIn: 5, Quality: 45},
	{Id: 19, Name: "Backstage passes item 7", SellIn: -1, Quality: 44},
	{Id: 20, Name: "Backstage passes item 8", SellIn: -1, Quality: 0},
	{Id: 21, Name: "Backstage passes item 9", SellIn: 0, Quality: 0},
	{Id: 22, Name: "Conjured item 1", SellIn: 1, Quality: 50},
	{Id: 23, Name: "Conjured item 2", SellIn: 2, Quality: 30},
	{Id: 24, Name: "Conjured item 3", SellIn: 3, Quality: 0},
	{Id: 25, Name: "Conjured item 4", SellIn: 3, Quality: 1},
	{Id: 26, Name: "Conjured item 5", SellIn: -3, Quality: 3},
	{Id: 27, Name: "Conjured item 6", SellIn: -3, Quality: 40},
}

func (repository *itemRepository) Get() ([]*entities.Item, error) {
	return items, nil
}

func (repository *itemRepository) UpdateInventory(days int) ([]*entities.Item, error) {
	UpdateQuality(items, days)

	return repository.Get()
}

// UpdateQuality updates Gilded Rose Inn's inventory
func UpdateQuality(items []*entities.Item, days int) {
	for _, item := range items {
		item.HandleItem()
	}
}
