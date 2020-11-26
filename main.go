package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

type Segment struct {
	Timestamp float64 `json:"timestamp"`
	Duration float64 `json:"duration"`
	Fingerprint string `json:"fingerprint"`
}

func main() {
	url := "https://radiom2o-lh.akamaihd.net/i/RadioM2o_Live_1@42518/master.m3u8"
	chunk := strconv.Itoa(3)
	// length := strconv.Itoa(20)

	cmd := exec.Command("fpcalc", "-chunk", chunk, "-overlap", "-json", url)
	cmdR, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	
	sc := bufio.NewScanner(cmdR)
	go func() {
		for sc.Scan() {
			var segment Segment
			err := json.Unmarshal(sc.Bytes(), &segment)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(segment.Fingerprint)
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}