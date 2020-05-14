package iso8583

import (
	"bytes"
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

func TestBitmapExUnpacked(t *testing.T) {
	fields := make(map[int]struct{})
	fields[2] = struct{}{}
	fields[5] = struct{}{}
	fields[6] = struct{}{}
	b, err := encodeBitmap(fields, false, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("% X", b)
	expected := make([]byte, 8)
	expected[0] = 0x4c
	if !bytes.Equal(b, expected) {
		t.Error("not equal")
	}
}

func TestBitmapExPacked(t *testing.T) {
	fields := make(map[int]struct{})
	fields[2] = struct{}{}
	fields[5] = struct{}{}
	fields[6] = struct{}{}
	fields[2+8] = struct{}{}
	fields[5+8] = struct{}{}
	fields[6+8] = struct{}{}
	b, err := encodeBitmap(fields, false, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
	expected := make([]byte, 8)
	expected[0] = 0x4c
	expected[1] = 0x4c

	if !bytes.Equal(b, expected) {
		t.Error("not equal")
	}
}

func TestBitmapBalanceInquiryAnATMSample(t *testing.T) {
	fields := make(map[int]struct{})
	fields[2] = struct{}{}
	fields[3] = struct{}{}
	fields[7] = struct{}{}
	fields[11] = struct{}{}
	fields[12] = struct{}{}
	fields[18] = struct{}{}
	fields[22] = struct{}{}
	fields[24] = struct{}{}
	fields[32] = struct{}{}
	fields[35] = struct{}{}
	fields[41] = struct{}{}
	fields[42] = struct{}{}
	fields[43] = struct{}{}
	fields[49] = struct{}{}
	fields[52] = struct{}{}

	b, err := encodeBitmap(fields, true, false)
	if err != nil {
		t.Fatal(err)
	}
	expected := make([]byte, 8)
	expected[0] = 0x62
	expected[1] = 0x30
	expected[2] = 0x45
	expected[3] = 0x01
	expected[4] = 0x20
	expected[5] = 0xe0
	expected[6] = 0x90
	expected[7] = 0x00

	if !bytes.Equal(b, expected) {
		t.Error("not equal")
	}
}

func TestBitmapBalanceInquiryResponseSample(t *testing.T) {
	fields := make(map[int]struct{})
	fields[2] = struct{}{}
	fields[3] = struct{}{}
	fields[7] = struct{}{}
	fields[11] = struct{}{}
	fields[12] = struct{}{}
	fields[32] = struct{}{}
	fields[39] = struct{}{}
	fields[54] = struct{}{}
	fields[102] = struct{}{}

	b, err := encodeBitmap(fields, true, true)
	if err != nil {
		t.Fatal(err)
	}
	expected := make([]byte, 16)
	expected[0] = 0xe2
	expected[1] = 0x30
	expected[2] = 0x00
	expected[3] = 0x01
	expected[4] = 0x02
	expected[5] = 0x00
	expected[6] = 0x04
	expected[7] = 0x00

	expected[0+8] = 0x00
	expected[1+8] = 0x00
	expected[2+8] = 0x00
	expected[3+8] = 0x00
	expected[4+8] = 0x04
	expected[5+8] = 0x00
	expected[6+8] = 0x00
	expected[7+8] = 0x00

	if !bytes.Equal(b, expected) {
		t.Logf("% b", b)
		t.Logf("% b", expected)

		t.Error("not equal")
	}
}

func TestBitmapPurchaseWithCashBackSample(t *testing.T) {
	fields := make(map[int]struct{})
	fields[2] = struct{}{}
	fields[3] = struct{}{}
	fields[4] = struct{}{}
	fields[5] = struct{}{}
	fields[7] = struct{}{}

	fields[11] = struct{}{}
	fields[12] = struct{}{}
	fields[18] = struct{}{}

	fields[22] = struct{}{}
	fields[24] = struct{}{}
	fields[26] = struct{}{}
	fields[28] = struct{}{}

	fields[32] = struct{}{}
	fields[33] = struct{}{}
	fields[35] = struct{}{}
	fields[37] = struct{}{}

	fields[41] = struct{}{}
	fields[42] = struct{}{}
	fields[43] = struct{}{}
	fields[46] = struct{}{}
	fields[49] = struct{}{}

	fields[54] = struct{}{}
	fields[62] = struct{}{}
	fields[63] = struct{}{}
	fields[100] = struct{}{}

	b, err := encodeBitmap(fields, true, true)
	if err != nil {
		t.Fatal(err)
	}
	expected := make([]byte, 16)
	expected[0] = 0xfa
	expected[1] = 0x30
	expected[2] = 0x45
	expected[3] = 0x51
	expected[4] = 0xa8
	expected[5] = 0xe4
	expected[6] = 0x84
	expected[7] = 0x06

	expected[0+8] = 0x00
	expected[1+8] = 0x00
	expected[2+8] = 0x00
	expected[3+8] = 0x00
	expected[4+8] = 0x10
	expected[5+8] = 0x00
	expected[6+8] = 0x00
	expected[7+8] = 0x00

	if !bytes.Equal(b, expected) {
		t.Logf("% b", b)
		t.Logf("% b", expected)

		t.Error("not equal")
	}
}
