// 4. Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из
// канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.
package main

//чтобы выбрать количество воркеров,
//ЗАПУСТИТЕ ПРОГРАММУ С ФЛАГОМ -w *число*
//где число это количество воркеров

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

// функция воркера
func worker(num int, stream chan int) {
	//используем тут оператор range потому что при закрытии канала он автоматически завершает свою работу
	for v := range stream {
		fmt.Printf("Воркер %d считал со стрима: %d\n", num, v)
	}
}

func main() {
	// парсим нужное кол-во воркеров через флаг
	numOfWorkers := flag.Int("w", 1, "a number of workers")
	flag.Parse()
	// логгируем
	fmt.Printf("Запускаем воркеры: %d штук(и)...\n", *numOfWorkers)

	// подписываемся на уведомление о сигналах
	ctx, closefunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer closefunc()
	// делаем канал для бесконечной записи
	stream_main := make(chan int)

	// запускаем нужное количество воркеров
	for w := 1; w <= *numOfWorkers; w++ {
		go worker(w, stream_main)
	}
	// счетчик для отправки значений в канал
	var counter int
	//используем классический паттерн for select case
	for {
		select {
		case <-ctx.Done():
			//просто для красоты чтобы в терминале выглядело нормально
			fmt.Printf("\n")
			fmt.Println("Получили сигнал. Закрываемся..")
			//имитируем graceful shutdown
			time.Sleep(1 * time.Second)
			//закрываем канал. оператор range в воркерках сигнализирует о том, что каналы закрыты, и перестает выполняться
			//после этого воркеры закрываются
			close(stream_main)
			return
		default:
			stream_main <- counter
			//инкрементим счетчик
			counter++
			//немного ждем чтобы не было мешанины в терминале
			time.Sleep(500 * time.Millisecond)
		}
	}

}
