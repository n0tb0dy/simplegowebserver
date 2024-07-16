package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Проверяем наличие аргументов командной строки
	if len(os.Args) < 3 {
		log.Fatal("Usage: simplegowebserver <port> <response message>")
	}

	// Получаем порт и сообщение из аргументов командной строки
	port := os.Args[1]
	responseMessage := os.Args[2]

	// Обработчик запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Логируем информацию о запросе
		fmt.Printf("Received request from IP: %s\n", r.RemoteAddr)
		fmt.Println("Request headers:")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Отправляем ответ
		fmt.Fprintln(w, responseMessage)
	})

	// Выводим информацию о запущенном сервере
	fmt.Printf("Server is listening on port %s\n", port)

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(":"+port, nil))
}