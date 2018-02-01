package sib

import (
	"fmt"
	"testing"
)

func TestAccountDetail(t *testing.T) {
	// recommendation: set API key as system environment variable
	apiKey := ""

	// Create SendInBlue Client
	sibClient, err := NewClient(apiKey)
	if err != nil {
		t.Error("Cannot create client")
	}

	accountDetails, err := sibClient.GetAccount()
	if err != nil {
		t.Error("Cannot get account")
	}

	fmt.Println("emails left:", accountDetails.SMTPCreditCount())
}

func TestAccountSMTPStatus(t *testing.T) {
	apiKey := ""

	// Create SendInBlue Client
	sibClient, err := NewClient(apiKey)
	if err != nil {
		t.Error("Cannot create client")
	}

	smtpDetails, err := sibClient.GetSMTPDetails()
	if err != nil {
		t.Error("Cannot get account details")
	}

	fmt.Printf("%+v\n", smtpDetails)
}
