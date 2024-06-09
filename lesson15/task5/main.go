package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type MeteoType struct {
	temp string
	mu   sync.Mutex
}

func (m *MeteoType) ReadTemp() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return fmt.Sprintf("%s°C", m.temp)
}

func (m *MeteoType) ChangeTemp(v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.temp = v
}

func main() {
	temp := &MeteoType{temp: "20"}
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()
			temp.ChangeTemp(strconv.Itoa(rand.Intn(20) + 20))
		}()
		go func() {
			defer wg.Done()
			fmt.Println("Current temperature: ", temp.ReadTemp())
		}()
	}

	wg.Wait()
}
