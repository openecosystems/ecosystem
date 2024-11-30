package main

import (
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"libs/poc/go/generated-wsdl/gen"
	"net/http"
	"time"
)

var done = make(chan struct{})

func client() {
	client := soap.NewClient("http://127.0.0.1:8000")
	service := gen.NewPrepaidServicesSoap(client)
	resp, err := service.AccountCreation(&gen.AccountCreation{
		XMLName: xml.Name{
			Space: "",
			Local: "",
		},
		PrePaidAccountCreationRequest: &gen.AccountCreationRequest{
			RecordType:                 "",
			RecordNumber:               "",
			RequestType:                "",
			ProductType:                "",
			SubProductType:             "",
			FromCardRange:              "",
			ToCardRange:                "",
			PAN:                        "",
			SecondaryPAN:               "",
			AccountCreationDate:        "",
			AccountCreationTime:        "",
			Amount:                     "",
			ReferenceNo:                "",
			ExpirationDate:             "",
			CurrencyCode:               "",
			MerchantGroup:              "",
			MCC:                        "",
			TerminalID:                 "",
			AccountType:                "",
			Name1:                      "",
			Name2:                      "",
			Address1:                   "",
			Address2:                   "",
			City:                       "",
			State:                      "",
			PostalCode:                 "",
			GovernmentIDType:           "",
			GovernmentID:               "",
			CountryOfIssue:             "",
			PhoneNumber:                "",
			WorkPhoneNumber:            "",
			MobilePhoneNumber:          "",
			OtherPhoneNumber:           "",
			EmailID:                    "",
			EmailTwo:                   "",
			DateOfBirth:                "",
			BankRouting:                "",
			BankAcctNumber:             "",
			Comment:                    "",
			PIN:                        "",
			KeyLabel:                   "",
			Filler:                     "",
			ProductId:                  "",
			FollowUpDate:               "",
			Employername:               "",
			EmployerContactName:        "",
			EmployerContactPhoneNumber: "",
			EmployerContactFaxNumber:   "",
			Memos:                      "",
			CustRefName:                "",
			StoreName:                  "",
			GovtIDIssueDate:            "",
			GovtIDExpirationDate:       "",
			GovtIDCountryofIssuance:    "",
			CIPType:                    "",
			CIPNumber:                  "",
			CIPStatus:                  "",
			GovtIDIssueState:           "",
			SSN:                        "",
			TaxId:                      "",
			DeliveryMechanism:          "",
			EmbossingLine4:             "",
			OtherIDDescription:         "",
			EmbossingHotStamp:          "",
			Title:                      "",
			SecondLastName:             "",
			NameOnCard:                 "",
			MotherMaidenName:           "",
			HomeFaxNumber:              "",
			WorkFaxNumber:              "",
			LanguageIndicator:          "",
			MiddleName:                 "",
		},
	})
	fmt.Println(resp.AccountNumber, err)
	//done <- struct{}{}
}

// use fixtures/test.wsdl
func main() {
	http.HandleFunc("/", gen.Endpoint)
	go func() {
		time.Sleep(time.Second * 1)
		client()
	}()
	go func() {
		http.ListenAndServe(":8000", nil)
	}()
	<-done
}
