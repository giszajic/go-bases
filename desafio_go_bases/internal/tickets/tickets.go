package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          string
	Name        string
	Email       string
	Destination string
	Time        string
	Price       string
}

// Func aux para abrir y cerrar el archivo CSV
func openAndCloseCSVFile() (*os.File, *csv.Reader, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(file)
	return file, reader, nil
}

func GetTotalTickets(destination string) (int, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	//contador para el pais de destino
	count := 0

	//itera sobre las lineas del archivo CSV
	for {
		//lee una linea del archivo
		record, err := reader.Read()
		if err != nil {
			break
		}

		// compara el país de destino en la línea actual con el destino dado
		if record[3] == destination {
			count++
		}
	}

	return count, nil
}

// Func aux para verificar si una hora está en un período dado
func isTimeInPeriod(time string, period string) bool {
	hour, _ := strconv.Atoi(strings.Split(time, ":")[0])

	switch period {
	case "madrugada":
		return hour >= 0 && hour < 6
	case "mañana":
		return hour >= 7 && hour < 12
	case "tarde":
		return hour >= 13 && hour < 20
	case "noche":
		return hour >= 20 && hour <= 23
	default:
		return false
	}
}

func GetCountByPeriod(period string) (int, error) {

	// Abre el archivo CSV
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Crea un lector CSV
	reader := csv.NewReader(file)

	// Inicializa un contador para el período dado
	count := 0

	// Itera sobre las líneas del archivo CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		time := record[4]

		if isTimeInPeriod(time, period) {
			count++
		}
	}
	return count, nil
}

// Func aux para convertir un precio str a int
func parsePrice(price string) int {
	parsedPrice, _ := strconv.Atoi(price)
	return parsedPrice
}

func AverageDestination(destination string, total int) (float64, error) {

	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0.0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Inicializa un contador para el país de destino
	count := 0
	totalPassengers := 0

	// Itera sobre las líneas del archivo CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Compara el país de destino en la línea actual con el destino dado
		if strings.TrimSpace(record[3]) == destination {
			totalPassengers += parsePrice(record[5]) // Agrega el precio al total
			count++
		}
	}

	// Calcula el promedio si hay al menos una entrada para el destino
	if count > 0 {
		average := float64(totalPassengers) / float64(count)
		return average, nil
	}

	return 0.0, nil // Retorna 0.0 si no hay datos para el destino
}
