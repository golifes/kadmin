package main

import (
	"fmt"
	"github.com/golang/text/encoding/simplifiedchinese"
	"github.com/golang/text/transform"

	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error : status code", resp.StatusCode)
		return
	}

	reader := transform.NewReader(resp.Body, simplifiedchinese.GBK)
	all, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)

}
