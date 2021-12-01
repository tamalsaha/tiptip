package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mailgun/mailgun-go/v4"
	"gomodules.xyz/mailer"
	"google.golang.org/api/option"

	gdrive "gomodules.xyz/gdrive-utils"
	"google.golang.org/api/sheets/v4"
)

// https://docs.google.com/spreadsheets/d/1evwv2ON94R38M-Lkrw8b6dpVSkRYHUWsNOuI7X0_-zA/edit#gid=584220329
const (
	spreadsheetId = "10Jx3-1Ww2UQ7xNjs9-CRvJX4iIA22EDu-EsLKoHp1hc"
	sheetName     = "NEW_SIGNUP"
	// header        = "email"
)

func main_date() {
	now := time.Now()
	t2 := Timestamp{now}
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
	now := time.Now()

	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	srv, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	campaign := getDripCampaign()

	type ContactData struct {
		Name    string
		Product string
	}

	c := Contact{
		Email: "tamal@appscode.com",
		Data: toJson(ContactData{
			Name:    "Tamal Saha",
			Product: "KubeDB",
		}),
	}
	campaign.Prepare(&c, now)

	w := gdrive.NewWriter(srv, spreadsheetId, sheetName)
	err = gocsv.MarshalCSV([]*Contact{&c}, w)
	if err != nil {
		panic(err)
	}
}

func main() {
	now := time.Now()

	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	var mg mailgun.Mailgun

	srv, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	campaign := getDripCampaign()

	reader, err := gdrive.NewReader(srv, spreadsheetId, sheetName, 1)
	if err != nil {
		panic(err)
	}
	var contacts []*Contact
	err = gocsv.UnmarshalCSV(reader, &contacts)
	if err != nil {
		panic(err)
	}
	for _, c := range contacts {
		fmt.Println(c)
		if c.Stop {
			continue
		}
		if c.Step_0_Timestamp.IsZero() || now.After(c.Step_0_Timestamp.Time) && !c.Step_0_Done {
			processStep(srv, mg, 0, campaign.Steps[0], *c)
			continue
		}
		if c.Step_1_Timestamp.IsZero() || now.After(c.Step_1_Timestamp.Time) && !c.Step_1_Done {
			processStep(srv, mg, 1, campaign.Steps[1], *c)
			continue
		}
		if c.Step_2_Timestamp.IsZero() || now.After(c.Step_2_Timestamp.Time) && !c.Step_2_Done {
			processStep(srv, mg, 2, campaign.Steps[2], *c)
			continue
		}
		if c.Step_3_Timestamp.IsZero() || now.After(c.Step_3_Timestamp.Time) && !c.Step_3_Done {
			processStep(srv, mg, 3, campaign.Steps[3], *c)
			continue
		}
		if c.Step_4_Timestamp.IsZero() || now.After(c.Step_4_Timestamp.Time) && !c.Step_4_Done {
			processStep(srv, mg, 4, campaign.Steps[4], *c)
			continue
		}
	}
}

func getDripCampaign() DripCampaign {
	campaign := DripCampaign{
		Steps: []DripCampaignStep{
			{
				WaitTime: 0,
				Mailer: mailer.Mailer{
					Sender:          "",
					BCC:             "",
					ReplyTo:         "",
					Subject:         "Welcome to {{.Product}}",
					Body:            "Hey {{.Name}}, Thanks for using {{.Product}}!",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  false,
				},
			},
			{
				WaitTime: 10 * time.Second,
				Mailer: mailer.Mailer{
					Sender:          "",
					BCC:             "",
					ReplyTo:         "",
					Subject:         "How are things with {{.Product}}",
					Body:            "Hey {{.Name}}, How are things going with {{.Product}}. If you need help, contact support@appscode.com",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  false,
				},
			},
			{
				WaitTime: 30 * time.Second,
				Mailer: mailer.Mailer{
					Sender:          "",
					BCC:             "",
					ReplyTo:         "",
					Subject:         "Your trial ending soon",
					Body:            "Hey {{.Name}}, your trial of {{.Product}} is ending soon",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  false,
				},
			},
		},
	}
	return campaign
}

func processStep(srv *sheets.Service, mg mailgun.Mailgun, stepIndex int, step DripCampaignStep, c Contact) error {
	var data Contact
	if err := json.Unmarshal([]byte(c.Data), &data); err != nil {
		return err
	}

	m := step.Mailer
	m.Params = &data
	err := m.SendMail(mg, c.Email, "", nil)
	if err != nil {
		return err
	}

	switch stepIndex {
	case 0:
		c.Step_0_Done = true
	case 1:
		c.Step_1_Done = true
	case 2:
		c.Step_2_Done = true
	case 3:
		c.Step_3_Done = true
	case 4:
		c.Step_4_Done = true
	}

	w := gdrive.NewRowWriter(srv, spreadsheetId, sheetName, &gdrive.Filter{
		Header: "email",
		By: func(v []interface{}) (int, error) {
			for idx, entry := range v {
				if entry.(string) == c.Email {
					return idx, nil
				}
			}
			return -1, fmt.Errorf("missing email %s", c.Email)
		},
	})
	return gocsv.MarshalCSV([]*Contact{&c}, w)
}

func toJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
