package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const name = "Ho Sy Binh"

type Person struct {
	Name string
	Age  int
}

type PersonCollection []Person



func (pc PersonCollection) Len() int {
	return len(pc)
}

func (pc PersonCollection) Swap(i, j int) {
	pc[i], pc[j] = pc[j], pc[i]
}

func (pc PersonCollection) Less(i, j int) bool {
	return pc[i].Age < pc[j].Age
}

func main() {

	fmt.Println("Hello world!")
	fmt.Println(name)
	num := 12
	fmt.Println(num)

	var one, two, three, four, five int = 1, 2, 3, 4, 5
	fmt.Println(one, two, three, four, five)

	one = 10
	fmt.Println(one)

	const (
		a = 10
		b = 20
		c = 30
		d = 40
	)
	fmt.Println(a, b, c, d)

	isGolang := false
	isJavascript := true
	fmt.Println(isGolang, isJavascript)

	// Operators
	fmt.Println(a + b)
	fmt.Println(a / b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a == b)
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)
	num += 3
	fmt.Println(num)
	num -= 3
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 3
	fmt.Println(num)
	num %= 3
	fmt.Println(num)

	// string
	fmt.Println("Stringggggggggggg")
	fruit := "Orange"
	info := `Hello everyone 
					 my name is Binh`
	fmt.Println(fruit)
	fmt.Println(info)
	fmt.Printf("%c", fruit[0])
	fmt.Println()

	one1 := "one"
	two2 := "one"
	result := strings.Compare(one1, two2)
	fmt.Println(result)

	golang := "golang is a language"
	fmt.Println(strings.Contains(golang, "golang"))
	fmt.Println(strings.ToLower(golang))
	fmt.Println(strings.ToUpper(golang))
	fmt.Println(strings.Count(golang, "g"))
	fmt.Println(strings.Join([]string{"binhbinh", "hello"}, "bihbinh"))

	myBoo := "my boo"
	ageBoo := 21
	message := fmt.Sprintf("%s is %d years old", myBoo, ageBoo)
	fmt.Println(message)

	values := []interface{}{
		true,
		int8(8),
	}
	fmt.Println(values)
	for _, value := range values {
		fmt.Println(reflect.TypeOf(value))
	}

	array := []byte{1, 2}
	if array != nil {
		fmt.Println("array not empty")
	}

	if len(array) == 2 {
		fmt.Println("length is 2")
	} else if len(array) == 1 {
		fmt.Println("length is 1")
	} else {
		fmt.Println("length is other")
	}

	switch len(array) {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("case default")
	}

	arrayInt := []int{1, 2, 3, 4, 5}
	fmt.Println(arrayInt)

	clone := make([]int, len(arrayInt))
	copy(clone, arrayInt)
	fmt.Println(clone)

	sub := arrayInt[2:4]
	fmt.Println(sub)

	concat := append(arrayInt, []int{6, 7, 8, 9}...)
	fmt.Println(concat)

	arrayString := []string{"a", "b", "c"}

	var filtered []string
	for i, value := range arrayString {
		fmt.Println(i, value)
		filtered = append(filtered, strings.ToUpper(value)+" Binh")
	}

	fmt.Println(filtered)

	///
	arrayStringg := []string{"a", "b", "c"}
	for i, value := range arrayStringg {
		fmt.Println(i, value)
	}
	fmt.Println(arrayString, "arrayString")

	mapped := make([]string, len(arrayStringg))
	for i, value := range arrayString {
		mapped[i] = strings.ToUpper(value)
	}
	fmt.Println(mapped, "mapped")

	intList := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
	sort.Ints(intList)
	fmt.Println(intList)

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)

	stringList := []string{"a", "d", "z", "b", "c", "y"}
	sort.Strings(stringList)
	fmt.Println(stringList)

	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(stringList)

	collection := []Person{
		{"L", 1},
		{"C", 10},
		{"W", 19},
		{"Y", 99},
	}

	sort.Sort(PersonCollection(collection))
	fmt.Println(collection)

	sort.Sort(sort.Reverse(PersonCollection(collection)))
	fmt.Println(collection)
}
