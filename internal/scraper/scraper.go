package scraper

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/burhanudinyahya/putmedown/internal/utils"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func TiktokScraper(url string) (string, string, error) {

	chromeManager := utils.GetChromeManager()
	allocCtx := chromeManager.GetContext()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*page.EventFrameSubtreeWillBeDetached); ok {
			log.Println("Ignoring frame detach event...")
		}
	})

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var videoSrc string
	var cookies []*network.Cookie

	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
		chromedp.WaitVisible(`video`, chromedp.ByQuery),
		chromedp.AttributeValue(`video > source`, "src", &videoSrc, nil),
	)
	if err != nil {
		log.Println(err)
	}

	cookieStrings := make([]string, len(cookies))
	for index, cookie := range cookies {
		cookieStrings[index] = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
	}

	cookiesJoined := strings.Join(cookieStrings, "; ")

	return videoSrc, cookiesJoined, nil
}
