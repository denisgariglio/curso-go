package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		number++
// 		time.Sleep(300 * time.Millisecond)
// 		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
// 	})
// 	http.ListenAndServe(":3000", nil)
// }

// Lock com Mutex
// func main() {
// 	m := sync.Mutex{}
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		m.Lock()
// 		number++
// 		m.Unlock()
// 		time.Sleep(300 * time.Millisecond)
// 		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
// 	})
// 	http.ListenAndServe(":3000", nil)
// }

// para verificar se existe concorrencia go run -race main.go
// soma atomica
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
