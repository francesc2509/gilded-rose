package com.gildedrose;

import java.util.Arrays;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.function.Function;

class GildedRose {
    Item[] items;

    private final int DEFAULT_STEP = 1;

    private final Map<String,Function<Item, Integer>> itemTypes;


    public GildedRose(Item[] items) {
        this.items = items;

        this.itemTypes = Map.ofEntries(
                Map.entry("Aged Brie", this::handleAgedBrie),
                Map.entry("Backstage passes", this::handleBackstagePass),
                Map.entry("Conjured", this::handleConjured)
        );
    }

    public void updateQuality() {
        Arrays.stream(this.items)
                .filter(Objects::nonNull)
                .forEach(this::handleItem);
    }

    private void handleItem(Item item) {
        if (item == null) {
            return;
        }

        if (item.name.startsWith("Sulfuras")) {
            return;
        }

        item.sellIn--;

        int step = DEFAULT_STEP;
        Function<Item, Integer> handler = getHandlerByItemType(item);

        if (handler != null) {
            step = handler.apply(item);
        } else {
            step = getAllowedStep(item.quality, -doDoubleStep(item, step));
        }

        item.quality += step;
    }

    private Function<Item, Integer> getHandlerByItemType(Item item) {
        Optional<Function<Item, Integer>> handler = this.itemTypes
                .entrySet()
                .stream()
                .filter(entry -> item.name.startsWith(entry.getKey()))
                .map(Map.Entry::getValue)
                .findFirst();

        return handler.orElse(null);
    }

    private int handleAgedBrie(Item item) {
        return getAllowedStep(item.quality, doDoubleStep(item, DEFAULT_STEP));
    }

    private int handleBackstagePass(Item item) {
        if (item.sellIn < 0) {
            if (item.quality > 0) {
                item.quality = 0;
            }
            return 0;
        }

        if (item.sellIn > 10) {
            return getAllowedStep(item.quality, DEFAULT_STEP);
        }

        if (item.sellIn > 5) {
            return getAllowedStep(item.quality, 2);
        }

        return getAllowedStep(item.quality, 3);
    }

    private int handleConjured(Item item) {
        return getAllowedStep(item.quality, -doDoubleStep(item, 2*DEFAULT_STEP));
    }

    private int getAllowedStep(int quality, int step) {
        int qualityDiff;
        final int minQuality = 0;
        final int maxQuality = 50;

        if (step > 0) {
            if (quality >= maxQuality) {
                return 0;
            }

            qualityDiff = maxQuality - quality;
            return Math.min(step, qualityDiff);
        }

        if (step < 0) {
            if (quality <= minQuality) {
                return 0;
            }

            qualityDiff = quality - minQuality;
            if (qualityDiff >= -step) {
                return step;
            }

            return -qualityDiff;
        }

        return 0;
    }

    private int doDoubleStep(Item item, int step) {
        if (item.sellIn < 0) {
            return step * 2;
        }

        return step;
    }
}