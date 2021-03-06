package main

import (
	"time"

	"gomodules.xyz/mailer"
)

type ContactData struct {
	Name    string
	Product string
}

type Contact struct {
	Email                   string    `csv:"email"`
	Data                    string    `csv:"data"` // json format
	Stop                    bool      `csv:"stop"`
	Step_0_Timestamp        Timestamp `csv:"s0_timestamp"`
	Step_0_WaitForCondition bool      `csv:"s0_wait_for_cond"`
	Step_0_Done             bool      `csv:"s0_done"`
	Step_1_Timestamp        Timestamp `csv:"s1_timestamp"`
	Step_1_WaitForCondition bool      `csv:"s1_wait_for_cond"`
	Step_1_Done             bool      `csv:"s1_done"`
	Step_2_Timestamp        Timestamp `csv:"s2_timestamp"`
	Step_2_WaitForCondition bool      `csv:"s2_wait_for_cond"`
	Step_2_Done             bool      `csv:"s2_done"`
	Step_3_Timestamp        Timestamp `csv:"s3_timestamp"`
	Step_3_WaitForCondition bool      `csv:"s3_wait_for_cond"`
	Step_3_Done             bool      `csv:"s3_done"`
	Step_4_Timestamp        Timestamp `csv:"s4_timestamp"`
	Step_4_WaitForCondition bool      `csv:"s4_wait_for_cond"`
	Step_4_Done             bool      `csv:"s4_done"`
}

type Timestamp struct {
	time.Time
}

func (date *Timestamp) MarshalCSV() (string, error) {
	if date.IsZero() {
		return "", nil
	}
	return date.Time.UTC().Format("01/02/2006 15:04:05"), nil
}

func (date *Timestamp) String() string {
	return date.Time.UTC().Format(time.RFC3339) // Redundant, just for example
}

func (date *Timestamp) UnmarshalCSV(csv string) (err error) {
	if csv != "" {
		date.Time, err = time.Parse("01/02/2006 15:04:05", csv)
	}
	return err
}

type DripCampaign struct {
	Steps []DripCampaignStep
}

type DripCampaignStep struct {
	WaitTime time.Duration
	Mailer   mailer.Mailer
}

func (dc DripCampaign) Prepare(c *Contact, t time.Time) {
	for idx, step := range dc.Steps {
		switch idx {
		case 0:
			c.Step_0_Timestamp = Timestamp{t.Add(step.WaitTime)}
		case 1:
			c.Step_1_Timestamp = Timestamp{t.Add(step.WaitTime)}
		case 2:
			c.Step_2_Timestamp = Timestamp{t.Add(step.WaitTime)}
		case 3:
			c.Step_3_Timestamp = Timestamp{t.Add(step.WaitTime)}
		case 4:
			c.Step_4_Timestamp = Timestamp{t.Add(step.WaitTime)}
		}
	}
}
