package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

type PingRequestV1 struct {
	ThingID string   `json:"user_id"`
	Secret  string   `json:"secret"`
	Status  string   `json:"status"`
	Tags    []string `json:"tags"`
}

var (
	secret      = os.Getenv("THING_SECRET")
	id          = os.Getenv("THING_ID")
	status      = flag.String("s", "online", "status to set for device")
	tags        = flag.String("t", strings.Join(defaultTags, ","), "comma seperated list of tags")
	addr        = flag.String("a", "http://127.0.0.1:3000", "location of things server")
	defaultTags = []string{"os:" + runtime.GOOS, "go_version:" + runtime.Version(), fmt.Sprintf("num_cpu:%d", runtime.NumCPU()), "time:" + time.Now().Format(time.RFC3339)}
)

func main() {
	flag.Parse()
	if secret == "" || id == "" {
		fmt.Println("please set THING_SECRET and THING_ID")
		os.Exit(1)
	}

	pr := PingRequestV1{
		ThingID: id,
		Secret:  secret,
		Status:  *status,
		Tags:    strings.Split(*tags, ","),
	}
	d, err := json.Marshal(pr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Post(*addr+"/v1/ping", "application/json", bytes.NewBuffer(d))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if res.StatusCode != 200 {
		fmt.Println(res)
		os.Exit(1)
	}
	io.Copy(os.Stdout, res.Body)
}
