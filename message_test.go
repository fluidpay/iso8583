package iso8583

import (
	"bytes"
	"testing"
)

func TestMessageWithFields234(t *testing.T) {
	m := &Message{
		DE2: NewNumeric("123"),
		DE3: NewNumeric("11"),
		DE4: NewNumeric("12"),
	}
	m.Mti = "1100"
	b, _ := m.Encode()
	if !bytes.Equal(b, []byte("1100700000000000000003123000011000000000012")) {
		t.Error("invalid encoded string")
	}
}

func TestNewMessageWithSecondaryBitmap(t *testing.T) {
	m := &Message{
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
	m.encoder = ASCII
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "F230040102A000000000000004000000" {
		t.Error("invalid bitmap")
	}
	if string(b) != "1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234" {
		t.Error("invalid encoded string")
	}
}

func TestBalanceInquiryFromAnATM(t *testing.T) {
	m := &Message{
		DE2:  NewNumeric("00000000000000"),                                         // Primary Account Number
		DE3:  NewNumeric("312000"),                                                 // Processing Code
		DE7:  NewNumeric("0108204503"),                                             // Date And Time, Transmission
		DE11: NewNumeric("007530"),                                                 // Systems Trace Audit Number
		DE12: NewNumeric("950108144500"),                                           // Date And Time, Local Transaction
		DE18: NewNumeric("6011"),                                                   // Merchant Type
		DE22: NewAlphanumeric("21120121014C"),                                      // Point Of Service Data Code
		DE24: NewNumeric("100"),                                                    // Function Code
		DE32: NewNumeric("10111111118"),                                            // Acquiring Institution Identification Code
		DE35: NewTrack2Code("56258101223070=99120041947"),                          // Track 2 Data
		DE41: NewANS("NY030400"),                                                   // Card Acceptor Terminal Identification
		DE42: NewANS(""),                                                           // Card Acceptor Identification Code
		DE43: NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE49: NewNumeric("840"),                                                    // Currency Code, Transaction
		DE52: NewBinaryHex("CD2C09CDCA80244C"),                                     // Personal Identification Number Data
	}
	m.Mti = "1100"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if bitmapHex(m.bitmapPrimary) != "6230450120E09000" {
		t.Error("invalid bitmap")
	}
	expected := "11006230450120E0900014000000000000003120000108204503007530950108144500601121120121014C10011101111111182656258101223070=99120041947NY030400               58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS840CD2C09CDCA80244C"
	if string(b) != expected {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestBalanceInquiryResponse(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("00000000000000"),                       // Primary Account Number
		DE3:   NewNumeric("312000"),                               // Processing Code
		DE7:   NewNumeric("0108204506"),                           // Date And Time, Transmission
		DE11:  NewNumeric("7530"),                                 // Systems Trace Audit Number
		DE12:  NewNumeric("950108144500"),                         // Date And Time, Local Transaction
		DE32:  NewNumeric("10111111118"),                          // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),                                  // Action Code
		DE54:  NewANS("2001840C0000007000002002840C000000600000"), // Amounts, Additional
		DE102: NewANS("00000012456184"),                           // Account Identification 1
	}
	m.encoder = ASCII
	m.Mti = "1110"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "E2300001020004000000000004000000" {
		t.Error("invalid bitmap")
	}
	expected := "1110E23000010200040000000000040000001400000000000000312000010820450600753095010814450011101111111180000402001840C0000007000002002840C0000006000001400000012456184"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestPurchaseWithCashBackRequest(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("0000000000000000"),                                       // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("20000"),                                                  // Amount, Transaction
		DE5:   NewNumeric("20000"),                                                  // Amount, Reconciliation
		DE7:   NewNumeric("0123205001"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("30402"),                                                  // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),                                           // Date And Time, Local Transaction
		DE18:  NewNumeric("5912"),                                                   // Merchant Type
		DE22:  NewAlphanumeric("21010121314C"),                                      // Point of Service Data Code
		DE24:  NewNumeric("200"),                                                    // Function Code
		DE26:  NewNumeric("5912"),                                                   // Card Acceptor Business Code
		DE28:  NewNumeric("950123"),                                                 // Date, Reconciliation
		DE32:  NewNumeric("10076401251"),                                            // Acquiring Institution Identification Code
		DE33:  NewNumeric("10111111118"),                                            // Forwarding Institution Identification Code
		DE35:  NewTrack2Code("54212248887288158=99120010109"),                       // Track 2 Data
		DE37:  NewANP("012401"),                                                     // Retrieval Reference Number
		DE41:  NewANS("NJ020111"),                                                   // Card Acceptor Terminal Identification
		DE42:  NewANS("73420"),                                                      // Card Acceptor Identification Code
		DE43:  NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE46:  NewANS("000"),                                                        // Amounts, Fees
		DE49:  NewNumeric("840"),                                                    // Currency Code, Transaction
		DE54:  NewANS("2040840D00000050002041840D0000015000"),                       // Amounts, Additional
		DE62:  NewNumeric("777001"),                                                 // Network Identifier
		DE63:  NewNumeric("0123"),                                                   //  Network Settlement Date
		DE100: NewNumeric("10222222226"),                                            // Receiving Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1110"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "FA304551A8E484060000000010000000" {
		t.Error("invalid bitmap")
	}
	expected := "1110FA304551A8E4840600000000100000001600000000000000000920000000000200000000000200000123205001030402950123154952591221010121314C2005912950123111007640125111101111111182954212248887288158=99120010109012401      NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400362040840D00000050002041840D000001500077700101231110222222226"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestPurchaseWithCashBackPartialApprovalToAquirer(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("0000000000000000"),                                       // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("15000"),                                                  // Amount, Transaction
		DE5:   NewNumeric("15000"),                                                  // Amount, Reconciliation
		DE7:   NewNumeric("0123205001"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("030402"),                                                 // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),                                           // Date And Time, Local Transaction
		DE18:  NewNumeric("5912"),                                                   // Merchant Type
		DE22:  NewAlphanumeric("21010121314C"),                                      // Point of Service Data Code
		DE24:  NewNumeric("200"),                                                    // Function Code
		DE26:  NewNumeric("5912"),                                                   // Card Acceptor Business Code
		DE28:  NewNumeric("950123"),                                                 // Date, Reconciliation
		DE30:  NewNumeric("000000020000000000020000"),                               // Amounts, Original
		DE32:  NewNumeric("10076401251"),                                            // Acquiring Institution Identification Code
		DE33:  NewNumeric("10111111118"),                                            // Forwarding Institution Identification Code
		DE35:  NewTrack2Code("54212248887288158=99120010109"),                       // Track 2 Data
		DE37:  NewANP("012401"),                                                     // Retrieval Reference Number
		DE39:  NewNumeric("002"),                                                    // Action code
		DE41:  NewANS("NJ020111"),                                                   // Card Acceptor Terminal Identification
		DE42:  NewANS("73420"),                                                      // Card Acceptor Identification Code
		DE43:  NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE46:  NewANS("000"),                                                        // Amounts, Fees
		DE49:  NewNumeric("840"),                                                    // Currency Code, Transaction
		DE62:  NewNumeric("777001"),                                                 // Network Identifier
		DE63:  NewNumeric("0123"),                                                   //  Network Settlement Date
		DE100: NewNumeric("10222222226"),                                            // Receiving Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1210"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "FA304555AAE480060000000010000000" {
		t.Log("FA304555AAE480060000000010000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "1210FA304555AAE4800600000000100000001600000000000000000920000000000150000000000150000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS00300084077700101231110222222226"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestPurchaseWithCashBackPartialApprovalFromIssuer(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("0000000000000000"),                                       // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("000000000000"),                                           // Amount, Transaction
		DE5:   NewNumeric("000000000000"),                                           // Amount, Reconciliation
		DE7:   NewNumeric("0123205001"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("030402"),                                                 // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),                                           // Date And Time, Local Transaction
		DE18:  NewNumeric("5912"),                                                   // Merchant Type
		DE22:  NewAlphanumeric("21010121314C"),                                      // Point of Service Data Code
		DE24:  NewNumeric("200"),                                                    // Function Code
		DE26:  NewNumeric("5912"),                                                   // Card Acceptor Business Code
		DE28:  NewNumeric("950123"),                                                 // Date, Reconciliation
		DE30:  NewNumeric("000000020000000000020000"),                               // Amounts, Original
		DE32:  NewNumeric("10076401251"),                                            // Acquiring Institution Identification Code
		DE33:  NewNumeric("10111111118"),                                            // Forwarding Institution Identification Code
		DE35:  NewTrack2Code("54212248887288158=99120010109"),                       // Track 2 Data
		DE37:  NewANP("012401"),                                                     // Retrieval Reference Number
		DE39:  NewNumeric("002"),                                                    // Action code
		DE41:  NewANS("NJ020111"),                                                   // Card Acceptor Terminal Identification
		DE42:  NewANS("73420"),                                                      // Card Acceptor Identification Code
		DE43:  NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE46:  NewANS("000"),                                                        // Amounts, Fees
		DE49:  NewNumeric("840"),                                                    // Currency Code, Transaction
		DE54:  NewANS("0202017840D000000015000"),                                    // Amounts, Additional
		DE62:  NewNumeric("777001"),                                                 // Network Identifier
		DE63:  NewNumeric("0123"),                                                   //  Network Settlement Date
		DE100: NewNumeric("10222222226"),                                            // Receiving Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1210"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "FA304555AAE484060000000010000000" {
		t.Log("FA304555AAE484060000000010000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "1210FA304555AAE4840600000000100000001600000000000000000920000000000000000000000000000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400230202017840D00000001500077700101231110222222226"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestPurchaseAuthorization(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("0000000000000000"), // Primary Account Number
		DE3:   NewNumeric("092000"),           // Processing Code
		DE4:   NewNumeric("20000"),            // Amount, Transaction
		DE7:   NewNumeric("0123205007"),       // Date And Time, Transmission
		DE11:  NewNumeric("030402"),           // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),     // Date And Time, Local Transaction
		DE32:  NewNumeric("10076401251"),    // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),              // Action code
		DE41:  NewANS("NJ020111"),             // Card Acceptor Terminal Identification
		DE49:  NewNumeric("840"),              // Currency Code, Transaction
		DE54:  NewANS(""),                     // Amounts, Additional
		DE102: NewANS("1234567890"),           // Account Identification 1
	}
	m.encoder = ASCII
	m.Mti = "1210"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "F2300001028084000000000004000000" {
		t.Log("F2300001028084000000000004000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "1210F230000102808400000000000400000016000000000000000009200000000002000001232050070304029501231549521110076401251000NJ020111840000101234567890"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}
