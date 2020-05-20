package iso8583

import (
	"reflect"
	"testing"
)

func TestAddField(t *testing.T) {
	var scenarios = []struct {
		result string
		bits   []uint8
	}{
		{
			result: "42000104021C0078",
			bits:   []uint8{2, 7, 24, 30, 39, 44, 45, 46, 58, 59, 60, 61},
		},
		{
			result: "4200010002140068",
			bits:   []uint8{2, 7, 24, 39, 44, 46, 58, 59, 61},
		},
		{
			result: "0000010002140068",
			bits:   []uint8{24, 39, 44, 46, 58, 59, 61},
		},
		{
			result: "A888010002143068",
			bits:   []uint8{1, 3, 5, 9, 13, 24, 39, 44, 46, 51, 52, 58, 59, 61},
		},
	}

	for _, scenario := range scenarios {
		var fields uint64
		for _, bit := range scenario.bits {
			fields = addField(fields, bit)
		}
		if result := bitmapHex(fields); scenario.result != result {
			t.Errorf("bitmap should be %s, instead of %s", scenario.result, result)
		}
	}
}

func TestIsBitSet(t *testing.T) {
	var scenarios = []struct {
		bits []uint8
	}{
		{
			bits: []uint8{2, 7, 24, 30, 39, 44, 45, 46, 58, 59, 60, 61},
		},
		{
			bits: []uint8{2, 7, 24, 39, 44, 46, 58, 59, 61},
		},
		{
			bits: []uint8{24, 39, 44, 46, 58, 59, 61},
		},
		{
			bits: []uint8{1, 3, 5, 9, 13, 24, 39, 44, 46, 51, 52, 58, 59, 61},
		},
		{
			bits: []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 61, 62, 63, 64},
		},
	}

	for _, scenario := range scenarios {
		var fields uint64
		// create bitset
		for _, bit := range scenario.bits {
			fields = addField(fields, bit)
		}

		bits := make([]uint8, 0)
		for i := uint8(1); i <= 64; i++ {
			if isBitSet(fields, i) {
				bits = append(bits, i)
			}
		}

		if !reflect.DeepEqual(scenario.bits, bits) {
			t.Errorf("bitmap should be %#v, instead of %#v", scenario.bits, bits)
		}
	}
}

func TestDecodeHexString(t *testing.T) {
	var scenarios = []struct {
		bitmapString string
		bits         []uint8
	}{
		{
			bitmapString: "42000104021C0078",
			bits:         []uint8{2, 7, 24, 30, 39, 44, 45, 46, 58, 59, 60, 61},
		},
		{
			bitmapString: "4200010002140068",
			bits:         []uint8{2, 7, 24, 39, 44, 46, 58, 59, 61},
		},
		{
			bitmapString: "0000010002140068",
			bits:         []uint8{24, 39, 44, 46, 58, 59, 61},
		},
		{
			bitmapString: "A888010002143068",
			bits:         []uint8{1, 3, 5, 9, 13, 24, 39, 44, 46, 51, 52, 58, 59, 61},
		},
		{
			bitmapString: "FFFFFF000000000F",
			bits:         []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 61, 62, 63, 64},
		},
	}

	for _, scenario := range scenarios {
		var expected uint64
		for _, bit := range scenario.bits {
			expected = addField(expected, bit)
		}

		if result, err := decodeHexString(scenario.bitmapString); err != nil {
			t.Error(err)
		} else if expected != result {
			t.Errorf("bitmap should be %d, instead of %d", expected, result)
		}
	}
}
