# Golang ISO-8583

ISO-8583 support library for Golang

## Usage

```
// Encode
m := iso8583.Message{
  DE2:   NewNumeric("4846811212"),        // Primary Account Number
  DE3:   NewNumeric("201234"),            // Processing Code
  DE4:   NewNumeric("10000000"),          // Amount, Transaction
  DE7:   NewNumeric("1107221800"),        // Date And Time, Transmission
  DE11:  NewNumeric("000001"),            // Systems Trace Audit Number
  DE12:  NewNumeric("161204171926"),      // Date And Time, Local Transaction
  DE22:  NewAlphanumeric("FABCDE123ABD"), // Point Of Service Data Code
  DE32:  NewNumeric("414243"),            // Acquiring Institution Identification Code
  DE39:  NewNumeric("000"),               // Action Code
  DE41:  NewANS("termid12"),              // Card Acceptor Terminal Identification
  DE43:  NewANS("Community1"),            // Card Acceptor Name/Location
  DE102: NewANS("12341234234"),           // Account Identification 1
}
m.Mti = "1200"
m.SetEncoder(iso8583.ASCII)

b, _ := m.Encode()
//  Output: 1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234
```

```
// Decode
msg := "1210FA304555AAE4840600000000100000001600000000000000000920000000000000000000000000000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400230202017840D00000001500077700101231110222222226"
	m := Message{}
	m.SetEncoder(iso8583.ASCII)
	_ := m.Decode([]byte(msg))
```
