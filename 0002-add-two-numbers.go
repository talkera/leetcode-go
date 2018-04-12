package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	v int
	next *ListNode
}

func main(){
	l1 := createListNode([]int{2,4,3})
	l2 := createListNode([]int{5,6,4})

	res := AddTwoNumbers(l1,l2)
	res.Print()
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, newNode *ListNode
	var over int
	for{
		tmpSum := over
		over = 0
		if l1 !=nil{
			tmpSum += l1.v
			l1 = l1.next
		}
		if l2 !=nil{
			tmpSum += l2.v
			l2 = l2.next
		}
		if tmpSum > 9{
			tmpSum -= 10
			over = 1
		}

		if res == nil{
			newNode = &ListNode{
				v:tmpSum,
			}
			res = newNode //res指向第一个node
		}else{
			newNode.next = &ListNode{
				v:tmpSum,
			}
			newNode = newNode.next //new总指向最新的node
		}
		if l1 == nil && l2 == nil {
			if over > 0{
				newNode.next = &ListNode{
					v:over,
				}
			}
			break
		}
	}
	return res
}
func (l *ListNode) Print() {
	var v string
	for {
		v = strconv.Itoa(l.v) + v
		if l.next==nil {
			break
		}
		l = l.next
	}
	fmt.Println(v)
}
func createListNode(num []int) *ListNode{
	if len(num)==0{
		return &ListNode{}
	}
	var pre *ListNode
	for _,v := range num{
		tmp := &ListNode{
			v:v,
		}
		if pre!=nil{
			tmp.next = pre
		}
		pre = tmp
	}

	return  pre
}