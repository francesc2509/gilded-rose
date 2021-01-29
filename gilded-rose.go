package main

import (
	"strings"
)

const MIN_QUALITY int = 0
const MAX_QUALITY int = 50
const DECREASE_AMOUNT int = 1

type Item struct {
	name            string
	sellIn, quality int
}

type ItemHandler func(item *Item) *Item

var itemTypes map[string]ItemHandler = map[string]ItemHandler{
	"Aged Brie":        agedBrieHandler,
	"Backstage passes": backstagePassHandler,
	"Conjured":         conjuredHandler,
}

var conjuredHandler ItemHandler = func(item *Item) *Item {
	qualityDecrease := 2 * DECREASE_AMOUNT

	if item.quality == 1 {
		item.quality--

		return item
	}

	if item.sellIn < 0 {
		item.quality -= (qualityDecrease * 2)

		return item
	}

	if item.quality == 1 {
		qualityDecrease = 1
	}
	item.quality -= qualityDecrease

	return item
}

var backstagePassHandler ItemHandler = func(item *Item) *Item {
	if item.sellIn < 0 && item.quality > 0 {
		item.quality = 0
		return item
	}

	if item.quality == 49 {
		item.quality += 1
		return item
	}

	if item.sellIn > 10 {
		item.quality += DECREASE_AMOUNT
	}

	if item.sellIn > 5 {
		item.quality += 2
		return item
	}

	item.quality += 3
	return item
}

var agedBrieHandler ItemHandler = func(item *Item) *Item {
	item.quality++

	return item
}

func (item *Item) update() *Item {
	item.sellIn -= DECREASE_AMOUNT

	return item
}

func (item *Item) handle() {
	if strings.HasPrefix(item.name, "Sulfuras") {
		return
	}

	item.update()

	if item.quality >= MAX_QUALITY || item.quality <= MIN_QUALITY {
		return
	}

	handler := getHandlerByItemType(item)
	if handler != nil {
		handler(item)
		return
	}

	item.quality -= DECREASE_AMOUNT
	return
}

func getHandlerByItemType(item *Item) ItemHandler {
	for key, value := range itemTypes {
		if strings.HasPrefix(item.name, key) {
			return value
		}
	}

	return nil
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		handleItem(item)
	}
}

func handleItem(item *Item) {
	if item == nil {
		return
	}

	item.handle()
}
