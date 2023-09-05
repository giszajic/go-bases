package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"log"
)

func main() {

	// Req 1: Calcular cant personas viajan a un país determinado
	destination := "China"
	totalTickets, err := tickets.GetTotalTickets(destination)
	if err != nil {
		log.Fatalf("Error getting total tickets to %s: %v", destination, err)
	}
	fmt.Printf("Total tickets to %s: %d\n", destination, totalTickets)

	// Req 2: Calcular cant personas viajan en diferentes períodos del día
	period := "mañana"
	countByPeriod, err := tickets.GetCountByPeriod(period)
	if err != nil {
		log.Fatalf("Error getting count by %s: %v", period, err)
	}
	fmt.Printf("Count of passengers in the %s: %d\n", period, countByPeriod)

	// Req 3: Calcular promedio de personas que viajan a un país determinado en un día
	destination2 := "Czech Republic"
	total := 0 // Reemplaza con el total real de tickets
	average, err := tickets.AverageDestination(destination2, total)
	if err != nil {
		log.Fatalf("Error calculating average for %s: %v", destination2, err)
	}
	fmt.Printf("Average passengers to %s: %.2f\n", destination2, average)
}
