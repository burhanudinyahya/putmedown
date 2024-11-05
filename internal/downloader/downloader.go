package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	scaper "github.com/burhanudinyahya/putmedown/internal/scraper"
	"github.com/burhanudinyahya/putmedown/internal/utils"
)

func TiktokDownloader(url string) error {
	videoSrc, cookie, err := scaper.TiktokScraper(url)
	if err != nil {
		return fmt.Errorf("failed to scrap the video: %w", err)
	}

	filename, err := utils.FormatTikTokURL(url)
	if err != nil {
		return fmt.Errorf("failed to get filename: %w", err)
	}

	out, err := os.Create(filename + ".mp4")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	client := &http.Client{
		Timeout: 30 * time.Second, // Set a timeout to avoid hanging indefinitely
	}

	req, err := http.NewRequest("GET", videoSrc, nil)
	req.Header.Set("Cookie", cookie)
	if err != nil {
		return fmt.Errorf("failed to fetch the video: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download video: received status code %d", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save video: %w", err)
	}

	fmt.Printf("Successfully downloaded the video to %s\n", filename)
	return nil
}
