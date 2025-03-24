package publisher

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyRegisterd = errors.New("already registered")
	ErrNotRegisterd     = errors.New("not registered")
)

// Observer Interface (Subscribers)
type IHashtagObserver interface {
	Notify(articleID, author string)
}

var allHashtags map[string]*hashtag

func init() {
	allHashtags = make(map[string]*hashtag)
}

func GetHashtag(name string) *hashtag {
	return allHashtags[name]
}
func AddNewHashtag(name string) {
	allHashtags[name] = &hashtag{
		name:      name,
		observers: make(map[string]IHashtagObserver),
		articles:  make(map[string]bool),
	}
}

// Hashtag acts as a subject (Publisher)
type hashtag struct {
	observers map[string]IHashtagObserver
	name      string
	articles  map[string]bool
}

// Register an observer (User)
func (tag *hashtag) Register(name string, observer IHashtagObserver) {
	tag.observers[name] = observer
}

// Deregister an observer (User)
func (tag *hashtag) Deregister(name string) {
	delete(tag.observers, name)
}

// Publish a new article
func (tag *hashtag) NewArticlePublished(articleID, author string) {
	fmt.Printf("\n📢 Notifying subscribers of #%s about new article (ID: %s) by %s\n", tag.name, articleID, author)
	tag.notify(articleID, author)
}

// Notify all observers
func (tag *hashtag) notify(articleID, author string) {
	for _, observer := range tag.observers {
		observer.Notify(articleID, author)
	}
}
