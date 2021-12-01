package main

import "time"

type Contact struct {
	Email            string   `csv:"email"`
	Name             string   `csv:"name"`
	Stop             bool      `csv:"stop"`
	Step_0_Timestamp Timestamp `csv:"step_0_timestamp"`
	Step_0_Done      bool      `csv:"step_0_done"`
	Step_1_Timestamp Timestamp `csv:"step_1_timestamp"`
	Step_1_Done      bool      `csv:"step_1_done"`
	Step_2_Timestamp Timestamp `csv:"step_2_timestamp"`
	Step_2_Done      bool      `csv:"step_2_done"`
	Step_3_Timestamp Timestamp `csv:"step_3_timestamp"`
	Step_3_Done      bool      `csv:"step_3_done"`
	Step_4_Timestamp Timestamp `csv:"step_4_timestamp"`
	Step_4_Done      bool      `csv:"step_5_done"`
}

type Timestamp struct {
	time.Time
}

func (date *Timestamp) MarshalCSV() (string, error) {
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
