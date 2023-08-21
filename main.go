// server.go

package main

import (
	"fmt"
	"io"
	"net/http"
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

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// read the data from the request
		d, err := io.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops!"))

			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
		// rw.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
