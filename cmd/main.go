package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FrancoJalil/cli-stocks-price-viewer/internal/api"
	"github.com/FrancoJalil/cli-stocks-price-viewer/internal/styles"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(styles.GetUsageMessageStyle().Render(
			"Usage: price <TICKER> [x]",
		))
		return
	}

	ticket := os.Args[1]

	// Verifica si se debe agregar ".BA"
	if len(os.Args) <= 2 || os.Args[2] != "x" {
		ticket = ticket + ".BA"
	}

	// Define y aplica estilos
	style := styles.GetTickerStyle()
	secondStyle := styles.GetPriceStyle()
	errorStyle := styles.GetErrorStyle()

	// Obtiene el precio actual del ticker

	price, err := api.GetCurrentPrice(ticket)
	if err != nil {
		fmt.Println(errorStyle.Render(
			"Error:", err.Error(),
		))
		return
	}

	// Renderiza el ticker estilizado
	fmt.Print(style.Render(fmt.Sprintf("%s ", strings.ToUpper(ticket))))

	// Renderiza el precio con el segundo estilo
	priceLabel := "USD $"
	if strings.HasSuffix(ticket, ".BA") {
		priceLabel = "ARS $"
	}

	priceStr := fmt.Sprintf("%s%.2f", priceLabel, price)
	fmt.Println(secondStyle.Render(priceStr))
}
