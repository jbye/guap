package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jbye/heroes/dota"

	logger "github.com/Sirupsen/logrus"
	r "github.com/dancannon/gorethink"
)

var gLeagueID = "2733"

func main() {
	err := initializeDatabase("nas1.local:28015", "", "", "heroes")

	if err != nil {
		log.Fatalln(err.Error())
	}

	//processMatchHistory()
	processHeroes()
}

func processHeroes() {
	var heroesResponse = dota.GetHeroesResponse("en_us")

	for _, heroVO := range heroesResponse.Result.Heroes {
		heroModel := Hero{
			Name:   heroVO.Name,
			HeroID: heroVO.ID,
		}

		insertHero(heroModel)
	}
}

func processMatchHistory() {
	var matchHistoryResponse = dota.GetMatchHistoryResponse(3)

	fmt.Printf("Num results: %d\n", matchHistoryResponse.Result.NumResults)

	for _, matchHistory := range matchHistoryResponse.Result.Matches {
		var matchDetails = dota.GetMatchDetailsResponse(matchHistory.MatchID)
		fmt.Printf("MatchDetails: %s\n", strconv.FormatBool(matchDetails.Result.RadiantWin))
	}
}

func parseMatchHistory(matchHistory dota.MatchHistory) {
	var doc = map[string]interface{}{
		"MatchId":   matchHistory.MatchID,
		"StartTime": matchHistory.StartTime,
	}

	result, err := r.Table("MatchHistory").Insert(doc).RunWrite(gSession)

	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Debug("Inserted MatchId %d with key: %s\n", matchHistory.MatchID, result.GeneratedKeys[0])
}
