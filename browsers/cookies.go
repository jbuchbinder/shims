package browsers

import (
	"net/http"

	"github.com/chromedp/cdproto/network"
)

// ImportChromeDPCookies imports a passed set of ChromeDP cookies into a
// http.Request object so that the simpler golang-native http.Request
// mechanism can be used to fetch data.
func ImportChromeDPCookies(req http.Request, cookies []*network.Cookie) {
	for _, c := range cookies {
		req.AddCookie(&http.Cookie{
			Name:     c.Name,
			Domain:   c.Domain,
			Value:    c.Value,
			Path:     c.Path,
			Secure:   c.Secure,
			HttpOnly: c.HTTPOnly,
			Expires:  TimestampFromFloat64(c.Expires).Time,
		})
	}
}
