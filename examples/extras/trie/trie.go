package trie

import "sync"

// Trie Holder
type Trie struct {
	letter   rune
	mutex    sync.RWMutex
	children []*Trie
	meta     map[string]interface{}
	isLeaf   bool
	//parent   *Trie
}

// NewTrie creates a new Trie and initialize it with default options
// May be used to create a new root trie.
func NewTrie() *Trie {
	newTrie := &Trie{}
	newTrie.children = []*Trie{}
	newTrie.meta = make(map[string]interface{})

	return newTrie
}

func (tr *Trie) hasChild(a rune) (bool, *Trie) {
	for _, child := range tr.children {
		if child.letter == a {
			return true, child
		}
	}

	return false, nil
}

func (tr *Trie) addChild(a rune) *Trie {
	tr.mutex.Lock()

	nw := NewTrie()
	nw.letter = a
	//nw.parent = tr
	tr.children = append(tr.children, nw)

	tr.mutex.Unlock()

	return nw
}

// Add a word to a trie
func (tr *Trie) Add(word string) *Trie {
	letters, node, i := []rune(word), tr, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			node = node.addChild(letters[i])
		}

		i++

		if i == n {
			tr.mutex.Lock()
			node.isLeaf = true
			tr.mutex.Unlock()
		}
	}

	return node
}

// FindNode returns the node whether it is a word or not.
func (tr *Trie) FindNode(word string) *Trie {
	letters, node, i := []rune(word), tr, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			return nil
		}

		i++
	}

	return node
}

// Get metadata belonging to a node.
func (tr *Trie) Get(key string) (interface{}, bool) {
	if tr == nil {
		return nil, false
	}

	if _, ok := tr.meta[key]; ok {
		return tr.meta[key], true
	}

	return nil, false
}

// Set metadata for a word
func (tr *Trie) Set(key string, value interface{}) {
	if tr == nil {
		return
	}

	tr.mutex.Lock()
	tr.meta[key] = value
	tr.mutex.Unlock()
}

func (tr *Trie) GetPathsFromHere() []string {
	children := []string{}
	if tr.isLeaf {
		children = append(children, string(tr.letter))
	}
	for _, c := range tr.children {
		for _, paths := range c.GetPathsFromHere() {
			children = append(children, string(tr.letter)+paths)
		}
	}
	return children
}
