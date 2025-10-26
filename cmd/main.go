package main

import (
	"cmp"
	"errors"
	"fmt"
	"learning-go/binarytree"
	"learning-go/exercises"
	"learning-go/list"
)

func main() {
	exercises.Exercise_12_3()
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
	list := list.NewList[int]()
	list.Add(2)
	list.Add(4)
	list.Add(10)
	list.Insert(0, 23)
	list.Insert(1, 17)
	list.Insert(4, 18)
	fmt.Println("List Length:", list.Len())
	for it := list.Front(); it != nil; it = it.Next() {
		fmt.Println("Value:", it.Value)
	}
	fmt.Println("Index of 23:", list.Index(23))
}

func testBinaryTree() {
	tree := binarytree.NewTree(cmp.Compare[int])
	tree.Add(40)
	tree.Add(60)
	tree.Add(90)
	fmt.Printf("Tree contains 60: %v\n", tree.Contains(60))
}

func test9_1() {
	itemDesc, err := exercises.FindItem_9_1("")
	if errors.Is(err, exercises.ErrInvalidID) {
		fmt.Println("Oh ohh. You screwed up")
		return
	}
	fmt.Println("Surprise mf", itemDesc)
}

func test9_2() {
	saved, err := exercises.SaveEmployeeDescription_9_2(exercises.Employee{
		Name: "Zhalna",
	})
	if err != nil {
		var fieldMissingErr exercises.MissingFieldError
		if errors.As(err, &fieldMissingErr) {
			fmt.Println("Field Missing Error: ", fieldMissingErr.FieldName)
		}
		return
	}
	if saved {
		fmt.Println("Employee saved successfully")
	}
}

func test11_1() {
	exercises.PrintUDHRText()
}
