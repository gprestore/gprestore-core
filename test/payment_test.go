package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/pkg/flip"
)

var flipClient *flip.FlipClient

func init() {
	config.Load()
	flipClient = flip.NewFlipClient()
}

func TestCreateBill(t *testing.T) {
	request := &flip.FlipBillRequest{
		Title:          "Minecraft Full Access",
		Amount:         150000,
		Type:           flip.FlipBillTypeSingle,
		Step:           "3",
		SenderName:     "Agil Ghani Istikmal",
		SenderEmail:    "agilistikmal3@gmail.com",
		SenderBank:     "qris",
		SenderBankType: "wallet_account",
	}

	bill, err := flipClient.CreatePayment(request)
	if err != nil {
		log.Fatal(err)
	}
	billJson, _ := json.Marshal(bill)
	fmt.Println(string(billJson))
}

func TestGetBills(t *testing.T) {
	bills, err := flipClient.GetBills()
	if err != nil {
		log.Fatal(err)
	}
	billsJson, _ := json.Marshal(bills)
	fmt.Println(string(billsJson))
}

func TestGetBill(t *testing.T) {
	bill, err := flipClient.GetBill(128572)
	if err != nil {
		log.Fatal(err)
	}
	billJson, _ := json.Marshal(bill)
	fmt.Println(string(billJson))
}

func TestGetPayment(t *testing.T) {
	bill, err := flipClient.GetPayment(128572)
	if err != nil {
		log.Fatal(err)
	}
	billJson, _ := json.Marshal(bill)
	fmt.Println(string(billJson))
}
