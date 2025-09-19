package teams

type ApiResponse struct {
	TeamInfo []teamInfo `json:"response"`
}

type teamInfo struct {
	Team team `json:"team"`
}

type team struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}
