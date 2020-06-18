package iso8583

import (
	"bytes"
	"reflect"
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

/*
	example of a Balance Inquiry from an ATM message where the processor is the source of the message.
	In this message, the checking account balance is requested.
*/
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
		DE52: NewBinary64Hex("CD2C09CDCA80244C"),                                   // Personal Identification Number Data
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

/*
	example of Balance Inquiry Response Message where the issuer processor is
	the source of the message. Message approves the transaction and returns ledger and available
	balance, as well as the customer’s checking account number.
*/

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

/*
	EPOC is the source of the message.
	The following message contains an example of a Purchase with Cash Back Request Message. In this
	message the customer requests $150.00 worth of goods and $50.00 cash-back.
*/

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

/*
	the following message contains an example of a Purchase with Cash Back Response Message where
	EPOC is sending the response to the acquirer. In this message the customer attempted to buy
	$150.00 worth of goods with $50.00 cash-back and is approved for the purchase amount
	only, $150.00.
*/

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

/*
	the following message contains a Purchase with Cash Back Response Message where the issuer is
	responding to EPOC. The customer attempted to buy $150.00 worth of goods with $50.00
	cash-back and is approved for the purchase amount only, $150.00.
*/

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

/*
	the following message contains an example of a Purchase Authorization message where the processor
	is the source of the message. This is an approval message that returns the checking account
	number but does not return account balances.
*/

func TestPurchaseAuthorization(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("0000000000000000"), // Primary Account Number
		DE3:   NewNumeric("092000"),           // Processing Code
		DE4:   NewNumeric("20000"),            // Amount, Transaction
		DE7:   NewNumeric("0123205007"),       // Date And Time, Transmission
		DE11:  NewNumeric("030402"),           // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),     // Date And Time, Local Transaction
		DE32:  NewNumeric("10076401251"),      // Acquiring Institution Identification Code
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

/*
	the following message contains an example of a Reversal Advice message where EPOC is the source
	of the message. In this message the customer canceled the transaction resulting in a full reversal.
*/

func TestReversalAdvice(t *testing.T) {
	m := &Message{
		DE2:   NewNumeric("000000000000000000"),                                     // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("20000"),                                                  // Amount, Transaction
		DE5:   NewNumeric("20000"),                                                  // Amount, Reconciliation
		DE7:   NewNumeric("0123205206"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("075809"),                                                 // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),                                           // Date And Time, Local Transaction
		DE18:  NewNumeric("5912"),                                                   // Merchant Type
		DE22:  NewAlphanumeric("21010121314C"),                                      // Point of Service Data Code
		DE24:  NewNumeric("400"),                                                    // Function Code
		DE26:  NewNumeric("5912"),                                                   // Card Acceptor Business Code
		DE28:  NewNumeric("950123"),                                                 // Date, Reconciliation
		DE32:  NewNumeric("10076401251"),                                            // Acquiring Institution Identification Code
		DE33:  NewNumeric("10111111118"),                                            // Forwarding Institution Identification Code
		DE35:  NewTrack2Code("5421224887288158=99120010109"),                        // Track 2 Data
		DE37:  NewANP("012401"),                                                     // Retrieval Reference Number
		DE41:  NewANS("NJ020111"),                                                   // Card Acceptor Terminal Identification
		DE42:  NewANS("73420"),                                                      // Card Acceptor Identification Code
		DE43:  NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE46:  NewANS("000"),                                                        // Amounts, Fees
		DE49:  NewNumeric("840"),                                                    // Currency Code, Transaction
		DE54:  NewANS("2040840D0000000050002041840D000000015000"),                   // Amounts, Additional
		DE56:  NewNumeric("12000304029501231549521110076401251"),                    // Original Data Elements
		DE62:  NewNumeric("777001"),                                                 // Network Identifier
		DE63:  NewNumeric("0123"),                                                   //  Network Settlement Date
		DE100: NewNumeric("10222222226"),                                            // Receiving Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1420"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "FA304551A8E485060000000010000000" {
		t.Log("FA304551A8E485060000000010000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "1420FA304551A8E485060000000010000000180000000000000000000920000000000200000000000200000123205206075809950123154952591221010121314C400591295012311100764012511110111111118285421224887288158=99120010109012401      NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400402040840D0000000050002041840D000000015000351200030402950123154952111007640125177700101231110222222226"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

/*
	the following message contains an example of a Reversal Advice Response Message where the
	processor is the source of the message.
*/

func TestReversalAdviceResponse(t *testing.T) {
	m := &Message{
		DE2:  NewNumeric("000000000000000000"), // Primary Account Number
		DE3:  NewNumeric("092000"),             // Processing Code
		DE4:  NewNumeric("20000"),              // Amount, Transaction
		DE7:  NewNumeric("0123210209"),         // Date And Time, Transmission
		DE11: NewNumeric("075809"),             // Systems Trace Audit Number
		DE12: NewNumeric("950123210135"),       // Date And Time, Local Transaction
		DE32: NewNumeric("10076401251"),        // Acquiring Institution Identification Code
		DE33: NewNumeric("10222222226"),        // Forwarding Institution Identification Code
		DE39: NewNumeric("400"),                // Action code
	}
	m.encoder = ASCII
	m.Mti = "1430"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary) != "7230000182000000" {
		t.Log("FA304551A8E485060000000010000000")
		t.Log(bitmapHex(m.bitmapPrimary))
		t.Error("invalid bitmap")
	}
	expected := "1430723000018200000018000000000000000000092000000000020000012321020907580995012321013511100764012511110222222226400"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

/*
	The following message contains an example of a Network Management Request message, initiated by
	EPOC (sign-on).
*/

func TestNetworkManagementRequest(t *testing.T) {
	m := &Message{
		DE7:  NewNumeric("0124081908"),   // Date And Time, Transmission
		DE11: NewNumeric("031972"),       // Systems Trace Audit Number
		DE12: NewNumeric("950124081904"), // Date And Time, Local Transaction
		DE24: NewNumeric("801"),          // Function Code
		DE93: NewNumeric("00000000001"),  // Transaction Destination Institution Identification Code
		DE94: NewNumeric("00000000002"),  // Transaction Originator Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1804"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "82300100000000000000000C00000000" {
		t.Log("82300100000000000000000C00000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "180482300100000000000000000C00000000012408190803197295012408190480111000000000011100000000002"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

/*
	the following message contains an example of a Network Management Request Response Message,
	initiated by the processor (sign-on accepted).
*/

func TestNetworkManagementRequestResponse(t *testing.T) {
	m := &Message{
		DE7:  NewNumeric("0124081920"),   // Date And Time, Transmission
		DE11: NewNumeric("031972"),       // Systems Trace Audit Number
		DE12: NewNumeric("950124081904"), // Date And Time, Local Transaction
		DE24: NewNumeric("801"),          // Function Code
		DE39: NewNumeric("800"),          // Action Code
		DE93: NewNumeric("10222222226"),  // Transaction Destination Institution Identification Code
		DE94: NewNumeric("10999999992"),  // Transaction Originator Institution Identification Code
	}
	m.encoder = ASCII
	m.Mti = "1804"
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != "82300100020000000000000C00000000" {
		t.Log("82300100020000000000000C00000000")
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Error("invalid bitmap")
	}
	expected := "180482300100020000000000000C00000000012408192003197295012408190480180011102222222261110999999992"
	if expected != string(b) {
		t.Log(expected)
		t.Log(string(b))
		t.Error("invalid encoding")
	}
}

func TestMessageDecode(t *testing.T) {
	msgToDecode := "1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("4846811212"),        // Primary Account Number
		DE3:   NewNumeric("201234"),            // Processing Code
		DE4:   NewNumeric("000010000000"),      // Amount, Transaction
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
	expectedMsg.Mti = "1200"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeBalanceInquiryFromAnATM(t *testing.T) {
	msgToDecode := "11006230450120E0900014000000000000003120000108204503007530950108144500601121120121014C10011101111111182656258101223070=99120041947NY030400               58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS840CD2C09CDCA80244C"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
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
		DE52: NewBinary64Hex("CD2C09CDCA80244C"),                                   // Personal Identification Number Data
	}
	expectedMsg.Mti = "1100"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeBalanceInquiryResponse(t *testing.T) {
	msgToDecode := "1110E23000010200040000000000040000001400000000000000312000010820450600753095010814450011101111111180000402001840C0000007000002002840C0000006000001400000012456184"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("00000000000000"),                       // Primary Account Number
		DE3:   NewNumeric("312000"),                               // Processing Code
		DE7:   NewNumeric("0108204506"),                           // Date And Time, Transmission
		DE11:  NewNumeric("007530"),                               // Systems Trace Audit Number
		DE12:  NewNumeric("950108144500"),                         // Date And Time, Local Transaction
		DE32:  NewNumeric("10111111118"),                          // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),                                  // Action Code
		DE54:  NewANS("2001840C0000007000002002840C000000600000"), // Amounts, Additional
		DE102: NewANS("00000012456184"),                           // Account Identification 1
	}
	expectedMsg.Mti = "1110"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodePurchaseWithCashBackRequest(t *testing.T) {
	msgToDecode := "1110FA304551A8E4840600000000100000001600000000000000000920000000000200000000000200000123205001030402950123154952591221010121314C2005912950123111007640125111101111111182954212248887288158=99120010109012401      NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400362040840D00000050002041840D000001500077700101231110222222226"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("0000000000000000"),                                       // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("000000020000"),                                           // Amount, Transaction
		DE5:   NewNumeric("000000020000"),                                           // Amount, Reconciliation
		DE7:   NewNumeric("0123205001"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("030402"),                                                 // Systems Trace Audit Number
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
	expectedMsg.Mti = "1110"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodePurchaseWithCashBackPartialApprovalToAquirer(t *testing.T) {
	msgToDecode := "1210FA304555AAE4800600000000100000001600000000000000000920000000000150000000000150000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS00300084077700101231110222222226"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("0000000000000000"),                                       // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("000000015000"),                                           // Amount, Transaction
		DE5:   NewNumeric("000000015000"),                                           // Amount, Reconciliation
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
	expectedMsg.Mti = "1210"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodePurchaseWithCashBackPartialApprovalFromIssuer(t *testing.T) {
	msgToDecode := "1210FA304555AAE4840600000000100000001600000000000000000920000000000000000000000000000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400230202017840D00000001500077700101231110222222226"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
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
	expectedMsg.Mti = "1210"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodePurchaseAuthorization(t *testing.T) {
	msgToDecode := "1210F230000102808400000000000400000016000000000000000009200000000002000001232050070304029501231549521110076401251000NJ020111840000101234567890"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("0000000000000000"), // Primary Account Number
		DE3:   NewNumeric("092000"),           // Processing Code
		DE4:   NewNumeric("000000020000"),     // Amount, Transaction
		DE7:   NewNumeric("0123205007"),       // Date And Time, Transmission
		DE11:  NewNumeric("030402"),           // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),     // Date And Time, Local Transaction
		DE32:  NewNumeric("10076401251"),      // Acquiring Institution Identification Code
		DE39:  NewNumeric("000"),              // Action code
		DE41:  NewANS("NJ020111"),             // Card Acceptor Terminal Identification
		DE49:  NewNumeric("840"),              // Currency Code, Transaction
		DE54:  NewANS(""),                     // Amounts, Additional
		DE102: NewANS("1234567890"),           // Account Identification 1
	}
	expectedMsg.Mti = "1210"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeReversalAdvice(t *testing.T) {
	msgToDecode := "1420FA304551A8E485060000000010000000180000000000000000000920000000000200000000000200000123205206075809950123154952591221010121314C400591295012311100764012511110111111118285421224887288158=99120010109012401      NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400402040840D0000000050002041840D000000015000351200030402950123154952111007640125177700101231110222222226"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:   NewNumeric("000000000000000000"),                                     // Primary Account Number
		DE3:   NewNumeric("092000"),                                                 // Processing Code
		DE4:   NewNumeric("000000020000"),                                           // Amount, Transaction
		DE5:   NewNumeric("000000020000"),                                           // Amount, Reconciliation
		DE7:   NewNumeric("0123205206"),                                             // Date And Time, Transmission
		DE11:  NewNumeric("075809"),                                                 // Systems Trace Audit Number
		DE12:  NewNumeric("950123154952"),                                           // Date And Time, Local Transaction
		DE18:  NewNumeric("5912"),                                                   // Merchant Type
		DE22:  NewAlphanumeric("21010121314C"),                                      // Point of Service Data Code
		DE24:  NewNumeric("400"),                                                    // Function Code
		DE26:  NewNumeric("5912"),                                                   // Card Acceptor Business Code
		DE28:  NewNumeric("950123"),                                                 // Date, Reconciliation
		DE32:  NewNumeric("10076401251"),                                            // Acquiring Institution Identification Code
		DE33:  NewNumeric("10111111118"),                                            // Forwarding Institution Identification Code
		DE35:  NewTrack2Code("5421224887288158=99120010109"),                        // Track 2 Data
		DE37:  NewANP("012401"),                                                     // Retrieval Reference Number
		DE41:  NewANS("NJ020111"),                                                   // Card Acceptor Terminal Identification
		DE42:  NewANS("73420"),                                                      // Card Acceptor Identification Code
		DE43:  NewANS("NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS"), // Card Acceptor Name/Location
		DE46:  NewANS("000"),                                                        // Amounts, Fees
		DE49:  NewNumeric("840"),                                                    // Currency Code, Transaction
		DE54:  NewANS("2040840D0000000050002041840D000000015000"),                   // Amounts, Additional
		DE56:  NewNumeric("12000304029501231549521110076401251"),                    // Original Data Elements
		DE62:  NewNumeric("777001"),                                                 // Network Identifier
		DE63:  NewNumeric("0123"),                                                   //  Network Settlement Date
		DE100: NewNumeric("10222222226"),                                            // Receiving Institution Identification Code
	}
	expectedMsg.Mti = "1420"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeReversalAdviceResponse(t *testing.T) {
	msgToDecode := "1430723000018200000018000000000000000000092000000000020000012321020907580995012321013511100764012511110222222226400"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE2:  NewNumeric("000000000000000000"), // Primary Account Number
		DE3:  NewNumeric("092000"),             // Processing Code
		DE4:  NewNumeric("000000020000"),       // Amount, Transaction
		DE7:  NewNumeric("0123210209"),         // Date And Time, Transmission
		DE11: NewNumeric("075809"),             // Systems Trace Audit Number
		DE12: NewNumeric("950123210135"),       // Date And Time, Local Transaction
		DE32: NewNumeric("10076401251"),        // Acquiring Institution Identification Code
		DE33: NewNumeric("10222222226"),        // Forwarding Institution Identification Code
		DE39: NewNumeric("400"),                // Action code
	}
	expectedMsg.Mti = "1430"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeNetworkManagementRequest(t *testing.T) {
	msgToDecode := "180482300100000000000000000C00000000012408190803197295012408190480111000000000011100000000002"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE7:  NewNumeric("0124081908"),   // Date And Time, Transmission
		DE11: NewNumeric("031972"),       // Systems Trace Audit Number
		DE12: NewNumeric("950124081904"), // Date And Time, Local Transaction
		DE24: NewNumeric("801"),          // Function Code
		DE93: NewNumeric("00000000001"),  // Transaction Destination Institution Identification Code
		DE94: NewNumeric("00000000002"),  // Transaction Originator Institution Identification Code
	}
	expectedMsg.Mti = "1804"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}

func TestDecodeNetworkManagementRequestResponse(t *testing.T) {
	msgToDecode := "180482300100020000000000000C00000000012408192003197295012408190480180011102222222261110999999992"
	m := &Message{}
	m.encoder = ASCII
	err := m.Decode([]byte(msgToDecode))
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := &Message{
		DE7:  NewNumeric("0124081920"),   // Date And Time, Transmission
		DE11: NewNumeric("031972"),       // Systems Trace Audit Number
		DE12: NewNumeric("950124081904"), // Date And Time, Local Transaction
		DE24: NewNumeric("801"),          // Function Code
		DE39: NewNumeric("800"),          // Action Code
		DE93: NewNumeric("10222222226"),  // Transaction Destination Institution Identification Code
		DE94: NewNumeric("10999999992"),  // Transaction Originator Institution Identification Code
	}
	expectedMsg.Mti = "1804"
	expectedMsg.encoder = ASCII
	expectedMsg.Encode()

	if bitmapHex(m.bitmapPrimary)+bitmapHex(m.DE1) != bitmapHex(expectedMsg.bitmapPrimary)+bitmapHex(expectedMsg.DE1) {
		t.Log(bitmapHex(m.bitmapPrimary) + bitmapHex(m.DE1))
		t.Log(bitmapHex(expectedMsg.bitmapPrimary) + bitmapHex(expectedMsg.DE1))
		t.Error("invalid bitmap")
	}
	if !reflect.DeepEqual(m, expectedMsg) {
		t.Error("not equal")
	}
}
