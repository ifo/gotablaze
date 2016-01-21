package main

import (
	"time"
)

type Dota2ApiResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Games  []Game `json:"games"`
	Status int    `json:"status"`
}

type Game struct {
	DireSeriesWins int `json:"dire_series_wins" gorethink:"dire_series_wins,omitempty"`
	LeagueGameID   int `json:"league_game_id" gorethink:"league_game_id,omitempty"`
	LeagueID       int `json:"league_id" gorethink:"league_id,omitempty"`
	LeagueSeriesID int `json:"league_series_id" gorethink:"league_series_id,omitempty"`
	LeagueTier     int `json:"league_tier" gorethink:"league_tier,omitempty"`
	LobbyID        int `json:"lobby_id" gorethink:"lobby_id,omitempty"`
	MatchID        int `json:"match_id" gorethink:"match_id"`
	Players        []struct {
		AccountID int    `json:"account_id" gorethink:"account_id,omitempty"`
		HeroID    int    `json:"hero_id" gorethink:"hero_id,omitempty"`
		Name      string `json:"name" gorethink:"name,omitempty"`
		Team      int    `json:"team" gorethink:"team,omitempty"`
	} `json:"players" gorethink:"players,omitempty"`
	RadiantSeriesWins int `json:"radiant_series_wins" gorethink:"radiant_series_wins,omitempty"`
	RadiantTeam       struct {
		Complete bool   `json:"complete" gorethink:"complete,omitempty"`
		TeamID   int    `json:"team_id" gorethink:"team_id,omitempty"`
		TeamLogo int    `json:"team_logo" gorethink:"team_logo,omitempty"`
		TeamName string `json:"team_name" gorethink:"team_name,omitempty"`
	} `json:"radiant_team" gorethink:"radiant_team,omitempty"`
	DireTeam struct {
		Complete bool   `json:"complete" gorethink:"complete,omitempty"`
		TeamID   int    `json:"team_id" gorethink:"team_id,omitempty"`
		TeamLogo int    `json:"team_logo" gorethink:"team_logo,omitempty"`
		TeamName string `json:"team_name" gorethink:"team_name,omitempty"`
	} `json:"dire_team" gorethink:"dire_team,omitempty"`
	Scoreboard   Scoreboard `json:"scoreboard" gorethink:"scoreboard,omitempty"`
	SeriesType   int        `json:"series_type" gorethink:"series_type,omitempty"`
	Spectators   int        `json:"spectators" gorethink:"spectators,omitempty"`
	StreamDelayS float64    `json:"stream_delay_s" gorethink:"stream_delay_s,omitempty"`
	TimeStamp    time.Time  `json:"-" gorethink:"time_stamp"`
}

type Scoreboard struct {
	Dire               Side    `json:"dire" gorethink:"dire,omitempty"`
	Duration           float64 `json:"duration" gorethink:"duration,omitempty"`
	Radiant            Side    `json:"radiant" gorethink:"radiant,omitempty"`
	RoshanRespawnTimer int     `json:"roshan_respawn_timer" gorethink:"roshan_respawn_timer,omitempty"`
}

type Side struct {
	BarracksState int      `json:"barracks_state" gorethink:"barracks_state,omitempty"`
	Players       []Player `json:"players" gorethink:"players,omitempty"`
	Score         int      `json:"score" gorethink:"score,omitempty"`
	TowerState    int      `json:"tower_state" gorethink:"tower_state,omitempty"`
}

type Player struct {
	AccountID        int     `json:"account_id" gorethink:"account_id,omitempty"`
	Assists          int     `json:"assists" gorethink:"assists,omitempty"`
	Death            int     `json:"death" gorethink:"death,omitempty"`
	Denies           int     `json:"denies" gorethink:"denies,omitempty"`
	Gold             int     `json:"gold" gorethink:"gold,omitempty"`
	GoldPerMin       int     `json:"gold_per_min" gorethink:"gold_per_min,omitempty"`
	HeroID           int     `json:"hero_id" gorethink:"hero_id,omitempty"`
	Item0            int     `json:"item0" gorethink:"item0,omitempty"`
	Item1            int     `json:"item1" gorethink:"item1,omitempty"`
	Item2            int     `json:"item2" gorethink:"item2,omitempty"`
	Item3            int     `json:"item3" gorethink:"item3,omitempty"`
	Item4            int     `json:"item4" gorethink:"item4,omitempty"`
	Item5            int     `json:"item5" gorethink:"item5,omitempty"`
	Kills            int     `json:"kills" gorethink:"kills,omitempty"`
	LastHits         int     `json:"last_hits" gorethink:"last_hits,omitempty"`
	Level            int     `json:"level" gorethink:"level,omitempty"`
	NetWorth         int     `json:"net_worth" gorethink:"net_worth,omitempty"`
	PlayerSlot       int     `json:"player_slot" gorethink:"player_slot,omitempty"`
	PositionX        float64 `json:"position_x" gorethink:"position_x,omitempty"`
	PositionY        float64 `json:"position_y" gorethink:"position_y,omitempty"`
	RespawnTimer     int     `json:"respawn_timer" gorethink:"respawn_timer,omitempty"`
	UltimateCooldown int     `json:"ultimate_cooldown" gorethink:"ultimate_cooldown,omitempty"`
	UltimateState    int     `json:"ultimate_state" gorethink:"ultimate_state,omitempty"`
	XpPerMin         int     `json:"xp_per_min" gorethink:"xp_per_min,omitempty"`
}
