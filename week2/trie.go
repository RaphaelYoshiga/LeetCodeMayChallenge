package main;

type Trie struct {
    Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	isEndOfWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{ Root: &TrieNode{ Children: make(map[rune]*TrieNode) } };
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	currentNode := this.Root;

	l := len(word);
	for i, c := range word{ 
		newNode, alreadyExists := currentNode.Children[c];
		if alreadyExists{
 			if i == l-1{
				newNode.isEndOfWord = true;
			}
			currentNode = newNode;
		}else{
			
			node := &TrieNode{ Children: make(map[rune]*TrieNode), isEndOfWord: i == l-1}
			currentNode.Children[c] = node;
			currentNode = node;

		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	currentNode := this.Root;

	l := len(word);
	for i, c := range word{ 
		newNode, alreadyExists := currentNode.Children[c];
		if alreadyExists{
 			if i == l-1{
				return newNode.isEndOfWord;
			}
			currentNode = newNode;
		}else{
			return false;
		}
	}
	return false;
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	currentNode := this.Root;
	for _, c := range prefix{ 
		newNode, alreadyExists := currentNode.Children[c];
		if !alreadyExists{
			return false;
		}
		currentNode = newNode;
	}
	return true;
}

func main(){
	trie:= Constructor();
	trie.Insert("apple");
	trie.Search("apple");  // returns true
	trie.Search("app");     // returns false
	trie.StartsWith("app"); // returns true
	trie.Insert("app");   
	trie.Insert("app"); 
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */