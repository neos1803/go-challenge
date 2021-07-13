//Without interface
// package main

// import "fmt"

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func (englishBot) getGreeting() string {
// 	return "Hi There!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//With interface

package main

import "fmt"

//Defining a new interface type
//In Which all function with different type of argument, gets treated similarly
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func (englishBot) getGreeting() string {
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
