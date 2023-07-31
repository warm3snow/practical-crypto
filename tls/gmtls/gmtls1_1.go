/**
 * @Author: xueyanghan
 * @File: gmtls1_1.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/29 13:05
 */

package main

import (
	tls "github.com/warm3snow/gmsm/gmtls"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received connection: %s", r.RemoteAddr)
	w.Write([]byte("Hello, world!"))
}

func main() {
	// 设置handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	// 加载服务端证书和私钥
	signKeyPair, err := tls.LoadX509KeyPair("../../openssl/gmcerts/server_sign.crt", "../../openssl/gmcerts/server_sign.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair: %v", err)
	}
	encKeyPair, err := tls.LoadX509KeyPair("../../openssl/gmcerts/server_enc.crt", "../../openssl/gmcerts/server_enc.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair: %v", err)
	}
	// 设置tls配置项
	tlsCfg := &tls.Config{
		GMSupport:    tls.NewGMSupport(),
		Certificates: []tls.Certificate{signKeyPair, encKeyPair},
		MaxVersion:   tls.VersionTLS12,
	}
	ln, err := tls.Listen("tcp", ":8443", tlsCfg)
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}

	err = http.Serve(ln, mux)
	if err != nil {
		log.Fatalf("Serve: %v", err)
	}
}
