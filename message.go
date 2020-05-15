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
	//Fields Fields
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

type fieldInfo struct {
	index  int
	format string
	length int
	field  field
}

func parseFields(m *Message) (map[int]*fieldInfo, error) {
	fields := make(map[int]*fieldInfo)

	v := reflect.Indirect(reflect.ValueOf(m))
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() != reflect.Ptr || v.Field(i).IsNil() {
			continue
		}

		sf := t.Field(i)
		fName := sf.Name
		ind, err := strconv.Atoi(strings.Trim(fName, "DE"))
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(sf.Tag.Get("length"))
		if err != nil {
			return nil, err
		}
		format := sf.Tag.Get("format")

		field := v.Field(i).Interface().(field)
		fields[ind] = &fieldInfo{
			index:  ind,
			format: format,
			length: l,
			field:  field,
		}
	}
	return fields, nil
}
