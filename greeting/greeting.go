package greeting

import "fmt"

type Salutation struct {
	Name, Greeting string
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

type Salutations []Salutation

type Printer func(string)

func (salutations Salutations) Greet(do func(string), isFormal bool, times int) {

	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Joe")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}

//func Print(s string) {
//	fmt.Print(s)
//}
//
//func PrintLine(s string) {
//	fmt.Println(s)
//}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s, custom)
	}
}
