package battlerite

type Team struct{
	Id string `jsonapi:"primary,team"`
	Name string `jsonapi:"attr,name"`
	TitleId string `jsonapi:"attr,titleId"`
	ShardId string `jsonapi:"attr,shardId"`
	Assets []*Asset `jsonapi:"relatin,assets"`
}
