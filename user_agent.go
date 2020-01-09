package models

import (
	"github.com/ua-parser/uap-go/uaparser"
	"strconv"
)

var uapGoParser *uaparser.Parser = uaparser.NewFromSaved()

// IsSameSiteCookieSupported returns true if through the user agent we have
// detected a browser that supports SameSite=None behaviour. It returns
// false otherwise.
func IsSameSiteCookieSupported(userAgent string) bool {
	parser := &uaParser{
		client: uapGoParser.Parse(userAgent),
	}
	return !parser.hasWebKitSameSiteBug() && !parser.dropsUnrecognizedSameSiteCookies()
}

// UserAgent contains a parsed user agent string to be able to perform checks on it.
type uaParser struct {
	client *uaparser.Client
}

func (p *uaParser) hasWebKitSameSiteBug() bool {
	return p.isIosVersion("12") || (p.isMacosxVersion("10", "14") && (p.isSafari() || p.isMacEmbeddedBrowser()))
}

func (p *uaParser) dropsUnrecognizedSameSiteCookies() bool {
	if p.isUcBrowser() {
		return !p.isUcBrowserVersionAtLeast(12, 13, 2)
	}
	if p.isChromiumBased() {
		return p.isChromiumVersionAtLeast(51) && !p.isChromiumVersionAtLeast(67)
	}
	return false
}

func (p *uaParser) isIosVersion(major string) bool {
	isIOS := p.client.Os.Family == "iOS"
	isMajorVersion := p.client.Os.Major == major
	return isIOS && isMajorVersion
}

func (p *uaParser) isMacosxVersion(major string, minor string) bool {
	isMacOSX := p.client.Os.Family == "Mac OS X"
	isMajorVersion := p.client.Os.Major == major
	isMinorVersion := p.client.Os.Minor == minor
	return isMacOSX && isMajorVersion && isMinorVersion
}

func (p *uaParser) isSafari() bool {
	return p.client.UserAgent.Family == "Safari"
}

func (p *uaParser) isMacEmbeddedBrowser() bool {
	return p.client.UserAgent.Family == "Apple Mail"
}

func (p *uaParser) isChromiumBased() bool {
	isChrome := p.client.UserAgent.Family == "Chrome"
	isChromium := p.client.UserAgent.Family == "Chromium"
	return isChrome || isChromium
}

func (p *uaParser) isChromiumVersionAtLeast(major int) bool {
	majorVersion, _ := strconv.Atoi(p.client.UserAgent.Major)
	return majorVersion >= major
}

func (p *uaParser) isUcBrowser() bool {
	return p.client.UserAgent.Family == "UC Browser"
}

func (p *uaParser) isUcBrowserVersionAtLeast(major int, minor int, build int) bool {
	majorVersion, _ := strconv.Atoi(p.client.UserAgent.Major)
	if majorVersion != major {
		return majorVersion > major
	}

	minorVersion, _ := strconv.Atoi(p.client.UserAgent.Minor)
	if minorVersion != minor {
		return minorVersion > minor
	}

	buildVersion, _ := strconv.Atoi(p.client.UserAgent.Patch)
	return buildVersion >= build
}
