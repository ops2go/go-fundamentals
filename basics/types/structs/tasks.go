package structs

import "html/template"

type Task struct {
	Id           int
	Title        string
	Content      template.HTML
	Created      string
	Priority     string
	Category     string
	Referer      string
	Comments     []Comment
	IsOverdue    bool
	IsHidden     int
	CompletedMsg string
}

type Tasks []Tasks

type Comment struct {
	ID       int
	Content  string
	Created  string
	Username string
}
