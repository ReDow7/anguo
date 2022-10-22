package main

import (
	"anguo/infra/request/tushare"
	"anguo/scene"
	"flag"
	"fmt"
	"os"
)

var (
	Token = ""
	Scene = ""
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
	tushare.InitClient(Token)
	if Scene == "all" {
		_, err := scene.CompareAllStockValueOfAssessmentWithPriceNow(1, 10000)
		if err != nil {
			fmt.Printf("error in all %v\n", err)
			os.Exit(1)
		}
	}
}
