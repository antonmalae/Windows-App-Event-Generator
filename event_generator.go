package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"golang.org/x/sys/windows/svc/eventlog"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: goappsecurityevent <eps> <eventID>")
		return
	}

	eps, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid EPS value: %sn", os.Args[1])
		return
	}

	eventID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid Event ID value: %sn", os.Args[2])
		return
	}

	// События безопасности для генерации
	securityEvents := []string{
		"Unauthorized access attempt detected",
		"User authentication success",
		"User authentication failure",
		"Unexpected file modification detected",
		"Firewall rule change detected",
	}

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	const source = "GoAppSecurityEvent"

	elog, err := eventlog.Open(source)
	if err != nil {
		fmt.Printf("Failed to open event log: %sn", err)
		return
	}
	defer elog.Close()

	fmt.Printf("Starting to write security events at %v EPS with Event ID %d. Press Ctrl+C to stop.n", eps, eventID)

	for {
		// Выбираем случайное событие безопасности
		randomEvent := securityEvents[rand.Intn(len(securityEvents))]

		// Записываем событие в журнал событий Windows с произвольным EventID
		err := elog.Info(uint32(eventID), randomEvent)
		if err != nil {
			fmt.Printf("Failed to write to event log: %sn", err)
			return
		}

		// Рассчитываем задержку в зависимости от EPS
		time.Sleep(time.Duration(float64(time.Second) / eps))
	}
}
