package cmd

import "time"

type TimingClientss struct {
	api_key string
}

type NewProjectRequest struct {
	Title             string `json:"title"`
	Parent            string `json:"parent"`
	Color             string `json:"color"`
	ProductivityScore int    `json:"productivity_score"`
	IsArchived        bool   `json:"is_archived"`
	TeamID            string `json:"team_id"`
}

type UpdateProjectRequest struct {
	Title             string `json:"title"`
	Parent            string `json:"parent"`
	Color             string `json:"color"`
	ProductivityScore int    `json:"productivity_score"`
	IsArchived        bool   `json:"is_archived"`
}

type NewTimerRequest struct {
	StartDate time.Time `json:"start_date"`
	Project   string    `json:"project"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes"`
}

type TimeEntryRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Project   string    `json:"project"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes"`
}
