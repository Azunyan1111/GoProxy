package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"

	"io/ioutil"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Println(os.Getenv("PORT"))

	resp, _ := http.Get("https://ipinfo.io")
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), proxy))
}
