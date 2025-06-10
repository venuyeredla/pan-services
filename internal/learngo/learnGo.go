package learngo

import (
	"fmt"
	"sort"
)

// String to primitive data types by "strconv"  package
// strings for string manipuations
// sort package for sorting primiitve and abstract

type Person struct {
	FName string
	LName string
	Age   uint
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func SortPrint() {
	var persons []Person = []Person{
		{"venu", "gopal", 36},
		{"Ananya", "Reddy", 32},
		{"Yashika", "Tanvi", 5},
	}
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].FName < persons[j].FName
	})
	fmt.Println(persons)

	sort.Sort(ByAge(persons))
	fmt.Println(persons)
}
