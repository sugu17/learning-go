package main

import (
	"cmp"
	"fmt"
	"learning-go/binarytree"
	"learning-go/exercises"
)

func main() {
	test8_3()
}

func test8_1() {
	fmt.Printf("Double value: %v\n", exercises.Double(89))
}

type OptimusPrime struct {
	name string
}

func (o OptimusPrime) String() string {
	return "I am Optimus prime...Epic BGM and you are my brothers!"
}

func test8_2() {
	optimus := OptimusPrime{name: "Optimus Prime"}
	exercises.GenericPrint(optimus)
}

func test8_3() {
	list := exercises.NewList[int]()
	list.Add(2)
	list.Add(4)
	list.Add(10)
	fmt.Println("List Length:", list.Len())
	for it := list.Front(); it != nil; it = it.Next() {
		fmt.Println("Value:", it.Value)
	}
}

func testBinaryTree() {
	tree := binarytree.NewTree(cmp.Compare[int])
	tree.Add(40)
	tree.Add(60)
	tree.Add(90)
	fmt.Printf("Tree contains 60: %v\n", tree.Contains(60))
}
