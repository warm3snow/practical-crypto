/**
 * @Author: xueyanghan
 * @File: notls.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/5 14:46
 */

package main

import (
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received connection: %s", r.RemoteAddr)
	w.Write([]byte("Hello, world!"))
}

func main() {
	// 设置&启动tls服务
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	srv := &http.Server{
		Addr:    ":8443",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServeTLS: %v", err)
	}
}
