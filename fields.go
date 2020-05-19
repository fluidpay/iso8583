package iso8583

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

const (
	// ASCII is ASCII encoding
	ASCII = iota
	// BCDIC is BCDIC encoding
	BCDIC
)

type field interface {
	Encode(encoder, length int, format, validator string) ([]byte, error)

	Decode(b []byte, encoder, length int, format, validator string) (int, error)

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
	if format == "" {
		if len(val) < length {
			val = append(bytes.Repeat([]byte("0"), length-len(val)), val...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		if len(val) > length {
			return nil, errors.New("invalid value length")
		}
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

func (n *N) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	var nextFieldOffset int

	switch encoder {
	case BCDIC:
	case ASCII:
		if format == "" {
			//n.value = bytes.TrimLeft(raw[:length], "0")
			n.value = raw[:length]
			nextFieldOffset = length
		} else {
			l, lenOfLen, err := getFieldLength(raw, encoder, length, format)
			if err != nil {
				return 0, err
			}
			n.value =  raw[lenOfLen : (l+lenOfLen)]
			nextFieldOffset = lenOfLen + l
		}
	}

	err := validate(string(n.value), validator)
	return nextFieldOffset, err
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
	if format == "" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		if len(val) > length {
			return nil, errors.New("invalid value length")
		}
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

func (an *AN) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	var nextFieldOffset int
	switch encoder {
	case BCDIC:
	case ASCII:
		if format == "" {
			an.value = bytes.TrimRight(raw[:length], " ")
			nextFieldOffset = length
		} else {
			l, lenOfLen, err := getFieldLength(raw, encoder, length, format)
			if err != nil {
				return 0, err
			}
			an.value = raw[lenOfLen : l+lenOfLen]
			nextFieldOffset = lenOfLen + l
		}
	}

	err := validate(string(an.value), validator)
	return nextFieldOffset, err
}

func (an *AN) isEmpty() bool {
	return len(an.value) == 0
}

type B struct {
	value []byte
}

// New Binary converts 64bit binary to Hexadecimal form, each 4 bits to 1 hexadecimal character
func NewBinary(value string) *B {
	b, _ := strconv.ParseUint(value, 10, 64)
	return &B{[]byte(bitmapHex(b))}
}

// New Binary converts 64bit binary to Hexadecimal form, each 4 bits to 1 hexadecimal character
func NewBinaryUint64(bin uint64) *B {
	return &B{[]byte(bitmapHex(bin))}
}

func NewBinaryHex(value string) *B {
	return &B{[]byte(value)}
}

func (b *B) Encode(encoder, length int, format, validator string) ([]byte, error) {
	val := b.value
	if err := validate(string(val), validator); err != nil {
		return []byte{}, err
	}

	switch encoder {
	case BCDIC:
		panic("implement me")
	default: //ASCII encoding
		return val, nil
	}
}

func (b *B) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	switch encoder {
	case BCDIC:
	case ASCII:
		b.value = raw[:length/4]
		err := validate(string(b.value), validator)
		return length / 4, err
	}
	return 0, nil
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
	if format == "" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		if len(val) > length {
			return nil, errors.New("invalid value length")
		}
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

func (z *Z) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	var nextFieldOffset int
	switch encoder {
	case BCDIC:
	case ASCII:
		l, lenOfLen, err := getFieldLength(raw, encoder, length, format)
		if err != nil {
			return 0, err
		}
		z.value = raw[lenOfLen : l+lenOfLen]
		nextFieldOffset = lenOfLen + l
	}

	err := validate(string(z.value), validator)
	return nextFieldOffset, err
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
	if format == "" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		if len(val) > length {
			return nil, errors.New("invalid value length")
		}
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

func (anp *ANP) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	var nextFieldOffset int
	switch encoder {
	case BCDIC:
	case ASCII:
		if format == "" {
			anp.value = bytes.TrimRight(raw[:length], " ")
			nextFieldOffset = length
		} else {
			l, lenOfLen, err := getFieldLength(raw, encoder, length, format)
			if err != nil {
				return 0, err
			}
			anp.value = raw[lenOfLen : l+lenOfLen]
			nextFieldOffset = lenOfLen + l
		}
	}

	err := validate(string(anp.value), validator)
	return nextFieldOffset, err
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
	if format == "" {
		if len(val) < length {
			val = append(val, bytes.Repeat([]byte(" "), length-len(val))...)
		}
		if len(val) != length {
			return nil, errors.New("invalid value length")
		}
	} else {
		if len(val) > length {
			return nil, errors.New("invalid value length")
		}
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

func (ans *ANS) Decode(raw []byte, encoder, length int, format, validator string) (int, error) {
	var nextFieldOffset int
	switch encoder {
	case BCDIC:
	case ASCII:
		if format == "" {
			ans.value = bytes.TrimRight(raw[:length], " ")
			nextFieldOffset = length
		} else {
			l, lenOfLen, err := getFieldLength(raw, encoder, length, format)
			if err != nil {
				return 0, err
			}
			ans.value = raw[lenOfLen : l+lenOfLen]
			nextFieldOffset = lenOfLen + l
		}
	}

	err := validate(string(ans.value), validator)
	return nextFieldOffset, err
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
	case "":
		return []byte{}, nil
	default:
		return []byte{}, errors.New("invalid format")
	}
}

func getFieldLength(raw []byte, encoder, maxLength int, format string) (length, lenOfLen int, err error) {
	switch format {
	case "LLVAR":
		lenOfLen = 2
		length, err = strconv.Atoi(string(raw[:2]))
		if err != nil {
			return length, lenOfLen, err
		}
		if length > maxLength {
			return length, lenOfLen, errors.New("invalid length")
		}
		return length, lenOfLen, err
	case "LLLVAR":
		lenOfLen = 3
		length, err := strconv.Atoi(string(raw[:3]))
		if err != nil {
			return length, lenOfLen, err

		}
		if length > maxLength {
			return length, lenOfLen, errors.New("invalid length")
		}
		return length, lenOfLen, err

	case "LLLLVAR":
		lenOfLen = 4
		length, err := strconv.Atoi(string(raw[4]))
		if err != nil {
			return length, lenOfLen, err
		}
		if length > maxLength {
			return length, lenOfLen, errors.New("invalid length")
		}
		return length, lenOfLen, err

	case "LLLLLVAR":
		lenOfLen = 5
		length, err := strconv.Atoi(string(raw[:5]))
		if err != nil {
			return length, lenOfLen, err
		}
		if length > maxLength {
			return length, lenOfLen, errors.New("invalid length")
		}
		return length, lenOfLen, err

	case "":
		lenOfLen = 0
		return length, lenOfLen, err

	default:
		return length, lenOfLen, errors.New("invalid format")
	}
}

func validate(value, validator string) error {
	switch validator {
	case "N":
		if !numberRegex.MatchString(value) {
			return errors.New("invalid number value format: " + value)
		}
	case "B":
		if !binaryRegex.MatchString(value) {
			return errors.New("invalid binary value format: " + value)
		}
	case "AN":
		if !alphaNumericRegex.MatchString(value) {
			return errors.New("invalid alphanumeric value format: " + value)
		}
	case "Z":
		if !track2Regex.MatchString(value) {
			return errors.New("only Track 2 code set characters (0–9, =, D) are allowed,invalid value format: " + value)
		}
	case "ANP":
		if !anpRegex.MatchString(value) {
			return errors.New("alphabetic, numeric, and special characters are allowed,invalid value format: " + value)
		}
	case "ANS":
		if !ansRegex.MatchString(value) {
			return errors.New("alphabetic, numeric, and special characters(ASCII printable characters) are allowed,invalid value format: " + value)
		}
	case "YYMMDDHHMMSS":
		if !yymmddhhmmssRegex.MatchString(value) {
			return errors.New("invalid YYMMDDHHMMSS value format: " + value)
		}
	case "MMDDHHMMSS":
		if !mmddhhmmssRegex.MatchString(value) {
			return errors.New("invalid MMDDHHMMSS value format: " + value)
		}
	case "YYMM":
		if !yymmRegex.MatchString(value) {
			return errors.New("invalid YYMM value format: " + value)
		}
	case "MMDD":
		if !mmddRegex.MatchString(value) {
			return errors.New("invalid MMDD value format: " + value)
		}
	case "YYMMDD":
		if !yymmddRegex.MatchString(value) {
			return errors.New("invalid YYMMDD value format: " + value)
		}
	}
	return nil
}
