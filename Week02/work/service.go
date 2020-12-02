package work

import (
	"fmt"
)

func service() ([]byte, error) {
	ret, err := dao()
	if err == nil {
		return ret, nil
	} else {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
}
