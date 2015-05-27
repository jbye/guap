package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jbye/guap/dota"

	logger "github.com/Sirupsen/logrus"
	r "github.com/dancannon/gorethink"
)

var gLeagueID = "2733"
var gMatchHistoryInterval time.Duration = 5 * 1000
var gMatchesInterval time.Duration = 12222222 * 60

// -
var WorkQueue = make(chan WorkRequest, 100)

func main() {
	loadConfig()

	err := initializeDatabase(Config.Rethink.Host, Config.Rethink.Database)

	if err != nil {
		log.Fatalln(err.Error())
	}

	StartDispatcher(4)

	heroesChan := time.NewTicker(time.Millisecond * Config.Intervals.Heroes).C
	leaguesChan := time.NewTicker(time.Millisecond * Config.Intervals.Leagues).C

	for {
		select {
		case <-heroesChan:
			WorkQueue <- WorkRequest{
				Name:  "Heroes request",
				Delay: 100,
			}
		case <-leaguesChan:
			WorkQueue <- WorkRequest{
				Name:  "League request",
				Delay: 100,
			}
		}
	}
}

func processHeroes() {
	logger.Info("Started processing heroes")

	var heroesResponse = dota.GetHeroesResponse(Config.Steam.Key, "en_us")

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

	var leaguesResponse = dota.GetLeaguesResponse(Config.Steam.Key)

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
	var matchHistoryResponse = dota.GetMatchHistoryResponse(Config.Steam.Key, 3)

	fmt.Printf("Num results: %d\n", matchHistoryResponse.Result.NumResults)

	for _, matchHistory := range matchHistoryResponse.Result.Matches {
		var matchDetails = dota.GetMatchDetailsResponse(Config.Steam.Key, matchHistory.ID)
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
