package battlerite

type Asset struct{
	Id string `jsonapi:"primary,asset"`
	URL string `jsonapi:"attr,URL"`
	CreatedAt string `jsonapi:"attr,createdAt"`
	Description string `jsonapi:"attr,description"`
	Name string `jsonapi:"attr,name"`
}

