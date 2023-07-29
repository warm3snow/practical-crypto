/**
 * @Author: xueyanghan
 * @File: tls1.3.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/5 12:09
 */

package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received connection: %s", r.RemoteAddr)
	w.Write([]byte("Hello, world!"))
}

func main() {
	// 加载服务端证书和私钥
	keyPair, err := tls.LoadX509KeyPair("../../openssl/tlscerts/server.crt", "../../openssl/tlscerts/server.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair: %v", err)
	}

	// 设置tls配置项
	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{keyPair},
		MaxVersion:   tls.VersionTLS13,
		MinVersion:   tls.VersionTLS13,
	}

	// 设置&启动tls服务
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	srv := &http.Server{
		TLSConfig: tlsCfg,
		Addr:      ":8443",
		Handler:   mux,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ListenAndServeTLS: %v", err)
	}
}
