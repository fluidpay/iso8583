package iso8583

import (
	"bytes"
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

func equals(t *testing.T, actual, expected, description string) {
	if actual != expected {
		t.Errorf("Data should be %s, instead of %s [%s]", expected, actual, description)
	}
}
