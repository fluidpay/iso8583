package iso8583

import (
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
	DE2   *N     `format:"LLVAR" length:"19"`
	DE3   *N     `format:"" length:"6"`
	DE4   *N     `format:"" length:"12"`
	DE5   *N     `format:"" length:"12"`
	DE6   *N     `format:"" length:"12"`
	DE7   *N     `format:"MMDDHHMMSS" length:"10"`
	DE9   *N     `format:"" length:"8"`
	DE10  *N     `format:"" length:"8"`
	DE11  *N     `format:"" length:"6"`
	DE12  *N     `format:"YYMMDDHHMMSS" length:"12"`
	DE14  *N     `format:"YYMM" length:"4"`
	DE16  *N     `format:"MMDD" length:"4"`
	DE17  *N     `format:"MMDD" length:"4"`
	DE18  *N     `format:"" length:"4"`
	DE19  *N     `format:"" length:"3"`
	DE22  *AN    `format:"" length:"12"`
	DE23  *N     `format:"" length:"3"`
	DE24  *N     `format:"" length:"3"`
	DE25  *N     `format:"" length:"4"`
	DE26  *N     `format:"" length:"4"`
	DE28  *N     `format:"YYMMDD" length:"6"`
	DE30  *N     `format:"" length:"24"`
	DE32  *N     `format:"LLVAR" length:"11"`
	DE33  *N     `format:"LLVAR" length:"11"`
	DE34  *N     `format:"LLVAR" length:"28"`
	DE35  *Z     `format:"LLVAR" length:"37"`
	DE37  *ANP   `format:"" length:"12"`
	DE38  *ANP   `format:"" length:"6"`
	DE39  *N     `format:"" length:"3"`
	DE41  *ANS   `format:"" length:"8"`
	DE42  *ANS   `format:"" length:"15"`
	DE43  *ANS   `format:"LLVAR" length:"99"`
	DE46  *ANS   `format:"LLLVAR" length:"186"`
	DE47  *ANS   `format:"LLLVAR" length:"999"`
	DE48  *ANS   `format:"LLLVAR" length:"999"`
	DE49  *N     `format:"" length:"3"`
	DE50  *N     `format:"" length:"3"`
	DE51  *N     `format:"" length:"3"`
	DE52  *B     `format:"" length:"64"`
	DE54  *ANS   `format:"LLLVAR" length:"120"`
	DE56  *N     `format:"LLVAR" length:"35"`
	DE57  *N     `format:"" length:"3"`
	DE58  *N     `format:"LLVAR" length:"11"`
	DE59  *ANS   `format:"LLLVAR" length:"999"`
	DE62  *N     `format:"" length:"6"`
	DE63  *N     `format:"MMDD" length:"4"`
	DE66  *ANS   `format:"LLLVAR" length:"204"`
	DE72  *ANS   `format:"LLLVAR" length:"999"`
	DE93  *N     `format:"LLVAR" length:"11"`
	DE94  *N     `format:"LLVAR" length:"11"`
	DE95  *ANS   `format:"LLLLVAR" length:"9999"`
	DE96  *ANS   `format:"LLLVAR" length:"100"`
	DE100 *N     `format:"LLVAR" length:"11"`
	DE101 *ANS   `format:"LLVAR" length:"17"`
	DE102 *ANS   `format:"LLVAR" length:"28"`
	DE103 *ANS   `format:"LLVAR" length:"28"`
	DE111 *ANS   `format:"LLLLVAR" length:"9999"`
	DE123 *ANS   `format:"LLLVAR" length:"999"`
	DE124 *ANS   `format:"LLLVAR" length:"999"`
	DE125 *ANS   `format:"LLLVAR" length:"999"`
	DE126 *ANS   `format:"LLLVAR" length:"999"`
	DE127 *ANS   `format:"LLLLVAR" length:"9999"`
	DE128 *ANS   `format:"LLLLLVAR" length:"99999"`
}

func (m *Message) Encode() ([]byte, error) {
	res := make([]byte, 0)
	// append mti
	mti, err := encodeMti(m.Mti)
	if err != nil {
		return nil, err
	}
	res = append(res, mti...)

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

		f := v.Field(i).Interface().(field)

		if index <= 64 {
			bitmapPrimary = addField(bitmapPrimary, uint8(index))
		} else {
			// if we need secondary bitmap, set first bit in primary bitmap
			bitmapPrimary |= 1 << 63
			bitmapSecondary = addField(bitmapSecondary, uint8(index-64))
		}

		// encode field, append it to data
		d, err := f.Encode(m.encoder, length, format)
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
