// 5. Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
package main

import (
	"fmt"
	"sync"
	"time"
)

// функция при помощи range читает значения из канала, пока они поступают. как только канал закрывается, range прекращает работу
func ReadValuesFromChannel(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d from channel!\n", x)
	}
}

func main() {
	// делаем вейтгруппу чтобы подождать, пока наша горутина завершит работу (чтобы main не завершился раньше горутины)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// делаем канал, в который посылаем значения, которые будет читать другая горутина
	ch := make(chan int)
	defer close(ch)

	// делаем таймер. по истечении таймера (N секунд по условиям задачи) мы получаем сигнал о закрытии канала, тем самым понимая, что таймер истек
	timer := time.NewTimer(time.Second * 5)
	// запускаем функцию чтения канала в горутине
	go ReadValuesFromChannel(ch)
	// делаем переменную-счетчик
	var x int
	// запускаем горутину, в которой отправляем значения в канал и проверяем, не истек ли таймер
	go func() {
		for {
			select {
			// если таймер истек, мы останавливаем выполнение программы
			case <-timer.C:
				fmt.Println("Timer stopped. Exiting...")
				// декрементим счетчик
				wg.Done()
				return
				// если не истек, посылаем значения в канал
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
				// инкрементим счетчик
				x++
			}
		}
	}()
	// ждем, пока счетчик вейтгруппы станет 0
	wg.Wait()
}