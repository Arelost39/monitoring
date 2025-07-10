package models

type DSPbyHour struct {
	Data	[]SingleDSP	`json:"data"`
}

type DSPbyDate struct {
	Data	[]SingleDSP	`json:"data"`
	Total	struct {
		Amount		float64 `json:"amount"`
	}	`json:"total"`
}

type SingleDSP struct {
	Amount		float64	`json:"amount"`
	Hour		uint	`json:"hour"`
	DSPname		string	`json:"dsp_name"`
	Requests	uint	`json:"requests"`
	Responses	uint	`json:"responses"`
}