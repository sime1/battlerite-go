package battlerite

type Team struct{
	Id string `jsonapi:"primary,team"`
	Name string `jsonapi:"attr,name"`
	TitleId string `jsonapi:"attr,titleId"`
	Stats Stats `jsonapi:"attr,stats"`
	ShardId string `jsonapi:"attr,shardId"`
	Assets []*Asset `jsonapi:"relation,assets"`
}
