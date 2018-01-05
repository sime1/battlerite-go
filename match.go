package battlerite

import (
	"errors"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

//MatchTags represetns the possible tags a Match object can have. Due to Go JSON library restricitons, it has to be a map.
//Possible keys for this map are:
//	"rankingType": indicates the rankig type of the match ("RANKED", "UNRANKED" or "NONE" in case of private matches). Its value is a string
//	"serverType": indicates the type of
type MatchTags map[string]interface{}

type Match struct{
	Id string `jsonapi:"primary,match"`
	CreatedAt string `jsonapi:"attr,createdAt"`
	Duration int `jsonapi:"attr,duration"`
	GameMode string `jsonapi:"attr,gameMode"`
	PatchVersion string `jsonapi:"attr,patchVersion"`
	ShardId string `jsonapi:"attr,shardId"`
	TitleId string `jsonapi:"attr,titleId"`
	Stats Stats `jsonapi:"attr,stats"`
	Tags MatchTags `jsonapi:"attr,tags"`
	Assets []*Asset `jsonapi:"relation,assets"`
	Rosters []*Roster `jsonapi:"relation,rosters"`
	Rounds []*Round `jsonapi:"relation,rounds"`
	Spectators []*Participant `jsonapi:"relation,spectators"`
}

func (match *Match) GetTelemetry() (*[]Telemetry, error){
	if match == nil || len(match.Assets) == 0 {
		return nil, errors.New("Invalid match")
	}
	url := match.Assets[0].URL
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	telemetry := new([]Telemetry)
	if err := json.Unmarshal(rawData, telemetry); err != nil {
		return nil, err
	}
	return telemetry, nil
}