package main

import "time"

func main() {
	ticker := time.NewTicker(time.Second * 5)
	stopChan := make(chan struct{})

	runApp(ticker, stopChan)
}

// printHello — функция для принтинга hello world
func printHello() {
	println("Hello github actions!")
}

// runApp — запускаем приложение
func runApp(timer *time.Ticker, stopChan chan struct{}) {
	for {
		select {
		// новый тик
		case <-timer.C:
			printHello()
		// сигнал о завершении тика
		case <-stopChan:
			return
		}
	}

}
