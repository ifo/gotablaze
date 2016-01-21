package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var apikey = os.Getenv("APIKEY")

func main() {
	setupRethinkDB()
	makeApiCallsForDuration(2, 16)
}

// TODO setup real logging
func makeApiCallsForDuration(numberOfCalls int, sleep time.Duration) {
	if numberOfCalls < 1 {
		numberOfCalls = 5
	}
	if sleep < 2 {
		sleep = 16 // the magic Dota2 API number
	}

	prev, err := makeApiCall()
	if err != nil {
		log.Println(err)
	}

	if len(prev) == 0 {
		log.Fatalln("makeApiCall failed for the first call")
	}

	for _, game := range prev {
		err := saveGame(game, session)
		if err != nil {
			log.Fatalln("saveGame failed: ", err)
		}
	}

	// loop: make api call
	itr, start := make([]struct{}, numberOfCalls), time.Now()
	for i := range itr {
		curr, err := makeApiCall()
		// either there was a problem or no games to process
		// TODO if response was EOF, retry sooner
		if len(curr) == 0 || err != nil {
			time.Sleep(sleep*time.Second*time.Duration(i+1) - time.Since(start))
			continue
		}

		// compare with previous request
		oldGames, curGames, newGames := alignGames(prev, curr)

		// create new games
		for _, game := range newGames {
			err := saveGame(game, session)
			if err != nil {
				log.Println("saveGame failed: ", err)
			}
		}

		// diff and save games (loop)
		for i := range oldGames {
			sb := diffScoreboard(oldGames[i].Scoreboard, curGames[i].Scoreboard)
			if !areIdenticalScoreboard(oldGames[i].Scoreboard,
				curGames[i].Scoreboard) {
				game := Game{MatchID: oldGames[i].MatchID, Scoreboard: sb,
					TimeStamp: curGames[i].TimeStamp}
				err := saveGame(game, session)
				if err != nil {
					log.Println("saveGame failed: ", err)
				}
			}
		}

		prev = curr

		// wait the remainder of 'sleep' seconds
		time.Sleep(sleep*time.Second*time.Duration(i+1) - time.Since(start))
	}
}

func makeApiCall() ([]Game, error) {
	resp, err := http.Get("https://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v0001/?key=" + apikey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiresponse Dota2ApiResponse
	if err := json.Unmarshal(body, &apiresponse); err != nil {
		return nil, err
	}

	curTime := time.Now().UTC()
	for i := range apiresponse.Result.Games {
		apiresponse.Result.Games[i].TimeStamp = curTime
	}

	return apiresponse.Result.Games, nil
}

// TODO improve the O(n^2) algorithm (but n is small so take your time)
func alignGames(old, cur []Game) (prev, current, begin []Game) {
	newGamesIndices := []int{}
OUTER:
	for _, pg := range old {
		for i, cg := range cur {
			if pg.MatchID == cg.MatchID {
				prev, current = append(prev, pg), append(current, cg)
				newGamesIndices = append(newGamesIndices, i)
				continue OUTER
			}
		}
		// game had no match, so it has ended
	}

	// get all new games
	if len(newGamesIndices) != len(current) {
		for i := range cur {
			if !Contains(newGamesIndices, i) {
				begin = append(begin, cur[i])
			}
		}
	}
	return
}

func Contains(haystack []int, needle int) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}