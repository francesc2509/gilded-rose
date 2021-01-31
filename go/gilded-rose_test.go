package main

import "testing"

// TestUpdateQuality tests UpdateQuality function
func TestUpdateQuality(t *testing.T) {
	expectedItems := setupUpdateQuality()

	var items = []*Item{
		nil,
		{"Ordinary item 1", 1, 50},
		{"Ordinary item 2", 2, 1},
		{"Ordinary item 3", 3, 0},
		{"Ordinary item 4", -3, 40},
		{"Ordinary item 5", -3, -4},
		{"Sulfuras item 1", 34, 80},
		{"Sulfuras item 2", 1, 50},
		{"Sulfuras item 3", 4, 100},
		{"Aged Brie item 1", 0, 49},
		{"Aged Brie item 2", 3, 48},
		{"Aged Brie item 3", -1, 47},
		{"Aged Brie item 4", 0, 54},
		{"Backstage passes item 1", 13, 49},
		{"Backstage passes item 2", 10, 49},
		{"Backstage passes item 3", 5, 48},
		{"Backstage passes item 4", 2, 50},
		{"Backstage passes item 5", 10, 46},
		{"Backstage passes item 6", 5, 45},
		{"Backstage passes item 7", -1, 44},
		{"Backstage passes item 8", -1, 0},
		{"Backstage passes item 9", 0, 0},
		{"Conjured item 1", 1, 50},
		{"Conjured item 2", 2, 30},
		{"Conjured item 3", 3, 0},
		{"Conjured item 4", 3, 1},
		{"Conjured item 5", -3, 3},
		{"Conjured item 5", -3, 40},
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
			item.name != expected.name ||
			item.sellIn != expected.sellIn ||
			item.quality != expected.quality {
			t.Errorf("ERROR - Item %d - expected %s, but got %s", pos, expected, item)
			return
		}

		t.Logf("OK - Item %d - got %s", pos, item)
	}
}

// setupUpdateQuality sets up the expected values that should be returned by UpdateQuality
func setupUpdateQuality() []*Item {
	return []*Item{
		nil,
		{"Ordinary item 1", 0, 49},
		{"Ordinary item 2", 1, 0},
		{"Ordinary item 3", 2, 0},
		{"Ordinary item 4", -4, 38},
		{"Ordinary item 5", -4, -4},
		{"Sulfuras item 1", 34, 80},
		{"Sulfuras item 2", 1, 50},
		{"Sulfuras item 3", 4, 100},
		{"Aged Brie item 1", -1, 50},
		{"Aged Brie item 2", 2, 49},
		{"Aged Brie item 3", -2, 49},
		{"Aged Brie item 4", -1, 54},
		{"Backstage passes item 1", 12, 50},
		{"Backstage passes item 2", 9, 50},
		{"Backstage passes item 3", 4, 50},
		{"Backstage passes item 4", 1, 50},
		{"Backstage passes item 5", 9, 48},
		{"Backstage passes item 6", 4, 48},
		{"Backstage passes item 7", -2, 0},
		{"Backstage passes item 8", -2, 0},
		{"Backstage passes item 9", -1, 0},
		{"Conjured item 1", 0, 48},
		{"Conjured item 2", 1, 28},
		{"Conjured item 3", 2, 0},
		{"Conjured item 4", 2, 0},
		{"Conjured item 5", -4, 0},
		{"Conjured item 5", -4, 36},
	}
}
