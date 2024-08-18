package styles

import "github.com/charmbracelet/lipgloss"

// Define y retorna el estilo para el ticker
func GetTickerStyle() lipgloss.Style {
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
		Background(lipgloss.Color("#593dad")).
		Bold(true).
		Padding(0, 1)
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
