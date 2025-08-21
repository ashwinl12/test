package test

import "fmt"


func Greet(name string) string {
	return fmt.Sprintf("Hi there, %s !!\n%s", name, Intro())
}

func Intro() string {
	return "Welcome to demo !!"
}
