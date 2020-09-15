/*******************************************************************************
 * // Copyright AnchyTec Corp. All Rights Reserved.
 * // SPDX-License-Identifier: Apache-2.0
 * // Author: shaozhiming
 ******************************************************************************/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("service beginning...")

	l := &node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next: &node{
					value: 4,
					next:  nil,
				},
			},
		},
	}

	printList(l)

	fmt.Println("reverse list")

	newList := reverseList(l)

	printList(newList)

}

type node struct {
	value int
	next  *node
}

func printList(l *node) {
	for {
		fmt.Println(l.value)
		if l.next == nil {
			break
		}

		l = l.next
	}
}

func reverseList(head *node) *node {
	var newHead *node
	var cur *node

	for head != nil {
		cur = head
		head = head.next

		cur.next = newHead
		newHead = cur
	}

	return newHead
}

