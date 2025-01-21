package main

import (
	"fmt"
	"maps"
	"reflect"
)

func AllMaps() {
	m := map[string]int{"one": 1, "two": 2}
	it := maps.All(m)

	fmt.Println("type of map All:", reflect.TypeOf(it))
	for k, v := range it {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
}

func KeysMap() {
	m := map[string]int{"one": 1, "two": 2}
	it := maps.Keys(m)
	fmt.Println("type of map Keys:", reflect.TypeOf(it))
	for k := range it {
		fmt.Printf("Key: %s\n", k)
	}

}

func ValuesMap() {
	m := map[string]int{"one": 1, "two": 2}
	it := maps.Values(m)
	fmt.Println("type of map Values:", reflect.TypeOf(it))
	for v := range it {
		fmt.Printf("Value: %d\n", v)
	}
}

func InsertMap() {
	m := map[string]int{"one": 1}
	it := maps.All(map[string]int{"two": 2, "three": 3})
	maps.Insert(m, it)
	fmt.Println("type of map with insert:", reflect.TypeOf(m))
	fmt.Println(m)
}

func CollectMap() {
	it := maps.All(map[string]int{"one": 1, "two": 2})
	newMap := maps.Collect(it)
	fmt.Println("type of collected maps:", reflect.TypeOf(newMap))
	fmt.Println(newMap)
}
