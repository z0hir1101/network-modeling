package web

import (
	"context"
	"log"
	"strings"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func extract_id(url string) string {
	parts := strings.Split(url, "/")

	for i, part := range parts {
		if part == "comments" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

func Parse_reddit(url string) *reddit.PostAndComments {
	post_id := extract_id(url)
	if post_id == "" {
		log.Fatalf("reddit_parse.go[extract_id]: Invalid reddit post\n")
	}

	post, _, err := reddit.DefaultClient().Post.Get(context.Background(), post_id)
	if err != nil {
		log.Fatalf("reddit_parse.go[Parse_post]: %v\n", err)
	}
	return post
}

func Find_posts(query string) []*reddit.Post {
	client := reddit.DefaultClient()
	posts, _, err := client.Subreddit.SearchPosts(context.Background(), query, "all", &reddit.ListPostSearchOptions{
		Sort: "relevance",
		ListPostOptions: reddit.ListPostOptions{
			ListOptions: reddit.ListOptions{Limit: 20},
		},
	})
	if err != nil {
		log.Fatalf("reddit_parse.go[Find_post]: %v\n", err)
	}

	return posts
}
