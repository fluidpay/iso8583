package iso8583

import (
	"reflect"
	"strconv"
	"strings"
)

type SubMessage struct {
	encoder       int
	bitmapPrimary uint64

	SE1 uint64 `format:"" length:"64"  json:",omitempty"`
	SE2 *ANS   `format:"" length:"29" validator:"ANS"  json:",omitempty"`
	SE3 *ANS   `format:"" length:"5" validator:"ANS"  json:",omitempty"`
	SE4 *N     `format:"" length:"10" validator:"N"  json:",omitempty"`
	SE5 *Reserved `json:",omitempty"` 
	SE6 *Reserved `json:",omitempty"`
	SE7 *ANS `format:"" length:"3" validator:"ANS"  json:",omitempty"`
	SE8 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE9 *AN  `format:"" length:"3" validator:"AN"  json:",omitempty"`
	SE10 *Reserved `json:",omitempty"`
	SE11 *AN `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE12 *Reserved `json:",omitempty"`
	SE13 *Reserved `json:",omitempty"`
	SE14 *Reserved `json:",omitempty"`
	SE15 *ANS `format:"" length:"2" validator:"ANS"  json:",omitempty"`
	SE16 *ANS `format:"" length:"2" validator:"ANS"  json:",omitempty"`
	SE17 *Reserved `json:",omitempty"`
	SE18 *ANS `format:"" length:"16" validator:"ANS"  json:",omitempty"`
	SE19 *ANS `format:"" length:"999" validator:"ANS"  json:",omitempty"`
	SE20 *ANS `format:"" length:"2" validator:"ANS"  json:",omitempty"`
	SE21 *ANS `format:"" length:"194" validator:"ANS"  json:",omitempty"` // ANS 255 if Tagged field 0008 is included
	SE22 *ANS `format:"LLLVAR" length:"255" validator:"ANS"  json:",omitempty"`
	SE23 *Reserved `json:",omitempty"`
	SE24 *ANS `format:"LLVAR" length:"99" validator:"ANS"  json:",omitempty"`
	SE25 *ANS `format:"LLVAR" length:"99" validator:"ANS"  json:",omitempty"`
	SE26 *N   `format:"" length:"3" validator:"N"  json:",omitempty"`
	SE27 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE28 *Reserved `json:",omitempty"`
	SE29 *ANS `format:"" length:"9" validator:"ANS"  json:",omitempty"`
	SE30 *N   `format:"" length:"4" validator:"N"  json:",omitempty"`
	SE31 *ANS `format:"LLLVAR" length:"255" validator:"ANS"  json:",omitempty"`
	SE32 *Reserved `json:",omitempty"`
	SE33 *Reserved `json:",omitempty"`
	SE34 *N   `format:"" length:"11" validator:"N"  json:",omitempty"`
	SE35 *N   `format:"" length:"11" validator:"N"  json:",omitempty"`
	SE36 *ANS `format:"" length:"15" validator:"ANS"  json:",omitempty"`
	SE37 *AN  `format:"" length:"7" validator:"AN"  json:",omitempty"`
	SE38 *ANP `format:"" length:"15" validator:"ANP"  json:",omitempty"`
	SE39 *AN  `format:"LLLVAR" length:"120" validator:"AN"  json:",omitempty"`
	SE40 *N   `format:"" length:"2" validator:"N"  json:",omitempty"`
	SE41 *BN  `format:"LLLVAR" length:"100" validator:"BN"  json:",omitempty"`
	SE42 *N   `format:"" length:"3" validator:"N"  json:",omitempty"`
	SE43 *Reserved `json:",omitempty"`
	SE44 *Reserved `json:",omitempty"`
	SE45 *Reserved `json:",omitempty"`
	SE46 *Reserved `json:",omitempty"`
	SE47 *Reserved `json:",omitempty"`
	SE48 *Reserved `json:",omitempty"`
	SE49 *Reserved `json:",omitempty"`
	SE50 *Reserved `json:",omitempty"`
	SE51 *Reserved `json:",omitempty"`
	SE52 *Reserved `json:",omitempty"`
	SE53 *Reserved `json:",omitempty"`
	SE54 *ANS `format:"" length:"6" validator:"ANS"  json:",omitempty"`
	SE55 *AN  `format:"LLLVAR" length:"120" validator:"AN"  json:",omitempty"`
	SE56 *N   `format:"LLVAR" length:"45" validator:"N"  json:",omitempty"`
	SE57 *AN  `format:"LLVAR" length:"10" validator:"AN"  json:",omitempty"`
	SE58 *N   `format:"LLVAR" length:"30" validator:"N"  json:",omitempty"`
	SE59 *AN  `format:"LLVAR" length:"10" validator:"AN"  json:",omitempty"`
	SE60 *N   `format:"" length:"1" validator:"N"  json:",omitempty"`
	SE61 *N   `format:"" length:"4" validator:"N"  json:",omitempty"`
	SE62 *N   `format:"" length:"1" validator:"N"  json:",omitempty"`
	SE63 *ANS `format:"" length:"94" validator:"ANS"  json:",omitempty"`
	SE64 *AN  `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE65 *AN  `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE66 *AN  `format:"" length:"6" validator:"AN"  json:",omitempty"`
	SE67 *AN  `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE68 *ANS `format:"" length:"40" validator:"ANS"  json:",omitempty"`
	SE69 *ANP `format:"" length:"6" validator:"ANP"  json:",omitempty"`
	SE70 *ANS `format:"" length:"15" validator:"ANS"  json:",omitempty"`
	SE71 *ANS `format:"LLLVAR" length:"999" validator:"ANS"  json:",omitempty"`
	SE72 *ANS `format:"LLLVAR" length:"999" validator:"ANS"  json:",omitempty"`
	SE73 *ANS `format:"LLLVAR" length:"999" validator:"ANS"  json:",omitempty"`
	SE74 *ANS `format:"LLLVAR" length:"999" validator:"ANS"  json:",omitempty"`
	SE75 *ANS `format:"" length:"3" validator:"ANS"  json:",omitempty"`
	SE76 *ANS `format:"" length:"23" validator:"ANS"  json:",omitempty"`
	SE77 *ANS `format:"" length:"12" validator:"ANS"  json:",omitempty"`
	SE78 *ANS `format:"" length:"15" validator:"ANS"  json:",omitempty"`
	SE79 *ANS `format:"" length:"4" validator:"ANS"  json:",omitempty"`
	SE80 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE81 *ANS `format:"" length:"2" validator:"ANS"  json:",omitempty"`
	SE82 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE83 *AN  `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE84 *ANS `format:"" length:"11" validator:"ANS"  json:",omitempty"`
	SE85 *BN  `format:"LLLVAR" length:"256" validator:"BN"  json:",omitempty"`
	SE86 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE87 *Reserved `json:",omitempty"`
	SE88 *ANS `format:"" length:"1" validator:"ANS"  json:",omitempty"`
	SE89 *ANS `format:"" length:"15" validator:"ANS"  json:",omitempty"`
	SE90 *ANS `format:"" length:"6" validator:"ANS"  json:",omitempty"`
	SE91 *ANS `format:"LLLVAR" length:"255" validator:"ANS"  json:",omitempty"`
	SE92 *Reserved `json:",omitempty"`
	SE93 *N  `format:"" length:"4" validator:"N"  json:",omitempty"`
	SE94 *N  `format:"LLVAR" length:"19" validator:"N"  json:",omitempty"`
	SE95 *AN `format:"" length:"2" validator:"AN"  json:",omitempty"`
	SE96 *N  `format:"LLVAR" length:"19" validator:"N"  json:",omitempty"`
	SE97 *N  `format:"" length:"11" validator:"N"  json:",omitempty"`
	SE98 *AN `format:"" length:"1" validator:"AN"  json:",omitempty"`
	SE99 *AN `format:"LLVAR" length:"99" validator:"AN"  json:",omitempty"`
}

func (m *SubMessage) Encode() ([]byte, error) {
	res := make([]byte, 0)

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
		// get field index, e.g. for SE2 index=2
		index, err := strconv.Atoi(strings.Trim(sf.Name, "SE"))
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

		f := v.Field(i).Interface().(field)

		if index <= 64 {
			bitmapPrimary = addField(bitmapPrimary, uint8(index))
		} else {
			// if we need secondary bitmap, set first bit in primary bitmap
			bitmapPrimary |= 1 << 63
			bitmapSecondary = addField(bitmapSecondary, uint8(index-64))
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
		m.SE1 = bitmapSecondary
		res = append(res, []byte(bitmapHex(bitmapSecondary))...)
	}

	// append iso data elements to result
	res = append(res, data...)
	return res, nil
}

func (m *SubMessage) Decode(bytes []byte) error {
	var err error

	// decode MTI

	// it is an iterator, watching where we are currently in the iteration,
	// which byte will be the starting position of the next decode
	it := 0

	// decode bitmaps
	//decode primary bitmap
	m.bitmapPrimary, err = decodeHexString(string(bytes[it : it+16]))
	if err != nil {
		return err
	}
	it += 16

	// if first bit is 1, it means that we have secondary bitmap, decode secondary bitmap
	if isBitSet(m.bitmapPrimary, 1) {
		m.SE1, err = decodeHexString(string(bytes[it : it+16]))
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
		index, err := strconv.Atoi(strings.Trim(sf.Name, "SE"))
		if err != nil {
			return err
		}
		// if index < 64, search in primary bitmap if it is set
		if index <= 64 && !isBitSet(m.bitmapPrimary, uint8(index)) {
			continue
		}

		// if index > 64, search in secondary bitmap if it is set
		if index > 64 && !isBitSet(m.SE1, uint8(index-64)) {
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

		f := v.Field(i).Interface().(field)
		nextFieldOffset, err := f.Decode(bytes[it:], m.encoder, length, format, validator)
		if err != nil {
			return err
		}
		it += nextFieldOffset
	}
	return nil
}
