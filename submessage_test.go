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

func TestMessageWithSubMessageEncodeWithExtendedBitmap(t *testing.T) {
	m := &Message{
		DE2: NewNumeric("4846811212"), // Primary Account Number
		DE125: &SubMessage{
			SE2:  NewANS("Test Address"),     // AVS Cardholder Address
			SE3:  NewANS("12345"),            // AVS Additional Response Data
			SE4:  NewNumeric("9876543210"),   // Shared Branch National Point of Service Condition Code
			SE7:  NewANS("123"),              // Interchange Fee Indicator
			SE8:  NewANS("A"),                // Market–Specific Indicator
			SE9:  NewAlphanumeric("123"),     // Transaction Type Indicator
			SE11: NewAlphanumeric("1"),       // International Service Assessment (ISA) Indicator
			SE15: NewANS("12"),               // Multiple Clearing Sequence Number
			SE16: NewANS("12"),               // Multiple Clearing Sequence Count
			SE18: NewANS("1111111100000000"), // Shazam Transaction ID
			SE20: NewANS("12"),               // Business Application Identifier (BAI)
			SE98: NewAlphanumeric("1"),       // Authorization Type
		},
	}

	m.Mti = "1200"
	m.encoder = ASCII
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	hexExpected := "C0000000000000000000000000000008"
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != hexExpected {
		t.Errorf("bitmapHex should be %s, instead of %s", hexExpected, bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1))
	}
	subMessagehexExpected := "F3A35000000000000000000040000000"
	if bitmapHex(m.DE125.bitmapPrimary) + bitmapHex(m.DE125.SE1) != subMessagehexExpected {
		t.Errorf("bitmapHex should be %s, instead of %s", subMessagehexExpected, bitmapHex(m.DE125.bitmapPrimary) + bitmapHex(m.DE125.SE1))
	}

	expectedMsg := "1200C0000000000000000000000000000008104846811212107F3A35000000000000000000040000000Test Address                 123459876543210123A123112121111111100000000121"
	if string(b) != expectedMsg {
		t.Errorf("Encoded should be %s, instead of %s", expectedMsg, string(b))
	}

	sm,_ := m.DE125.Encode()
	expectedSubmessage := "F3A35000000000000000000040000000Test Address                 123459876543210123A123112121111111100000000121"
	if string(sm) != expectedSubmessage {
		t.Errorf("Encoded should be %s, instead of %s", expectedSubmessage, string(sm))
	}
}

func TestMessageWithSubMessageDecodeWithExtendedBitmap(t *testing.T) {
	msgToDecode := "1200C0000000000000000000000000000008104846811212107F3A35000000000000000000040000000Test Address                 123459876543210123A123112121111111100000000121"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2: NewNumeric("4846811212"), // Primary Account Number
		DE125: &SubMessage{
			SE2:  NewANS("Test Address"),     // AVS Cardholder Address
			SE3:  NewANS("12345"),            // AVS Additional Response Data
			SE4:  NewNumeric("9876543210"),   // Shared Branch National Point of Service Condition Code
			SE7:  NewANS("123"),              // Interchange Fee Indicator
			SE8:  NewANS("A"),                // Market–Specific Indicator
			SE9:  NewAlphanumeric("123"),     // Transaction Type Indicator
			SE11: NewAlphanumeric("1"),       // International Service Assessment (ISA) Indicator
			SE15: NewANS("12"),               // Multiple Clearing Sequence Number
			SE16: NewANS("12"),               // Multiple Clearing Sequence Count
			SE18: NewANS("1111111100000000"), // Shazam Transaction ID
			SE20: NewANS("12"),               // Business Application Identifier (BAI)
			SE98: NewAlphanumeric("1"),       // Authorization Type
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
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}
