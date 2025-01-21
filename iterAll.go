package main

import (
	"fmt"
	"iter"
	"reflect"
	"slices"
)

// AllFn Slices All
func AllFn() {
	slice := []int{1, 2, 3}
	it := slices.All(slice)
	fmt.Println("type of All: ", reflect.TypeOf(it))
	for i, itr := range it {
		fmt.Printf("Index: %d, Value: %d\n", i, itr)
	}
}

// AppSeq Values and AppendSeq
func AppSeq() {
	existing := []int{1, 2}
	it := slices.Values([]int{3, 4})
	fmt.Println("type of Values", reflect.TypeOf(it))
	combined := slices.AppendSeq(existing, it)
	fmt.Println("type of AppendSeq", reflect.TypeOf(it))
	fmt.Println(combined)
}

// Coll Collect
func Coll() {
	it := slices.Values([]int{1, 2, 3})
	newSlice := slices.Collect(it)
	fmt.Println("type of Collect: ", reflect.TypeOf(it))
	fmt.Println(newSlice)
}

// Bckwd Backward
func Bckwd() {
	slice := []int{1, 2, 3}
	it := slices.Backward(slice)
	fmt.Println("type of Backward: ", reflect.TypeOf(it))
	for i, i2 := range it {
		fmt.Println(i, i2)
	}
}

// Srt Sorted
func Srt() {
	slice := []int{5, 6, 7, 1, 2, 3}
	fmt.Println("type of slice: ", reflect.TypeOf(slice))
	it := slices.Values(slice)
	sort := slices.Sorted(it)
	fmt.Println("type of Sorted: ", reflect.TypeOf(it))
	fmt.Println(sort)
}

// Chnk Chunk
func Chnk() {
	slice := []int{1, 2, 3, 4, 5}
	it := slices.Chunk(slice, 3)
	fmt.Println("type of Chunk: ", reflect.TypeOf(it))
	slicePrint("chunk ", it)
}

func slicePrint(keyWord string, it iter.Seq[[]int]) {
	for x := range it {
		fmt.Println(keyWord, x)
	}
}

// Val Values
func Val() {
	slice := []int{1, 2, 3}
	it := slices.Values(slice)
	fmt.Println("type of Values: ", reflect.TypeOf(it))

	for x := range it {
		fmt.Println("val ", x)
	}
}
