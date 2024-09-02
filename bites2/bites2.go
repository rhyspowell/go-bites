package main

var myInteger = 29
var myFloat float32 = 400.68
var myString = "I love data types!"

func bites2() (int, float32, string) {
	return myInteger, myFloat, myString
}

func main() {
	println(bites2())
}
