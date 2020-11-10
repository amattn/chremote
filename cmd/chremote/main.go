package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/amattn/chremote/internal/util"
	"github.com/amattn/chremote/pkg/chremotelib"
	"github.com/amattn/deeperror"
)

func main() {
	log.Println(util.CurrentFunction(), "entering")
	defer util.Trace(util.CurrentFunction(), time.Now())

	// var err error
	log.Printf("Starting %v", chremotelib.VersionInfo())
	log.Printf("os.Args: %v", os.Args)
	log.Printf("Go (runtime:%v) (GOMAXPROCS:%d) (NumCPUs:%d)\n", runtime.Version(), runtime.GOMAXPROCS(-1), runtime.NumCPU())

	// things that will eventually be cli flags and/or config variables:
	browserBootstrapURL := "http://localhost:9222/json"

	payloadHandler := func(tracer int64, payload interface{}) {
		log.Println(1675213581, "payloadHandler: incoming payload", tracer, payload)
	}
	errorHandler := func(tracer int64, err error) {
		log.Println(9282934429, "errorHandler: unexpected error", tracer, err)
	}

	client := chremotelib.NewClient(chremotelib.Chrome, browserBootstrapURL, payloadHandler, errorHandler)
	if client == nil {
		derr := deeperror.New(4064709656, "unexpected nil client failure", nil)
		derr.AddDebugField("browserWebSocketURL", browserBootstrapURL)
		log.Fatal(derr)
	}

	err := client.Connect()
	if err != nil {
		derr := deeperror.New(4134345954, "client.Connect failure:", err)
		derr.AddDebugField("client", client)
		log.Fatal(derr)
		return
	}

	thing := map[string]interface{}{
		"hi": "world",
		"数字": 1234,
	}

	id, err := client.SendJSON(thing)
	if err != nil {
		derr := deeperror.New(2479404338, "client.SendJson failure:", err)
		derr.AddDebugField("id", id)
		log.Println(derr)
		return
	}

	i := 0
	go func() {
		for {
			i++
			time.Sleep(5 * time.Second)

			websites := []string{
				"https://amattn.com/",
				"https://github.com/amattn",
				"https://twitter.com/amattn",
			}
			idx := i % len(websites)
			wurl := websites[idx]
			log.Println("navigate to", i, idx, wurl)

			id, err := client.NavigateTo(wurl)
			if err != nil {
				derr := deeperror.New(2479404338, "client.SendJson failure:", err)
				derr.AddDebugField("id", id)
				log.Println(derr)
				return
			}
		}
	}()

	err = client.Listen()
	log.Fatal(2274101116, err)
}
