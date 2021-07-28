package common

import (
	"encoding/json"
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

func LogToJSON(data interface{}) string {
	dataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("LogToJSON error: ", err)
		return ""
	}
	return string(dataByte)
}
