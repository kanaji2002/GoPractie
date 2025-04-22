package main

import (
	"fmt"
	"log"
	"net/http"
)

// ハンドラー関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Hello, Go API!")
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/hello", helloHandler) // URLパスとハンドラーを関連づける

	fmt.Println("🚀 サーバー起動中: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
