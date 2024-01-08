package main

import "time"

const Templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Producing output with a template is a two-step process.
// First we must parse the template into a suitable internal representation, and then execute it on specific inputs
