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

	isEmpty() bool
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

func (n *N) isEmpty() bool {
	return len(n.value) == 0
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

func (an *AN) isEmpty() bool {
	return len(an.value) == 0
}

type B struct {
	value []byte
}

func NewBinary() *B {
	return &B{}
}

func (b *B) Encode(encoder, length int) ([]byte, error) {
	panic("implement me")
}

func (b *B) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

func (b *B) isEmpty() bool {
	return len(b.value) == 0
}

type Z struct {
	value []byte
}

func NewTrack2Code() *Z {
	return &Z{}
}

func (z *Z) Encode(encoder, length int) ([]byte, error) {
	panic("implement me")
}

func (z *Z) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}
func (z *Z) isEmpty() bool {
	return len(z.value) == 0
}

type ANP struct {
	value []byte
}

func NewANP() *ANP {
	return &ANP{}
}

func (A *ANP) Encode(encoder, length int) ([]byte, error) {
	panic("implement me")
}

func (A *ANP) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

func (anp *ANP) isEmpty() bool {
	return len(anp.value) == 0
}

type ANS struct {
	value []byte
}

func NewANS() *ANS {
	return &ANS{}
}

func (A *ANS) Encode(encoder, length int) ([]byte, error) {
	panic("implement me")
}

func (A *ANS) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

func (ans *ANS) isEmpty() bool {
	return len(ans.value) == 0
}
