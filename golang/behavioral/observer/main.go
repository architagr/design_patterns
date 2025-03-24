package main

import (
	"observer/publisher"
	"observer/user"
)

func main() {
	var (
		tagCleanCode           = "Clean Code"
		tagSoftwareDevelopment = "Software Development"
	)
	publisher.AddNewHashtag(tagSoftwareDevelopment)
	publisher.AddNewHashtag(tagCleanCode)

	// Create users and subscribe them to hashtags
	user.InitUser("1", "Alice", "alice@example.com", "alice_slack", []string{tagSoftwareDevelopment}, nil)
	user.InitUser("2", "Bob", "bob@example.com", "bob_slack", []string{tagSoftwareDevelopment, tagCleanCode}, nil)
	user3 := user.InitUser("3", "Charlie", "charlie@example.com", "charlie_slack", []string{tagSoftwareDevelopment, tagCleanCode}, []string{"author1"})

	// Publish new articles under hashtags
	publisher.GetHashtag(tagSoftwareDevelopment).NewArticlePublished("article1", "author2")
	publisher.GetHashtag(tagSoftwareDevelopment).NewArticlePublished("article2", "author1") // Charlie blocked author1
	publisher.GetHashtag(tagCleanCode).NewArticlePublished("article3", "author1")

	// Charlie unsubscribes from "Clean Code" hashtag
	user3.RemoveHashtag(tagCleanCode)
	publisher.GetHashtag(tagCleanCode).NewArticlePublished("article4", "author2") // Charlie won't be notified
}
