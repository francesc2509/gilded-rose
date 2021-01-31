package repositories

import (
	"testing"

	"github.com/francesc2509/glided-rose/entities"
)

// TestUpdateQuality tests UpdateQuality function
func TestUpdateQuality(t *testing.T) {
	expectedItems := setupUpdateQuality()

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
		{Name: "Conjured item 5", SellIn: -3, Quality: 40},
	}

	UpdateQuality(items)

	for pos, item := range items {
		expected := expectedItems[pos]

		if item == expected {
			t.Logf("OK - Item %d - got %s", pos, item)
			continue
		}

		if (item == nil && expected != nil) ||
			(item != nil && expected == nil) ||
			item.Name != expected.Name ||
			item.SellIn != expected.SellIn ||
			item.Quality != expected.Quality {
			t.Errorf("ERROR - Item %d - expected %s, but got %s", pos, expected, item)
			return
		}

		t.Logf("OK - Item %d - got %s", pos, item)
	}
}

// setupUpdateQuality sets up the expected values that should be returned by UpdateQuality
func setupUpdateQuality() []*entities.Item {
	return []*entities.Item{
		nil,
		{Name: "Ordinary item 1", SellIn: 0, Quality: 49},
		{Name: "Ordinary item 2", SellIn: 1, Quality: 0},
		{Name: "Ordinary item 3", SellIn: 2, Quality: 0},
		{Name: "Ordinary item 4", SellIn: -4, Quality: 38},
		{Name: "Ordinary item 5", SellIn: -4, Quality: -4},
		{Name: "Sulfuras item 1", SellIn: 34, Quality: 80},
		{Name: "Sulfuras item 2", SellIn: 1, Quality: 50},
		{Name: "Sulfuras item 3", SellIn: 4, Quality: 100},
		{Name: "Aged Brie item 1", SellIn: -1, Quality: 50},
		{Name: "Aged Brie item 2", SellIn: 2, Quality: 49},
		{Name: "Aged Brie item 3", SellIn: -2, Quality: 49},
		{Name: "Aged Brie item 4", SellIn: -1, Quality: 54},
		{Name: "Backstage passes item 1", SellIn: 12, Quality: 50},
		{Name: "Backstage passes item 2", SellIn: 9, Quality: 50},
		{Name: "Backstage passes item 3", SellIn: 4, Quality: 50},
		{Name: "Backstage passes item 4", SellIn: 1, Quality: 50},
		{Name: "Backstage passes item 5", SellIn: 9, Quality: 48},
		{Name: "Backstage passes item 6", SellIn: 4, Quality: 48},
		{Name: "Backstage passes item 7", SellIn: -2, Quality: 0},
		{Name: "Backstage passes item 8", SellIn: -2, Quality: 0},
		{Name: "Backstage passes item 9", SellIn: -1, Quality: 0},
		{Name: "Conjured item 1", SellIn: 0, Quality: 48},
		{Name: "Conjured item 2", SellIn: 1, Quality: 28},
		{Name: "Conjured item 3", SellIn: 2, Quality: 0},
		{Name: "Conjured item 4", SellIn: 2, Quality: 0},
		{Name: "Conjured item 5", SellIn: -4, Quality: 0},
		{Name: "Conjured item 5", SellIn: -4, Quality: 36},
	}
}
