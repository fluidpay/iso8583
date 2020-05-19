package iso8583

import "regexp"

const (
	binaryRegexString       = "^[0-1]{64}|[0-9A-F]{16}|[0-9A-F]{8}$"
	numberRegexString       = "^[0-9]+$"
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
	anpRegexString          = "^[a-zA-Z0-9]+\\s*$"
	ansRegexString          = "^[ -~]*$"
	yymmddhhmmssRegexString = "^([0-9]{2})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])([0-1][0-9]|2[0-3])([0-5][0-9]){2}$"
	mmddhhmmssRegexString   = "^(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])([0-1][0-9]|2[0-3])([0-5][0-9]){2}$"
	yymmRegexString         = "^([0-9]{2})(0[1-9]|1[0-2])$"
	mmddRegexString         = "^^(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])$"
	yymmddRegexString       = "^([0-9]{2})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])$"
	track2RegexString       = "^[0-9=D]+$"
)

var (
	binaryRegex       = regexp.MustCompile(binaryRegexString)
	numberRegex       = regexp.MustCompile(numberRegexString)
	alphaNumericRegex = regexp.MustCompile(alphaNumericRegexString)
	anpRegex          = regexp.MustCompile(anpRegexString)
	ansRegex          = regexp.MustCompile(ansRegexString)
	yymmddhhmmssRegex = regexp.MustCompile(yymmddhhmmssRegexString)
	mmddhhmmssRegex   = regexp.MustCompile(mmddhhmmssRegexString)
	yymmRegex         = regexp.MustCompile(yymmRegexString)
	mmddRegex         = regexp.MustCompile(mmddRegexString)
	yymmddRegex       = regexp.MustCompile(yymmddRegexString)
	track2Regex       = regexp.MustCompile(track2RegexString)
)
