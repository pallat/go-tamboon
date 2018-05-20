package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type work struct{}

func (work) Do(v interface{}) error {
	w := v.([]string)

	if len(w) < 6 {
		log.Println("wrong data:", w)
		return nil
	}

	client, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	month, err := strconv.Atoi(w[4])
	if err != nil {
		log.Fatal(err)
		return err
	}

	year, err := strconv.Atoi(w[5])
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Creates a token from a test card.
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            w[0],
		Number:          w[2],
		ExpirationMonth: time.Month(month),
		ExpirationYear:  year,
	}
	if e := client.Do(token, createToken); e != nil {
		log.Fatal(e)
	}

	amount, err := strconv.Atoi(w[1])
	if err != nil {
		return err
	}

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(amount), // ฿ 1,000.00
		Currency: "thb",
		Card:     token.ID,
	}

	client.Client.Transport.(*http.Transport).DisableKeepAlives = true

	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	fmt.Println("finish", w)

	return nil
}

type mySimpler struct {
	row [][]string
}

func (m mySimpler) Len() int {
	return len(m.row)
}

func (m mySimpler) Pop(i int) interface{} {
	return m.row[i]
}
