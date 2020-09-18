package linkedList

import (
	"fmt"
	"testing"
)

var listNull = &List{}

func TestList_IsEmpty(t *testing.T) {
	fmt.Println(listNull.IsEmpty())

	listNull.Add(0)
	fmt.Println(listNull.IsEmpty())
}

func TestList_Length(t *testing.T) {
	listNull.Add(0)
	listNull.Add(0)
	listNull.Add(0)

	fmt.Println(listNull.Length())
}

func TestList_Append(t *testing.T) {
	listNull.Show()

	listNull.Append(0)
	listNull.Show()

	listNull.Append(1)
	listNull.Show()
}

func TestList_Add(t *testing.T) {
	listNull.Show()

	listNull.Add(0)
	listNull.Show()

	listNull.Add(1)
	listNull.Show()
}

func TestList_AddAppend(t *testing.T) {
	listNull.Add(3)
	listNull.Append(4)
	listNull.Append(8)
	listNull.Add(5)
	listNull.Append(6)

	//5-> 3-> 4-> 8-> 6
	listNull.Show()
}

func TestList_Insert(t *testing.T) {
	listNull.Append(4)
	listNull.Append(8)
	listNull.Append(6)
	listNull.Show()

	listNull.Insert(-1, 2)
	listNull.Show()

	listNull.Insert(0, 3)
	listNull.Show()

	listNull.Insert(1, 9)
	listNull.Show()


	listNull.Insert(2, 7)
	listNull.Show()

	listNull.Insert(7, 10)
	listNull.Show()
}

func TestList_Remove(t *testing.T) {
	listNull.Append(4)
	listNull.Append(6)
	listNull.Append(8)
	listNull.Show()

	listNull.Remove(4)
	listNull.Show()

	listNull.Append(10)
	listNull.Append(12)
	listNull.Show()

	listNull.Remove(8)
	listNull.Show()

	listNull.Remove(12)
	listNull.Show()
}

func TestList_RemoveAtIndex(t *testing.T) {
	listNull.Append(4)
	listNull.Append(6)
	listNull.Append(8)
	listNull.Append(10)
	listNull.Show()

	listNull.RemoveAtIndex(-1)
	listNull.Show()

	listNull.RemoveAtIndex(1)
	listNull.Show()

	listNull.RemoveAtIndex(1)
	listNull.Show()

	listNull.Append(14)
	listNull.Append(16)
	listNull.Append(18)
	listNull.Append(110)
	listNull.Show()

	listNull.RemoveAtIndex(5)
	listNull.Show()

}

func TestList_Contain(t *testing.T) {
	listNull.Append(4)
	listNull.Append(6)
	listNull.Append(8)
	listNull.Append(10)

	fmt.Println(listNull.Contain(4), listNull.Contain(10), listNull.Contain(11))
}

func TestList_Reverse(t *testing.T) {
	listNull.Append(4)
	listNull.Append(6)
	listNull.Append(8)
	listNull.Append(10)

	listNull.Reverse()
	listNull.Show()
}