package com.gildedrose;

import java.util.Arrays;
import java.util.Map;
import java.util.Optional;
import java.util.function.BiFunction;

class GildedRose {
    Item[] items;

    private final int DEFAULT_STEP = 1;

    private final Map<String, BiFunction<Item, Integer, Integer>> itemTypes;

    public GildedRose(Item[] items) {
        this.items = items;

        this.itemTypes = Map.ofEntries(
                Map.entry("Aged Brie", this::handleAgedBrie),
                Map.entry("Backstage passes", this::handleBackstagePass),
                Map.entry("Conjured", this::handleConjured)
        );
    }

    public void updateQuality(int days) {
        if (days == 0) {
            return;
        }

        int reverse = days > 0 ? 1: -1;

        Arrays.stream(this.items)
                .forEach((item) -> this.handleItem(item, days, reverse));
    }

    private void handleItem(Item item, int days, int reverse) {
        if (item == null) {
            return;
        }

        if (item.name.startsWith("Sulfuras")) {
            return;
        }

        BiFunction<Item, Integer, Integer> handler = getHandlerByItemType(item);

        int daysLimit = Math.abs(days);
        for(int day = 0; day < daysLimit; day++) {
            item.sellIn--;
            int step = handler.apply(item, reverse);

            item.quality += step;
        }
    }

    private BiFunction<Item, Integer, Integer> getHandlerByItemType(Item item) {
        Optional<BiFunction<Item, Integer, Integer>> handler = this.itemTypes
                .entrySet()
                .stream()
                .filter(entry -> item.name.startsWith(entry.getKey()))
                .map(Map.Entry::getValue)
                .findFirst();

        return handler.orElse((it, reverse) -> getMaxStep(item.quality, -doDoubleStep(it, DEFAULT_STEP), reverse));
    }

    private int handleAgedBrie(Item item, int reverse) {
        return getMaxStep(item.quality, doDoubleStep(item, DEFAULT_STEP), reverse);
    }

    private int handleBackstagePass(Item item, int reverse) {
        if (item.sellIn < 0) {
            if (item.quality > 0) {
                item.quality = 0;
            }
            return 0;
        }

        if (item.sellIn > 10) {
            return getMaxStep(item.quality, DEFAULT_STEP, reverse);
        }

        if (item.sellIn > 5) {
            return getMaxStep(item.quality, 2, reverse);
        }

        return getMaxStep(item.quality, 3, reverse);
    }

    private int handleConjured(Item item, int reverse) {
        return getMaxStep(item.quality, -doDoubleStep(item, 2*DEFAULT_STEP), reverse);
    }

    private int getMaxStep(int quality, int step, int reverse) {
        int qualityDiff;
        final int minQuality = 0;
        final int maxQuality = 50;

        step *= reverse;

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