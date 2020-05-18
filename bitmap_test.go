package iso8583

import (
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
