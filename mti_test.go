package iso8583

import (
	"bytes"
	"testing"
)

func TestMti(t *testing.T)  {
	b,err := encodeMti("1100")
	if err != nil {
		t.Fatal(err)
	}
	bytes.Equal(b,[]byte("1100"))
}
