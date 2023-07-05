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
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	tlsVersion = flag.String("tlsVersion", "1.2", "specified tls version")
	addr       = flag.String("addr", "https://localhost:8443", "addr")
	skipVerify = flag.Bool("skipVerify", true, "skipVerify")
)

func main() {
	flag.Parse()

	Dial(*tlsVersion, *addr, *skipVerify)
}

func Dial(tlsVersion string, addr string, skipVerify bool) {
	versionTLS := GetTLSVersion(tlsVersion)
	client := http.DefaultClient
	if versionTLS != 0 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipVerify,
				MinVersion:         versionTLS,
				MaxVersion:         versionTLS,
			},
		}
		client = &http.Client{Transport: tr}
	}
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
	return 0
}
