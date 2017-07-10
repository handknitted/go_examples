package main

import (
	"github.com/handknitted/testbed/interface_one"
	"fmt"
)

func main() {
	interface_one.GetCountUp().Click()
	fmt.Printf("My Instance: %v\n",interface_one.GetCountUp())
	interface_one.GetCountUp().Click()
	fmt.Printf("My Instance: %v\n",interface_one.GetCountUp())
	interface_one.GetCountUp().UnClick()
	fmt.Printf("My Instance: %v\n",interface_one.GetCountUp())
	interface_one.GetCountUp().UnClick()
	fmt.Printf("My Instance: %v\n",interface_one.GetCountUp())
}