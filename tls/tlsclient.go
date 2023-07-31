/**
 * @Author: xueyanghan
 * @File: tlsclient.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/5 11:34
 */

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/warm3snow/gmsm/gmtls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
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
		if versionTLS != gmtls.VersionGMSSL {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: skipVerify,
					MinVersion:         versionTLS,
					MaxVersion:         versionTLS,
				},
			}
			client = &http.Client{Transport: tr}
		} else {
			config := &gmtls.Config{
				InsecureSkipVerify: skipVerify,
				GMSupport:          gmtls.NewGMSupport(),
			}
			client = &http.Client{
				Transport: &http.Transport{
					DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
						dialer := &net.Dialer{}
						conn, err := gmtls.DialWithDialer(dialer, network, addr, config)
						if err != nil {
							return nil, err
						}

						return conn, nil
					},
					Dial: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).Dial,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          100,
					TLSHandshakeTimeout:   10 * time.Second,
					IdleConnTimeout:       90 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				},
			}
		}
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
	case "gmtls1.1":
		return gmtls.VersionGMSSL
	}
	return 0
}
