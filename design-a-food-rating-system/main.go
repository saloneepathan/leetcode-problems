package main

import (
	"container/heap"
	"fmt"
)

// Food struct holds the data
type Food struct {
	name    string
	cuisine string
	rating  int
}

// Item for priority queue
type Item struct {
	food   string
	rating int
}

// MaxHeap with custom comparator (highest rating first, then lexicographically smaller name)
type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food // lexicographically smaller first
	}
	return h[i].rating > h[j].rating
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// FoodRatings system
type FoodRatings struct {
	foodMap    map[string]*Food    // food name -> Food struct
	cuisineMap map[string]*MaxHeap // cuisine -> heap of foods
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodMap := make(map[string]*Food)
	cuisineMap := make(map[string]*MaxHeap)

	for i := range foods {
		f := &Food{
			name:    foods[i],
			cuisine: cuisines[i],
			rating:  ratings[i],
		}
		foodMap[foods[i]] = f

		if _, ok := cuisineMap[cuisines[i]]; !ok {
			cuisineMap[cuisines[i]] = &MaxHeap{}
		}
		heap.Push(cuisineMap[cuisines[i]], Item{food: foods[i], rating: ratings[i]})
	}

	return FoodRatings{foodMap: foodMap, cuisineMap: cuisineMap}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	f := fr.foodMap[food]
	f.rating = newRating
	// Push new entry (lazy deletion strategy)
	heap.Push(fr.cuisineMap[f.cuisine], Item{food: f.name, rating: newRating})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.cuisineMap[cuisine]
	for {
		top := (*h)[0]
		f := fr.foodMap[top.food]
		if f.rating == top.rating {
			return top.food
		}
		// otherwise discard outdated entry
		heap.Pop(h)
	}
}

// Example usage
func main() {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	fr := Constructor(foods, cuisines, ratings)
	fmt.Println(fr.HighestRated("korean"))   // kimchi
	fmt.Println(fr.HighestRated("japanese")) // ramen

	fr.ChangeRating("sushi", 16)
	fmt.Println(fr.HighestRated("japanese")) // sushi

	fr.ChangeRating("ramen", 16)
	fmt.Println(fr.HighestRated("japanese")) // ramen
}
