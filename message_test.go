package iso8583

import (
	"bytes"
	"testing"
)

func TestMessageWithFields234(t *testing.T) {
	m := &Message{
		DE2: NewNumeric("123") ,
		DE3: NewNumeric("11"),
		DE4: NewNumeric("12"),
	}
	m.Mti = "1100"
	b,_:= m.Encode()
	if !bytes.Equal(b,[]byte("1100700000000000000003123000011000000000012")) {
		t.Error("invalid encoded string")
	}
}

func TestNewMessageWithSecondaryBitmap(t *testing.T) {
	m := &Message{
		DE2: NewNumeric("4846811212"),
		DE3: NewNumeric("201234"),
		DE4: NewNumeric("10000000"),
		DE7: NewNumeric("1107221800"),
		DE11: NewNumeric("000001"),
		DE12: NewNumeric("161204171926"),
		DE22: NewAlphanumeric("FABCDE123ABD"),
		DE32: NewNumeric("414243"),
		DE39: NewNumeric("000"),
		DE41: NewANS("termid12"),
		DE43: NewANS("Community1"),
		DE102: NewANS("12341234234"),
	}
	m.Mti = "1200"
	m.encoder = ASCII
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1) != "F230040102A000000000000004000000" {
		t.Error("invalid bitmap")
	}
	if string(b) != "1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234" {
		t.Error("invalid encoded string")
	}
}



