package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Task17.3")
	fmt.Println("http://localhost:8080/hello")

	logFile, err := os.OpenFile("serverlogs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	l := log.New(logFile, "", log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello(l))

	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(l, logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}
func hello(l *log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		msg := "Hello, Go!"
		l.Println("resp:", msg)
		fmt.Fprint(res, msg)
	}
}
func authHandler(l *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			l.Println("Alarm! Unauthorised user")
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
