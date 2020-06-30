package iso8583

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNumericFixed(t *testing.T) {
	n := NewNumeric("12345")
	b, err := n.Encode(ASCII, 5, "", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("12345")) {
		t.Error("bad encoding")
	}

	n = NewNumeric("12345")
	b, err = n.Encode(ASCII, 6, "", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("012345")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12345", "string should be the same")

	n = NewNumeric("12345")
	b, err = n.Encode(ASCII, 10, "", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("0000012345")) {
		t.Error("bad encoding")
	}
	str = n.String()
	equals(t, str, "12345", "string should be the same")
}

func TestNumericLLVAR(t *testing.T) {
	n := NewNumeric("12345")
	b, err := n.Encode(ASCII, 19, "LLVAR", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("0512345")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12345", "string should be the same")

	b, err = n.Encode(ASCII, 4, "LLVAR", "N")
	if err == nil {
		t.Error("expecting error, length 4 < len(n)")
	}

	b, err = n.Encode(ASCII, 5, "LLVAR", "N")
	if err != nil {
		t.Error(err)
	}
}

func TestNumericLLLVAR(t *testing.T) {
	n := NewNumeric("12345")
	b, err := n.Encode(ASCII, 19, "LLLVAR", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("00512345")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12345", "string should be the same")

	b, err = n.Encode(ASCII, 4, "LLLVAR", "N")
	if err == nil {
		t.Error("expecting error, length 4 < len(n)")
	}

	b, err = n.Encode(ASCII, 5, "LLLVAR", "N")
	if err != nil {
		t.Error(err)
	}
}

func TestAlphaNumericFixed(t *testing.T) {
	n := NewAlphanumeric("12AN")
	b, err := n.Encode(ASCII, 10, "", "AN")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("12AN      ")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12AN", "string should be the same")
}

func TestAlphaNumericLLVAR(t *testing.T) {
	n := NewAlphanumeric("12ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS")
	b, err := n.Encode(ASCII, 99, "LLVAR", "AN")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("4812ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS", "string should be the same")
}

func TestAlphaNumericLLLVAR(t *testing.T) {
	n := NewAlphanumeric("12ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS")
	b, err := n.Encode(ASCII, 99, "LLLVAR", "AN")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("04812ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS")) {
		t.Error("bad encoding")
	}
	str := n.String()
	equals(t, str, "12ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS", "string should be the same")
}

func TestNDecode(t *testing.T) {
	b := []byte("0512345")
	n := NewNumeric("")
	_, err := n.Decode(b, ASCII, 19, "LLVAR", "N")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(n.Value))

	b = []byte("0000012356")
	nn, err := n.Decode(b, ASCII, 10, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(n.Value))
	t.Log(nn)
}

func TestMarshalJSON(t *testing.T) {
	m := Message{
		DE2: NewNumeric("1234412"),
		DE22: NewAlphanumeric("3123"),
		DE35: NewTrack2Code("latrack2"),
		DE37: NewANP("affa32"),
		DE41: NewANS("ans433"),
		DE52: NewBinary64("433"),
		DE125: &SubMessage{
			SE85: NewBN("123"),
		},
	}
	result, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"DE2":"1234412","DE22":"3123","DE35":"latrack2","DE37":"affa32","DE41":"ans433","DE52":"00000000000001B1","DE125":{"SE85":"123"}}`
	equals(t, string(result), expected, "")
}

func TestUnmarshalJSON(t *testing.T) {
	m := Message{}
	expected := `{"DE2":"1234412","DE22":"3123","DE35":"latrack2","DE37":"affa32","DE41":"ans433","DE52":"00000000000001B1","DE125":{"SE85":"123"}}`
	// expected := `{"DE2":"1234412"}`
	err := json.Unmarshal([]byte(expected), &m)
	if err != nil {
		t.Fatal(err)
	}

	equals(t, string(m.DE2.Value), "1234412", "")
	equals(t, string(m.DE22.Value), "3123", "")
	equals(t, string(m.DE35.Value), "latrack2", "")
	equals(t, string(m.DE37.Value), "affa32", "")
	equals(t, string(m.DE41.Value), "ans433", "")
	equals(t, string(m.DE52.Value), "00000000000001B1", "")
	equals(t, string(m.DE125.SE85.Value), "123", "")

	//equals(t, string(result), expected, "")
}

func equals(t *testing.T, actual, expected, description string) {
	if description == "" {
		description = "empty description"
	}
	if actual != expected {
		t.Errorf("Data should be %s, instead of %s [%s]", expected, actual, description)
	}
}
