package delivery

type getPtListRequest struct {
	Period string `schema:"period" validate:"required,oneof='1h' '1d' '1mo' '1y'"`
	Tz     string `schema:"tz" validate:"required,validateTimezone"`
	T1     string `schema:"t1" validate:"required,validateTimestampFormat"`
	T2     string `schema:"t2" validate:"required,validateTimestampFormat"`
}

type getPtListErrorResponse struct {
	Status      string `json:"status,omitempty"`
	Description string `json:"desc,omitempty"`
}
