package main

import (
	"anguo/infra/request/tushare"
	"anguo/scene"
	"flag"
	"fmt"
	"os"
)

var (
	Token  = ""
	Scene  = ""
	Output = ""
)

func main() {
	parseTokenFromCmdParams()
}

func parseTokenFromCmdParams() {
	flag.StringVar(&Token, "token", "", "tushare token ")
	if Token == "" {
		panic("can not run with no tushare token")
	}
	flag.StringVar(&Scene, "scene", "", "application scene ")
	if Token == "" {
		panic("can not run without a scene given")
	}
	flag.StringVar(&Output, "output", "", "file path to output result ")
	tushare.InitClient(Token)
	if Scene == "all" {
		scene.OutputFile = Output
		_, err := scene.CompareAllStockValueOfAssessmentWithPriceNow(1, 10000)
		if err != nil {
			fmt.Printf("error in all %v\n", err)
			os.Exit(1)
		}
	}
}
