package battlerite

type Player struct{
	Id string `jsonapi:"primary,player"`
	Name string `jsonapi:"attr,name"`
	PatchVersion string `jsonapi:"attr,patchVersion"`
	ShardId string `jsonapi:"attr,shardId"`
	Stats Stats `jsonapi:"attr,stats"`
	TitleId string `jsonapi:"attr,titleId"`
	Assets []*Asset `jsonapi:"relation,assets"`
}
