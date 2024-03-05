package get_name

import "fmt"

func GetName() string {
	var name string

	fmt.Scanln(&name)

	return name
}
