package main

import (
	"ex-010-go-web-simple/functions"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/pprof"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", functions.BasicResponse).Methods("GET")
	r.HandleFunc("/books/", functions.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", functions.GetBook).Methods("GET")
	r.HandleFunc("/books/", functions.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", functions.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", functions.DeleteBook).Methods("DELETE")

	//Профилирование
	//1. Создать нагрузку:
	// > ab -n 1000000 -c10 http://localhost:8180/books/
	//2. Сбор данных в файл
	//2.1. Сохранение в файл:
	// > wget http://127.0.0.1:8180/debug/pprof/profile?seconds=5 -O profile
	//2.2. Просмотр файла-результата:
	// > go tool pprof profile
	//3. Прямой сбор данных
	// > go tool pprof ex-010-go-web-simple http://127.0.0.1:8180/debug/pprof/profile
	//4. Внутри средства вызвать генерацию диаграмм:
	// > > web
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	log.Fatal(http.ListenAndServe("localhost:8180", r))
}