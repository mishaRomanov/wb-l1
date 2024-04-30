// 25. Реализовать собственную функцию sleep.
package main

import (
	"log"
	"time"
)

// для имитации функции сна мы используем конструкцию select case и тип Timer
// из пакета time, который по истечении времени посылает время в канал.
// как только мы получаем информацию из канала, мы возвращаем управление функции main
// если мы не получаем ничего из канала, то просто продолжаем ждать
func MySleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	for {
		select {
		case <-timer.C:
			return
		default:
			continue
		}
	}
}

func main() {
	log.Println("Doing some work...")
	log.Println("I decided to sleep for 5 seconds for some reason")
	// вызываем функцию сна
	MySleep(time.Second * 5)
	log.Println("I decided to stop sleeping")
}
