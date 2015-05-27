package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/jbye/guap/dota"

	logger "github.com/Sirupsen/logrus"
	r "github.com/dancannon/gorethink"
)

var gLeagueID = "2733"
var gHeroesInterval time.Duration = 15 * 1000
var gLeaguesInterval time.Duration = 30 * 1000
var gMatchesInterval time.Duration = 12222222 * 60

func main() {
	err := initializeDatabase("nas1.local:28015", "", "", "heroes")

	if err != nil {
		log.Fatalln(err.Error())
	}

	// Heroes
	heroesChan := time.NewTicker(time.Millisecond * gHeroesInterval).C
	leaguesChan := time.NewTicker(time.Millisecond * gLeaguesInterval).C
	matchHistoryChan := time.NewTicker(time.Millisecond * gMatchesInterval).C
	doneChan := make(chan bool)

	var heroesWg sync.WaitGroup

	for {
		select {
		case <-heroesChan:
			go func() {
				processHeroes()
				defer heroesWg.Done()
			}()
		case <-leaguesChan:
			go processLeagues()
		case <-matchHistoryChan:
			fmt.Println("Checking matches")
		case <-doneChan:
			fmt.Println("Done")
		}
	}

	/*

		leagues := []string{"2733"}

		for _, leagueID := range leagues {
			fmt.Printf("Running for leagueID: %s\n", leagueID)
		}

		var leaguesResponse = dota.GetLeaguesResponse()
		for _, league := range leaguesResponse.Result.Leagues {
			fmt.Printf("LeagueName: %s\n", league.Name)
		}*/

	//historyChan := make(chan Match)

	//processMatchHistory()
	//processHeroes()
}

func processHeroes() {
	logger.Info("Started processing heroes")

	var heroesResponse = dota.GetHeroesResponse("en_us")

	for _, heroVO := range heroesResponse.Result.Heroes {
		hero := Hero{
			Name:   heroVO.Name,
			HeroID: heroVO.ID,
		}

		insertHero(hero)
	}

	logger.Info("Completed processing leagues")
}

func processLeagues() {
	logger.Info("Started processing leagues")

	var leaguesResponse = dota.GetLeaguesResponse()

	for _, leagueVO := range leaguesResponse.Result.Leagues {
		league := League{
			Name:     leagueVO.Name,
			LeagueID: leagueVO.ID,
		}

		insertLeague(league)
	}

	logger.Info("Completed processing leagues")
}

func processMatchHistory() {
	var matchHistoryResponse = dota.GetMatchHistoryResponse(3)

	fmt.Printf("Num results: %d\n", matchHistoryResponse.Result.NumResults)

	for _, matchHistory := range matchHistoryResponse.Result.Matches {
		var matchDetails = dota.GetMatchDetailsResponse(matchHistory.ID)
		fmt.Printf("MatchDetails: %s\n", strconv.FormatBool(matchDetails.Result.RadiantWin))
	}
}

func parseMatchHistory(matchHistory dota.MatchHistory) {
	var doc = map[string]interface{}{
		"MatchId":   matchHistory.ID,
		"StartTime": matchHistory.StartTime,
	}

	result, err := r.Table("MatchHistory").Insert(doc).RunWrite(gSession)

	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Debug("Inserted MatchId %d with key: %s\n", matchHistory.ID, result.GeneratedKeys[0])
}
