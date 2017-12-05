package battlerite

import (
	"net/http"
	"github.com/google/jsonapi"
	"reflect"
	"errors"
	"strconv"
	"strings"
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
func (client Client) GetMatchList( page *Page, sort string, filter *Filter) ([]*Match, error){
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
//It takes the id as a parameter, and return the Match. In case of success, error is nil, else
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

//Page lets you specify pagination options for the getMatchList() method
type Page struct{
	Offset *int
	Limit *int
}

//Filter lets you decide how to filter the matches retrived by getMatchList()
type Filter struct{
	CreatedAtStart *string
	CreatedAtEnd *string
	PlayerIds *[]string
	PlayerNames *[]string
	TeamNames *[]string
	GameMode *[]string
}

func createMatchListParams(page *Page, sort string, filter *Filter) string {
	urlParams := make([]string, 0)
	//page params
	if page != nil{
		if page.Offset != nil {
			urlParams = append(urlParams, "page[offset]=" + strconv.Itoa(*page.Offset))
		}
		if page.Limit != nil {
			urlParams = append(urlParams, "page[limit]=" + strconv.Itoa(*page.Limit))
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
		if filter.PlayerIds != nil{
			urlParams = append(urlParams, "filter[playerIds]=" + strings.Join(*filter.PlayerIds, ","))
		}
		if filter.PlayerNames != nil{
			urlParams = append(urlParams, "filter[playerNames]=" + strings.Join(*filter.PlayerNames, ","))
		}
		if filter.TeamNames != nil{
			urlParams = append(urlParams, "filter[teamNames]=" + strings.Join(*filter.TeamNames, ","))
		}
		if filter.GameMode != nil{
			urlParams = append(urlParams, "filter[gameMode]=" + strings.Join(*filter.GameMode, ","))
		}
	}
	if len(urlParams) > 0{
		return "?" + strings.Join(urlParams, "&")
	}
	return ""
}
