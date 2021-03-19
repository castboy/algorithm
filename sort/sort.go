package main

import "fmt"

//时间复杂度 O(n^2),空间复杂度 O(1)
//从头到尾,走一步,比较一下,如果前面的比后面的大,互换位置,然后继续比较
//这样 最后一个是最大的,然后继续循环

func Bubble(nums []int) []int {
	numLen := len(nums)
	for i := 0; i < numLen; i++ {
		for j := 0; j < numLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

//时间复杂度 O(n^2),空间复杂度 O(1)
//第一次,先找出最小(大)的值,放到第一的位置
//第二次,找出最小(大)的值,放到第二的位置
func Selection(nums []int) []int {
	numLen := len(nums)
	var min, k int

	for i := 0; i < numLen; i++ {
		for j := i; j < numLen; j++ {
			if j == i || nums[j] < min {
				min = nums[j]
				k = j
			}
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	return nums
}

//时间复杂度 O(n^2),空间复杂度 O(1)
//理解很简单,写起来有点绕
//逐步往后选取一个数,然后把该数与之前的排好的进行比较,如果比某个大,就插入到该数的后面
func Insertion(nums []int) []int {
	numLen := len(nums)
	for i := 1; i < numLen; i++ {
		tmp := nums[i]
		for j := i -1; j >= 0; j-- {
			if nums[j] > tmp {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = tmp
				break
			}
			// 如果都移动了,然后遍历到0了,说明当前值是最小的,把当前值放到最小的位置
			if j == 0 {
				nums[j] = tmp
			}
		}
	}

	return nums
}
// ref: http://interview.wzcu.com/算法/经典排序/insertion.html
func Insertion2(nums []int) []int {
	for i := range nums {
		preIndex := i - 1
		current := nums[i]
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex -= 1
		}

		nums[preIndex+1] = current
	}

	return nums
}

//时间复杂度 O(nlogn),空间复杂度 O(1)
//希尔排序是 插入排序的升级版
//将插入排序拆分多个区间,然后区间里进行排序
// ref: https://m.runoob.com/data-structures/shell-sort.html  && https://www.jianshu.com/p/d730ae586cf3
func Shell(nums []int) []int {
	numsLen := len(nums)
	for gap := numsLen/2; gap > 0; gap /= 2 {
		for endIndex := gap; endIndex < numsLen; endIndex++ { // endIndex 作为每一组比较数据的最后一个元素角标
			// just insertion bellow 9 lines, similar with Insertion2.
			for i := endIndex; i >= 0; i -= gap { //j:代表与i同一组的数组元素角标
				preIndex := i - gap
				current := nums[i]
				for preIndex >=0 && nums[preIndex] > current {
					nums[preIndex+gap] = nums[preIndex]
					preIndex -= gap
				}
				nums[preIndex+gap] = current
			}
		}
	}

	return nums
}

func Merge(nums []int) []int {
	numsLen := len(nums)

	if numsLen < 2 {
		return nums
	}

	mid := numsLen / 2

	return mergeTwo(Merge(nums[0:mid]), Merge(nums[mid:]))
}

func mergeTwo(nums1, nums2 []int) []int {
	res1Len := len(nums1)
	res2Len := len(nums2)

	index1 := 0
	index2 := 0

	res := make([]int, res1Len+res2Len)
	index := 0

	for index1 < res1Len && index2 < res2Len {
		if nums1[index1] < nums2[index2] {
			res[index] = nums1[index1]
			index1++
		} else {
			res[index] = nums2[index2]
			index2++
		}

		index++
	}

	for index1 < res1Len {
		res[index] = nums1[index1]
		index1++
		index++
	}

	for index2 < res2Len {
		res[index] = nums2[index2]
		index2++
		index++
	}

	return res
}

func main() {
	nums := []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Bubble(nums))

	nums = []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Selection(nums))

	nums = []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Insertion(nums))

	nums = []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Insertion2(nums))

	nums = []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Shell(nums))

	nums = []int{1, 2, 6, 4, 3, 0, -1}
	fmt.Println(Merge(nums))

	nums1 := []int{1, 4, 5}
	nums2 := []int{2, 3, 6}

	fmt.Println(mergeTwo(nums1, nums2))
}
