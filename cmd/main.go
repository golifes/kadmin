package main

import (
	"flag"
	"fmt"
	"kadmin/pkg/config"
	"os"
)

var dir string

func init() {

	flag.StringVar(&dir, "c", "config/db.json", "config dir")
}

func main() {
	flag.Parse()
	usage()

	if dir == "" {
		panic("请指定默认配置文件")
	}

	err := config.InitConfig(dir)
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `main version:1.0.0
Usage: bash main [-c filename]

Options: 
`)
	flag.PrintDefaults()
}

//once   sync.Once
