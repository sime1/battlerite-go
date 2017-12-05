package battlerite

type Participant struct{
	Id string `jsonapi:"primary,participant"`
	Actor string `jsonapi:"attr,actor"`
	ShardId string `jsonapi:"attr,shardId"`
	Stats Stats `jsonapi:"attr,stats"`
	Player *Player `jsonapi:"relation,player"`
}