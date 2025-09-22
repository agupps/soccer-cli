package leagues

type ApiResponse struct {
	Leagues []leagueInfo `json:"response"`
}

type leagueInfo struct {
	League  league  `json:"league"`
	Country country `json:"country"`
	Seasons any     `json:"seasons"`
}

type league struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Type string `json:"type"`
}

type country struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

func (a *ApiResponse) GetPremierLeagueID() int {
	for _, leagueInfo := range a.Leagues {
		if leagueInfo.League.Name == "Premier League" {
			return leagueInfo.League.Id
		}
	}
	return 0
}
