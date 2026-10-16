
package main

import "fmt"

const MOD int64 = 1_000_000_007

type BIT struct {
	size   int
	_tree1 []int64
	_tree2 []int64
}

func NewBIT(n int) *BIT {
	return &BIT{
		size:   n,
		_tree1: make([]int64, n+2),
		_tree2: make([]int64, n+2),
	}
}

func (bit *BIT) _lowbit(index int) int {
	return index & -index
}

func (bit *BIT) add(left, right, delta int) {
	bit._add(left, int64(delta))
	bit._add(right+1, int64(-delta))
}

func (bit *BIT) query(left, right int) int64 {
	return bit._query(right) - bit._query(left-1)
}

func (bit *BIT) _add(index int, delta int64) {
	index++

	rawIndex := int64(index)

	for index <= bit.size {
		bit._tree1[index] += delta
		bit._tree2[index] += rawIndex * delta
		index += bit._lowbit(index)
	}
}

func (bit *BIT) _query(index int) int64 {
	index++

	rawIndex := int64(index + 1)

	if index > bit.size {
		index = bit.size
	}

	var res int64

	for index > 0 {
		res += rawIndex*bit._tree1[index] - bit._tree2[index]
		index -= bit._lowbit(index)
	}

	return res
}

func sumCounts(nums []int) int {
	n := len(nums)

	var res, cur int64

	bit := NewBIT(n + 10)
	seen := make(map[int]int)

	for i, a := range nums {
		j := -1

		if val, ok := seen[a]; ok {
			j = val
		}

		j++

		q := bit.query(j, i)

		cur = (cur + q*2 + int64(i-j+1)) % MOD
		res = (res + cur) % MOD

		bit.add(j, i, 1)
		seen[a] = i
	}

	return int(res)
}

func main() {
	nums1 := []int{1, 2, 1}
	fmt.Println(sumCounts(nums1)) // Output: 15

	nums2 := []int{2, 2}
	fmt.Println(sumCounts(nums2)) // Output: 3

	nums3 := []int{1, 2, 3}
	fmt.Println(sumCounts(nums3)) // Output: 20
}
