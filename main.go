package main

import (
	api2 "github.com/guil95/grpcApi/api"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("pkg/db/products.json")


	api2.Run(file)
}
