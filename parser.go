package main

import (
	"regexp"
	"strconv"
)

var (
	iosVersion12Regexp             *regexp.Regexp = regexp.MustCompile(`\(iP.+; CPU .*OS 12[_\d]*.*\) AppleWebKit\/`)
	macosxVersionTenFourteenRegexp *regexp.Regexp = regexp.MustCompile(`\(Macintosh;.*Mac OS X 10_14[_\d]*.*\)`)
	safariRegexp                   *regexp.Regexp = regexp.MustCompile(`Version\/.* Safari\/`)
	macEmbeddedBrowserRegexp       *regexp.Regexp = regexp.MustCompile(`^Mozilla\/[\.\d]+ \(Macintosh;.*Mac OS X [_\d]+\) AppleWebKit\/[\.\d]+ \(KHTML, like Gecko\)$`)
	chromiumBasedRegexp            *regexp.Regexp = regexp.MustCompile(`Chrom(e|ium)/(\d+)`)
	chromiumVersionRegexp          *regexp.Regexp = regexp.MustCompile(`Chrom[^ \/]+\/(\d+)[\.\d]*`)
	ucbrowserRegexp                *regexp.Regexp = regexp.MustCompile(`UCBrowser\/`)
	ucbrowserVersionRegexp         *regexp.Regexp = regexp.MustCompile(`UCBrowser\/(\d+)\.(\d+)\.(\d+)[\.\d]* `)
)

// Parser is an struct with the logic to extract information from an user agent string.
type Parser struct {
	userAgent string
}

// NewParser returns a parser based on custom regexp expressions.
func NewParser(userAgent string) *Parser {
	return &Parser{
		userAgent: userAgent,
	}
}

func (c *Parser) hasWebKitSameSiteBug() bool {
	return c.isIosVersionTwelve() || (c.isMacosxVersionTenFourteen() && (c.isSafari() || c.isMacEmbeddedBrowser()))
}

func (c *Parser) dropsUnrecognizedSameSiteCookies() bool {
	if c.isChromiumBased() {
		return c.isChromiumVersionBetweenFiftyOneandSixtySeven()
	}
	return c.isUcBrowser() && !c.isUcBrowserVersionAtLeastTwelveThirteenTwo()
}

func (c *Parser) isIosVersionTwelve() bool {
	return iosVersion12Regexp.MatchString(c.userAgent)
}

func (c *Parser) isMacosxVersionTenFourteen() bool {
	return macosxVersionTenFourteenRegexp.MatchString(c.userAgent)
}

func (c *Parser) isSafari() bool {
	return safariRegexp.MatchString(c.userAgent) && !c.isChromiumBased()
}

func (c *Parser) isMacEmbeddedBrowser() bool {
	return macEmbeddedBrowserRegexp.MatchString(c.userAgent)
}

func (c *Parser) isChromiumBased() bool {
	return chromiumBasedRegexp.MatchString(c.userAgent)
}

func (c *Parser) isChromiumVersionBetweenFiftyOneandSixtySeven() bool {
	res := chromiumVersionRegexp.FindAllStringSubmatch(c.userAgent, -1)
	if len(res) == 1 {
		version, _ := strconv.Atoi(res[0][1])
		return (version >= 51) && (version < 67)
	}
	return true
}

func (c *Parser) isUcBrowser() bool {
	return ucbrowserRegexp.MatchString(c.userAgent)
}

func (c *Parser) isUcBrowserVersionAtLeastTwelveThirteenTwo() bool {
	twelve := 12
	thirteen := 13
	two := 2
	res := ucbrowserVersionRegexp.FindAllStringSubmatch(c.userAgent, -1)
	if len(res) == 1 {
		majorVersion, _ := strconv.Atoi(res[0][1])
		if majorVersion != twelve {
			return majorVersion > twelve
		}

		minorVersion, _ := strconv.Atoi(res[0][2])
		if minorVersion != thirteen {
			return minorVersion > thirteen
		}

		buildVersion, _ := strconv.Atoi(res[0][3])
		return buildVersion >= two
	}
	return false
}
