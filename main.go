package main

import (
	"fmt"
	"os"
	"strconv"
)

var version string = "v0.0.0"

func getArg(index int) string {
	if index >= len(os.Args) {
		return ""
	}
	return os.Args[index]
}

func cliFloat(index int, helpMsg string) float64 {
	str := getArg(index)
	if str == "" {
		fmt.Fprintln(os.Stderr, helpMsg)
		os.Exit(1)
	}
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, helpMsg + "\n" + err.Error())
		os.Exit(1)
	}
	return v
}

func printHelp(arg string, helpMessage string) {
	switch arg {
	case "help", "-h", "--help":
		fmt.Fprintln(os.Stderr, helpMessage)
		os.Exit(0)
	}
}

func satToGold(btcUsdPrice, goldOzUsdPrice, satAmount float64) float64 {
	satOz := btcUsdPrice / goldOzUsdPrice / 100_000_000
	return satAmount * satOz
}

func goldToSat(btcUsdPrice, goldOzUsdPrice, ozAmount float64) float64 {
	satOz := btcUsdPrice / goldOzUsdPrice / 100_000_000
	return ozAmount / satOz
}

func cli() error {
	var (
		helpMsg = "usage: <s-to-g|g-to-s>"
		arg = getArg(1)
	)
	switch arg {
	case "s-to-g":
		var (
			helpMsg = "usage: s-to-g <BTCUSD> <XAUUSD> <satoshiAmount>"
			btcusd, gldusd, satAmount float64
		)
		printHelp(getArg(2), helpMsg)
		btcusd = cliFloat(2, helpMsg)
		gldusd = cliFloat(3, helpMsg)
		satAmount = cliFloat(4, helpMsg)
		fmt.Printf("%.4f gold oz\n", satToGold(btcusd, gldusd, satAmount))
	case "g-to-s":
		var (
			helpMsg = "usage: g-to-s <BTCUSD> <XAUUSD> <goldOzAmount>"
			btcusd, gldusd, ozAmount float64
		)
		printHelp(getArg(2), helpMsg)
		btcusd = cliFloat(2, helpMsg)
		gldusd = cliFloat(3, helpMsg)
		ozAmount = cliFloat(4, helpMsg)
		fmt.Printf("%8.0f sat\n", goldToSat(btcusd, gldusd, ozAmount))
	case "-v", "--version", "version":
		fmt.Println(version)
	default:
		fmt.Fprintln(os.Stderr, helpMsg)
		os.Exit(1)
	}
	return nil
}

func main(){
	err := cli()
	if err != nil { 
		panic(err)
	}
}
