package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AddContextMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add context to incoming request
		ctx := context.Background()
		uuidObj := uuid.New()
		ctx = context.WithValue(ctx, "RequestID", uuidObj.String())
		next(w, r.WithContext(ctx))
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("SayHello > RequestID is %v and Method is %v and Path is %v and ", r.Context().Value("RequestID"), r.Method, r.URL.Path)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func main() {
	log.Println("Basic Middleware....!")

	r := mux.NewRouter()
	r.HandleFunc("/status", AddContextMiddleWare(SayHello)).Methods("GET")
	srv := http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	err := srv.ListenAndServe()
	log.Println("Started server at port", srv.Addr)
	if err != nil {
		log.Fatalln("Server Error")
	}
}
