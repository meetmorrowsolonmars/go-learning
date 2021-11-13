package main

import "fmt"

func main() {
	fmt.Println("Slice types")

	// Empty slice declaration. The value of an uninitialized slice is nil.
	var emptySlice []int
	fmt.Printf("Empty slice %v, is nil %t, len %d, cap %d\n", emptySlice, emptySlice == nil,
		len(emptySlice), cap(emptySlice))

	// A slice literal describes the entire underlying array literal.
	// Thus the length and capacity of a slice literal are the maximum
	// element index plus one. A slice literal has the form
	var literalSlice = []int{1, 2, 3}
	fmt.Printf("Literal slice %v, is nil %t, len %d, cap %d\n", literalSlice, literalSlice == nil,
		len(literalSlice), cap(literalSlice))

	// A new, initialized slice value for a given element type T is made
	// using the built-in function make([]T, length, capacity), which takes
	// a slice type and parameters specifying the length and optionally the
	// capacity.
	// A slice created with make always allocates a new, hidden array to which
	// the returned slice value refers. That is, executing
	var makeSlice = make([]int, 3, 3)
	fmt.Printf("Make slice %v, is nil %t, len %d, cap %d\n", makeSlice, makeSlice == nil,
		len(makeSlice), cap(makeSlice))

	// Produces the same slice as allocating an array and slicing it, so these
	// two expressions are equivalent:
	var newSlice = new([3]int)[0:3]
	fmt.Printf("New slice %v, is nil %t, len %d, cap %d\n", newSlice, newSlice == nil,
		len(newSlice), cap(newSlice))

	//type Slice struct {
	//	Array [Cap]interface{} // C++ void* Array
	//	Len int
	//	Cap int
	//}

	// A little brain fuck.
	fmt.Println("\nA little brain fuck.")

	parentSlice := make([]int, 0, 3)
	parentSlice = append(parentSlice, 1, 2, 3)

	childSliceA := parentSlice

	childSliceB := parentSlice
	childSliceB = append(childSliceB, 4)

	childSliceC := parentSlice
	childSliceC = append(childSliceC, 5)

	fmt.Printf("Parent slice %v, len %d, cap %d\n", parentSlice, len(parentSlice), cap(parentSlice))
	fmt.Printf("Child slice A %v\n", childSliceA)
	fmt.Printf("Child slice B %v\n", childSliceB)
	fmt.Printf("Child slice C %v\n", childSliceC)

	//type Slice struct {
	//	Array [Cap]interface{} // C++ void* Array
	//	Len int
	//	Cap int
	//}

	// And more
	fmt.Println("\nAnd more")
	parentSlice = make([]int, 0, 6)
	parentSlice = append(parentSlice, 1, 2, 3)

	childSliceA = parentSlice

	childSliceB = parentSlice
	childSliceB = append(childSliceB, 4)

	childSliceC = parentSlice
	childSliceC = append(childSliceC, 5)

	fmt.Printf("Parent slice %v, len %d, cap %d\n", parentSlice, len(parentSlice), cap(parentSlice))
	fmt.Printf("Child slice A %v\n", childSliceA)
	fmt.Printf("Child slice B %v\n", childSliceB)
	fmt.Printf("Child slice C %v\n", childSliceC)

	//type Slice struct {
	//	Array [Cap]interface{} // C++ void* Array
	//	Len int
	//	Cap int
	//}

	// And more
	fmt.Println("\nAnd more")
	parentSlice = make([]int, 0, 6)
	parentSlice = append(parentSlice, 1, 2, 3, 4, 5, 6)

	childSliceA = parentSlice[2:4] // change 4 to 6
	childSliceA[0] = -1
	childSliceA[1] = -2
	childSliceA = append(childSliceA, -3, -4)

	fmt.Printf("Parent slice %v, len %d, cap %d\n", parentSlice, len(parentSlice), cap(parentSlice))
	fmt.Printf("Parent slice[0:2] %v\n", parentSlice[0:2])
	fmt.Printf("Parent slice[2:4] %v\n", parentSlice[2:4])
	fmt.Printf("Parent slice[4:6] %v\n", parentSlice[4:6])
	fmt.Printf("Child slice A %v\n", childSliceA)

	//type Slice struct {
	//	Array [Cap]interface{} // C++ void* Array
	//	Len int
	//	Cap int
	//}

	// And more
	fmt.Println("\nAnd more")
	type Bag struct {
		Things []int
	}

	bags := []Bag{
		{Things: []int{1, 2}},   // Bag 1
		{Things: []int{}},       // Bag 2
		{Things: []int{10, 20}}, // Bag 3
	}

	fmt.Printf("Bags: %+v\n", bags)
	for i, bag := range bags {
		fmt.Printf("Bag %d before append: %+v\n", i+1, bag)
		bag.Things = append(bag.Things, 3, 4, 5)
		fmt.Printf("Bag %d after append: %+v\n\n", i+1, bag)
	}

	fmt.Printf("Bags after first loop: %+v\n", bags)
	for i := range bags {
		fmt.Printf("Bag %d before append: %+v\n", i+1, bags[i])
		bags[i].Things = append(bags[i].Things, 3, 4, 5)
		fmt.Printf("Bag %d after append: %+v\n\n", i+1, bags[i])
	}

	fmt.Printf("Bags: %+v\n", bags)
	// HOMEWORK: Create the same list of bags, but keep pointers to the Bag
	// object in it and do the same append operations yourself.
}
