package store

import (
	"github.com/ncarlier/feedpushr/autogen/app"
)

// FeedRepository interface to manage feeds
type FeedRepository interface {
	ListFeeds(page, limit int) (*app.FeedCollection, error)
	ExistsFeed(url string) bool
	GetFeed(id string) (*app.Feed, error)
	DeleteFeed(id string) (*app.Feed, error)
	SaveFeed(feed *app.Feed) error
	ForEachFeed(cb func(*app.Feed) error) error
}
