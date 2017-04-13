package main

import (
	"fmt"

	"os"

	"bitbucket.org/mundipagg/boletoapi/api"
	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
)

func main() {
	logo()
	installLog()
	bank.InstallBanks()
	api.InstallRestAPI()
}

func logo() {
	l := `
  ____        _      _                      _ 
 |  _ \      | |    | |         /\         (_)
 | |_) | ___ | | ___| |_ ___   /  \   _ __  _ 
 |  _ < / _ \| |/ _ \ __/ _ \ / /\ \ | '_ \| |
 | |_) | (_) | |  __/ || (_) / ____ \| |_) | |
 |____/ \___/|_|\___|\__\___/_/    \_\ .__/|_|
                                     | |      
                                     |_|      
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
