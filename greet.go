package test

import "fmt"


func Greet(name string) string {
	return fmt.Sprintf("Hi there, %s !!\n%s", name, intro())
}

func intro() string {
	return "welcome to demo !!"
}