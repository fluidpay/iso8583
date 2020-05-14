package iso8583

import (
	"bytes"
	"errors"
)

const (
	// ASCII is ASCII encoding
	ASCII = iota
	// BCDIC is BCDIC encoding
	BCDIC
)

type field interface {
	Encode(encoder, length int) ([]byte, error)

	Decode(raw []byte, encoder, length int)
}

type N struct {
	value []byte
}

func NewNumeric(value string) *N {
	return &N{value: []byte(value)}
}

func (n *N) Encode(encoder, length int) ([]byte, error) {
	val := n.value
	if len(val) < length {
		val = append(bytes.Repeat([]byte("0"), length-len(val)), val...)
	}
	if len(val) != length {
		return nil, errors.New("invalid value length")
	}
	switch encoder {

	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
}

func (n *N) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

type AN struct {
	value []byte
}

func NewAlphanumeric(value string) *AN {
	return &AN{value: []byte(value)}
}

func (an *AN) Encode(encoder, length int) ([]byte, error) {
	val := an.value
	if len(val) < length {
		val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
	}
	if len(val) != length {
		return nil, errors.New("invalid value length")
	}
	return val, nil
}

func (an *AN) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}
