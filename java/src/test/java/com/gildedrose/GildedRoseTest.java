package com.gildedrose;

import org.junit.jupiter.api.Test;

import java.util.TreeMap;
import java.util.stream.IntStream;

import static org.junit.jupiter.api.Assertions.*;

class GildedRoseTest {

    @Test
    void updateQuality() throws Exception {
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

        GildedRose app = new GildedRose(items);

        TreeMap<Integer, Item[]> tests = new TreeMap<>((value1, value2) -> -Integer.compare(value1, value2));
        tests.put(2, setupUpdateQuality());
        tests.put(-2, items.clone());

        app.updateQuality(0);

        tests.forEach((value, expectedItems) -> {
            app.updateQuality(value);

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
        });
    }
    
    private Item[] setupUpdateQuality() {
        return new Item[] {
                null,
                new Item("Ordinary item 1", -1, 47),
                new Item("Ordinary item 2", 0, 0),
                new Item("Ordinary item 3", 1, 0),
                new Item("Ordinary item 4", -5, 36),
                new Item("Ordinary item 5", -5, -4),
                new Item("Sulfuras item 1", 34, 80),
                new Item("Sulfuras item 2", 1, 50),
                new Item("Sulfuras item 3", 4, 100),
                new Item("Aged Brie item 1", -2, 50),
                new Item("Aged Brie item 2", 1, 50),
                new Item("Aged Brie item 3", -3, 50),
                new Item("Aged Brie item 4", -2, 54),
                new Item("Backstage passes item 1", 11, 50),
                new Item("Backstage passes item 2", 8, 50),
                new Item("Backstage passes item 3", 3, 50),
                new Item("Backstage passes item 4", 0, 50),
                new Item("Backstage passes item 5", 8, 50),
                new Item("Backstage passes item 6", 3, 50),
                new Item("Backstage passes item 7", -3, 0),
                new Item("Backstage passes item 8", -3, 0),
                new Item("Backstage passes item 9", -2, 0),
                new Item("Conjured item 1", -1, 44),
                new Item("Conjured item 2", 0, 26),
                new Item("Conjured item 3", 1, 0),
                new Item("Conjured item 4", 1, 0),
                new Item("Conjured item 5", -5, 0),
                new Item("Conjured item 6", -5, 32),


        };
    }
}