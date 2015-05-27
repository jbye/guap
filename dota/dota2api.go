package dota

import (
	"fmt"

	"github.com/franela/goreq"
)

var gAPIKey = "632EDF23098CDBB4E13564B96831FD34"

// GetMatchHistoryResponse -
func GetMatchHistoryResponse(matchesRequested int) *MatchHistoryResponse {
	var baseURI = "https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v001/"
	/*var query = map[string]interface{}{
		"key":               gAPIKey,
		"matches_requested": matchesRequested,
	}*/

	query := struct {
		Key              string `url:"key"`
		MatchesRequested int    `url:"matches_requested"`
	}{
		gAPIKey,
		matchesRequested,
	}

	res, err := goreq.Request{
		Uri:         baseURI,
		QueryString: query,
	}.Do()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var responseObj MatchHistoryResponse
	res.Body.FromJsonTo(&responseObj)

	return &responseObj
}

// GetMatchDetailsResponse -
func GetMatchDetailsResponse(matchID int) *MatchDetailsResponse {
	var baseURI = "https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/V001/"
	query := struct {
		Key     string `url:"key"`
		MatchID int    `url:"match_id"`
	}{
		gAPIKey,
		matchID,
	}

	res, err := goreq.Request{
		Uri:         baseURI,
		QueryString: query,
	}.Do()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var responseObj MatchDetailsResponse
	res.Body.FromJsonTo(&responseObj)

	return &responseObj
}

// GetHeroesResponse -
func GetHeroesResponse(language string) *HeroesResponse {
	var baseURI = "https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/"
	query := struct {
		Key      string `url:"key"`
		Language string `url:"language"`
	}{
		gAPIKey,
		"en_us",
	}

	res, err := goreq.Request{
		Uri:         baseURI,
		QueryString: query,
	}.Do()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var responseObj HeroesResponse
	res.Body.FromJsonTo(&responseObj)

	return &responseObj
}

// MatchHistoryResponse -
type MatchHistoryResponse struct {
	Result struct {
		NumResults   int `json:"num_results"`
		TotalResults int `json:"total_results"`
		Matches      []MatchHistory
	}
}

// MatchHistory -
type MatchHistory struct {
	MatchID   int `json:"match_id"`
	StartTime int `json:"start_time"`
}

// HeroesResponse -
type HeroesResponse struct {
	Result struct {
		Heroes []HeroDetails `json:"heroes"`
		Status int           `json:"status"`
		Count  int           `json:"count"`
	}
}

// HeroDetails -
type HeroDetails struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localized_name"`
}

// MatchDetailsResponse -
type MatchDetailsResponse struct {
	Result struct {
		Players               []PlayerDetails  `json:"players"`
		PickBans              []PickBanDetails `json:"picks_bans"`
		RadiantWin            bool             `json:"radiant_win"`
		Duration              int              `json:"duration"`
		StartTime             int              `json:"start_time"`
		MatchID               int              `json:"match_id"`
		MatchSeqNum           int              `json:"match_seq_num"`
		TowerStatusRadiant    int              `json:"tower_status_radiant"`
		TowerStatusDire       int              `json:"tower_status_dire"`
		BarracksStatusRadiant int              `json:"barracks_status_radiant"`
		BarracksStatusDire    int              `json:"barracks_status_dire"`
		Cluster               int              `json:"cluster"`
		FirstBloodTime        int              `json:"first_blood_time"`
		LobbyType             int              `json:"lobby_type"`
		HumanPlayers          int              `json:"human_players"`
		LeagueID              int              `json:"leagueid"`
		PositiveVotes         int              `json:"positive_votes"`
		NegativeVotes         int              `json:"negative_votes"`
		GameMode              int              `json:"game_mode"`
		RadiantTeamID         int              `json:"radiant_team_id"`
		RadiantName           string           `json:"radiant_name"`
		RadiantLogo           int              `json:"radiant_logo"`
		RadiantTeamComplete   int              `json:"radiant_team_complete"`
		DireTeamID            int              `json:"dire_team_id"`
		DireName              string           `json:"dire_name"`
		DireLogo              int              `json:"dire_logo"`
		DireTeamComplete      int              `json:"dire_team_complete"`
		RadiantCaptain        int              `json:"radiant_captain"`
		DireCaptain           int              `json:"dire_captain"`
	}
}

// PlayerDetails -
type PlayerDetails struct {
	AccountID  string `json:"account_id"`
	PlayerSlot int    `json:"player_slot"`
	HeroID     int    `json:"hero_id"`
	Item0      int    `json:"item_0"`
	Item1      int    `json:"item_1"`
	Item2      int    `json:"item_2"`
	Item3      int    `json:"item_3"`
	Item4      int    `json:"item_4"`
	Item5      int    `json:"item_5"`
	Kills      int    `json:"kills"`
	Deaths     int    `json:"deaths"`
	Assists    int    `json:"assists"`
}

// PickBanDetails -
type PickBanDetails struct {
	IsPick bool `json:"is_pick"`
	HeroID int  `json:"hero_id"`
	Team   int  `json:"team"`
	Order  int  `json:"order"`
}
