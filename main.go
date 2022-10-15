package main

import (
	"flag"
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
}
