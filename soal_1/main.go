package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MatchResponse struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Match `json:"data"`
}

type Match struct {
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

func fetchGoals(team string, year int, position string) (int, error) {
	totalGoals := 0
	page := 1

	for {
		url := fmt.Sprintf(
			"https://jsonmock.hackerrank.com/api/football_matches?year=%d&%s=%s&page=%d",
			year, position, team, page,
		)

		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}

		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var mr MatchResponse
		err = json.Unmarshal(body, &mr)
		if err != nil {
			return 0, err
		}

		// Process matches
		for _, m := range mr.Data {
			if position == "team1" {
				g, _ := strconv.Atoi(m.Team1Goals)
				totalGoals += g
			} else {
				g, _ := strconv.Atoi(m.Team2Goals)
				totalGoals += g
			}
		}

		// Last page?
		if page >= mr.TotalPages {
			break
		}

		page++
	}

	return totalGoals, nil
}

func main() {
	var team string
	var year int

	fmt.Scanln(&team)
	fmt.Scanln(&year)

	team1Goals, _ := fetchGoals(team, year, "team1")
	team2Goals, _ := fetchGoals(team, year, "team2")

	fmt.Println(team1Goals + team2Goals)
}
