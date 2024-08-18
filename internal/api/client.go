package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseURL = "https://query1.finance.yahoo.com/v8/finance/chart/"

// Estructuras para deserializar el JSON
type ChartMeta struct {
	RegularMarketPrice float64 `json:"regularMarketPrice"`
}

type Result struct {
	Meta ChartMeta `json:"meta"`
}

type Chart struct {
	Result []Result `json:"result"`
}

type Data struct {
	Chart Chart `json:"chart"`
}

func GetCurrentPrice(ticket string) (float64, error) {
	// ticker uppercase y construye la URL
	ticket = strings.ToUpper(ticket)
	url := baseURL + ticket + "?events=capitalGain%7Cdiv%7Csplit&formatted=true&includeAdjustedClose=true&interval=1d&symbol=" + ticket + "&userYfid=true&range=1d&lang=en-US&region=US"

	// Crea una nueva solicitud con un User-Agent personalizado
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")

	// Envía la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Lee el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Deserializa los datos JSON
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	// Retorna el precio actual
	if len(data.Chart.Result) > 0 {
		return data.Chart.Result[0].Meta.RegularMarketPrice, nil
	}

	return 0, fmt.Errorf("no se encontró el ticket")
}
