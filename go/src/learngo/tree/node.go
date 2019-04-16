package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value} // here is a local variable in c++, it can not be used at main function in c++,
									// but it can be used at go, the position(stack or pile) to be stored is decide by go
}

func (node Node) Print() { // the code before function name is receiver, it like a parameter, usage: node.print()
	fmt.Print(node.Value)
	fmt.Print(" ")
}

func (node *Node) SetValue(value int)  { // if we wanna change value, we need to set the pointer,
											  // but when we use, it does not matter
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value

}








