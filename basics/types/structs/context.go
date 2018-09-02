package structs

type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
	CSRFToken  string
	Categories []CategoryCount
	Referer    string
}
