package main

import (
	"fmt"
	"errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}) {
	if len(stack) == 0 {
		return nil
	} else {
		t := stack[len(stack)-1]
		return t
	}
}

func (stack *Stack) Pop() (interface{}, error){
	theStack := *stack
	if len(theStack) == 0{
		return nil, errors.New("out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}



func main(){
	var str string
	var mystack Stack
	fmt.Scan(&str)
	fmt.Printf("INPUT :%s\n", str)
	fmt.Printf("%d\n", len(str))
	for i:= 0; i<len(str); i++{
		if str[i]=='{'||str[i]=='['||str[i]=='(' {
			mystack.Push(string(str[i]))
			fmt.Println(mystack.Top())
		} else if string(str[i])=="}" && mystack.Top()==string("{") {
				mystack.Pop()
		} else if string(str[i])=="]" && mystack.Top()==string("[") {
				mystack.Pop()
		} else if string(str[i])==")" && mystack.Top()==string("(") {
				mystack.Pop()
		}

	}
	if mystack.IsEmpty() {
		fmt.Printf("match")
	}else{
		fmt.Printf("notmatch")
	}

}
