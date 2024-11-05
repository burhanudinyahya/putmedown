package utils

import (
	"fmt"
	"regexp"
)

func FormatTikTokURL(url string) (string, error) {
	re := regexp.MustCompile(`https://www\.tiktok\.com/@([a-zA-Z0-9_]+)/video/(\d+)`)

	matches := re.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", fmt.Errorf("invalid TikTok URL format")
	}

	username := matches[1]
	videoID := matches[2]

	return username + videoID, nil
}
