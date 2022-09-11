package entity

import (
	"encoding/json"
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewEvent(date time.Time, name, description string) *Event {
	return &Event{
		Date:        date,
		Name:        name,
		Description: description,
	}
}

func (e Event) Validate() bool {
	if e.Date.IsZero() || e.Name == "" || e.Description == "" || e.Id < 0 {
		return false
	}
	return true
}

func (e Event) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
