package main

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

var gSession *r.Session

func initializeDatabase(host string, username string, password string, database string) error {
	var err error
	gSession, err = r.Connect(r.ConnectOpts{
		Address:  host,
		Database: database,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func insertHero(hero Hero) {
	// Check if hero is inserted already
	res, err := r.Table("Heroes").Filter(r.Row.Field("hero_id").Eq(hero.HeroID)).Run(gSession)

	if err != nil {
		fmt.Println(err)
		return
	}

	var existing *Hero
	res.One(&existing)

	if existing != nil {
		fmt.Printf("Updating hero with ID: %s\n", existing.ID)
		_, err = r.Table("Heroes").Get(existing.ID).Update(hero).Run(gSession)

	} else {
		fmt.Printf("Inserting hero\n")
		_, err = r.Table("Heroes").Insert(hero).Run(gSession)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}

/*********** Models *************/

// Hero -
type Hero struct {
	ID     string `gorethink:"id,omitempty"`
	HeroID int    `gorethink:"hero_id"`
	Name   string `gorethink:"name"`
}

// Match -
type Match struct {
	ID          string    `gorethink:"id,omitempty"`
	MatchID     int       `gorethink:"match_id"`
	LeagueID    int       `gorethink:"league_id"`
	Processed   bool      `gorethink:"processed"`
	ProcessedAt time.Time `gorethink:"processed_at"`
}

// HeroScore -
type HeroScore struct {
	ID      string `gorethink:"id,omitempty"`
	MatchID int    `gorethink:"match_id"`
	HeroID  int    `gorethink:"hero_id"`
	Score   int    `gorethink:"score"`
}

// League -
type League struct {
	ID       string `gorethink:"id,omitempty"`
	LeagueID int    `gorethink:"league_id"`
	Observed bool   `gorethink:"observed"`
}

/*
cursor, err := r.Table("people").Run(session)
if err != nil {
    fmt.Println(err)
    return
}

var person interface{}
cursor.One(&person)
cursor.Close()

printStr("*** Fetch one record: ***")
printObj(person)
printStr("\n")
*/
