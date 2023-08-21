// server.go

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TezzBhandari/go_web_server/handlers"
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

	hello_handler := handlers.NewHello(log)

	sm := http.NewServeMux()

	sm.Handle("/", hello_handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
