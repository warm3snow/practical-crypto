/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const (
	CERT_CONFIG_REPLACE_BEGIN = 24
	CERT_CONFIG_REPLACE_END   = 60
)

var (
	rootCrt            *string
	certConfigFilePath *string
)

// renewRootPubCmd represents the renewRootPub command
var renewRootPubCmd = &cobra.Command{
	Use:   "renewRootPub",
	Short: "update cert_config.h root_public_key from root.crt",
	Long:  "update cert_config.h root_public_key from root.crt",
	Run: func(cmd *cobra.Command, args []string) {
		cFormedPubKey, err := extractCFromedPubFromCrt()
		if err != nil {
			panic(err)
		}
		err = replacePub(cFormedPubKey)
		if err != nil {
			panic(err)
		}
		log.Printf("update cert_config.h root_public_key success")
	},
}

func init() {
	rootCmd.AddCommand(renewRootPubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCrt = renewRootPubCmd.PersistentFlags().StringP("rootCrt", "r",
		"root.crt", "root crt file")
	certConfigFilePath = renewRootPubCmd.PersistentFlags().StringP("certConfigFilePath", "c",
		"cert_config.h", "cert config file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renewRootPubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
