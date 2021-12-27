package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strconv"
)

func add(x,y int)(z1 int,z2 int){
	z1 = x + y
	z2 = x-y
	return
}

// func test(x int){
// 	fmt.Println("Hello", x)
// }

// type Point struct{
// 	x int
// 	y int
// }

type Student struct{
	name string
	grade []int
	age int
}

func (s *Student) setAge(age int){
	s.age = age
}

func (s Student) getAverageStudent() float32{
	sum:=0

	for i:=0; i<len(s.grade); i++{
		sum +=s.grade[i]
	}

	return float32(sum)/float32(len(s.grade))


}



func main(){
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born: ")
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("you will be %d years old at the end of 2020: ", 2020-input)

	// num1 := 10
	// num2 := 15
	// answer := num1 + num2

	// fmt.Printf("The answeer is %d", answer)

	// x :="tim"
	// y:="Tim"
	// val := x==y

	// fmt.Printf("%t", val)

	// age := 17

	// fmt.Println("Before the if Statement")

	// if  age >= 18 {
	// 	fmt.Println("You can ride")
	// }else{
	// 	fmt.Println("You cannot ride")
	// }

	// fmt.Println("After the if statement")

	// x:=0

	// for x<=5{
	// 	fmt.Println(x)
	// 	x++
	// }

	// for x=0; x<=1000; x++{
	// 	if x!=0 && x%3 ==0 && x%7 == 0 && x%9 == 0{
	// 		fmt.Println(x)
	// 		continue
	// 	}

	// 	fmt.Println("N")
	// }

	// x := 10

	// switch x {
	// case 7:
	// 	fmt.Println("I am in 7")
		
	// default:
	// 	fmt.Println("Not a case")
	// }

	// var x []int = []int{1,2,3,4,5}
	// var y []int = []int{5,6,7}
	// b := append(x, 20,30)
	// c := append(y,11,18)

	// fmt.Println(b,c)

	// var a []int = []int{54,232,32,42,45,23,5,2,32,34,43}

	// for i, element := range a{



	// 	fmt.Printf("%d: %d\n", i,element)
	// }

	// var mp map[string]int = map[string]int{
	// 	"apple": 5, 
	// 	"pear": 6,
	// 	"orange": 9,
	// }

	// val, ok := mp["tim"]
	// fmt.Println(val, ok)


	// ans1, ans2 := add(7,18)
	// fmt.Println(ans1, ans2)

	// x:=test
	// x(5)

	// p1 := &Point{1,2}
	// var p2 Point = Point{x:3}

	// fmt.Println(p1, p2)
	

	s1 := Student{"Tim", []int{89,43,43,23,53}, 19}
	fmt.Println(s1)

	s1.setAge(22)
	fmt.Println(s1)
	fmt.Println(s1.getAverageStudent())
	

	




}