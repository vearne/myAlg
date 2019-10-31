package main

import (
	"fmt"
	"math/rand"
)

const MaxLevel = 10
const Probability = 0.5 // 基于时间与空间综合 best practice 值, 越上层概率越小

func init() {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
}

func randomLevel() int {
	level := 1
	for rand.Float32() < Probability && level <= MaxLevel {
		level++
	}
	return level
}

type Node struct {
	forward []*Node
	key     int
}

type SkipList struct {
	level int
	head  *Node
}

func NewNode(key int, level int) *Node {
	return &Node{forward: make([]*Node, level), key: key}
}

func NewSkipList() *SkipList {
	return &SkipList{level: 1, head: NewNode(-1, MaxLevel)}
}

func (sl *SkipList) Insert(key int) bool {
	var p *Node
	prev := make([]*Node, MaxLevel)

	for i := sl.level - 1; i >= 0; i-- {
		p = sl.head
		for p.forward[i] != nil && p.forward[i].key <= key {
			p = p.forward[i]
		}
		prev[i] = p
		if p.key == key {
			// 不允许重复
			return false
		}
	}

	level := randomLevel()
	fmt.Println("random level=", level)
	if level > sl.level {
		for i := level - 1; i > sl.level-1; i-- {
			prev[i] = sl.head
		}
		sl.level = level
	}
	node := NewNode(key, level)

	for i := level - 1; i >= 0; i-- {
		node.forward[i] = prev[i].forward[i]
		prev[i].forward[i] = node
		fmt.Printf("[execute] level %d: %d->%d\n", i+1, prev[i].key, node.key)
	}

	return true
}

func (sl *SkipList) Search(key int) bool {
	var p *Node
	p = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for p.forward[i] != nil && p.forward[i].key <= key {
			p = p.forward[i]
		}
		if p.key == key {
			return true
		}
	}
	return false
}

func (sl *SkipList) Delete(key int) bool {
	var p *Node
	var result bool = false
	prev := make([]*Node, sl.level)
	for i := sl.level - 1; i >= 0; i-- {
		p = sl.head
		for p.forward[i] != nil && p.forward[i].key < key {
			p = p.forward[i]
		}
		prev[i] = p
		node := prev[i].forward[i]
		if node != nil && node.key == key {
			prev[i].forward[i] = node.forward[i]
			result = true
		}
		if sl.head.forward[i] == nil {
			sl.level--
		}
	}
	return result
}

func (sl *SkipList) Print() {
	fmt.Println()
	var p *Node

	for i := sl.level - 1; i >= 0; i-- {
		p = sl.head
		fmt.Printf("Level %d:", i+1)
		for p != nil {
			if p.key != -1 {
				fmt.Print(p.key)
				fmt.Print(" ")
			}

			p = p.forward[i]
		}
		fmt.Println("")
		fmt.Println("############")
	}
}

func main() {
	var value int
	var result bool
	length := 0
	list := NewSkipList()
	for i := 0; i < 20; i++ {
		fmt.Println("\n")
		value := rand.Intn(100)
		fmt.Println("insert", value)
		result = list.Insert(value)
		fmt.Println("insert", value, result)
		if result {
			length++
		}
		list.Print()
		fmt.Println("\n--------------------------------------", length)

	}
	//list.Print()

	fmt.Println("\n--------------------------------------")
	//value = 88
	//list.Delete(value)
	//fmt.Println("delete", value)
	//list.Print()
	//
	fmt.Println("\n--------------------------------------")
	//list.Delete(6)
	//list.Print()
	//
	//list.Delete(12)
	//list.Print()
	//
	//	//list.Delete(94)
	//	//list.Print()
	value = 88
	result = list.Search(value)
	fmt.Println("value=%d", value, result)

	//value = 88
	//list.Delete(value)
	//fmt.Println("delete", value)

	value = 88
	result = list.Search(value)
	fmt.Println("value=%v", value, result)

	value = 100
	result = list.Search(value)
	fmt.Println("value=%v", value, result)

	value = 87
	result = list.Search(value)
	fmt.Println("value=%v", value, result)

}
