package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(message string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
}

type Italian struct {
}

func (s Italian) LanguageName() string {
	return "Italian"
}

func (s Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct {
}

func (s Portuguese) LanguageName() string {
	return "Portuguese"
}

func (s Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}
