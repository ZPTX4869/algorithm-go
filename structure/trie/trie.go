package trie

const AlphabetSize = 26

type TrieNode struct {
	isEnd    bool
	children [AlphabetSize]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func New() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (t Trie) Put() (ok bool, err error) {
	return
}

func (t Trie) Get() (ok bool, err error) {
	return
}
