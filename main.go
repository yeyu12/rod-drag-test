/**
 *
 * @Author: lenovo
 * @Email: lingchenyiduz@gmail.com
 * @QQ: 3334241893
 * @Date: 2022-08-17
 */

package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/util/gconv"
	"math/rand"
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

	pt := iframe.MustElement("#slider").MustShape().OnePointInside()

	x := page.MustEval(JsGetElementViewX(), "tcaptcha_transform").Str()
	y := page.MustEval(JsGetElementViewY(), "tcaptcha_transform").Str()
	ix := iframe.MustEval(JsGetElementViewX(), "sliderbg").Str()
	iy := iframe.MustEval(JsGetElementViewY(), "sliderbg").Str()
	startWidth := pt.X - 20
	startHeight := pt.Y - 20
	endWidth := 280
	endHeight := 40
	iframe.Overlay(startWidth, startHeight, float64(endWidth), float64(endHeight), "iframe")
	rand.Seed(time.Now().Unix())
	startX := gconv.Float64(x) + gconv.Float64(ix) + gconv.Float64(rand.Intn(30)+10)
	startY := gconv.Float64(y) + gconv.Float64(iy) + gconv.Float64(rand.Intn(30)+10)
	iframe.Mouse.MustMove(startX, startY)
	iframe.Mouse.MustClick("left")
	time.Sleep(3 * time.Second)
	iframe.Mouse.MustMove(startX+50, startY)
	iframe.Mouse.MustUp("left")

	time.Sleep(30 * time.Second)
}

func JsGetElementViewX() string {
	return `(path)=>document.getElementById(path).getBoundingClientRect().left + document.documentElement.scrollLeft`
}

func JsGetElementViewY() string {
	return `(path)=>document.getElementById(path).getBoundingClientRect().top + document.documentElement.scrollTop`
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
