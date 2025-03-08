package main

import (
	"fmt"
	"learnGo/dictionary"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpperV1(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpperV2(name string) (length int, upper string) { // Naked Return
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func canIDrinkV1(age int) bool {
	defer fmt.Println("canIDrink: I'm done")  // executed after function returns
	if koreanAge := age + 2; koreanAge < 18 { // variable expression
		return false
	}
	return true
}

func canIDrinkV2(age int) bool { // Switch & variable expression
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func superAdd(numbers ...int) (addIndex, ans int) { // 가변 인자 처리
	for number := range numbers { // 이경우 index를 더함
		addIndex += number
	}

	for _, number := range numbers { // index 말고 실제 값처리 해줘야함
		ans += number
	}

	return
}

type person struct {
	name    string
	age     int
	favFood []string
}

// main for Basic of Go
//func main() {
//ages := [5]int{10, 20, 30}
//ages[3] = 40
//ages[4] = 50
//fmt.Println(ages)
//
//names := []string{"tony", "thomas", "taewan"} // slice
//names = append(names, "taiwan")               // append()좀 특이함
//fmt.Println(names)
//
//userAge := map[string]int{"tony": 10, "thomas": 20, "taewan": 25} // map
//fmt.Println(userAge)
//
//// Initializing struct
//favFood := []string{"ramen", "unagi"}
//tony := person{"tony", 24, favFood}
//thomas := person{name: "tony", age: 24, favFood: favFood}
//fmt.Println(tony)
//fmt.Println(thomas)
//}

// main for simple bank application
//func main() {
//	tonyAccount := bank.NewAccount("Tony")
//	tonyAccount.Deposit(1000)
//	fmt.Println(tonyAccount.Balance())
//	fmt.Println(tonyAccount.Withdraw(2000))
//
//	tonyAccount.Withdraw(500)
//	tonyAccount.ChangeOwner("taewan")
//
//	fmt.Println(tonyAccount)
//}

func main() {
	dictionary := dictionary.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
