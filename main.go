package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gres"
	"time"
)

func main() {
	laun := launcher.New().
		Headless(false).
		Devtools(true)
	browser := rod.New().ControlURL(
		laun.MustLaunch(),
	).Timeout(3 * time.Minute).Trace(true).NoDefaultDevice().MustConnect().MustIncognito().MustIgnoreCertErrors(true)

	page, _ := stealth.Page(browser)

	page.MustNavigate("file:///" + GetCurrentAbPathByCaller() + "/template/index.html")
	iframe := page.MustElement("#tcaptcha_iframe").MustWaitVisible().MustFrame()
	iframe.MustElement("#slider").MustWaitVisible()

	pt := iframe.MustElement("#slider").MustShape().Box()

	iframe.Overlay(pt.X, pt.Y, float64(280), pt.Height, "iframe")
	iframe.Mouse.MustMove(pt.X+5, pt.Y+8)
	iframe.Mouse.MustClick("left")
	time.Sleep(3 * time.Second)
	iframe.Mouse.MustMove(pt.X+50, pt.Y)
	// iframe.Mouse.MustUp("left")

	time.Sleep(30 * time.Second)
}

func GetCurrentAbPathByCaller() string {
	if !gres.Contains(".") {
		if p, err := gfile.Search("."); err != nil {
			return ""
		} else {
			return p
		}
	}

	return ""
}
