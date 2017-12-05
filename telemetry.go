package battlerite

type Telemetry struct {
	Cursor int `json:"cursor"`
	Type string `json:"type"`
	Data map[string]interface{} `json:"dataObject"`
}

type RegionSamples struct{
	Region string `json:"region"`
	LatencyMS int `json:latencyMS`
}

type TelemetryData struct {
	Time int `json:"time"`
	AccountId string `json:"accountId"`
	MatchId string `json:"matchId"`
	ExternalMatchId string `json:"externalMatchId"`
	Round int `json:"round"`
	Character int `json:"character"`
	TypeId int `json:"typeId"`
	SourceTypeId int `json:"sourceTypeId"`
	ScoreType string `json:"scoreType"`
	Value int `json:"value"`
	TimeIntoRound int `json:"timeIntoRound"`
	UserId string `json:"userID"`
	Type string `json:"type"`
	ServerType string `json:"serverType"`
	CharacterLevel int `json:"characterLevel"`
	TeamId string `json:"teamId"`
	TotalTimePlayed int `json:"totalTimePlayed"`
	CharacterTimePlayed int `json:"characterTimePlayed"`
	Team int `json:"team"`
	RankingType string `json:"rankingType"`
	Mount int `json:"mount"`
	Attachment int `json:"mount"`
	Outfit int `json:"outfit"`
	Emote int `json:"emote"`
	League int `json:"league"`
	Division int `json:"division"`
	DivisionRating int `json:"divisionRationg"`
	SeasonId int `json:"seasonId"`
	SessionId string `json:"sessionId"`
	EventType string `json:"eventType"`
	TimeJoinedQueue string `json:"timeJoinedQueue"`
	TimeInQueue float64 `json:"timeInQueue"`
	CharacterArchetype int `json:"CharacterArchetype"`
	QueueTypes []string `json:"queueTypes"`
	LimitMatchMakingRange bool `json:"limitMatchMakingRange"`
	RegionSamples []RegionSamples `json:"regionSamples"`
	PreferredRegion string `json:"preferredRegion"`
	TeamSize int `json:"teamSize"`
	TeamMembers []int `json:"teamMembers"`
	PlacementGamesLeft int `json:"placementGamesLeft"`
	MatchRegion string `json:"matchRegion"`
	TeamSide int `json:"teamSide"`
	AutoMatchMaking bool `json:"autoMatchMaking"`
}