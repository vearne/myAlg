package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func build(a []int, t []int, v int, i, j int) {
	if i >= j {
		t[v] = a[i]
	} else {
		m := (i + j) / 2
		build(a, t, v*2+1, i, m)
		build(a, t, v*2+2, m+1, j)
		t[v] = t[v*2+1] + t[v*2+2]
	}
}

func printTree(t []int) {
	fmt.Println(t)
	i := 0
	k := 1
	for i < len(t) {
		index := int(math.Exp2(float64(k))) - 1
		//fmt.Println("--", i, index)
		if i == index {
			fmt.Println()
			k++
		}
		fmt.Print(t[i], " ")
		i++
	}
}

func sum(t []int, v int, tl, tr int, l, r int) int {
	//fmt.Println("--sum--", "v=", v, "tl=", tl, "tr=", tr, "l=", l, "r=", r)
	if l > r || r < tl || l > tr {
		return 0
	}
	if tl == l && tr == r {
		return t[v]
	}

	tm := (tl + tr) / 2
	//fmt.Printf("sum(t, v*2+1, tl, tm, l, m)--sum(t, %v, %v, %v, %v, %v)\n", v*2+1, tl, tm, l, min(r, tm))
	//fmt.Printf("sum(t, v*2+2, tm+1, tr, m+1, r)--sum(t, %v, %v, %v, %v, %v)\n", v*2+2, tm+1, tr, max(tm+1, l), r)
	return sum(t, v*2+1, tl, tm, l, min(tm, r)) + sum(t, v*2+2, tm+1, tr, max(tm+1, l), r)
}

func update(t []int, v int, tl, tr int, pos int, val int) {
	if tl == tr {
		t[v] = val
		return
	}
	m := (tl + tr) / 2
	if pos <= m {
		update(t, v*2+1, tl, m, pos, val)
	} else {
		update(t, v*2+2, m+1, tr, pos, val)
	}
	t[v] = t[v*2+1] + t[v*2+2]
}

func buildTree(a []int) []int {
	n := len(a)
	t := make([]int, n*4)
	build(a, t, 0, 0, len(a)-1)
	return t
}

func main() {
	//var res, target int = 0
	//a := []int{3, 4, 5, 6}
	a := []int{3, 4, 5}
	t := buildTree(a)
	printTree(t)
	fmt.Println()
	//res = sum(t, 0, 0, len(a)-1, 0, 2)
	//target = 12
	//fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = sum(t, 0, 0, len(a)-1, 0, 1)
	//target = 7
	//fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = sum(t, 0, 0, len(a)-1, 1, 2)
	//target = 9
	//fmt.Println(res == target, "res=", res, "target=", target)
	//
	//res = sum(t, 0, 0, len(a)-1, 0, 3)
	//target = 12
	//fmt.Println(res == target, "res=", res, "target=", target)

	update(t, 0, 0, len(a)-1, 2, 10)
	printTree(t)
}
