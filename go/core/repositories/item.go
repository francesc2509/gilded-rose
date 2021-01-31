package repositories

import (
	"github.com/francesc2509/glided-rose/entities"
)

type itemRepository struct{}

var items = []*entities.Item{
	nil,
	{Name: "Ordinary item 1", SellIn: 1, Quality: 50},
	{Name: "Ordinary item 2", SellIn: 2, Quality: 1},
	{Name: "Ordinary item 3", SellIn: 3, Quality: 0},
	{Name: "Ordinary item 4", SellIn: -3, Quality: 40},
	{Name: "Ordinary item 5", SellIn: -3, Quality: -4},
	{Name: "Sulfuras item 1", SellIn: 34, Quality: 80},
	{Name: "Sulfuras item 2", SellIn: 1, Quality: 50},
	{Name: "Sulfuras item 3", SellIn: 4, Quality: 100},
	{Name: "Aged Brie item 1", SellIn: 0, Quality: 49},
	{Name: "Aged Brie item 2", SellIn: 3, Quality: 48},
	{Name: "Aged Brie item 3", SellIn: -1, Quality: 47},
	{Name: "Aged Brie item 4", SellIn: 0, Quality: 54},
	{Name: "Backstage passes item 1", SellIn: 13, Quality: 49},
	{Name: "Backstage passes item 2", SellIn: 10, Quality: 49},
	{Name: "Backstage passes item 3", SellIn: 5, Quality: 48},
	{Name: "Backstage passes item 4", SellIn: 2, Quality: 50},
	{Name: "Backstage passes item 5", SellIn: 10, Quality: 46},
	{Name: "Backstage passes item 6", SellIn: 5, Quality: 45},
	{Name: "Backstage passes item 7", SellIn: -1, Quality: 44},
	{Name: "Backstage passes item 8", SellIn: -1, Quality: 0},
	{Name: "Backstage passes item 9", SellIn: 0, Quality: 0},
	{Name: "Conjured item 1", SellIn: 1, Quality: 50},
	{Name: "Conjured item 2", SellIn: 2, Quality: 30},
	{Name: "Conjured item 3", SellIn: 3, Quality: 0},
	{Name: "Conjured item 4", SellIn: 3, Quality: 1},
	{Name: "Conjured item 5", SellIn: -3, Quality: 3},
	{Name: "Conjured item 6", SellIn: -3, Quality: 40},
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
