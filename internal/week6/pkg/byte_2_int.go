package pkg

import (
	// 标准包
	"bytes"
	"encoding/binary"
	"fmt"
)

func Bytes2IntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

func Bytes2Int16(b []byte) (int16, error) {
	if nil == b || len(b) < 2 {
		return -1, fmt.Errorf("字节转换int16参数不合法")
	}
	return int16(b[1]) | int16(b[0])<<8, nil
}
