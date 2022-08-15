package switcher

type Message struct {
	Id       int    `json:"id"`
	Callback bool   `json:"callback"`
	Root     string `json:"root"`
	Service  string `json:"service"`
	Data     Data   `json:"data"`
}

type Data struct {
	Content any `json:"content"`
}
