package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	gdrive "gomodules.xyz/gdrive-utils"
	"gomodules.xyz/mailer"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// https://docs.google.com/spreadsheets/d/1evwv2ON94R38M-Lkrw8b6dpVSkRYHUWsNOuI7X0_-zA/edit#gid=584220329
const (
	spreadsheetId = "10Jx3-1Ww2UQ7xNjs9-CRvJX4iIA22EDu-EsLKoHp1hc"
	sheetName     = "NEW_SIGNUP"
	// header        = "email"

	// MailLicenseSender  = "license-issuer@mail.appscode.com"
	MailLicenseTracker = "issued-license-tracker@appscode.com"
	MailSupport        = "support@appscode.com"
	// MailSales          = "sales@appscode.com"
	MailHello          = "hello@appscode.com"
)

func main_date() {
	now := time.Now()
	t2 := mailer.Timestamp{now}
	fmt.Println(now.UTC())

	s, _ := t2.MarshalCSV()
	t, err := time.Parse("01/02/2006 15:04:05", s)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.String())
	// os.Exit(1)
}

func main_add_contact() {
	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	srv, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		panic(err)
	}

	campaign := getDripCampaign(srv, mg)

	c := mailer.Contact{
		Email: "tamal@appscode.com",
		Data: toJson(ContactData{
			Name:    "Tamal Saha",
			Product: "KubeDB",
		}),
	}
	err = campaign.AddContact(c)
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		panic(err)
	}

	srv, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	campaign := getDripCampaign(srv, mg)

	err = campaign.ProcessCampaign()
	if err != nil {
		panic(err)
	}
}

func getDripCampaign(srv *sheets.Service, mg mailgun.Mailgun) *mailer.DripCampaign {
	return &mailer.DripCampaign{
		Name: "New Signup",
		Steps: []mailer.CampaignStep{
			{
				WaitTime: 0,
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "Welcome to {{.Product}}",
					Body:            "Hey {{.Name}}, Thanks for using {{.Product}}!",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 10 * time.Second,
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "How are things with {{.Product}}",
					Body:            "Hey {{.Name}}, How are things going with {{.Product}}. If you need help, contact support@appscode.com",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 30 * time.Second,
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "Your trial ending soon",
					Body:            "Hey {{.Name}}, your trial of {{.Product}} is ending soon",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
		},
		M:             mg,
		SheetService:  srv,
		SpreadsheetId: spreadsheetId,
		SheetName:     sheetName,
	}
}

func toJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
