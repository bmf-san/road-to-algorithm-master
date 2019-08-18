/////////////////////////////////////////
// Warning！
// This is an incomplete implementation.
/////////////////////////////////////////

package main

import (
	"fmt"
	"sort"
)

// Node is a node of tree.
type Node struct {
	key      string
	children []*Node
}

// hasChildren is check whether a node is leaf node.
func hasChildren(n *Node) bool {
	if len(n.children) == 0 {
		// This means leaf node
		return false
	}

	return true
}

// Insert is insert a node.
func (n *Node) Insert(word string) {
	if len(n.key) == 0 && len(n.children) == 0 {
		n.key = word

		return
	}

	recursiveInsert(word, n)
}

func recursiveInsert(word string, n *Node) {
	shorterWord, longerWord := []rune(word), []rune(n.key)
	if len(longerWord) < len(shorterWord) {
		shorterWord, longerWord = longerWord, shorterWord
	}

	for i := 0; i < len(shorterWord); i++ {
		if longerWord[i] == shorterWord[i] {
			// last loop
			if i == len(shorterWord)-1 {
				if len(shorterWord) == len(longerWord) {
					// exact match
					break
				} else {
					word = string(longerWord[i+1:])

					// single node prefix match
					if !hasChildren(n) {
						n.key = string(shorterWord)
						n.children = append(n.children,
							&Node{
								key: word,
							},
						)
					}

					for j, v := range n.children {
						if []rune(v.key)[0] == []rune(word)[0] {
							recursiveInsert(word, v)
						} else {
							if j == len(n.children)-1 {
								// mismatch
								n.children = append(n.children,
									&Node{
										key: word,
									},
								)
								return
							}
						}
					}
				}
			} else {
				continue
			}
		} else {
			if i == 0 {
				// mismatch
				s := []string{n.key, word}
				sort.Strings(s)

				n.children = append(n.children,
					&Node{
						key: n.key,
						children: []*Node{
							&Node{
								key: s[0],
							},
							&Node{
								key: s[1],
							},
						},
					},
				)
				n.key = ""

				return
			}

			if !hasChildren(n) {
				// double node prefix match
				n.key = string(shorterWord[:i])

				s := []string{string(shorterWord[i:]), string(longerWord[i:])}
				sort.Strings(s)

				n.children = append(n.children,
					&Node{
						key: string(s[0]),
					},
					&Node{
						key: string(s[1]),
					},
				)
				return
			}

			for k, v := range n.children {
				if []rune(v.key)[0] == []rune(word)[0] {
					recursiveInsert(word[:i], v)
				}

				if k == len(n.children)-1 {
					// mismatch
					n.children = append(n.children,
						&Node{
							key: word,
						},
					)
					return
				}
			}
		}
	}
}

func main() {
	n := &Node{}

	n.Insert("post")
	n.Insert("poem")
	n.Insert("roll")
	// n.Insert("word")
	// n.Insert("rd")
	// n.Insert("or")

	fmt.Printf("%v", n)
}
