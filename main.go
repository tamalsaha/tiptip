package main

import (
	"fmt"
	"log"
	"strings"

	gdrive "gomodules.xyz/gdrive-utils"
	"gomodules.xyz/sets"
	"google.golang.org/api/sheets/v4"
)

func main() {
	client, err := gdrive.DefaultClient(".")
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// https://docs.google.com/spreadsheets/d/1evwv2ON94R38M-Lkrw8b6dpVSkRYHUWsNOuI7X0_-zA/edit#gid=584220329
	const (
		spreadsheetId = "1evwv2ON94R38M-Lkrw8b6dpVSkRYHUWsNOuI7X0_-zA"
		sheetName     = "License Issue Log"
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
