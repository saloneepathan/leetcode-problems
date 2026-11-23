package main

import "fmt"

type ProductOfNumbers struct {
	prefix []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = []int{1}
		return
	}

	last := this.prefix[len(this.prefix)-1]
	this.prefix = append(this.prefix, last*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefix)

	if k >= n {
		return 0
	}

	return this.prefix[n-1] / this.prefix[n-1-k]
}

func main() {
	productOfNumbers := Constructor()

	productOfNumbers.Add(3)
	productOfNumbers.Add(0)
	productOfNumbers.Add(2)
	productOfNumbers.Add(5)
	productOfNumbers.Add(4)

	fmt.Println(productOfNumbers.GetProduct(2)) // 20
	fmt.Println(productOfNumbers.GetProduct(3)) // 40
	fmt.Println(productOfNumbers.GetProduct(4)) // 0

	productOfNumbers.Add(8)

	fmt.Println(productOfNumbers.GetProduct(2)) // 32
}
