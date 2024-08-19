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

	// Define el estilo para los errores
	errorStyle := styles.GetErrorStyle()

	// Obtiene el precio actual y el porcentaje de cambio del ticket
	price, percentChange, err := api.GetCurrentPrice(ticket)
	if err != nil {
		fmt.Println(errorStyle.Render(
			"Error:", err.Error(),
		))
		return
	}

	// Define los estilos para el ticket, el precio y el porcentaje
	ticketStyle := styles.GetTicketStyle()
	priceStyle := styles.GetPriceStyle()
	percentStyle := styles.GetPercentStyle(percentChange)

	// Renderiza el ticket estilizado
	fmt.Print(ticketStyle.Render(fmt.Sprintf("%s ", strings.ToUpper(ticket))))

	// Renderiza el precio con un label y su estilo
	priceLabel := "USD $"
	if strings.HasSuffix(ticket, ".BA") {
		priceLabel = "ARS $"
	}
	priceStr := fmt.Sprintf("%s%.2f", priceLabel, price)
	fmt.Print(priceStyle.Render(priceStr))

	// Renderiza el porcentaje de cambio con su estilo
	percentStr := fmt.Sprintf("%.2f%%", percentChange)
	fmt.Println(percentStyle.Render(percentStr))
}
