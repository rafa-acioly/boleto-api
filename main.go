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
	defer log.Close()
	logo1()
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
