package com.gildedrose;

import org.junit.jupiter.api.Test;

import java.util.Comparator;
import java.util.stream.IntStream;

import static org.junit.jupiter.api.Assertions.*;

class GildedRoseTest {

    @Test
    void updateQuality() {
        Item[] items = new Item[]{
            null,
            new Item("Ordinary item 1", 1, 50),
            new Item("Ordinary item 2", 2, 1),
            new Item("Ordinary item 3", 3, 0),
            new Item("Ordinary item 4", -3, 40),
            new Item("Ordinary item 5", -3, -4),
            new Item("Sulfuras item 1", 34, 80),
            new Item("Sulfuras item 2", 1, 50),
            new Item("Sulfuras item 3", 4, 100),
            new Item("Aged Brie item 1", 0, 49),
            new Item("Aged Brie item 2", 3, 48),
            new Item("Aged Brie item 3", -1, 47),
            new Item("Aged Brie item 4", 0, 54),
            new Item("Backstage passes item 1", 13, 49),
            new Item("Backstage passes item 2", 10, 49),
            new Item("Backstage passes item 3", 5, 48),
            new Item("Backstage passes item 4", 2, 50),
            new Item("Backstage passes item 5", 10, 46),
            new Item("Backstage passes item 6", 5, 45),
            new Item("Backstage passes item 7", -1, 44),
            new Item("Backstage passes item 8", -1, 0),
            new Item("Backstage passes item 9", 0, 0),
            new Item("Conjured item 1", 1, 50),
            new Item("Conjured item 2", 2, 30),
            new Item("Conjured item 3", 3, 0),
            new Item("Conjured item 4", 3, 1),
            new Item("Conjured item 5", -3, 3),
            new Item("Conjured item 6", -3, 40),
        };
        Item[] expectedItems = setupUpdateQuality();

        GildedRose app = new GildedRose(items);
        app.updateQuality();

        IntStream.range(1, items.length)
                .forEach((idx) -> {
                    Item item = items[idx];
                    Item expected = expectedItems[idx];

                    if (expected == item) {
                        return;
                    }

                    assertTrue(expected != null && item != null &&
                                    expected.name.equals(item.name) &&
                                    expected.sellIn == item.sellIn &&
                                    expected.quality == item.quality,
                            String.format("ERROR - Item %d - Name: expected %s, but got %s", idx, expected, item)
                    );
                });
    }
    
    private Item[] setupUpdateQuality() {
        return new Item[] {
            null,
            new Item("Ordinary item 1", 0, 49),
            new Item("Ordinary item 2", 1, 0),
            new Item("Ordinary item 3", 2, 0),
            new Item("Ordinary item 4", -4, 38),
            new Item("Ordinary item 5", -4, -4),
            new Item("Sulfuras item 1", 34, 80),
            new Item("Sulfuras item 2", 1, 50),
            new Item("Sulfuras item 3", 4, 100),
            new Item("Aged Brie item 1", -1, 50),
            new Item("Aged Brie item 2", 2, 49),
            new Item("Aged Brie item 3", -2, 49),
            new Item("Aged Brie item 4", -1, 54),
            new Item("Backstage passes item 1", 12, 50),
            new Item("Backstage passes item 2", 9, 50),
            new Item("Backstage passes item 3", 4, 50),
            new Item("Backstage passes item 4", 1, 50),
            new Item("Backstage passes item 5", 9, 48),
            new Item("Backstage passes item 6", 4, 48),
            new Item("Backstage passes item 7", -2, 0),
            new Item("Backstage passes item 8", -2, 0),
            new Item("Backstage passes item 9", -1, 0),
            new Item("Conjured item 1", 0, 48),
            new Item("Conjured item 2", 1, 28),
            new Item("Conjured item 3", 2, 0),
            new Item("Conjured item 4", 2, 0),
            new Item("Conjured item 5", -4, 0),
            new Item("Conjured item 6", -4, 36),
        };
    }
}