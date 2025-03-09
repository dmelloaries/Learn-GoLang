package main

import "fmt"

func main() {
	//type declared
	var age int = 30
	//shorthand method
	age2:=20

	age2=21

	//infered the type
	var age3=23

	//assigned later val
	var name string

	name="aries"

	fmt.Println("Age is ",age)
	fmt.Println("Age2 is ",age2)
	fmt.Println("Age3 is ",age3)
	fmt.Println("name is ",name)
	fmt.Println("Hello world")
	fmt.Println(1+1)
	fmt.Println(true)
	fmt.Println(10.5)

}