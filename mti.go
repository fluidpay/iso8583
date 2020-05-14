package iso8583

import "errors"

const (
	AuthorizationRequest         = "1100"
	AuthorizationRequestResponse = "1110"
	AuthorizationAdvice          = "1120"
	AuthorizationAdviceRepeat    = "1121"
	AuthorizationAdviceResponse  = "1130"

	FinancialRequest         = "1200"
	FinancialRequestResponse = "1210"
	FinancialAdvice          = "1220"
	FinancialAdviceRepeat    = "1221"
	FinancialAdviceResponse  = "1230"

	FileActionRequest  = "1304"
	FileActionResponse = "1314"

	ReversalAdvice         = "1420"
	ReversalAdviceRepeat   = "1421"
	ReversalAdviceResponse = "1430"

	AdministrativeAdvice         = "1624"
	AdministrativeAdviceResponse = "1634"

	NetworkManagementRequest         = "1804"
	NetworkManagementRequestResponse = "1814"
)

var mti = map[string]struct{}{
	AuthorizationRequest:         {},
	AuthorizationRequestResponse: {},
	AuthorizationAdvice:          {},
	AuthorizationAdviceRepeat:    {},
	AuthorizationAdviceResponse:  {},

	FinancialRequest:         {},
	FinancialRequestResponse: {},
	FinancialAdvice:          {},
	FinancialAdviceRepeat:    {},
	FinancialAdviceResponse:  {},

	FileActionRequest:  {},
	FileActionResponse: {},

	ReversalAdvice:         {},
	ReversalAdviceRepeat:   {},
	ReversalAdviceResponse: {},

	AdministrativeAdvice:         {},
	AdministrativeAdviceResponse: {},

	NetworkManagementRequest:         {},
	NetworkManagementRequestResponse: {},
}

func encodeMti(messageTypeId string) ([]byte, error) {
	if _, ok := mti[messageTypeId]; !ok {
		return nil, errors.New("invalid message type identifier")
	}
	return []byte(messageTypeId), nil
}
