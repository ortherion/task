package requests

type Task struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	SignsMails []string `json:"signs_mails"`
}
