package main

import (
	"fmt"
	"iter"
	"slices"
)

func AllFn() {
	slice := []int{1, 2, 3}
	it := slices.All(slice)
	for i, itr := range it {
		fmt.Printf("Index: %d, Value: %d\n", i, itr)
	}
}

func AppSeq() {
	existing := []int{1, 2}
	it := slices.Values([]int{3, 4})
	combined := slices.AppendSeq(existing, it)
	fmt.Println(combined)
}

func Coll() {
	it := slices.Values([]int{1, 2, 3})
	newSlice := slices.Collect(it)
	fmt.Println(newSlice)
}

func Bckwd() {
	slice := []int{1, 2, 3}
	it := slices.Backward(slice)
	for i, i2 := range it {
		fmt.Println(i, i2)
	}
}

func Srt() {
	slice := []int{5, 6, 7, 1, 2, 3}
	it := slices.Values(slice)
	sort := slices.Sorted(it)
	fmt.Println(sort)
}

func Chnk() {
	slice := []int{1, 2, 3, 4, 5}
	it := slices.Chunk(slice, 3)
	slicePrint("chunk ", it)
}

func Val() {
	slice := []int{1, 2, 3}
	it := slices.Values(slice)

	for x := range it {
		fmt.Println("val ", x)
	}
}

func slicePrint(keyWord string, it iter.Seq[[]int]) {
	for x := range it {
		fmt.Println(keyWord, x)
	}
}
