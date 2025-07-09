package models

import (

)

type DSPbyHour struct {
	Data	[]struct {
		Amount		float64	`json:"amount"`
		Hour		uint	`json:"hour"`
		DSPname		string	`json:"dsp_name"`
		Requests	uint	`json:"requests"`
		Responses	uint	`json:"responses"`
	}	`json:"data"`
}


type DSPbyDate struct {
	Data	[]struct {
		Amount		float64	`json:"amount"`
		Hour		uint	`json:"hour"`
		DSPid		uint	`json:"dsp_id"`
		Requests	uint	`json:"requests"`
		Responses	uint	`json:"responses"`
	}	`json:"data"`
	Total	struct {
		Amount		float64 `json:"amount"`
	}	`json:"total"`
}