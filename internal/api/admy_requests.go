package api

import (
	"monitoring/internal/models"
	"fmt"
	"net/http"
)

func DSPbyHour(token)  {

	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/byhour/%s?token=%s", date, token)
}

func DSPbyTraffic() {
	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/bydate/?token=%s&from=%s&to=%s", token, from, to)
}

func DSPtotal()  {
	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/bydate?token=%s&traffic_type=%s", date, token)
}