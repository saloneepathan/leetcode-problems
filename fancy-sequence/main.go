package main

import (
	"fmt"
)

const MOD = 1000000007

type Fancy struct {
	v []int
	a []int
	b []int
}

func Constructor() Fancy {
	return Fancy{
		v: []int{},
		a: []int{1},
		b: []int{0},
	}
}

func (this *Fancy) Append(val int) {
	this.v = append(this.v, val)
	this.a = append(this.a, this.a[len(this.a)-1])
	this.b = append(this.b, this.b[len(this.b)-1])
}

func (this *Fancy) AddAll(inc int) {
	last := len(this.b) - 1
	this.b[last] = (this.b[last] + inc) % MOD
}

func (this *Fancy) MultAll(m int) {
	last := len(this.a) - 1
	this.a[last] = (this.a[last] * m) % MOD
	this.b[last] = (this.b[last] * m) % MOD
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.v) {
		return -1
	}

	// ratio of multipliers
	ao := (inv(this.a[idx]) * this.a[len(this.a)-1]) % MOD

	// compute additive difference
	bo := (this.b[len(this.b)-1] - this.b[idx]*ao%MOD + MOD) % MOD

	return (ao*this.v[idx]%MOD + bo) % MOD
}

func quickMul(x, y int) int {
	res := 1
	cur := x

	for y > 0 {
		if y&1 == 1 {
			res = (res * cur) % MOD
		}
		cur = (cur * cur) % MOD
		y >>= 1
	}

	return res
}

func inv(x int) int {
	return quickMul(x, MOD-2)
}

func main() {

	fancy := Constructor()

	fancy.Append(2)
	fancy.AddAll(3)
	fancy.Append(7)
	fancy.MultAll(2)

	fmt.Println(fancy.GetIndex(0)) // 10

	fancy.AddAll(3)
	fancy.Append(10)
	fancy.MultAll(2)

	fmt.Println(fancy.GetIndex(0)) // 26
	fmt.Println(fancy.GetIndex(1)) // 34
	fmt.Println(fancy.GetIndex(2)) // 20
}
