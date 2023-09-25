package main

import (
	"fmt"
	"net/url"
	"strings"
)

// getYoutubeURL prompts the user to enter a YouTube URL and returns it.
func getYoutubeURL() string {
	var youtubeURL string
	fmt.Print("Enter the YouTube URL: ")
	fmt.Scanln(&youtubeURL)
	return youtubeURL
}

// extractVideoID extracts the video ID from a YouTube URL.
func extractVideoID(youtubeURL string) (string, error) {
	parsedURL, err := url.Parse(youtubeURL)
	if err != nil {
		return "", fmt.Errorf("Invalid URL: %v", err)
	}

	switch parsedURL.Host {
	case "youtu.be":
		pathParts := strings.Split(parsedURL.Path, "/")
		if len(pathParts) < 2 {
			return "", fmt.Errorf("Invalid YouTube URL")
		}
		return pathParts[1], nil
	case "www.youtube.com":
		queryValues, ok := parsedURL.Query()["v"]
		if !ok || len(queryValues) == 0 {
			return "", fmt.Errorf("Invalid YouTube URL")
		}
		return queryValues[0], nil
	default:
		return "", fmt.Errorf("Unsupported YouTube URL")
	}
}

// generatePipedURL generates a Piped URL from a YouTube video ID.
func generatePipedURL(videoID string) string {
	return fmt.Sprintf("https://piped.video/watch?v=%s", videoID)
}

func main() {
	youtubeURL := getYoutubeURL()

	videoID, err := extractVideoID(youtubeURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	pipedURL := generatePipedURL(videoID)
	fmt.Println("Piped URL:")
	fmt.Println(pipedURL)
}
