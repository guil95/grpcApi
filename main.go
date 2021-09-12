package main

import (
	"github.com/guil95/grpcApi/api"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("pkg/db/products.json")
	api.Run(file)
}
