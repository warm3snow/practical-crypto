/**
 * @Author: xueyanghan
 * @File: tlsclient.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/5 11:34
 */

package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	Dial("1.2", "https://localhost:8443")
}

func Dial(tlsVersion string, addr string) {
	versionTLS := GetTLSVersion(tlsVersion)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         versionTLS,
			MaxVersion:         versionTLS,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("resp:", string(msg))
}

// GetTLSVersion get tls version
func GetTLSVersion(tlsVersion string) uint16 {
	switch tlsVersion {
	case "1.0":
		return tls.VersionTLS10
	case "1.1":
		return tls.VersionTLS11
	case "1.2":
		return tls.VersionTLS12
	case "1.3":
		return tls.VersionTLS13
	}
	return tls.VersionTLS12
}
