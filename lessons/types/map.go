package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map type")

	// A map is an unordered group of elements of one type, called the element
	// type, indexed by a set of unique keys of another type, called the key
	// type. The value of an uninitialized map is nil.
	var emptyMap map[string]float64
	fmt.Printf("Empty map %+v, is nil %t, len %d\n", emptyMap, emptyMap == nil, len(emptyMap))

	// The comparison operators == and != must be fully defined for operands
	// of the key type; thus the key type must not be a function, map, or
	// slice. If the key type is an interface type, these comparison operators
	// must be defined for the dynamic key values; failure will cause a
	// run-time panic.
	// TODO: add examples when key type is struct and interface

	// A map literal describes an initial set of keys and values. Thus the
	// length of the map literal is equal to the number of key-value pairs.
	// The map literal has the form
	literalMap := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87,
	}
	fmt.Printf("Literal map %+v, is nil %t, len %d\n", literalMap, literalMap == nil, len(literalMap))

	// A new, empty map value is made using the built-in function make, which
	// takes the map type and an optional capacity hint as arguments:
	// make(map[string]int)
	// make(map[string]int, 100)
	// The initial capacity does not bound its size: maps grow to accommodate
	// the number of items stored in them, with the exception of nil maps.
	// A nil map is equivalent to an empty map except that no elements may be
	// added.
	newMap := make(map[string]int)
	newMap["Maxim"] = 100
	newMap["Bob"] = 200
	newMap["Alice"] = 300
	newMap["Frank"] = 400
	fmt.Printf("New map %v, is nil %t, len %d\n", newMap, newMap == nil, len(newMap))

	// The delete built-in function deletes the element with the specified
	// key (m[key]) from the map. If m is nil or there is no such element,
	// delete is a no-op.
	delete(newMap, "Frank")
	fmt.Printf("New map %v, is nil %t, len %d\n", newMap, newMap == nil, len(newMap))

	// Quoting from the source code: https://github.com/golang/go/blob/master/src/runtime/map.go
	// A map is just a hash table. The data is arranged
	// into an array of buckets. Each bucket contains up to
	// 8 key/elem pairs. The low-order bits of the hash are
	// used to select a bucket. Each bucket contains a few
	// high-order bits of each hash to distinguish the entries
	// within a single bucket.
	//
	// If more than 8 keys hash to a bucket, we chain on
	// extra buckets.
	//
	// When the hashtable grows, we allocate a new array
	// of buckets twice as big. Buckets are incrementally
	// copied from the old bucket array to the new bucket array.

	// #1
	for i := 0; i < 2; i++ {
		newMap := make(map[int]string)
		for j := 0; j < 10; j++ {
			newMap[j] = "some string"
		}

		fmt.Printf("%+v\n", newMap)
		for k, v := range newMap {
			fmt.Println(k, v)
		}
		fmt.Println("END")
	}

	// #2
	type Point struct {
		x, y float32
	}

	field := map[Point]bool{
		Point{1, 1}: false,
		Point{1, 2}: true,
		Point{1, 3}: true,
		Point{2, 1}: false,
		Point{2, 2}: false,
		Point{2, 3}: false,
	}

	fmt.Printf("%+v\n", field)
	fmt.Printf("Point{1, 1} == Point{1, 3} (%t)\n", Point{1, 1} == Point{1, 1})
	fmt.Printf("%+v\n", field[Point{1, 2}])

	// TODO: add about slice and map
	for point, v := range field {
		fmt.Println(point, v)
	}

	// #3
	{
		a := map[int]int{1: 10, 2: 20}
		b := a
		b[2] = 30

		for i := 4; i < 20; i++ {
			b[i] = i * 10
		}
		fmt.Println(a, "\n\n\nB:", b)
	}

	// #4
	{
		floors := map[int][]int{
			1: {101, 102, 103, 104, 105}, // floor 1
			2: {201, 202, 203, 204, 205}, // floor 2
			3: {301, 302, 303, 304, 305}, // floor 3
		}
		fmt.Println(floors)

		rooms := map[int]int{
			101: 1,
			102: 1,
			103: 1,
			104: 1,
			105: 1,
			201: 2,
			202: 2,
			203: 2,
			204: 2,
			205: 2,
			301: 3,
			302: 3,
			303: 3,
			304: 3,
			305: 3,
		}
		fmt.Println(rooms)

		type KV struct {
			Key    int
			Value  int
			Exists bool
		}
		list := []KV{
			{Key: 101, Value: 1, Exists: true},
			{Key: 102, Value: 1, Exists: true},
			{Key: 103, Value: 1, Exists: true},
			{Key: 104, Value: 1, Exists: true},
			{Key: 105, Value: 1, Exists: true},
			{Key: 201, Value: 2, Exists: true},
			{Key: 202, Value: 2, Exists: true},
			//{Key: 203, Value: 2, Exists: true},
			{Key: 204, Value: 2, Exists: true},
			{Key: 205, Value: 2, Exists: true},
			{Key: 301, Value: 3, Exists: true},
			{Key: 302, Value: 3, Exists: true},
			//{Key: 303, Value: 3, Exists: true},
			{Key: 304, Value: 3, Exists: true},
			{Key: 305, Value: 3, Exists: true},
		}

		hashMap := make([]KV, len(list))

		for _, v := range list {
			hashMap[v.Key%len(list)] = v
		}

		fmt.Println("\nhashMap[203]: ", hashMap[203%len(hashMap)].Exists)
		fmt.Println("\nhashMap[303]: ", hashMap[303%len(hashMap)].Exists)
		fmt.Println("\nhashMap[603]: ", hashMap[603%len(hashMap)])

		// 1. Получить список всех комнат на этаже
		rooms_, ok := floors[4]
		fmt.Println(rooms_, ok)

		rooms_, ok = floors[2]
		fmt.Println(rooms_, ok)

		// 2. Получить по номеру комнаты этаж
		floor, ok := rooms[203]
		fmt.Println(floor, ok)
	}

}
