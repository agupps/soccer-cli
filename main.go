package main

import (
	"encoding/json"
	"fmt"
	"intro-project/soccer-scores/leagues"
	"intro-project/soccer-scores/teams"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	apiKey := viper.GetString("apiKey")

	apiURL := "https://v3.football.api-sports.io/leagues"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var apiResponse = &leagues.ApiResponse{}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	premID := apiResponse.GetPremierLeagueID()

	apiURL = "https://v3.football.api-sports.io/teams"

	req, err = http.NewRequest(method, apiURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")
	// Create a new URL.Values map
	params := url.Values{}

	// Add your query parameters
	params.Add("league", strconv.Itoa(premID))
	params.Add("season", "2023")
	req.URL.RawQuery = params.Encode()
	res, err = client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var teamsResponse = &teams.ApiResponse{}

	err = json.Unmarshal(body, &teamsResponse)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Teams in the Premier League for the 2023 season:")
	for _, teamInfo := range teamsResponse.TeamInfo {
		fmt.Println(teamInfo.Team.Name)
	}
}
