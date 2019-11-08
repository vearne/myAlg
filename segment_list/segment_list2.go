package main

import (
	"container/list"
	"fmt"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Node struct {
	// tl, tr表示节点负责的范围
	tl    int
	tr    int
	value int
	lNode *Node
	rNode *Node
	level int
}

func BuildNode(a []int, tl, tr int, level int) *Node {
	if tl == tr {
		return &Node{tl, tr, a[tl], nil, nil, level}
	}
	currNode := &Node{tl, tr, 0, nil, nil, level}
	m := (tl + tr) / 2
	// create left node
	currNode.lNode = BuildNode(a, tl, m, level+1)
	// create right node
	currNode.rNode = BuildNode(a, m+1, tr, level+1)
	currNode.value = currNode.lNode.value + currNode.rNode.value
	return currNode
}

func BuildTree(a []int) *Node {
	return BuildNode(a, 0, len(a)-1, 1)
}

func PrintTree(root *Node) {
	var node *Node
	queue := NewQueue()
	queue.Push(root)
	level := 1
	for queue.Size() > 0 {
		node = queue.Pop()
		if node.level != level {
			fmt.Println(" ")
			level = node.level
		}
		fmt.Printf("%v[%v,%v]  ", node.value, node.tl, node.tr)
		if node.lNode != nil {
			queue.Push(node.lNode)
		}
		if node.rNode != nil {
			queue.Push(node.rNode)
		}
	}
	fmt.Printf("\n\n")
}

func SumRange(node *Node, l, r int) int {
	fmt.Printf("SumRange-[%v,%v], %v,%v\n", node.tl, node.tr, l, r)
	if l == node.tl && r == node.tr {
		return node.value
	}
	// case 1
	// |------------|------------|
	//    |-----|
	if r <= node.lNode.tr {
		return SumRange(node.lNode, l, r)
	}
	// case 2
	// |------------|------------|
	//    				|-----|
	if node.rNode.tl <= l {
		return SumRange(node.rNode, l, r)
	}
	// case 3
	// |------------|------------|
	//    		|--------|
	return SumRange(node.lNode, l, node.lNode.tr) + SumRange(node.rNode, node.rNode.tl, r)
}

func Update(node *Node, pos int, newVal int) {
	if node.tl == node.tr && node.tl == pos {
		node.value = newVal
		return
	}
	if pos <= node.lNode.tr {
		Update(node.lNode, pos, newVal)
	} else {
		Update(node.rNode, pos, newVal)
	}
	node.value = node.lNode.value + node.rNode.value
}

func main() {
	var res, target int
	a := []int{3, 4, 5}
	root := BuildTree(a)

	PrintTree(root)

	res = SumRange(root, 0, 1)
	target = 7
	fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = SumRange(root, 1, 3)
	//target = 15
	//fmt.Println(res == target, "res=", res, "target=", target)
	////
	//res = SumRange(root, 2, 3)
	//target = 11
	//fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = SumRange(root, 0, 3)
	//target = 18
	//fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = SumRange(root, 2, 2)
	//target = 5
	//fmt.Println(res == target, "res=", res, "target=", target)
	Update(root, 0, 10)
	PrintTree(root)
}

// ---

type Queue struct {
	list.List
}

func NewQueue() *Queue {
	return &Queue{}
}
func (q *Queue) Push(node *Node) {
	q.PushBack(node)
}

func (q *Queue) Pop() *Node {
	ele := q.Front()
	q.Remove(ele)
	return ele.Value.(*Node)
}

func (q *Queue) Size() int {
	return q.Len()
}
