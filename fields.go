package iso8583

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	// ASCII is ASCII encoding
	ASCII = iota
	// BCDIC is BCDIC encoding
	BCDIC
)

type field interface {
	Encode(encoder, length int, format, validator string) ([]byte, error)

	Decode(raw []byte, encoder, length int)

	isEmpty() bool
}

type N struct {
	value []byte
}

func NewNumeric(value string) *N {
	return &N{value: []byte(value)}
}

func (n *N) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := n.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add left padding with '0', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(bytes.Repeat([]byte("0"), length-len(val)), val...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
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

func (an *AN) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := an.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add right padding with ' ', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
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

func NewBinary(value string) *B {
	return &B{[]byte(value)}
}

func (b *B) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := b.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add right padding with ' ', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
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

func NewTrack2Code(value string) *Z {
	return &Z{[]byte(value)}
}

func (z *Z) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := z.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add right padding with ' ', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
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

func NewANP(value string) *ANP {
	return &ANP{[]byte(value)}
}

func (anp *ANP) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := anp.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add right padding with ' ', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
}

func (anp *ANP) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

func (anp *ANP) isEmpty() bool {
	return len(anp.value) == 0
}

type ANS struct {
	value []byte
}

func NewANS(value string) *ANS {
	return &ANS{[]byte(value)}
}

func (ans *ANS) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := ans.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}
	// if field has fixed length, add right padding with ' ', else
	// add length prefix in specific format
	if format == "FIXED" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		lInd, err := lengthIndicator(encoder, len(val), format)
		if err != nil {
			return nil, err
		}
		val = append(lInd, val...)
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
}

func (ans *ANS) Decode(raw []byte, encoder, length int) {
	panic("implement me")
}

func (ans *ANS) isEmpty() bool {
	return len(ans.value) == 0
}

func lengthIndicator(encoder, length int, format string) ([]byte, error) {
	switch format {
	case "LLVAR":
		if length < 0 || length > 99 {
			return nil, errors.New("invalid length for LLVAR")
		}
		return []byte(fmt.Sprintf("%02d", length)), nil
	case "LLLVAR":
		if length < 0 || length > 999 {
			return nil, errors.New("invalid length for LLLVAR")
		}
		return []byte(fmt.Sprintf("%03d", length)), nil

	case "LLLLVAR":
		if length < 0 || length > 9999 {
			return nil, errors.New("invalid length for LLLLVAR")
		}
		return []byte(fmt.Sprintf("%04d", length)), nil

	case "LLLLLVAR":
		if length < 0 || length > 99999 {
			return nil, errors.New("invalid length for LLLLLVAR")
		}
		return []byte(fmt.Sprintf("%05d", length)), nil
	case "FIXED":
		return []byte{}, nil
	default:
		return []byte{}, errors.New("invalid format")
	}
}

func validate(value, validator string) error {
	// TODO
	switch validator {
	case "N":
	case "B":
	case "AN":
	case "Z":
	case "ANP":
	case "ANS":
	case "YYMMDDHHMMSS":
	case "MMDDHHMMSS":
	case "YYMM":
	case "MMDD":
	case "YYMMDD":
	}
	return nil
}
