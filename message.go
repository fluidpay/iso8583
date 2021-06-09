package iso8583

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Message struct {
	Mti string `json:",omitempty"`

	packedBitmap bool
	packedMsg    bool
	encoder      int

	bitmapPrimary uint64

	SafeLog bool `json:"-"` // This determines whether or not to log DE2

	DE1   uint64      `format:"" length:"64" json:",omitempty"` //secondary bitmap
	DE2   *N          `format:"LLVAR" length:"19" validator:"N" json:",omitempty"`
	DE3   *N          `format:"" length:"6" validator:"N" json:",omitempty"`
	DE4   *N          `format:"" length:"12" validator:"N" json:",omitempty"`
	DE5   *N          `format:"" length:"12" validator:"N" json:",omitempty"`
	DE6   *N          `format:"" length:"12" validator:"N" json:",omitempty"`
	DE7   *N          `format:"" length:"10" validator:"MMDDHHMMSS" json:",omitempty"`
	DE9   *N          `format:"" length:"8" validator:"N" json:",omitempty"`
	DE10  *N          `format:"" length:"8" validator:"N" json:",omitempty"`
	DE11  *N          `format:"" length:"6" validator:"N" json:",omitempty"`
	DE12  *N          `format:"" length:"12" validator:"YYMMDDHHMMSS" json:",omitempty"`
	DE14  *N          `format:"" length:"4" validator:"YYMM" json:",omitempty"`
	DE16  *N          `format:"" length:"4" validator:"MMDD" json:",omitempty"`
	DE17  *N          `format:"" length:"4" validator:"MMDD" json:",omitempty"`
	DE18  *N          `format:"" length:"4" validator:"N" json:",omitempty"`
	DE19  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE22  *AN         `format:"" length:"12" validator:"AN" json:",omitempty"`
	DE23  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE24  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE25  *N          `format:"" length:"4" validator:"N" json:",omitempty"`
	DE26  *N          `format:"" length:"4" validator:"N" json:",omitempty"`
	DE28  *N          `format:"" length:"6" validator:"YYMMDD" json:",omitempty"`
	DE30  *N          `format:"" length:"24" validator:"N" json:",omitempty"`
	DE32  *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE33  *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE34  *N          `format:"LLVAR" length:"28" validator:"N" json:",omitempty"`
	DE35  *Z          `format:"LLVAR" length:"37" validator:"Z" json:",omitempty"`
	DE37  *ANP        `format:"" length:"12" validator:"ANP" json:",omitempty"`
	DE38  *ANP        `format:"" length:"6" validator:"ANP" json:",omitempty"`
	DE39  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE41  *ANS        `format:"" length:"8" validator:"ANS" json:",omitempty"`
	DE42  *ANS        `format:"" length:"15" validator:"ANS" json:",omitempty"`
	DE43  *ANS        `format:"LLVAR" length:"99" validator:"ANS" json:",omitempty"`
	DE46  *ANS        `format:"LLLVAR" length:"186" validator:"ANS" json:",omitempty"`
	DE47  *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE48  *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE49  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE50  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE51  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE52  *B64        `format:"" length:"64" validator:"B64" json:",omitempty"`
	DE54  *ANS        `format:"LLLVAR" length:"120" validator:"ANS" json:",omitempty"`
	DE56  *N          `format:"LLVAR" length:"35" validator:"N" json:",omitempty"`
	DE57  *N          `format:"" length:"3" validator:"N" json:",omitempty"`
	DE58  *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE59  *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE62  *N          `format:"" length:"6" validator:"N" json:",omitempty"`
	DE63  *N          `format:"" length:"4" validator:"MMDD" json:",omitempty"`
	DE66  *ANS        `format:"LLLVAR" length:"204" validator:"ANS" json:",omitempty"`
	DE72  *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE93  *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE94  *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE95  *ANS        `format:"LLLLVAR" length:"9999" validator:"ANS" json:",omitempty"`
	DE96  *ANS        `format:"LLLVAR" length:"100" validator:"ANS" json:",omitempty"`
	DE100 *N          `format:"LLVAR" length:"11" validator:"N" json:",omitempty"`
	DE101 *ANS        `format:"LLVAR" length:"17" validator:"ANS" json:",omitempty"`
	DE102 *ANS        `format:"LLVAR" length:"28" validator:"ANS" json:",omitempty"`
	DE103 *ANS        `format:"LLVAR" length:"28" validator:"ANS" json:",omitempty"`
	DE111 *ANS        `format:"LLLLVAR" length:"9999" validator:"ANS" json:",omitempty"`
	DE123 *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE124 *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE125 *SubMessage `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE126 *ANS        `format:"LLLVAR" length:"999" validator:"ANS" json:",omitempty"`
	DE127 *ANS        `format:"LLLLVAR" length:"9999" validator:"ANS" json:",omitempty"`
	DE128 *ANS        `format:"LLLLLVAR" length:"99999" validator:"ANS" json:",omitempty"`
}

func New() *Message {
	return &Message{}
}

func NewSafe() *Message {
	return &Message{SafeLog: true}
}

func (m *Message) Encode() ([]byte, error) {
	res := make([]byte, 0)

	// append mti
	if len(m.Mti) != 4 {
		return []byte{}, errors.New("invalid MTI length")
	}
	switch m.encoder {
	case BCDIC:
	case ASCII:
		res = append(res, []byte(m.Mti)...)
	}

	// initialize bitmaps
	var bitmapPrimary uint64
	var bitmapSecondary uint64

	data := make([]byte, 0, 512)

	v := reflect.Indirect(reflect.ValueOf(m))
	t := v.Type()
	// iterate through iso fields, if field is not empty,
	// encode and append it to data, and set the proper bit in bitmap
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() != reflect.Ptr || v.Field(i).IsNil() {
			continue
		}
		sf := t.Field(i)
		// get field index, e.g. for DE2 index=2
		index, err := strconv.Atoi(strings.Trim(sf.Name, "DE"))
		if err != nil {
			return nil, err
		}
		// field length
		length, err := strconv.Atoi(sf.Tag.Get("length"))
		if err != nil {
			return nil, err
		}
		// iso field format, "" means that field has fixed length, but
		// e.g. LLVAR means that length indicator is encoded in the first 2 byte
		// in this case length means "maximum length
		format := sf.Tag.Get("format")
		validator := sf.Tag.Get("validator")

		if index <= 64 {
			bitmapPrimary = addField(bitmapPrimary, uint8(index))
		} else {
			// if we need secondary bitmap, set first bit in primary bitmap
			bitmapPrimary |= 1 << 63
			bitmapSecondary = addField(bitmapSecondary, uint8(index-64))
		}

		// tweak of submessage
		f, ok := v.Field(i).Interface().(field)
		if !ok {
			if sm, ok := v.Field(i).Interface().(*SubMessage); ok {
				res, err := sm.Encode()
				if err != nil {
					return nil, err
				}
				d, err := NewANS(string(res)).Encode(m.encoder, length, format, validator)
				data = append(data, d...)
				continue
			} else {
				continue
			}
		}

		// encode field, append it to data
		d, err := f.Encode(m.encoder, length, format, validator)
		if err != nil {
			return nil, err
		}
		data = append(data, d...)
	}

	// append bitmaps to result
	m.bitmapPrimary = bitmapPrimary
	res = append(res, []byte(bitmapHex(bitmapPrimary))...)

	if bitmapSecondary != 0 {
		m.DE1 = bitmapSecondary
		res = append(res, []byte(bitmapHex(bitmapSecondary))...)
	}

	// append iso data elements to result
	res = append(res, data...)
	return res, nil
}

func (m *Message) Decode(bytes []byte) error {
	var err error

	// decode MTI

	// it is an iterator, watching where we are currently in the iteration,
	// which byte will be the starting position of the next decode
	it := 4
	m.Mti = string(bytes[:it])

	// decode bitmaps
	//decode primary bitmap
	m.bitmapPrimary, err = decodeHexString(string(bytes[it : it+16]))
	if err != nil {
		return err
	}
	it += 16

	// if first bit is 1, it means that we have secondary bitmap, decode secondary bitmap
	if isBitSet(m.bitmapPrimary, 1) {
		m.DE1, err = decodeHexString(string(bytes[it : it+16]))
		if err != nil {
			return err
		}
		it += 16
	}

	v := reflect.Indirect(reflect.ValueOf(m))
	t := v.Type()
	// iterate through iso fields, if bitmap is not empty at bit position i,
	// set field with index i with proper value
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() != reflect.Ptr {
			continue
		}
		sf := t.Field(i)
		// get field index, e.g. for DE2 index=2
		index, err := strconv.Atoi(strings.Trim(sf.Name, "DE"))
		if err != nil {
			return err
		}
		// if index < 64, search in primary bitmap if it is set
		if index <= 64 && !isBitSet(m.bitmapPrimary, uint8(index)) {
			continue
		}

		// if index > 64, search in secondary bitmap if it is set
		if index > 64 && !isBitSet(m.DE1, uint8(index-64)) {
			continue
		}

		// field (maximum) length
		length, err := strconv.Atoi(sf.Tag.Get("length"))
		if err != nil {
			return err
		}
		// iso field format, "" means that field has fixed length, but
		// e.g. LLVAR means that length indicator is encoded in the first 2 byte
		// in this case length means "maximum length
		format := sf.Tag.Get("format")
		validator := sf.Tag.Get("validator")

		// Decode field
		structField := v.Field(i)

		// skip if unexported
		if !structField.CanSet() {
			continue
		}
		// initialize field with empty struct
		fieldTyp := reflect.New(structField.Type().Elem())
		structField.Set(fieldTyp)

		var nextFieldOffset int
		f, ok := v.Field(i).Interface().(field)
		if !ok {
			if sm, ok := v.Field(i).Interface().(*SubMessage); ok {
				ans := NewANS("")
				nextFieldOffset, err = ans.Decode(bytes[it:], m.encoder, length, format, validator)
				err = sm.Decode(ans.Value)
				if err != nil {
					return err
				}
				continue
				//v.Field(i).Set(reflect.ValueOf(sm))
			} else {
				continue
			}
		}
		nextFieldOffset, err = f.Decode(bytes[it:], m.encoder, length, format, validator)
		if err != nil {
			return err
		}
		it += nextFieldOffset
	}
	return nil
}

// String will take in the message struct and output to a string
// if SafeLog is true clear out DE2 for safe logging
func (m *Message) String() string {
	out, _ := json.Marshal(m)
	outStr := string(out)

	if m.SafeLog {
		return strings.Replace(outStr, m.DE2.String(), strings.Repeat("x", len(m.DE2.String())), -1)
	}

	return outStr
}

func (m *Message) PackedBitmap(packed bool) {
	m.packedBitmap = packed
}

func (m *Message) PackedMessage(packed bool) {
	m.packedMsg = packed
}

func (m *Message) SetEncoder(encoder int) {
	m.encoder = encoder
}
