package model

import (
	"fmt"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err)
	}
}
