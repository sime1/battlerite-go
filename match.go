package battlerite

import (
	"errors"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Match struct{
	Id string `jsonapi:"primary,match"`
	CreatedAt string `jsonapi:"attr,createdAt"`
	Duration int `jsonapi:"attr,duration"`
	GameMode string `jsonapi:"attr,gameMode"`
	PatchVersion string `jsonapi:"attr,patchVersion"`
	ShardId string `jsonapi:"attr,shardId"`
	TitleId string `jsonapi:"attr,titleId"`
	Stats Stats `jsonapi:"attr,stats"`
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