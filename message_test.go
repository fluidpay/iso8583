package iso8583

import (
	"bytes"
	"testing"
)

func TestMessageWithFields234(t *testing.T) {
	m := &Message{
		DE2: NewNumeric("123"),
		DE3: NewNumeric("11"),
		DE4: NewNumeric("12"),
	}
	m.Mti = "1100"
	b, _ := m.Encode()
	if !bytes.Equal(b, []byte("1100700000000000000003123000011000000000012")) {
		t.Error("invalid encoded string")
	}
}

func TestNewMessageWithSecondaryBitmap(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("4846811212"),
		DE3:   NewNumeric("201234"),
		DE4:   NewNumeric("10000000"),
		DE7:   NewNumeric("1107221800"),
		DE11:  NewNumeric("000001"),
		DE12:  NewNumeric("161204171926"),
		DE22:  NewAlphanumeric("FABCDE123ABD"),
		DE32:  NewNumeric("414243"),
		DE39:  NewNumeric("000"),
		DE41:  NewANS("termid12"),
		DE43:  NewANS("Community1"),
		DE102: NewANS("12341234234"),
	}
	m.Mti = "1200"
	m.encoder = ASCII
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "F230040102A000000000000004000000" {
		t.Error("invalid bitmap")
	}
	if string(b) != "1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234" {
		t.Error("invalid encoded string")
	}
}

func TestBalanceInquiryFromAnATM(t *testing.T) {
	m := &Message{
		DE2:  NewNumeric("00000000000000"),
		DE3:  NewNumeric("312000"),
		DE7:  NewNumeric("0108204503"),
		DE11: NewNumeric("007530"),
		DE12: NewNumeric("950108144500"),
		DE18: NewNumeric("6011"),
		DE22: NewAlphanumeric("21120121014C"),
		DE24: NewNumeric("100"),
		DE32: NewNumeric("10111111118"),
		DE35: NewTrack2Code("56258101223070=99120041947"),
		DE41: NewANS("NY030400"),
		DE42: NewANS(""),
		DE43: NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"),
		DE49: NewNumeric("840"),
		DE52: NewBinaryHex("CD2C09CDCA80244C"),
	}
	m.Mti = "1100"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if bitmapHex(m.bitmapPrimary) != "6230450120E09000" {
		t.Error("invalid bitmap")
	}
	expected := "11006230450120E0900014000000000000003120000108204503007530950108144500601121120121014C10011101111111182656258101223070=99120041947NY030400               58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS840CD2C09CDCA80244C"
	if string(b) != expected {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestBalanceInquiryResponse(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("00000000000000"),
		DE3:   NewNumeric("312000"),
		DE7:   NewNumeric("0108204506"),
		DE11:  NewNumeric("7530"),
		DE12:  NewNumeric("950108144500"),
		DE32:  NewNumeric("10111111118"),
		DE39:  NewNumeric("000"),
		DE54:  NewANS("2001840C0000007000002002840C000000600000"),
		DE102: NewANS("00000012456184"),
	}
	m.encoder=ASCII
	m.Mti = "1110"
	b,err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "E2300001020004000000000004000000" {
		t.Error("invalid bitmap")
	}
	expected := "1110E23000010200040000000000040000001400000000000000312000010820450600753095010814450011101111111180000402001840C0000007000002002840C0000006000001400000012456184"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}
