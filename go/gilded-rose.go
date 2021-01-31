package main

import (
	"fmt"
	"strings"
)

const DEFAULT_STEP int = 1

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) String() string {
	// NOTE: The Java project has toString method included in the initial code,
	// so I've overridden it here as well to display the values of the items properly
	return fmt.Sprintf("{name: %s, sellIn: %d, quality: %d}", item.name, item.sellIn, item.quality)
}

type ItemHandler func(item *Item) int

var itemTypes map[string]ItemHandler = map[string]ItemHandler{
	"Aged Brie":        handleAgedBrie,
	"Backstage passes": handleBackstagePass,
	"Conjured":         handleConjured,
}

// UpdateQuality updates Gilded Rose Inn's inventory
func UpdateQuality(items []*Item) {
	for _, item := range items {
		handleItem(item)
	}
}

// handleItem manages the update of the provided item
func handleItem(item *Item) {
	if item == nil {
		return
	}

	if strings.HasPrefix(item.name, "Sulfuras") {
		return
	}

	item.sellIn--

	step := DEFAULT_STEP
	handler := getHandlerByItemType(item)
	if handler != nil {
		step = handler(item)
	} else {
		step = getMaxStep(item.quality, -doDoubleStep(item, step))
	}

	item.quality += step
	return
}

// getHandlerByItemType returns the handler associated to the type of the item provided
func getHandlerByItemType(item *Item) ItemHandler {
	for key, value := range itemTypes {
		if strings.HasPrefix(item.name, key) {
			return value
		}
	}

	return nil
}

// handleAgedBrie controls the specific changes that need to be applied to "Aged Brie" items
// when thet're updated
func handleAgedBrie(item *Item) int {
	// As far as I understand from the next statements:
	// 	-Once the sell by date has passed, Quality degrades twice as fast
	// 	-"Aged Brie" actually increases in Quality the older it gets
	// The quality of an "Aged Brie" item increases twice as fast when
	// its sell date has expired
	return getMaxStep(item.quality, doDoubleStep(item, DEFAULT_STEP))
}

// handleBackstagePass controls the specific changes that need to be applied to "Backstage Passes" items
// when thet're updated
func handleBackstagePass(item *Item) int {
	if item.sellIn < 0 {
		if item.quality > 0 {
			item.quality = 0
		}
		return 0
	}

	if item.sellIn > 10 {
		return getMaxStep(item.quality, DEFAULT_STEP)
	}

	if item.sellIn > 5 {
		return getMaxStep(item.quality, 2)
	}

	return getMaxStep(item.quality, 3)
}

// handleConjured controls the specific changes that need to be applied to "Conjured" items
// when thet're updated
func handleConjured(item *Item) int {
	return getMaxStep(item.quality, -doDoubleStep(item, 2*DEFAULT_STEP))
}

// getMaxStep returns the allowed step that not break min/max constraints
func getMaxStep(quality int, step int) int {
	minQuality := 0
	maxQuality := 50

	switch true {
	case step > 0:
		// From the next statements:
		// 	-"Just for clarification, an item can never have its Quality increase above 50"
		//	-"'Sulfuras' is a legendary item and as such its Quality is 80 and it never alters."
		// I understand that the quality can never increase above those values. Nevertheless,
		// they don't mention anything about the items whose initial quality is already greater
		// than the maximum, so I'm not sure if their value should be decreased to the maximum
		if quality >= maxQuality {
			return 0
		}

		qualityDiff := maxQuality - quality
		if step <= qualityDiff {
			return step
		}

		return qualityDiff
	case step < 0:
		// From the next statement:
		// 	-"The Quality of an item is never negative"
		// I understand that an item's quality cannot be lesser than 0. However...
		// what should we do with the items whose initial quality is already negative?
		if quality <= minQuality {
			return 0
		}

		qualityDiff := quality - minQuality
		if qualityDiff >= -step {
			return step
		}

		return -qualityDiff
	}

	return 0
}

// doDoubleStep checks whether the step must be increased twice and returns it
func doDoubleStep(item *Item, step int) int {
	if item.sellIn < 0 {
		return step * 2
	}

	return step
}
