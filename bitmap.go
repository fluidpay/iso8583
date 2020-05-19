package iso8583

import (
	"fmt"
	"strconv"
)

func addField(fields uint64, num uint8) uint64 {
	// Explenation:
	// (64-num) means that we want to set the bit from the start to the end
	// 1 << is a binary push with the 1 on the first place to the nth place
	// fields |= bitwise OR, so it will set the nth bit to 1 in the fields
	fields |= 1 << (64 - num)
	return fields
}

func bitmapHex(fields uint64) string {
	// Explenation:
	// %016X will convert the fields (uint64) to hexadecimal number
	// with left padding to length of 16
	return fmt.Sprintf("%016X", fields)
}

func decodeHexString(value string) (uint64, error) {
	return strconv.ParseUint(value, 16, 64)
}

func isBitSet(fields uint64, num int8) bool {
	return fields & (1 << (64 - num)) != 0
}
