package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pallat/queue"
)

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

func main() {
	works := decrypt()
	fmt.Println("start...")

	s := mySimpler{works}
	var w work
	ctx := context.Background()
	m := queue.NewManager(ctx, &w, s)

	for i := 0; i < 200; i++ {
		go m.Do()
	}

	<-m.End()
	fmt.Println("finish")
}

func decrypt() [][]string {
	b, _ := ioutil.ReadFile("fng.1000.csv.rot128")
	rot128(b)

	s := string(b)

	row := strings.Split(s, "\n")

	table := [][]string{}

	for i, r := range row {
		if i == 0 {
			continue
		}
		table = append(table, strings.Split(r, ","))
	}

	return table
}

func rot128(buf []byte) {
	for idx := range buf {
		buf[idx] -= 128
	}
}
