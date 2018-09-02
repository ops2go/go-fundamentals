package types

import (
	"database/sql"
	"html/template"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

/*
Package types is used to store the context struct which
is passed while templates are executed.
*/
var err error

////////////////////////////////////////////////////////////////////
//Task is the struct used to identify tasks
type Task struct {
	Id           int           `json:"id"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	ContentHTML  template.HTML `json:"content_html"`
	Created      string        `json:"created"`
	Priority     string        `json:"priority"`
	Category     string        `json:"category"`
	Referer      string        `json:"referer,omitempty"`
	Comments     []Comment     `json:"comments,omitempty"`
	IsOverdue    bool          `json:"isoverdue, omitempty"`
	IsHidden     int           `json:"ishidden, omitempty`
	CompletedMsg string        `json:"ishidden, omitempty"`
}

type Tasks []Task

//Comment is the struct used to populate comments per tasks
type Comment struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Created  string `json:"created_date"`
	Username string `json:"username"`
}

///////////////////////////////////////////////////////////////
//CategoryCount is the struct used to populate the sidebar
//which contains the category name and the count of the tasks
//in each category
type CategoryCount struct {
	Name  string
	Count int
}

//Category is the structure of the category table
type Category struct {
	ID      int    `json:"category_id"`
	Name    string `json:"category_name"`
	Created string `json:"created_date"`
}

//Categories will show
type Categories []Category

/////////////////////////////////////////////////////////////
//Database
var database Database
var taskStatus map[string]int

type Database struct {
	db *sql.DB
}

var Store = sessions.NewCookieStore([]byte("secret-password"))
var session *sessions.Session

//////////////////////////////////////////////////////////////
//Templates
var homeTemplate *template.Template
var deletedTemplate *template.Template
var completedTemplate *template.Template
var editTemplate *template.Template
var searchTemplate *template.Template
var templates *template.Template
var loginTemplate *template.Template

var message string //message will store the message to be shown as notification

//////////////////////////////////////////////////////////////////

//Status is the JSON struct to be returned
type Status struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var redirectUrl string

///////////////////////////////////////////////////////////////////

//Context is the struct passed to templates
type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
	CSRFToken  string
	Categories []CategoryCount
	Referer    string
}

/////////////////////////////////////////////////////////////////////

type Configuration struct {
	ServerPort string
}

var config Configuration

//////////////////////////////////////////////////////////////////////

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret")
