package getname

import "fmt"

func GetName() string {
	var name string

	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)

	return name
}
