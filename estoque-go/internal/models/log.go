package models

import "time"

type Log struct {
	Timestamp time.Time
	Action    string
	User      string
	ItemID    int
	Quantity  int
	Reason    string
}
