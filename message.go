package iso8583

type Message struct {
	Mti string
	bitmap []byte
	packedBitmap bool
	packedMsg bool
	//Fields Fields
}
