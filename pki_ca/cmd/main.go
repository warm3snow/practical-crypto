package main

import (
	"flag"
	"fmt"
	"github.com/warm3snow/practical-crypto/crypto"
	"github.com/warm3snow/practical-crypto/pki_ca/config"
	"github.com/warm3snow/practical-crypto/pki_ca/logger"
	"github.com/warm3snow/practical-crypto/pki_ca/storedb"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var configPath = flag.String("config", "./config.yaml", "config file path")
var printVersion = flag.Bool("version", false, "print version")

func main() {
	flag.Parse()

	if *printVersion {
		PrintVersion()
		return
	}

	// init config
	if err := config.InitConfig(*configPath); err != nil {
		log.Panicf("failed to initialize config: %v", err)
	}

	// init logger
	logger.Init(&config.Conf.Log)
	//logger.Infof("parse config success! config: %+v\n", config.Conf)
	logger.Info(logo())

	// init crypto
	var err error
	crypto.Csp, err = crypto.InitCrypto(&config.Conf.Crypto)
	if err != nil {
		log.Panicf("failed to initialize csp module: %v", err)
	}
	logger.Infof("init crypto success!")

	//init db
	storedb.DBAccess, err = storedb.InitDBAccess(&config.Conf.DB)
	if err != nil {
		log.Panicf("failed to initialize db module: %v", err)
	}
	logger.Infof("init db success!")

	// init service
	// service.ChainSrv = service.NewChainService()
	// service.ChainSrv.Start()

	// init github.com/warm3snow/crypto-service-backend/ service
	go func() {
		// start server
		r := gin.Default()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())

		//register handlers
		ctrl.InitRouter(r)

		if err := r.Run(":8888"); err != nil {
			panic(err)
		}
	}()
	logger.Info("Startup server success!")

	// wait for signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTSTP)
	select {
	case sig := <-sigs:
		logger.Infof("recv sigs: %d(%s)", sig, sig)
		// service.ChainSrv.Stop()
		logger.Info("Shutdown server success!")
		fmt.Fprintln(os.Stdout, "Shutdown server success!")
		os.Exit(0)
	}
}

func logo() string {
	fig := figure.NewFigure("crypto-service-backend", "slant", true)
	s := fig.String()
	fragment := "================================================================================="
	versionInfo := fmt.Sprintf("Version: %2s%s\n", "", config.CurrentVersion)
	versionInfo += fmt.Sprintf("Build Time:%6s%s\n", " ", config.BuildDateTime)
	versionInfo += fmt.Sprintf("Git Branch:%6s%s\n", " ", config.GitBranch)
	versionInfo += fmt.Sprintf("Git Commit:%6s%s", " ", config.GitCommit)

	return fmt.Sprintf("\n%s\n%s%s\n%s\n", fragment, s, fragment, versionInfo)
}

func PrintVersion() {
	fmt.Println(logo())
	fmt.Println()
}
