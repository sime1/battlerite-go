package battlerite

type Round struct{
	Id string `jsonapi:"primary,round"`
	Duration int `jsonapi:"attr,duration"`
	Ordinal int `jsonapi:"attr,ordinal"`
	Stats Stats `jsonapi:"attr,stats"`
}
