package service

type Data struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

type DataCreate struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

type DataUpdate struct {
	URL  *string `json:"url"`
	Slug *string `json:"slug"`
}

/*
{
	"url": "new url"
}
var req Link

ShouldBindJSON(&req)

var link Link
link{url: "old url", slug: "slug"}

if req.url != nil {
	link.url = *url
}

if req.slug != nil {
	link.slug = slug
}

link{url: "new url", slug: "slug"}


{id 1    slug: google}
{id 2    slug: youtube}

id  2    slug: google
*/