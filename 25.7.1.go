package main

import (
	"fmt"
	"log"
)

func main() {
	//Не придумал как вставить произвольные данные.
	//Пробовал через пустой интерфейс, но с ним scan не работает.
	//Буду благодарен за подсказку
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели данные: %v\n", n)
}
