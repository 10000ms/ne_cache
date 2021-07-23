package common

import (
	"fmt"
	"strconv"
)

func BytesStringToInt(bys []byte) int {
	s := string(bys)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("BytesStringToInt error: ", err)
		return 0
	}
	return i
}
