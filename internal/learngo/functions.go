package learngo

import "fmt"

/*
For getting type.

		fmt.Printf("%T\n" ,age)
		fmt.Println(reflect.TypeOf(age))
		  rune := 'A'
	   fmt.Printf("%d \n", rune)
	   fmt.Println(reflect.TypeOf(rune))
*/
func closureExample() {
	number := 10
	squareNum := func() int {
		temp := number
		number *= number
		fmt.Printf("Calculated in clousure : %v -> %v \n", temp, number)
		return number
	}
	squareNum()
	squareNum()
}

func CallBackExample(n int, callback func(n int) bool) {
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	callback(5)
}
func ReturnFunc() func(string) bool {
	return func(input string) bool {
		fmt.Printf("From Return function %v \n", input)
		return false
	}
}

func TestCallBack() {
	CallBackExample(3, func(n int) bool {
		fmt.Println("From callback - start")
		for i := 0; i < n; i++ {
			fmt.Print(i)
		}
		fmt.Println("From callback - end")
		return false
	})
}
