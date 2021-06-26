package main

import (
	"fmt"

	"github.com/jyebe9034/golearning/mydict"
)

// func main() { // Account
// 	account := accounts.NewAccount("hannah")
// 	account.Deposit(10)
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// fmt.Println(account.Balnace(), account.Owner())
// 	fmt.Println(account)
// }

func main() { // Dictionary
	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(definition)

	// word := "hello"
	// definition := "greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// def, _ := dictionary.Search("hello")
	// fmt.Println("found: ", word, ", definition: ", def)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// err := dictionary.Update("noExist", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
