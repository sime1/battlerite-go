package battlerite

type Roster struct{
	Id string `jsonapi:"primary,roster"`
	ShardId string `jsonapi:"attr,shardId"`
	Stats Stats `jsonapi:"attr,stats"`
	Participants []*Participant `jsonapi:"relation,participants"`
}
