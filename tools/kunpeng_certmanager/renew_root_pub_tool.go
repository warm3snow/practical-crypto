/**
 * @Author: xueyanghan
 * @File: extract_pub.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/10/13 15:25
 */

package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	CERT_CONFIG_FILE_PATH     = "cert_config.h"
	CERT_CONFIG_REPLACE_BEGIN = 24
	CERT_CONFIG_REPLACE_END   = 60
)

var (
	rootCrt            = flag.String("rootCrt", "root.crt", "root crt file")
	certConfigFilePath = flag.String("certConfigFilePath", CERT_CONFIG_FILE_PATH, "cert config file path")
)

func main() {
	flag.Parse()
	cFormedPubKey, err := extractCFromedPubFromCrt()
	if err != nil {
		panic(err)
	}
	err = replacePub(cFormedPubKey)
	if err != nil {
		panic(err)
	}
}

func extractCFromedPubFromCrt() (string, error) {
	//cmd := exec.Command("openssl", "x509", "-in", *rootCrt, "-pubkey", "-noout", "-C")
	shell := "openssl"
	args := []string{
		"x509",
		"-in",
		*rootCrt,
		"-pubkey",
		"-noout",
		"-C",
	}

	// 执行命令并获取输出
	output, err := exec.Command(shell, args...).Output()
	if err != nil {
		return "", err
	}
	//log.Printf("output: %s", output)

	keyList := strings.Split(string(output), ";")
	strList := strings.Split(strings.Split(keyList[1], "{")[1], "}")
	//log.Printf("pub key: %s", strList[0])

	return strList[0], nil
}

func replacePub(newPub string) error {
	inputFile, err := os.OpenFile(*certConfigFilePath, os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, "open cert config file failed")
	}
	defer inputFile.Close()

	pubKeyLineList := strings.Split(newPub, "\n")
	pubKeyLineList = pubKeyLineList[1 : len(pubKeyLineList)-1]
	pubKeyLineNum := 0

	inputScanner := bufio.NewScanner(inputFile)
	lineNum := 0
	var buf bytes.Buffer

	for inputScanner.Scan() {
		lineNum++
		if lineNum >= CERT_CONFIG_REPLACE_BEGIN && lineNum <= CERT_CONFIG_REPLACE_END {
			buf.WriteString(pubKeyLineList[pubKeyLineNum] + "\n")
			pubKeyLineNum++
			if lineNum == CERT_CONFIG_REPLACE_END && pubKeyLineNum < len(pubKeyLineList) {
				for pubKeyLineNum < len(pubKeyLineList) {
					buf.WriteString(pubKeyLineList[pubKeyLineNum] + "\n")
					pubKeyLineNum++
				}
			}
		} else {
			buf.WriteString(inputScanner.Text() + "\n")
		}
	}

	return ioutil.WriteFile(*certConfigFilePath, buf.Bytes(), 0644)
}
