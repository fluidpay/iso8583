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

func TestMessageWithSubMessage(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("4846811212"),        // Primary Account Number
		DE3:   NewNumeric("201234"),            // Processing Code
		DE4:   NewNumeric("10000000"),          // Amount, Transaction
		DE7:   NewNumeric("1107221800"),        // Date And Time, Transmission
		DE11:  NewNumeric("000001"),            // Systems Trace Audit Number
		DE12:  NewNumeric("161204171926"),      // Date And Time, Local Transaction
		DE22:  NewAlphanumeric("FABCDE123ABD"), // Point Of Service Data Code
		DE32:  NewNumeric("414243"),            // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),               // Action Code
		DE41:  NewANS("termid12"),              // Card Acceptor Terminal Identification
		DE43:  NewANS("Community1"),            // Card Acceptor Name/Location
		DE102: NewANS("12341234234"),           // Account Identification 1
		DE125: &SubMessage{
			SE2: NewANS("Test Address"),
		},
	}

	m.Mti = "1200"
	m.encoder = ASCII
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	hexExpected := "F230040102A000000000000004000008"
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != hexExpected {
		t.Errorf("bitmapHex should be %s, instead of %s", hexExpected, bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1))
	}
	expected := "1200F230040102A0000000000000040000081048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community111123412342340454000000000000000Test Address                 "
	if string(b) != expected {
		t.Errorf("Encoded should be %s, instead of %s", expected, string(b))
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

func TestMessageWithSubMessageDecode(t *testing.T) {
	msgToDecode := "1200F230040102A0000000000000040000081048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community111123412342340454000000000000000Test Address                 "
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("4846811212"),        // Primary Account Number
		DE3:   NewNumeric("201234"),            // Processing Code
		DE4:   NewNumeric("10000000"),          // Amount, Transaction
		DE7:   NewNumeric("1107221800"),        // Date And Time, Transmission
		DE11:  NewNumeric("000001"),            // Systems Trace Audit Number
		DE12:  NewNumeric("161204171926"),      // Date And Time, Local Transaction
		DE22:  NewAlphanumeric("FABCDE123ABD"), // Point Of Service Data Code
		DE32:  NewNumeric("414243"),            // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),               // Action Code
		DE41:  NewANS("termid12"),              // Card Acceptor Terminal Identification
		DE43:  NewANS("Community1"),            // Card Acceptor Name/Location
		DE102: NewANS("12341234234"),           // Account Identification 1
		DE125: &SubMessage{
			SE2: NewANS("Test Address"),
		},
	}
	expectedMsg.Mti = "1200"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if string(m.DE125.SE2.Value) != string(expectedMsg.DE125.SE2.Value) {
		t.Errorf("SE2 value should be %s, instead of %s", string(expectedMsg.DE125.SE2.Value), string(m.DE125.SE2.Value))
	}
}