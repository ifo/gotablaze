package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var apikey = os.Getenv("APIKEY")

func main() {
	cfg := defaultDBConfig()
	s := cfg.Setup()
	s.MakeApiCallsForever(16)
}

// TODO setup real logging
func (s *DBSession) MakeApiCallsForever(sleep time.Duration) {
	log.Println("api calls forever started")
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
		err := s.SaveGame(game)
		if err != nil {
			log.Fatalln("saveGame failed: ", err)
		}
	}

	for {
		start := time.Now()
		curr, err := makeApiCall()
		// either there was a problem or no games to process
		// TODO if response was EOF, retry sooner
		if len(curr) == 0 || err != nil {
			if err != nil {
				log.Println(err)
			}
			time.Sleep(sleep*time.Second - time.Since(start))
			continue
		}

		// compare with previous request
		oldGames, curGames, newGames := alignGames(prev, curr)

		// create new games
		for _, game := range newGames {
			err := s.SaveGame(game)
			if err != nil {
				log.Println("saveGame failed: ", err)
			}
		}

		// diff and save games (loop)
		for i := range oldGames {
			sb := curGames[i].Scoreboard
			if !sb.identical(oldGames[i].Scoreboard) {
				sbd := curGames[i].Scoreboard.diff(oldGames[i].Scoreboard)
				game := Game{
					MatchID:    oldGames[i].MatchID,
					Scoreboard: sbd,
					Timestamp:  curGames[i].Timestamp,
				}
				err := s.SaveGame(game)
				if err != nil {
					log.Println("saveGame failed: ", err)
				}
			}
		}

		prev = curr

		// wait the remainder of 'sleep' seconds
		time.Sleep(sleep*time.Second - time.Since(start))
	}
}

func makeApiCall() ([]Game, error) {
	resp, err := http.Get("https://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v0001/?format=XML&key=" + apikey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Result
	if err := xml.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	curTime := time.Now().UTC()
	for i := range res.Games {
		res.Games[i].Timestamp = curTime
	}

	return res.Games, nil
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
