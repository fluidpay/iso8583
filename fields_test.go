package iso8583

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
	str, err := n.String(ASCII, 6, "", "N")
	assert.Equal(t, str, "012345", "string should be the same.")

	n = NewNumeric("12345")
	b, err = n.Encode(ASCII, 10, "", "N")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte("0000012345")) {
		t.Error("bad encoding")
	}
	str, err = n.String(ASCII, 10, "", "N")
	if strings.Compare(str, "0000012345") != 0 {
		t.Error("string should be the same.")
	}
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
	str, err := n.String(ASCII, 19, "LLVAR", "N")
	if strings.Compare(str, "0512345") != 0 {
		t.Error("string should be the same.")
	}

	b, err = n.Encode(ASCII, 4, "LLVAR", "N")
	if err == nil {
		t.Error("expecting error, length 4 < len(n)")
	}
	str, err = n.String(ASCII, 4, "LLVAR", "N")
	if err == nil {
		t.Error("invalid value length")
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
	str, err := n.String(ASCII, 19, "LLLVAR", "N")
	if strings.Compare(str, "00512345") != 0 {
		t.Error("string should be the same.")
	}

	b, err = n.Encode(ASCII, 4, "LLLVAR", "N")
	if err == nil {
		t.Error("expecting error, length 4 < len(n)")
	}
	str, err = n.String(ASCII, 4, "LLLVAR", "N")
	if err == nil {
		t.Error("invalid value length")
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
	str, err := n.String(ASCII, 10, "", "AN")
	if strings.Compare(str, "12AN      ") != 0 {
		t.Error("string should be the same.")
	}
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
	str, err := n.String(ASCII, 99, "LLVAR", "AN")
	if strings.Compare(str, "4812ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS") != 0 {
		t.Error("string should be the same.")
	}
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
	str, err := n.String(ASCII, 99, "LLLVAR", "AN")
	if strings.Compare(str, "04812ANABCDEFGHTCASDASSAASCSACSACSACSACSACACSACSACS") != 0 {
		t.Error("string should be the same.")
	}
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
