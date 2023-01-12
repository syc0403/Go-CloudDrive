package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := 0, 0
	p := 1
	for l1 != nil {
		n1 += l1.Val * p
		p *= 10
		l1 = l1.Next
	}
	p = 1
	for l2 != nil {
		n2 = l2.Val * p
		p *= 10
		l2 = l2.Next
	}
	return

}
func main() {
	fmt.Println("")
}
