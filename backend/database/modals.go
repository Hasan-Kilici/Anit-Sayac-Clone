package database

type Incident struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	FullName   string   `json:"fullname"`
	Age        string   `json:"age"`
	Location   string   `json:"location"`
	Date       string   `json:"date"`
	Year	   string	`json:"year"`
	Reason     string   `json:"reason"`
	By         string   `json:"by"`
	Protection string   `json:"protection"`
	Method     string   `json:"method"`
	Status     string   `json:"status"`
	Source     []string `json:"source"`
	Image      string   `json:"image"`
	Url        string   `json:"url"`
}

type Detail struct {
	Name       string   `json:"name"`
	Age        string   `json:"age"`
	Location   string   `json:"location"`
	Date       string   `json:"date"`
	Reason     string   `json:"reason"`
	By         string   `json:"by"`
	Protection string   `json:"protection"`
	Method     string   `json:"method"`
	Status     string   `json:"status"`
	Source     []string `json:"source"`
	Image      string   `json:"image"`
}
