package user

import (
	"fmt"
	"observer/publisher"
)

// User struct acts as an Observer
type User struct {
	id            string
	name          string
	email         string
	slackUserID   string
	hashtags      map[string]bool
	blockedAuthor map[string]bool
}

// Add an author to the blocked list
func (u *User) AddBlockedAuthor(author string) {
	u.blockedAuthor[author] = true
}

// Remove an author from the blocked list
func (u *User) RemoveBlockedAuthor(author string) {
	delete(u.blockedAuthor, author)
}

// Subscribe to a hashtag
func (u *User) AddHashtag(tagName string) {
	u.hashtags[tagName] = true
	publisher.GetHashtag(tagName).Register(u.name, u)
}

// Unsubscribe from a hashtag
func (u *User) RemoveHashtag(tagName string) {
	delete(u.hashtags, tagName)
	publisher.GetHashtag(tagName).Deregister(u.name)
}

// Receive notifications for new articles
func (u *User) Notify(articleID, author string) {
	if u.blockedAuthor[author] {
		fmt.Printf("🚫 %s skipped notification for article (ID: %s) - blocked author: %s\n", u.name, articleID, author)
		return
	}
	fmt.Printf("\n📨 Sending notification to %s\n", u.name)
	if u.email != "" {
		fmt.Printf("📧 Email sent to %s for new article (ID: %s) by %s\n", u.email, articleID, author)
	}
	if u.slackUserID != "" {
		fmt.Printf("💬 Slack notification sent to %s for new article (ID: %s) by %s\n", u.slackUserID, articleID, author)
	}
}

// Initialize a new user
func InitUser(id, name, email, slackID string, hashtags, blockedAuthors []string) *User {
	u := &User{
		id:            id,
		name:          name,
		email:         email,
		slackUserID:   slackID,
		hashtags:      make(map[string]bool),
		blockedAuthor: make(map[string]bool),
	}
	for _, tag := range hashtags {
		u.AddHashtag(tag)
	}
	for _, author := range blockedAuthors {
		u.AddBlockedAuthor(author)
	}
	return u
}
