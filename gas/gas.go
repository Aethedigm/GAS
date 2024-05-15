package gas

import "sync"

type GAS struct {
	Lock sync.RWMutex

	Node trie
}

type trie struct {
	value    []string
	children map[rune]*trie
	word     string
}

func (t *trie) insert(word, value string) {
	node := t
	for _, r := range word {
		next := node.children[r]
		if next == nil {
			if node.children == nil {
				node.children = map[rune]*trie{}
			}
			next = new(trie)
			node.children[r] = next
		}
		node = next
	}

	if node.word != word {
		node.word = word
	}
	node.value = append(node.value, value)
}

func (t *trie) collectAll() []*TrieResult {
	var result []*TrieResult

	if t.word != "" && len(t.value) > 0 {
		result = append(result, &TrieResult{
			Word:   t.word,
			Values: t.value,
		})
	}

	for _, child := range t.children {
		result = append(result, child.collectAll()...)
	}

	return result
}

type TrieResult struct {
	Word   string   `json:"word"`
	Values []string `json:"values"`
}

func (t *trie) retrieve(word string) []*TrieResult {
	node := t
	for _, r := range word {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}

	return node.collectAll()
}

type ResultResponse struct {
	Query   string        `json:"query"`
	Results []*TrieResult `json:"results"`
}

func (g *GAS) AddResult(key, value string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()

	g.Node.insert(key, value)
}

func (g *GAS) GetResults(query string) ResultResponse {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	results := g.Node.retrieve(query)

	rr := ResultResponse{
		Query:   query,
		Results: results,
	}

	return rr
}
