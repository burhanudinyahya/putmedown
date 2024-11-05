package utils

import (
	"context"
	"sync"

	"github.com/chromedp/chromedp"
)

type ChromeManager struct {
	ctx    context.Context
	cancel context.CancelFunc
}

var instance *ChromeManager
var once sync.Once

func GetChromeManager() *ChromeManager {
	once.Do(func() {
		instance = &ChromeManager{}
		instance.Initialize()
	})
	return instance
}

func (c *ChromeManager) Initialize() {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("ignore-certificate-errors", true),
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	c.ctx = ctx
	c.cancel = cancel
}

func (c *ChromeManager) GetContext() context.Context {
	return c.ctx
}

func (c *ChromeManager) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}
