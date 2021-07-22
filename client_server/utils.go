package client_server

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	if err != nil {
		fmt.Println("BytesToInt error: ", err)
		return 0
	}
	return int(data)
}
