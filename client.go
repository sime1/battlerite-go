//	Package battlerite is an SDK for the official Battlerite API written in Go.
package battlerite

import (
	"net/http"
	"github.com/google/jsonapi"
	"reflect"
	"errors"
	"strconv"
	"strings"
	"fmt"
)

const URL = "https://api.dc01.gamelockerapp.com/shards/global"

//Client is the interface of the API sdk. Provides methods to query the API for data.
type Client struct {
	//Yor API key
	Key string
}

func (client Client) getRequest(url string) *http.Request{
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Authorization", "Bearer " + client.Key)
	return req
}

//GetMatchList retrieves a list of matches from the API.
func (client Client) GetMatchList( page *Page, sort string, filter *MatchFilter) ([]*Match, error){
	url := URL + "/matches" + createMatchListParams(page, sort, filter)
	req := client.getRequest(url)
	httpClient := &http.Client{}
	res,_ := httpClient.Do(req)
	defer res.Body.Close()
	matches := make([]*Match, 0)
	tempMatches,err := jsonapi.UnmarshalManyPayload(res.Body, reflect.TypeOf(new(Match)));
	if err != nil {
		return nil, err
	}
	for _,m := range tempMatches {
		match, success := m.(*Match);
		if !success {
			return nil, errors.New("Typecast error: Match");
		}
		matches = append(matches,match)
	}
	return matches, nil
}

//GetMatch lets you retrieve a single match from the API.
//It takes the id as a parameter, and returns the Match. In case of success, error is nil, else
//it contains an error object with data about the error
func (client Client) GetMatch(id string) (*Match, error) {
	url := URL + "/matches/" + id
	httpClient := &http.Client{}
	req := client.getRequest(url)
	res,_ := httpClient.Do(req)
	defer res.Body.Close()
	match := &Match{}
	err := jsonapi.UnmarshalPayload(res.Body, match)
	if err != nil {
		return nil, err
	}
	return match, nil
}

//GetPlayerList retrieves a list od players from the API.
func (client Client) GetPlayerList( filter *PlayerFilter) ([]*Player, error){
	url := URL + "/players" + createPlayerListParams(filter)
	req := client.getRequest(url)
	httpClient := &http.Client{}
	res,_ := httpClient.Do(req)
	defer res.Body.Close()
	players := make([]*Player, 0)
	tempPlayers,err := jsonapi.UnmarshalManyPayload(res.Body, reflect.TypeOf(new(Player)));
	if err != nil {
		return nil, err
	}
	for _,p := range tempPlayers {
		match, success := p.(*Player);
		if !success {
			return nil, errors.New("Typecast error: Player");
		}
		players = append(players, match)
	}
	return players, nil
}

//GetPlayer lets you retrieve a single player from the API.
//It takes the id as a parameter, and returns the Player. In case of success, error is nil, else
//it contains an error object with data about the error
func (client Client) GetPlayer(id string) (*Player, error) {
	url := URL + "/players/" + id
	httpClient := &http.Client{}
	req := client.getRequest(url)
	res,_ := httpClient.Do(req)
	defer res.Body.Close()
	player := &Player{}
	err := jsonapi.UnmarshalPayload(res.Body, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

//GetTeamList retrieves a list of teams from the API.
func (client Client) GetTeamList( tag *TeamTag) ([]*Team, error){
	if tag == nil || len(tag.PlayerIds) < 1 {
		return nil, errors.New("Error: missing required parameter")
	}
	url := fmt.Sprintf("%s/teams?tag[playerIds]=%s&tag[season]=%d", URL, strings.Join(tag.PlayerIds, ","), tag.Season)
	req := client.getRequest(url)
	httpClient := &http.Client{}
	res,_ := httpClient.Do(req)
	defer res.Body.Close()
	teams := make([]*Team, 0)
	tempTeams,err := jsonapi.UnmarshalManyPayload(res.Body, reflect.TypeOf(new(Team)));
	if err != nil {
		return nil, err
	}
	for _,t := range tempTeams {
		team, success := t.(*Team);
		if !success {
			return nil, errors.New("Typecast error: Team");
		}
		teams = append(teams,team)
	}
	return teams, nil
}

//Page lets you specify pagination options for the GetMatchList() method
type Page struct{
	Offset int
	Limit int
}

//MatchFilter lets you decide how to filter the matches retrived by GetMatchList()
type MatchFilter struct{
	CreatedAtStart *string
	CreatedAtEnd *string
	PlayerIds []string
	PatchVersion []string
	ServerType []string
	RankingType *string
}

//PlayerFilter lets you decide which players GetPlayerList should retrieve
type PlayerFilter struct {
	PlayerNames []string
	PlayerIds []string
	SteamIds []string
}

//TeamTag provides GetTeamList the required parametes for querying the API
type TeamTag struct {
	PlayerIds []string
	Season int
}


func createMatchListParams(page *Page, sort string, filter *MatchFilter) string {
	urlParams := make([]string, 0)
	//page params
	if page != nil{
		if page.Offset != 0 {
			urlParams = append(urlParams, "page[offset]=" + strconv.Itoa(page.Offset))
		}
		if page.Limit != 0 {
			urlParams = append(urlParams, "page[limit]=" + strconv.Itoa(page.Limit))
		}
	}
	//sort param
	if len(sort) > 0 {
		urlParams = append(urlParams, "sort=" +  sort)
	}
	//filter params
	if filter != nil{
		if filter.CreatedAtStart != nil{
			urlParams = append(urlParams, "filter[createdAt-start]="+ *filter.CreatedAtStart)
		}
		if filter.CreatedAtEnd != nil{
			urlParams = append(urlParams, "filter[createdAt-end]=" +  *filter.CreatedAtEnd)
		}
		if len(filter.PlayerIds) > 0 {
			urlParams = append(urlParams, "filter[playerIds]=" + strings.Join(filter.PlayerIds, ","))
		}
		if len(filter.PatchVersion) > 0 {
			urlParams = append(urlParams, "filter[patchVersion]=" + strings.Join(filter.PatchVersion, ","))
		}
		if len(filter.ServerType) > 0 {
			urlParams = append(urlParams, "filter[serverType]=" + strings.Join(filter.ServerType, ","))
		}
		if filter.RankingType != nil {
			urlParams = append(urlParams, "filter[rankingType]=" + *filter.RankingType)
		}
	}
	if len(urlParams) > 0{
		return "?" + strings.Join(urlParams, "&")
	}
	return ""
}

func createPlayerListParams(filter *PlayerFilter) string {
	urlParams := make([]string, 0)
	//filter params
	if filter != nil{
		if len(filter.PlayerNames) > 0{
			urlParams = append(urlParams, "filter[playerNames]="+ strings.Join(filter.PlayerNames, ","))
		}
		if len(filter.SteamIds) > 0{
			urlParams = append(urlParams, "filter[steamIds]=" +  strings.Join(filter.SteamIds, ","))
		}
		if len(filter.PlayerIds) > 0{
			urlParams = append(urlParams, "filter[playerIds]=" + strings.Join(filter.PlayerIds, ","))
		}
	}
	if len(urlParams) > 0{
		return "?" + strings.Join(urlParams, "&")
	}
	return ""
}
