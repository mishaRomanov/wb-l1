// 21. Реализовать паттерн «адаптер» на любом примере.
package main

import (
	"database/sql"
	"log"
	"sync"
)

// Паттерн проектирования "адаптер" нужен для того, чтобы использовать функционал
// стороннего класса, интерфейс которого не совместим с используемым нами интерфейсом

// создадим интерфейс, под который надо адаптировать сторонний класс (структуру)
type MainInterface interface {
	Write(id int, data string) error
}

// пример класса, который совместим с нашим интерфейсом
type InMemory struct {
	mu     *sync.Mutex
	memory map[int]string
}

// метод класса
func (i *InMemory) Write(id int, data string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.memory[id] = data
	return nil
}

// функция, работающая с нашим интерфейсом
func WriteDataToStorage(storage MainInterface, id int, data string) error {
	err := storage.Write(id, data)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	return nil
}

// тут пишем пример несовместимого с нашим интерфейсом класса
type Postgres struct {
	pg *sql.DB
}

// тот самый несовместимый метод
func (p *Postgres) InsertValues(id int, data string) error {
	_, err := p.pg.Exec(`INSERT INTO table (id,data) VALUES ($1,$2)`, id, data)
	if err != nil {
		return err
	}
	log.Println("Successfull insert.")
	return nil
}

// пишем адаптер для несовместимого класса
type Adapter struct {
	// встраиваем постгрес внутрь адаптера
	Postgres
}

// метод-обертка
func (a *Adapter) Write(id int, data string) error {
	err := a.InsertValues(id, data)
	if err != nil {
		return err
	}
	return nil
}
