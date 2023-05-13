package main

import (
	"bbst/helpers"
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

type Tree struct {
	Root *Node
}

var bstStr string

func main() {
	var count int
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&count)

	num := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&num[i])
	}
	sort.Ints(num)
	arr := helpers.RemoveDuplicates(num)

	fmt.Println("Numbers entered:", arr)

	bst := CreateBST(arr)

	fmt.Println("Balanced Binaray Tree is : ")
	string := PrintTree(bst, 0)

	helpers.WriteToFile("files/output.txt", string)
	fmt.Println("String saved to file: output.txt")

	var ins int
	fmt.Printf("Enter new unique number to insert into node ")
	fmt.Scan(&ins)

	arr = append(arr, ins)
	arr = helpers.RemoveDuplicates(arr)

	fmt.Println("New array after new insersion:", arr)

	bst = CreateBST(arr)
	bstStr = ""
	fmt.Println("Balanced Binaray Tree is : ")
	string = PrintTree(bst, 0)

	helpers.WriteToFile("files/output-after-insertion.txt", string)
	fmt.Println("String saved to file: output-after-insertion.txt")

	var new, old int
	fmt.Printf("Enter the number you want to replace : ")
	fmt.Scan(&old)
	fmt.Printf("Enter the number you want to replace with : ")
	fmt.Scan(&new)

	arr = helpers.UpdateArray(arr, old, new)

	arr = helpers.RemoveDuplicates(arr)

	fmt.Println("New array after new updating:", arr)

	bst = CreateBST(arr)
	bstStr = ""
	fmt.Println("Balanced Binaray Tree is : ")
	string = PrintTree(bst, 0)

	helpers.WriteToFile("files/output-after-updation.txt", string)
	fmt.Println("String saved to file: output-after-updation.txt")

	var del int
	fmt.Printf("Enter the number you want to delete : ")
	fmt.Scan(&del)

	arr = helpers.DeleteFromArray(arr, del)

	fmt.Println("New array after deletion:", arr)

	bst = CreateBST(arr)
	bstStr = ""
	fmt.Println("Balanced Binaray Tree is : ")
	string = PrintTree(bst, 0)

	helpers.WriteToFile("files/output-after-deletion.txt", string)
	fmt.Println("String saved to file: output-after-deletion.txt")

}

func CreateBST(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2

	root := &Node{
		Key:   arr[mid],
		Left:  CreateBST(arr[:mid]),
		Right: CreateBST(arr[mid+1:]),
	}

	return root
}

func PrintTree(root *Node, level int) string {
	if root == nil {
		return bstStr
	}

	PrintTree(root.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("    ")
		bstStr += "     "
	}

	bstStr += strconv.Itoa(root.Key) + "\n"
	fmt.Println(root.Key)

	PrintTree(root.Left, level+1)
	return bstStr
}
