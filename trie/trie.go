package trie

import "strings"

type node struct {
	Data     string
	Children [26]*node
	IsEnding bool
}

func newNode(char string) *node {
	node := &node{
		Data:     char,
		Children: [26]*node{},
		IsEnding: false,
	}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

type Trie struct {
	Root *node
}

func (t *Trie) NewTrie() *Trie {
	root := newNode("\000") //根节点存任意东西
	return &Trie{Root: root}
}

// 新增
func (t *Trie) Insert(word string) {
	cur := t.Root
	//去除单词中到空格
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'
		if cur.Children[index] == nil {
			node := newNode(string(strippedWord[i]))
			cur.Children[index] = node
		}
		cur = cur.Children[index]
	}
	cur.IsEnding = true
}

// 查找
func (t *Trie) Find(word string) bool {
	cur := t.Root
	//去除单词中到空格
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'
		if cur.Children[index] == nil {
			return false
		}
		cur = cur.Children[index]
	}
	return cur.IsEnding
}

// 初始化
func (t *Trie) Init(words []string) {
	for _, v := range words {
		t.Insert(v)
	}
}
