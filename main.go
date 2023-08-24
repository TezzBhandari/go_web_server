package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TezzBhandari/go_web_server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// 	listener, err := net.Listen("tcp", "localhost:8080")
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	defer listener.Close()

	// 	fmt.Println("Server listening on localhost:8080")

	// 	for {
	// 		conn, err := listener.Accept()
	// 		if err != nil {
	// 			fmt.Println("Error accepting connection:", err)
	// 			continue
	// 		}
	// 		go handleConnection(conn)
	// 	}
	// }

	// func handleConnection(conn net.Conn) {
	// 	defer conn.Close()

	// 	buffer := make([]byte, 1024)
	// 	n, err := conn.Read(buffer)
	// 	if err != nil {
	// 		fmt.Println("Error reading:", err)
	// 		return
	// 	}

	// 	message := string(buffer[:n])
	// 	fmt.Println("Received:", message)

	// 	_, err = conn.Write([]byte("Server received: " + message))
	// 	if err != nil {
	// 		fmt.Println("Error writing:", err)
	// 	}
	// }

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	// read the data from the request
	// 	d, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		// rw.WriteHeader(http.StatusBadRequest)
	// 		// rw.Write([]byte("Oops!"))

	// 		http.Error(rw, "Oops", http.StatusBadRequest)
	// 		return
	// 	}
	// 	fmt.Printf("%s\n", d)

	// 	fmt.Fprintf(rw, "Hello %s", d)
	// 	// rw.Write([]byte("hello world"))
	// })

	// http.ListenAndServe(":8080", nil)
	log := log.New(os.Stdout, "Product-Api ", log.LstdFlags)

	// hello_handler := handlers.NewHello(log)
	product_handler := handlers.NewProduct(log)

	// sm := http.NewServeMux()
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	// sm.Handle("/", hello_handler)
	getRouter.HandleFunc("/products", product_handler.GetProducts)
	// sm.Handle("/products", product_handler)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", product_handler.UpdateProducts)
	putRouter.Use(product_handler.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", product_handler.AddProduct)
	postRouter.Use(product_handler.MiddlewareProductValidation)

	// converting the function into handler
	time_handler := http.HandlerFunc(CustomHandler)

	// register time handler in serveMux
	// sm.Handle("/time", Middleware(time_handler))

	sm.Handle("/time", time_handler)

	// using middleware using gorilla mux router
	sm.Use(Middleware)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// catches all the signals for shutting down server
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	// blocking code. waits for signal
	sig := <-sigChan

	log.Println("Received Terminate, Graceful Shutdown ", sig)

	// this is required to forceful shutdown if the server doesn't shutdown after 30 seconds
	timeout_ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// gracefully shutdowns the server
	server.Shutdown(timeout_ctx)

}

// creating a function for handling request. later this function will be converted in handler
func CustomHandler(rw http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	rw.Write([]byte("The time is: " + tm))

}

func Middleware(next http.Handler) http.Handler {
	time_handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("fuck", "you")
		next.ServeHTTP(rw, r)
	})
	return time_handler
}
