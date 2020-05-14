package iso8583

import (
	"encoding/hex"
	"fmt"
)

//TODO unpacked format
func encodeBitmap(fields map[int]struct{}, packed, secondaryBitmap bool) ([]byte, error) {
	byteNum := 8
	if secondaryBitmap {
		byteNum = 16
	}
	b := make([]byte, byteNum)
	if secondaryBitmap {
		b[0] = b[0] | 1<<7 //if secondary encodeBitmap is present, first bit in byte[0] is 1
	}

	for i, _ := range fields {
		if _, ok := fields[i]; ok {
			byteInd := (i - 1) / 8
			b[byteInd] = b[byteInd] | 1<<((8-i%8)%8)
		}
	}
	if packed {
		return hex.DecodeString(fmt.Sprintf("%X", b))
	}
	return b, nil
}
