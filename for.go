package main

import "fmt"

func main() {

	sl := []int64{9, 8, 7}
	for key, value := range sl{
		fmt.Printf("key: %v, val: %v \n", key, value)
	}
	// key: 0, vale: 9
	// key: 1, val: 8
	// key: 2, val: 7

	for _, value := range sl {
		fmt.Println(value)
	}
	//9
	//8
	//7

	ages := map[string]int {
		"Андрей": 30,
		"Анастасия": 25,
	}
	for key, value := range ages {
		fmt.Println(key)
		fmt.Println(value)
	}
	// Андрей
	// 30
	// Анастасия
	// 25

	word := "слёрм"

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		fmt.Printf("%T", word[i])
	}
	//209
	//uint8129
	//uint8208
	//uint8187
	//uint8209
	//uint8145
	//uint8209
	//uint8128
	//uint8208
	//uint8188
	//uint8
}