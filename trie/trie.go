package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	Num       int32         //统计走过路径数量
	ChildNode [26]*TrieNode //26个字母就存在26个子点 思考：如果存在中文是不是使用[]rune
	Value     byte
	IsEnd     bool //钟点词
}

var rootNode *TrieNode

func (root *TrieNode) Insert(str string) {

	node := root
	tmpSlice := []byte(strings.ToLower(str))

	for _, v := range tmpSlice {
		i := v - 'a'
		if node.ChildNode[i] == nil {
			node.ChildNode[i] = &TrieNode{1, [26]*TrieNode{}, v, false}
			node = node.ChildNode[i]
		} else {
			node = node.ChildNode[i]
			node.Num++
		}

	}

	node.IsEnd = true

}

func (root *TrieNode) Search(str string) bool {

	node := root
	tmpSlice := []byte(strings.ToLower(str))
	for _, v := range tmpSlice {
		i := v - 'a'
		if node.ChildNode[i] == nil {
			return false
		}
		node = node.ChildNode[i]
	}

	return true
}

func (root *TrieNode) Path(i int8) {
	for _, node := range root.ChildNode {
		if node != nil {
			fmt.Printf("%s%c \n", strings.Repeat("|  ", int(i)), node.Value)
			node.Path(i + 1)
		}
	}
}

func main() {
	rootNode = new(TrieNode)
	rootNode.Insert("This")
	rootNode.Insert("is")
	rootNode.Insert("a")
	rootNode.Insert("little")
	rootNode.Insert("girl")
	fmt.Printf("%v \n", rootNode.Search("a"))
	fmt.Printf("%v \n", rootNode.Search("little"))
	fmt.Printf("%v \n", rootNode.Search("abc"))
	rootNode.Path(1)
}
