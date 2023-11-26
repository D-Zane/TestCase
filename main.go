package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Структура для представления формата JSON-файла.
type Numbers struct {
	Values []int `json:"values"`
}

func main() {
	// Открываем файл с числами
	file, err := os.Open("numbers.json")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Декодируем JSON
	var numbers Numbers
	if err := json.Unmarshal(byteValue, &numbers); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	// Считаем сумму чисел
	sum := 0
	for _, num := range numbers.Values {
		sum += num
	}

	// Логируем сумму чисел
	log.Printf("Sum of numbers: %d", sum)

	// Выполняем HTTP GET запрос
	url := "https://example.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed. Status: %s", resp.Status)
	}

	// Логируем успешный запрос
	log.Printf("HTTP request to %s was successful. Status: %s", url, resp.Status)
}
