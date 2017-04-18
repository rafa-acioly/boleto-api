package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"

	"os"

	"bitbucket.org/mundipagg/boletoapi/api"
	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
)

var (
	processID  = os.Getpid()
	totalProcs = runtime.NumCPU()
	devMode    = flag.Bool("dev", true, "-dev To run in dev mode")
)

func init() {
	createPIDfile()
	// Configure signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go handleSignal(c)
}

func handleSignal(c chan os.Signal) {
	<-c
	log.Info("Quiting BoletoApi")
	os.Exit(1)
}

func createPIDfile() {
	p := strconv.Itoa(processID)
	ioutil.WriteFile("boletoapi.pid", []byte(p), 0644)
}
func configFlags() {
	if *devMode {
		fmt.Println("Execute in DEV mode")
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://localhost:5341/")
		os.Setenv("SEQ_API_KEY", "") //4jZzTybZ9bUHtJiPdh6
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
	}
}
func main() {
	flag.Parse()
	configFlags()
	defer log.Close()
	logo1()
	installLog()
	bank.InstallBanks()
	api.InstallRestAPI()
}

func logo1() {
	l := `
$$$$$$$\            $$\            $$\                $$$$$$\            $$\ 
$$  __$$\           $$ |           $$ |              $$  __$$\           \__|
$$ |  $$ | $$$$$$\  $$ | $$$$$$\ $$$$$$\    $$$$$$\  $$ /  $$ | $$$$$$\  $$\ 
$$$$$$$\ |$$  __$$\ $$ |$$  __$$\\_$$  _|  $$  __$$\ $$$$$$$$ |$$  __$$\ $$ |
$$  __$$\ $$ /  $$ |$$ |$$$$$$$$ | $$ |    $$ /  $$ |$$  __$$ |$$ /  $$ |$$ |
$$ |  $$ |$$ |  $$ |$$ |$$   ____| $$ |$$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |$$ |
$$$$$$$  |\$$$$$$  |$$ |\$$$$$$$\  \$$$$  |\$$$$$$  |$$ |  $$ |$$$$$$$  |$$ |
\_______/  \______/ \__| \_______|  \____/  \______/ \__|  \__|$$  ____/ \__|
                                                               $$ |          
                                                               $$ |          
                                                               \__|          
	`
	fmt.Println(l)
	fmt.Println("Version: " + config.Get().Version)
}

func installLog() {
	err := log.Install()
	if err != nil {
		fmt.Println("Log SEQ Fails")
		os.Exit(-1)
	}
}
