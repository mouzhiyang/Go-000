package work

import "fmt"

func ApiInfo() {
	ret, err := service()
	if err == nil {
		fmt.Println(ret)
	} else {
		fmt.Println(err)
	}
}
