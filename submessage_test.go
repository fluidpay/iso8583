package iso8583

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSubMessageMessageWithFields234(t *testing.T) {
	m := &SubMessage{
		SE2: NewANS("Test Address"),
	}
	b, _ := m.Encode()
	expected := []byte("4000000000000000Test Address                 ")
	if !bytes.Equal(b, expected) {
		t.Errorf("Encoded should be %s, instead of %s", expected, b)
	}
}

func TestSubMessageDecode(t *testing.T) {
	msgToDecode := "4000000000000000Test Address                 "
	m := &SubMessage{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &SubMessage{
		SE2: NewANS("Test Address"),
	}
	//expectedMsg.Mti = "1200"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.SE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.SE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.SE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.SE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}