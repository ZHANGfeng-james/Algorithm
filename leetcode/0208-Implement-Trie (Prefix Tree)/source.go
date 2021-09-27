package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0208. 实现前缀树\n"
)

func init() {
	fmt.Println(topic)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{}, // 如果某个 index 存在，说明包含该位置索引对应的字符
		isEnd:    false,
	}
}

func (trie *Trie) Insert(word string) {
	node := trie
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (trie *Trie) SearchPrefix(prefix string) *Trie {
	node := trie
	for _, ch := range prefix {
		ch -= 'a'
		children := node.children[ch]
		if children == nil {
			return nil
		}
		node = children
	}
	return node
}

func (trie *Trie) Search(word string) bool {
	node := trie.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.SearchPrefix(prefix)
	return node != nil
}
