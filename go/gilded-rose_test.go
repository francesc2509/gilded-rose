package main

import (
	"fmt"
	"testing"
)

// TestUpdateQuality tests UpdateQuality function
func TestUpdateQuality(t *testing.T) {

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
		{"Conjured item 6", -3, 40},
	}

	itemsCopy := make([]*Item, len(items))
	fmt.Println(copy(itemsCopy, items))

	tests := map[int][]*Item{
		2:  setupUpdateQuality(),
		-2: itemsCopy,
	}

	UpdateQuality(items, 0)

	for days, expectedItems := range tests {

		UpdateQuality(items, days)
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
}

// setupUpdateQuality sets up the expected values that should be returned by UpdateQuality
func setupUpdateQuality() []*Item {
	return []*Item{
		nil,
		{name: "Ordinary item 1", sellIn: -1, quality: 47},
		{name: "Ordinary item 2", sellIn: 0, quality: 0},
		{name: "Ordinary item 3", sellIn: 1, quality: 0},
		{name: "Ordinary item 4", sellIn: -5, quality: 36},
		{name: "Ordinary item 5", sellIn: -5, quality: -4},
		{name: "Sulfuras item 1", sellIn: 34, quality: 80},
		{name: "Sulfuras item 2", sellIn: 1, quality: 50},
		{name: "Sulfuras item 3", sellIn: 4, quality: 100},
		{name: "Aged Brie item 1", sellIn: -2, quality: 50},
		{name: "Aged Brie item 2", sellIn: 1, quality: 50},
		{name: "Aged Brie item 3", sellIn: -3, quality: 50},
		{name: "Aged Brie item 4", sellIn: -2, quality: 54},
		{name: "Backstage passes item 1", sellIn: 11, quality: 50},
		{name: "Backstage passes item 2", sellIn: 8, quality: 50},
		{name: "Backstage passes item 3", sellIn: 3, quality: 50},
		{name: "Backstage passes item 4", sellIn: 0, quality: 50},
		{name: "Backstage passes item 5", sellIn: 8, quality: 50},
		{name: "Backstage passes item 6", sellIn: 3, quality: 50},
		{name: "Backstage passes item 7", sellIn: -3, quality: 0},
		{name: "Backstage passes item 8", sellIn: -3, quality: 0},
		{name: "Backstage passes item 9", sellIn: -2, quality: 0},
		{name: "Conjured item 1", sellIn: -1, quality: 44},
		{name: "Conjured item 2", sellIn: 0, quality: 26},
		{name: "Conjured item 3", sellIn: 1, quality: 0},
		{name: "Conjured item 4", sellIn: 1, quality: 0},
		{name: "Conjured item 5", sellIn: -5, quality: 0},
		{name: "Conjured item 6", sellIn: -5, quality: 32},
	}
}
