package styles

import "github.com/charmbracelet/lipgloss"

// Define y retorna el estilo para el ticker
func GetTicketStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Bold(true).
		Padding(0, 1)
}

// Define y retorna el estilo para el precio
func GetPriceStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#593DAD")).
		Bold(true).
		Padding(0, 1)
}

// Define y retorna el estilo para el porcentaje
func GetPercentStyle(percentChange float64) lipgloss.Style {
	style := lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Padding(0, 1).
		Foreground(lipgloss.Color("#FAFAFA"))

	switch {
	case percentChange > 0:
		style = style.Background(lipgloss.Color("#48732f")) // Verde para positivo
	case percentChange < 0:
		style = style.Background(lipgloss.Color("#c43939")) // Rojo para negativo
	default:
		style = style.Background(lipgloss.Color("#6e6e6e")) // Gris para cero
	}

	return style
}

// Define y retorna el estilo para los errores
func GetErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#AD3D3D")).
		Bold(true).
		Padding(0, 1)
}

// Define y retorna el estilo para el usage message
func GetUsageMessageStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("#3C3C3C")).
		Bold(true).
		Italic(true).
		Padding(0, 1)
}
