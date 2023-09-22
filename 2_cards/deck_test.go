package main

import (
	"os"
	"testing"
)

// 测试扑克牌对象
// 1. 对象长度是否达到16
// 2. 第一张扑克牌是否为"Ace of Spades"
// 3. 最后一张扑克牌是否为"Four of Clubs"
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20,but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// 测试读取的文件内容
// 1. 生成扑克牌对象
// 2. 保存为文件"_decktesing"
// 3. 读取内容到 loadedDeck
// 4. loadedDeck 的长度是否为16
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesing")

	deck := newDeck()
	deck.saveToFile("_decktesing")

	loadedDeck := newDeckFromFile("_decktesing")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 20,but got %v", len(loadedDeck))
	}

	os.Remove("_decktesing")
}
