package main

import (
	"fmt"
	"os"

	"github.com/google/go-querystring/query"
	"github.com/srcclr/example-go-modules/sub"
	"github.com/srcclr/example-go-modules/sub2"
)

type Options struct {
	Query   string `url:"q"`
	ShowAll bool   `url:"all"`
	Page    int    `url:"page"`
}

func main() {
	fmt.Println("Intro to Go!")
	opt := Options{"foo", true, 2}
	v, _ := query.Values(opt)
	fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"
	//Read string
	var str string
	fmt.Scanf("%s", &str)
	// fmt.Println(str)
	var path string
	path = "../k8s-info-collector/" + str
	fmt.Println(path)
	dat, err := os.ReadFile(path)
	fmt.Print(err)
	// check(err)
	fmt.Print(string(dat))

	var pwd string = "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"
	fmt.Print(pwd)
	sub.Foo()

	sub2.Bar()
}
