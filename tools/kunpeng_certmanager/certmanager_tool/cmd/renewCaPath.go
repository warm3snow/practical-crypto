package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	taSecPath             *string
	caCertManagerFilePath *string
)

// renewCaPathCmd represents the renewCaPath command
var renewCaPathCmd = &cobra.Command{
	Use:   "renewCaPath",
	Short: "renew TA path in CA",
	Long:  `renew TA path in cert_manager CA`,
	Run: func(cmd *cobra.Command, args []string) {
		err := replaceTaSec()
		if err != nil {
			panic(err)
		}
		log.Printf("update cert_config.h TA path success")
	},
}

func init() {
	rootCmd.AddCommand(renewCaPathCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	taSecPath = renewCaPathCmd.PersistentFlags().StringP("taSecPath", "s",
		"f68fd704-6eb1-4d14-b218-722850eb3ef0.sec", "ta sec file path")
	caCertManagerFilePath = renewCaPathCmd.PersistentFlags().StringP("caCertManagerFilePath", "c",
		"cert_manager.c", "ca cert_manager file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renewRootPubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func replaceTaSec() error {
	inputFile, err := os.OpenFile(*caCertManagerFilePath, os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, "open cert config file failed")
	}
	defer inputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)
	lineNum := 0
	var buf bytes.Buffer

	for inputScanner.Scan() {
		lineNum++

		//26 #define TA_PATH "/usr/bin/4acaf7c8-c652-4643-9b7a-cc07e7a3187a.sec"
		if lineNum == 26 {
			buf.WriteString("#define TA_PATH \"" + *taSecPath + "\"")
		} else if lineNum >= 257 && lineNum <= 258 {
			//256 static TEEC_UUID g_taId = {
			//257     0x4acaf7c8, 0xc652, 0x4643,
			//258     { 0x9b, 0x7a, 0xcc, 0x07, 0xe7, 0xa3, 0x18, 0x7a }
			//259 };
			taSec := filepath.Base(*taSecPath)
			fmt.Println("taSec: ", taSec)
			if lineNum == 257 {
				buf.WriteString("    0x" + taSec[0:8] + ", 0x" + taSec[9:13] + ", 0x" + taSec[14:18] + ",\n")
			} else if lineNum == 258 {
				buf.WriteString("    { 0x" + taSec[19:21] + ", 0x" + taSec[21:23] + ", 0x" + taSec[24:26] + ", 0x" + taSec[26:28] +
					", 0x" + taSec[28:30] + ", 0x" + taSec[30:32] + ", 0x" + taSec[32:34] + ", 0x" + taSec[34:36] + " }\n")
			}
		} else {
			buf.WriteString(inputScanner.Text() + "\n")
		}
	}

	return ioutil.WriteFile(*caCertManagerFilePath, buf.Bytes(), 0644)
}
