package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"strings"
	"time"

	gdrive "gomodules.xyz/gdrive-utils"
	"gomodules.xyz/sets"
	"google.golang.org/api/sheets/v4"
)

func main() {
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

	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	srv, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// https://docs.google.com/spreadsheets/d/1evwv2ON94R38M-Lkrw8b6dpVSkRYHUWsNOuI7X0_-zA/edit#gid=584220329
	const (
		spreadsheetId = "10Jx3-1Ww2UQ7xNjs9-CRvJX4iIA22EDu-EsLKoHp1hc"
		sheetName     = "NEW_SIGNUP"
		header        = "Email"
	)
	emails := ListEmails(err, srv, spreadsheetId, sheetName, header)
	fmt.Println(strings.Join(emails.List(), "\n"))
}

func ListEmails(err error, srv *sheets.Service, spreadsheetId string, sheetName string, header string) sets.String {
	reader, err := gdrive.NewColumnReader(srv, spreadsheetId, sheetName, header)
	if err != nil {
		panic(err)
	}
	cols, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	emails := sets.NewString()
	for _, row := range cols {
		emails.Insert(row...)
	}
	return emails
}
