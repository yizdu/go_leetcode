package main

import (
	"fmt"
//	"go_leetcode/src/algo"
	"go_leetcode/src/leetcode"
)

func main() {
//	t1:=&algo.TreeNode{Val:1}
//	t1.Left=&algo.TreeNode{Val:2}
//	t1.Right=&algo.TreeNode{Val:3}
//	t1.Left.Left=&algo.TreeNode{Val:4}
//	t1.Left.Right=&algo.TreeNode{Val:5}
//	t1.Right.Left=&algo.TreeNode{Val:6}
//	t1.Right.Right=&algo.TreeNode{Val:7}
//	print(t1,"\n")
	fmt.Println(leetcode.CountPrimes0(500000))
	fmt.Println(leetcode.CountPrimes1(500000))
}
