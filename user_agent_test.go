package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("SameSite cookie", func() {

	Describe("do not set for browsers with webkit samesite bug", func() {

		checker := func(line string, expected bool) {
			parser := &uaParser{
				client: uapGoParser.Parse(line),
			}
			Expect(parser.hasWebKitSameSiteBug()).To(Equal(expected))
		}

		DescribeTable("does not set samesite cookie for iOS major version 12",
			checker,
			Entry("version > 12", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", false),
			Entry("version = 12", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", true),
			Entry("version < 12", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", false),
			Entry("not ios", "Mozilla/5.0 (Macintosh; CPU Intel Mac OS X 12_1) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", false),
		)

		DescribeTable("does not set samesite cookie for Mac OSX version 10.14 and safari",
			checker,
			Entry("version empty", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) Version/4.0 Safari/419.3", false),
			Entry("version = 10.5.6", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) Version/4.0 Safari/528.16", false),
			Entry("version = 10.12.6", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) Version/11.1 Safari/605.1.15", false),
			Entry("version = 10.14", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) Version/11.1 Safari/605.1.15", true),
			Entry("version = 10.14.7", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) Version/11.1 Safari/605.1.15", true),
			Entry("not safari", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:58.0) Gecko/20100101 Firefox/58.0", false),
			Entry("not macintosh", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) Version/5.0.4 Safari/533.20.27", false),
		)

		DescribeTable("does not set samesite cookie for Mac OSX version 10.14 and embedded browser",
			checker,
			Entry("version empty", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) AppleWebKit/420+ (KHTML, like Gecko)", false),
			Entry("version = 10.5.6", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) AppleWebKit/528.16 (KHTML, like Gecko)", false),
			Entry("version = 10.12.6", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/605.1.15 (KHTML, like Gecko)", false),
			Entry("version = 10.14", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) AppleWebKit/605.1.15 (KHTML, like Gecko)", true),
			Entry("version = 10.14.7", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) AppleWebKit/605.1.15 (KHTML, like Gecko)", true),
			Entry("not embedded browser", "(Macintosh; Intel Mac OS X 10.14; rv:58.0) Gecko/20100101 Firefox/58.0", false),
			Entry("not macintosh", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko)", false),
		)
	})

	Describe("do not set for browsers with webkit samesite bug", func() {

		checker := func(line string, expected bool) {
			parser := &uaParser{
				client: uapGoParser.Parse(line),
			}
			Expect(parser.dropsUnrecognizedSameSiteCookies()).To(Equal(expected))
		}

		DescribeTable("UC browser drops samesite cookies prior version 12.13.2",
			checker,
			Entry("version < 12", "Mozilla/5.0 UCBrowser/11.4.5426.1034 Safari/537.36", true),
			Entry("version < 12.13", "Mozilla/5.0 UCBrowser/12.12.5426.1034 Safari/537.36", true),
			Entry("version < 12.13.2", "Mozilla/5.0 UCBrowser/12.13.1.1034 Safari/537.36", true),
			Entry("version = 12.13.2", "Mozilla/5.0 UCBrowser/12.13.2 Safari/537.36", false),
			Entry("version > 12.13.2", "Mozilla/5.0 UCBrowser/12.13.3 Safari/537.36", false),
			Entry("is not UC browser", "Mozilla/5.0 UBrowser/12.13.3.1034 Safari/537.36", false),
		)

		DescribeTable("is chromium based and major version between 51 and 67",
			checker,
			Entry("chromium", "Mozilla/5.0 (X11; Linux i686) Ubuntu/10.04 Chromium/51.0.813.0 Safari/535.1", true),
			Entry("chrome", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chrome/51.0.154.53 Safari/525.19", true),
			Entry("not chromium", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromiumnot/51.0.813.0 Safari/525.19", false),
			Entry("not chrome", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromenot/51.0.154.53 Safari/525.19", false),
			Entry("version < 51", "Mozilla/5.0 Ubuntu/10.04 Chromium/50.0.813.0", false),
			Entry("version = 51", "Mozilla/5.0 Ubuntu/10.04 Chromium/51.0.813.0", true),
			Entry("version > 51", "Mozilla/5.0 Ubuntu/10.04 Chromium/52.0.813.0", true),
			Entry("version < 67", "Mozilla/5.0 Ubuntu/10.04 Chromium/66.0.813.0", true),
			Entry("version = 67", "Mozilla/5.0 Ubuntu/10.04 Chromium/67.0.813.0", false),
			Entry("version > 67", "Mozilla/5.0 Ubuntu/10.04 Chromium/68.0.813.0", false),
		)
	})

	Describe("exposed IsSameSiteCookieSupported function works as expected", func() {

		checker := func(expected bool) func(string) {
			return func(line string) {
				Expect(IsSameSiteCookieSupported(line)).To(Equal(expected))
			}
		}

		DescribeTable("returns true for supported browsers",
			checker(true),
			Entry("supported", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1 like Mac OS X) Safari/605.1.15"),
			Entry("supported", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_1 like Mac OS X) Safari/605.1.15"),
			Entry("supported", "Mozilla/5.0 (Macintosh; CPU Intel Mac OS X 12_1) Safari/605.1.15"),
			Entry("supported", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) Version/4.0 Safari/419.3"),
			Entry("supported", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) Version/4.0 Safari/528.16"),
			Entry("supported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) Version/11.1 Safari/605.1.15"),
			Entry("supported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:58.0) Gecko/20100101 Firefox/58.0"),
			Entry("supported", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) Version/5.0.4 Safari/533.20.27"),
			Entry("supported", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) AppleWebKit/420+ (KHTML, like Gecko)"),
			Entry("supported", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) AppleWebKit/528.16 (KHTML, like Gecko)"),
			Entry("supported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/605.1.15 (KHTML, like Gecko)"),
			Entry("supported", "(Macintosh; Intel Mac OS X 10.14; rv:58.0) Gecko/20100101 Firefox/58.0"),
			Entry("supported", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko)"),
			Entry("supported", "Mozilla/5.0 UCBrowser/12.13.2 Safari/537.36"),
			Entry("supported", "Mozilla/5.0 UCBrowser/12.13.3 Safari/537.36"),
			Entry("supported", "Mozilla/5.0 UBrowser/12.13.3.1034 Safari/537.36"),
			Entry("supported", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromiumnot/51.0.813.0 Safari/525.19"),
			Entry("supported", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromenot/51.0.154.53 Safari/525.19"),
			Entry("supported", "Mozilla/5.0 Ubuntu/10.04 Chromium/50.0.813.0"),
			Entry("supported", "Mozilla/5.0 Ubuntu/10.04 Chromium/67.0.813.0"),
			Entry("supported", "Mozilla/5.0 Ubuntu/10.04 Chromium/68.0.813.0"),
		)

		DescribeTable("returns false for unsupported browsers",
			checker(false),
			Entry("unsupported", "Mozilla/5.0 (X11; Linux i686) Ubuntu/10.04 Chromium/51.0.813.0 Safari/535.1"),
			Entry("unsupported", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) Safari/605.1.15"),
			Entry("unsupported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) Version/11.1 Safari/605.1.15"),
			Entry("unsupported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) Version/11.1 Safari/605.1.15"),
			Entry("unsupported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) AppleWebKit/605.1.15 (KHTML, like Gecko)"),
			Entry("unsupported", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) AppleWebKit/605.1.15 (KHTML, like Gecko)"),
			Entry("unsupported", "Mozilla/5.0 UCBrowser/11.4.5426.1034 Safari/537.36"),
			Entry("unsupported", "Mozilla/5.0 UCBrowser/12.12.5426.1034 Safari/537.36"),
			Entry("unsupported", "Mozilla/5.0 UCBrowser/12.13.1.1034 Safari/537.36"),
			Entry("unsupported", "Mozilla/5.0 (X11; Linux i686) Ubuntu/10.04 Chromium/51.0.813.0 Safari/535.1"),
			Entry("unsupported", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chrome/51.0.154.53 Safari/525.19"),
			Entry("unsupported", "Mozilla/5.0 Ubuntu/10.04 Chromium/51.0.813.0"),
			Entry("unsupported", "Mozilla/5.0 Ubuntu/10.04 Chromium/52.0.813.0"),
			Entry("unsupported", "Mozilla/5.0 Ubuntu/10.04 Chromium/66.0.813.0"),
		)
	})
})
