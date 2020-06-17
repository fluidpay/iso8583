package iso8583

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Message struct {
	Mti string

	packedBitmap bool
	packedMsg    bool
	encoder      int

	bitmapPrimary uint64

	DE1   uint64 `format:"" length:"64"` //secondary bitmap
	DE2   *N     `format:"LLVAR" length:"19" validator:"N"`
	DE3   *N     `format:"" length:"6" validator:"N"`
	DE4   *N     `format:"" length:"12" validator:"N"`
	DE5   *N     `format:"" length:"12" validator:"N"`
	DE6   *N     `format:"" length:"12" validator:"N"`
	DE7   *N     `format:"" length:"10" validator:"MMDDHHMMSS"`
	DE9   *N     `format:"" length:"8" validator:"N"`
	DE10  *N     `format:"" length:"8" validator:"N"`
	DE11  *N     `format:"" length:"6" validator:"N"`
	DE12  *N     `format:"" length:"12" validator:"YYMMDDHHMMSS"`
	DE14  *N     `format:"" length:"4" validator:"YYMM"`
	DE16  *N     `format:"" length:"4" validator:"MMDD"`
	DE17  *N     `format:"" length:"4" validator:"MMDD"`
	DE18  *N     `format:"" length:"4" validator:"N"`
	DE19  *N     `format:"" length:"3" validator:"N"`
	DE22  *AN    `format:"" length:"12" validator:"AN"`
	DE23  *N     `format:"" length:"3" validator:"N"`
	DE24  *N     `format:"" length:"3" validator:"N"`
	DE25  *N     `format:"" length:"4" validator:"N"`
	DE26  *N     `format:"" length:"4" validator:"N"`
	DE28  *N     `format:"" length:"6" validator:"YYMMDD"`
	DE30  *N     `format:"" length:"24" validator:"N"`
	DE32  *N     `format:"LLVAR" length:"11" validator:"N"`
	DE33  *N     `format:"LLVAR" length:"11" validator:"N"`
	DE34  *N     `format:"LLVAR" length:"28" validator:"N"`
	DE35  *Z     `format:"LLVAR" length:"37" validator:"Z"`
	DE37  *ANP   `format:"" length:"12" validator:"ANP"`
	DE38  *ANP   `format:"" length:"6" validator:"ANP"`
	DE39  *N     `format:"" length:"3" validator:"N"`
	DE41  *ANS   `format:"" length:"8" validator:"ANS"`
	DE42  *ANS   `format:"" length:"15" validator:"ANS"`
	DE43  *ANS   `format:"LLVAR" length:"99" validator:"ANS"`
	DE46  *ANS   `format:"LLLVAR" length:"186" validator:"ANS"`
	DE47  *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE48  *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE49  *N     `format:"" length:"3" validator:"N"`
	DE50  *N     `format:"" length:"3" validator:"N"`
	DE51  *N     `format:"" length:"3" validator:"N"`
	DE52  *B     `format:"" length:"64" validator:"B"`
	DE54  *ANS   `format:"LLLVAR" length:"120" validator:"ANS"`
	DE56  *N     `format:"LLVAR" length:"35" validator:"N"`
	DE57  *N     `format:"" length:"3" validator:"N"`
	DE58  *N     `format:"LLVAR" length:"11" validator:"N"`
	DE59  *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE62  *N     `format:"" length:"6" validator:"N"`
	DE63  *N     `format:"" length:"4" validator:"MMDD"`
	DE66  *ANS   `format:"LLLVAR" length:"204" validator:"ANS"`
	DE72  *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE93  *N     `format:"LLVAR" length:"11" validator:"N"`
	DE94  *N     `format:"LLVAR" length:"11" validator:"N"`
	DE95  *ANS   `format:"LLLLVAR" length:"9999" validator:"ANS"`
	DE96  *ANS   `format:"LLLVAR" length:"100" validator:"ANS"`
	DE100 *N     `format:"LLVAR" length:"11" validator:"N"`
	DE101 *ANS   `format:"LLVAR" length:"17" validator:"ANS"`
	DE102 *ANS   `format:"LLVAR" length:"28" validator:"ANS"`
	DE103 *ANS   `format:"LLVAR" length:"28" validator:"ANS"`
	DE111 *ANS   `format:"LLLLVAR" length:"9999" validator:"ANS"`
	DE123 *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE124 *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE125 *SubMessage   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE126 *ANS   `format:"LLLVAR" length:"999" validator:"ANS"`
	DE127 *ANS   `format:"LLLLVAR" length:"9999" validator:"ANS"`
	DE128 *ANS   `format:"LLLLLVAR" length:"99999" validator:"ANS"`
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
				// TODO
				//continue
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

func (m *Message) PackedBitmap(packed bool) {
	m.packedBitmap = packed
}

func (m *Message) PackedMessage(packed bool) {
	m.packedMsg = packed
}

func (m *Message) SetEncoder(encoder int) {
	m.encoder = encoder
}
