package com.gildedrose;

import java.io.IOException;
import java.util.Arrays;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.function.BiFunction;
import java.util.function.Function;

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

        Arrays.stream(this.items)
                .forEach((item) -> this.handleItem(item, days));
    }

    private void handleItem(Item item, int days) {
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
            int step = handler.apply(item, days);

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

        return handler.orElse((it, days) -> getAllowedStep(item.quality, -doDoubleStep(it, DEFAULT_STEP), days));
    }

    private int handleAgedBrie(Item item, Integer days) {
        return getAllowedStep(item.quality, doDoubleStep(item, DEFAULT_STEP), days);
    }

    private int handleBackstagePass(Item item, Integer days) {
        if (item.sellIn < 0) {
            if (item.quality > 0) {
                item.quality = 0;
            }
            return 0;
        }

        if (item.sellIn > 10) {
            return getAllowedStep(item.quality, DEFAULT_STEP, days);
        }

        if (item.sellIn > 5) {
            return getAllowedStep(item.quality, 2, days);
        }

        return getAllowedStep(item.quality, 3, days);
    }

    private int handleConjured(Item item, Integer days) {
        return getAllowedStep(item.quality, -doDoubleStep(item, 2*DEFAULT_STEP), days);
    }

    private int getAllowedStep(int quality, int step, int days) {
        int qualityDiff;
        final int minQuality = 0;
        final int maxQuality = 50;

        if (days < 0) {
            step = -step;
        }

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