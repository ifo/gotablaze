package main

import (
	"time"
)

type Result struct {
	Games  []Game `xml:"games>game"`
	Status int    `xml:"status"`
}

type Game struct {
	DireSeriesWins int `xml:"dire_series_wins" gorethink:"dire_series_wins,omitempty"`
	GameNumber     int `xml:"game_number" gorethink:"game_number,omitempty"`
	LeagueGameID   int `xml:"league_game_id" gorethink:"league_game_id,omitempty"`
	LeagueID       int `xml:"league_id" gorethink:"league_id,omitempty"`
	LeagueSeriesID int `xml:"league_series_id" gorethink:"league_series_id,omitempty"`
	LeagueTier     int `xml:"league_tier" gorethink:"league_tier,omitempty"`
	LobbyID        int `xml:"lobby_id" gorethink:"lobby_id,omitempty"`
	MatchID        int `xml:"match_id" gorethink:"match_id"`
	Players        []struct {
		AccountID int    `xml:"account_id" gorethink:"account_id,omitempty"`
		HeroID    int    `xml:"hero_id" gorethink:"hero_id,omitempty"`
		Name      string `xml:"name" gorethink:"name,omitempty"`
		Team      int    `xml:"team" gorethink:"team,omitempty"`
	} `xml:"players>player" gorethink:"players,omitempty"`
	RadiantSeriesWins int `xml:"radiant_series_wins" gorethink:"radiant_series_wins,omitempty"`
	RadiantTeam       struct {
		Complete bool   `xml:"complete" gorethink:"complete,omitempty"`
		TeamID   int    `xml:"team_id" gorethink:"team_id,omitempty"`
		TeamLogo int    `xml:"team_logo" gorethink:"team_logo,omitempty"`
		TeamName string `xml:"team_name" gorethink:"team_name,omitempty"`
	} `xml:"radiant_team" gorethink:"radiant_team,omitempty"`
	DireTeam struct {
		Complete bool   `xml:"complete" gorethink:"complete,omitempty"`
		TeamID   int    `xml:"team_id" gorethink:"team_id,omitempty"`
		TeamLogo int    `xml:"team_logo" gorethink:"team_logo,omitempty"`
		TeamName string `xml:"team_name" gorethink:"team_name,omitempty"`
	} `xml:"dire_team" gorethink:"dire_team,omitempty"`
	Scoreboard   Scoreboard `xml:"scoreboard" gorethink:"scoreboard,omitempty"`
	SeriesID     int        `xml:"series_id" gorethink:"series_id,omitempty"`
	SeriesType   int        `xml:"series_type" gorethink:"series_type,omitempty"`
	Spectators   int        `xml:"spectators" gorethink:"spectators,omitempty"`
	StageName    string     `xml:"stage_name" gorethink:"stage_name,omitempty"`
	StreamDelayS float64    `xml:"stream_delay_s" gorethink:"stream_delay_s,omitempty"`
	Timestamp    time.Time  `xml:"-" gorethink:"timestamp"`
}

type Scoreboard struct {
	Dire               Side    `xml:"dire" gorethink:"dire,omitempty"`
	Duration           float64 `xml:"duration" gorethink:"duration,omitempty"`
	Radiant            Side    `xml:"radiant" gorethink:"radiant,omitempty"`
	RoshanRespawnTimer int     `xml:"roshan_respawn_timer" gorethink:"roshan_respawn_timer,omitempty"`
}

type Side struct {
	BarracksState int         `xml:"barracks_state" gorethink:"barracks_state,omitempty"`
	Players       []Player    `xml:"players>player" gorethink:"players,omitempty"`
	Score         int         `xml:"score" gorethink:"score,omitempty"`
	TowerState    int         `xml:"tower_state" gorethink:"tower_state,omitempty"`
	Picks         []HeroID    `xml:"picks>pick" gorethink:"picks,omitempty"`
	Bans          []HeroID    `xml:"bans>ban" gorethink:"bans,omitempty"`
	Abilities     []Abilities `xml:"abilities>ability" gorethink:"abilities,omitempty"`
}

type HeroID struct {
	HeroID int `xml:"hero_id" gorethink:"hero_id,omitempty"`
}

type Abilities struct {
	AbilityID    int `xml:"ability_id" gorethink:"ability_id,omitempty"`
	AbilityLevel int `xml:"ability_level" gorethink:"ability_level,omitempty"`
}

type Player struct {
	AccountID        int     `xml:"account_id" gorethink:"account_id,omitempty"`
	Assists          int     `xml:"assists" gorethink:"assists,omitempty"`
	Death            int     `xml:"death" gorethink:"death,omitempty"`
	Denies           int     `xml:"denies" gorethink:"denies,omitempty"`
	Gold             int     `xml:"gold" gorethink:"gold,omitempty"`
	GoldPerMin       int     `xml:"gold_per_min" gorethink:"gold_per_min,omitempty"`
	HeroID           int     `xml:"hero_id" gorethink:"hero_id,omitempty"`
	Item0            int     `xml:"item0" gorethink:"item0,omitempty"`
	Item1            int     `xml:"item1" gorethink:"item1,omitempty"`
	Item2            int     `xml:"item2" gorethink:"item2,omitempty"`
	Item3            int     `xml:"item3" gorethink:"item3,omitempty"`
	Item4            int     `xml:"item4" gorethink:"item4,omitempty"`
	Item5            int     `xml:"item5" gorethink:"item5,omitempty"`
	Kills            int     `xml:"kills" gorethink:"kills,omitempty"`
	LastHits         int     `xml:"last_hits" gorethink:"last_hits,omitempty"`
	Level            int     `xml:"level" gorethink:"level,omitempty"`
	NetWorth         int     `xml:"net_worth" gorethink:"net_worth,omitempty"`
	PlayerSlot       int     `xml:"player_slot" gorethink:"player_slot,omitempty"`
	PositionX        float64 `xml:"position_x" gorethink:"position_x,omitempty"`
	PositionY        float64 `xml:"position_y" gorethink:"position_y,omitempty"`
	RespawnTimer     int     `xml:"respawn_timer" gorethink:"respawn_timer,omitempty"`
	UltimateCooldown int     `xml:"ultimate_cooldown" gorethink:"ultimate_cooldown,omitempty"`
	UltimateState    int     `xml:"ultimate_state" gorethink:"ultimate_state,omitempty"`
	XpPerMin         int     `xml:"xp_per_min" gorethink:"xp_per_min,omitempty"`
}
