package api

import (
	m "monitoring/internal/models"
	"monitoring/internal/utils"
	
	"fmt"
	"net/http"
	// "time"
	"encoding/json"
	"io"
)

func DSPbyHour(token, date string) (m.DSPbyHour, error) {

	// loc, err := time.LoadLocation("Europe/Kaliningrad")
	// if err != nil {
    // 	utils.Log.Fatalf("Неизвестная timezone %v", err)
	// 	return m.DSPbyHour{}, err
	// }

	// date := time.Now().In(loc).Format("2006-01-02")
	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/byhour/%s?token=%s", date, token)
	resp, err := http.Get(url)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        bodyBytes, _ := io.ReadAll(resp.Body)
        return m.DSPbyHour{}, fmt.Errorf("Ошибка запроса к admy (DSPbyHour) %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var stat m.DSPbyHour
    if err := json.NewDecoder(resp.Body).Decode(&stat); err != nil {
        utils.Log.Errorf("не удалось распарсить JSON ответа: %w", err)
		return m.DSPbyHour{}, err
    }

	return stat, nil
}

// func DSPbyTraffic() {
// 	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/bydate/?token=%s&from=%s&to=%s", token, from, to)
// }

// func DSPtotal()  {
// 	url := fmt.Sprintf("https://backapi.admy.com/api/v4/api/dsp/report/bydate?token=%s&traffic_type=%s", date, token)
// }