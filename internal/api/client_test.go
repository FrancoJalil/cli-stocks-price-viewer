package api

import "testing"

func TestGetStockPriceFromYahooFinanceAPI_Success(t *testing.T) {

	// llama a la función que hace la petición a la API
	aaplCurrentPrice, err := GetCurrentPrice("AAPL")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// si la respuesta no es un float64, el test fallará
	if _, ok := interface{}(aaplCurrentPrice).(float64); !ok {
		t.Errorf("Expected float64, got %T", aaplCurrentPrice)
	}
}
