package main

import (
	"fmt"
	"strings"
)

const MIN_QUALITY int = 0
const MAX_QUALITY int = 50
const QUALITY_STEP int = 1

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) String() string {
	return fmt.Sprintf("Name: %s, SellIn: %d, Quality: %d", item.name, item.sellIn, item.quality)
}

type ItemHandler func(item *Item) int

var itemTypes map[string]ItemHandler = map[string]ItemHandler{
	"Aged Brie":        agedBrieHandler,
	"Backstage passes": backstagePassHandler,
	"Conjured":         conjuredHandler,
}

func updateItem(item *Item) *Item {
	item.sellIn -= QUALITY_STEP

	return item
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

	if strings.HasPrefix(item.name, "Sulfuras") {
		return
	}

	updateItem(item)

	step := QUALITY_STEP
	handler := getHandlerByItemType(item)
	if handler != nil {
		step = handler(item)
	} else {
		step = getAllowedStep(item.quality, -doDoubleStep(item, step))
	}

	item.quality += step
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

func agedBrieHandler(item *Item) int {
	return getAllowedStep(item.quality, doDoubleStep(item, QUALITY_STEP))
}

func backstagePassHandler(item *Item) int {
	if item.sellIn < 0 {
		if item.quality > 0 {
			item.quality = 0
		}
		return 0
	}

	if item.sellIn > 10 {
		return getAllowedStep(item.quality, QUALITY_STEP)
	}

	if item.sellIn > 5 {
		return getAllowedStep(item.quality, 2)
	}

	return getAllowedStep(item.quality, 3)
}

func conjuredHandler(item *Item) int {
	return getAllowedStep(item.quality, -doDoubleStep(item, 2*QUALITY_STEP))
}

func getAllowedStep(quality int, step int) int {

	switch true {
	case step > 0:
		if quality >= MAX_QUALITY {
			return 0
		}

		qualityDiff := MAX_QUALITY - quality
		if step <= qualityDiff {
			return step
		}

		return qualityDiff
	case step < 0:
		if quality <= MIN_QUALITY {
			return 0
		}

		qualityDiff := quality - MIN_QUALITY
		if qualityDiff >= -step {
			return step
		}

		return -qualityDiff
	}

	return 0
}

func doDoubleStep(item *Item, step int) int {
	if item.sellIn < 0 {
		return step * 2
	}

	return step
}
