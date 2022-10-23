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
	doInitUseParamsInput()
	choiceAndRunAScene()
}

func doInitUseParamsInput() {
	tushare.InitClient(Token)
}

func choiceAndRunAScene() {
	if Scene == "all" {
		scene.OutputFile = Output
		_, err := scene.CompareAllStockValueOfAssessmentWithPriceNow(1, 500)
		if err != nil {
			fmt.Printf("error in all %v\n", err)
			os.Exit(1)
		}
	}
}

func parseTokenFromCmdParams() {
	flag.StringVar(&Token, "token", "", "tushare token ")
	flag.StringVar(&Scene, "scene", "", "application scene ")
	flag.StringVar(&Output, "output", "", "file path to output result ")
	flag.Parse()
	if Token == "" {
		panic("can not run with no tushare token")
	}
	if Scene == "" {
		panic("can not run without a scene given")
	}
}
