package main

import (
	"log"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/amattn/deeperror"
	"github.com/amattn/wdc/internal/util"
	"github.com/amattn/wdc/pkg/wdclib"
)

func main() {
	log.Println(util.CurrentFunction(), "entering")
	defer util.Trace(util.CurrentFunction(), time.Now())

	// var err error
	log.Printf("Starting %v", util.VersionInfo())
	log.Printf("os.Args: %v", os.Args)
	log.Printf("Go (runtime:%v) (GOMAXPROCS:%d) (NumCPUs:%d)\n", runtime.Version(), runtime.GOMAXPROCS(-1), runtime.NumCPU())

	urlString := "ws://127.0.0.1:9222/example_socket"

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		derr := deeperror.New(3749032603, " failure:", err)
		derr.AddDebugField("urlString", urlString)
		log.Fatal(derr)
	}

	client := wdclib.NewClient(parsedURL)
	if client == nil {
		derr := deeperror.New(4064709656, "unexpected nil client failure", nil)
		derr.AddDebugField("parsedURL", parsedURL)
		log.Fatal(derr)
	}
}
