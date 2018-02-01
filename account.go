package sib

import (
	"math"
)

// AccountDetails conveys information and account limits
type AccountDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		PlanType   string `json:"plan_type,omitempty"`
		Credits    int    `json:"credits,omitempty"`
		CreditType string `json:"credit_type,omitempty"`
		ClientID   int    `json:"client_id,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Email      string `json:"email,omitempty"`
		Company    string `json:"company,omitempty"`
		Address    string `json:"address,omitempty"`
		City       string `json:"city,omitempty"`
		ZipCode    string `json:"zip_code,omitempty"`
		Country    string `json:"country,omitempty"`
	} `json:"data"`
}

// SMTPAccountDetails conveys SMTP account information and limits
type SMTPAccountDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TrackingData struct {
			SiteID      string `json:"site_id"`
			Email       string `json:"email"`
			Fname       string `json:"fname"`
			Lname       string `json:"lname"`
			CompanyName string `json:"company_name"`
		} `json:"tracking_data"`
		RelayData struct {
			Result  string `json:"result"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"relay_data"`
	} `json:"data"`
}

// SMTPCreditCount returns the number of credits available in the account
func (account AccountDetails) SMTPCreditCount() int {
	for _, v := range account.Data {
		switch v.PlanType {
		case "PAG":
			fallthrough
		case "FREE":
			fallthrough
		case "CREDIT_REC":
			return v.Credits
		case "UNLIMITED":
			return math.MaxInt32
		default:
			return 0
		}
	}
	return 0
}
