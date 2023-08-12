/**
 * @Author: xueyanghan
 * @File: tls_no_version.go
 * @Version: 1.0.0
 * @Description: if tls version is not specified, then it will be chosen by the tls handshake.
 * @Date: 2023/8/12 17:20
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
