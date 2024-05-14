package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sys/windows/svc/eventlog"
)

func main() {
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

	fmt.Println("Starting to write security events to log. Press Ctrl+C to stop.")

	for {
		// Выбираем случайное событие безопасности
		randomEvent := securityEvents[rand.Intn(len(securityEvents))]

		// Записываем событие в журнал событий Windows
		err := elog.Info(7777, randomEvent)
		if err != nil {
			fmt.Printf("Failed to write to event log: %sn", err)
			return // Или используйте continue, чтобы пытаться написать снова после ошибки
		}

		time.Sleep(1 * time.Second) // Ожидание 1 секунды перед записью следующего события
	}
}
