// 7. Реализовать конкурентную запись данных в map.
package main

import (
	"sync"
)

// для конкурентной записи в мапу мы можем использовать примитив синхронизации, который называется mutex

// делаем обертку - структуру, внутри которой сама мапа и мьютекс
// с помощью мьютекса мы блокируем доступ к критической секции на время записи
// критическая секция это участок исполнения кода, в котором существует доступ к общим ресурсам
// в данном случае - мапа
type customMap struct {
	storage map[int]string
	mu      *sync.Mutex
}

func (m *customMap) Add(key int, value string) {
	// блокируем доступ к критической секции
	m.mu.Lock()
	// деферим разблокирование мьютекса
	defer m.mu.Unlock()
	// записываем значение
	m.storage[key] = value
}
