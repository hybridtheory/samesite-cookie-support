package samesite

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsSameSiteCookieSupported", func() {

	userAgentChecker := func(userAgent string, isSupported bool) {
		Expect(IsSameSiteCookieSupported(userAgent)).To(Equal(isSupported))
	}

	DescribeTable("for UCBrowser", userAgentChecker,
		Entry("does not support v11", "Mozilla/5.0 UCBrowser/11.4.5426.1034 Safari/537.36", false),
		Entry("does not support v12.12", "Mozilla/5.0 UCBrowser/12.12.5426.1034 Safari/537.36", false),
		Entry("does not support v12.13.1", "Mozilla/5.0 UCBrowser/12.13.1.1034 Safari/537.36", false),
		Entry("supports v12.13.2", "Mozilla/5.0 UCBrowser/12.13.2 Safari/537.36", true),
		Entry("supports v12.13.3", "Mozilla/5.0 UCBrowser/12.13.3 Safari/537.36", true),
	)

	DescribeTable("for iOS", userAgentChecker,
		Entry("supports v13", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", true),
		Entry("supports v11", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", true),
		Entry("does not support v12", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", false),
		Entry("does not support v12 minimal user agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) Safari/605.1.15", false),
	)

	Describe("for Mac OSX", func() {
		DescribeTable("for Safari", userAgentChecker,
			Entry("supports empty OS X version", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) Version/4.0 Safari/419.3", true),
			Entry("supports v10.5.6", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) Version/4.0 Safari/528.16", true),
			Entry("supports v10.12.6", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) Version/11.1 Safari/605.1.15", true),
			Entry("does not support v10.14", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) Version/11.1 Safari/605.1.15", false),
			Entry("does not support v10.14.7", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) Version/11.1 Safari/605.1.15", false),
		)

		DescribeTable("for Embedded browsers on Mac OSX version 10.14", userAgentChecker,
			Entry("supports empty empty OS X version", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fi-fi) AppleWebKit/420+ (KHTML, like Gecko)", true),
			Entry("supports v10.5.6", "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) AppleWebKit/528.16 (KHTML, like Gecko)", true),
			Entry("supports v10.12.6", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/605.1.15 (KHTML, like Gecko)", true),
			Entry("does not support v10.14", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) AppleWebKit/605.1.15 (KHTML, like Gecko)", false),
			Entry("does not support v10.14.7", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) AppleWebKit/605.1.15 (KHTML, like Gecko)", false),
		)
	})

	DescribeTable("for Chromium based", userAgentChecker,
		Entry("supports v50", "Mozilla/5.0 Ubuntu/10.04 Chromium/50.0.813.0", true),
		Entry("does not support v51", "Mozilla/5.0 Ubuntu/10.04 Chromium/51.0.813.0", false),
		Entry("does not support Chromium v51", "Mozilla/5.0 (X11; Linux i686) Ubuntu/10.04 Chromium/51.0.813.0 Safari/535.1", false),
		Entry("does not support Chrome v51", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chrome/51.0.154.53 Safari/525.19", false),
		Entry("supports Chromiumnot v51", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromiumnot/51.0.813.0 Safari/525.19", true),
		Entry("supports Chromenot v51", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chromenot/51.0.154.53 Safari/525.19", true),
		Entry("does not support v52", "Mozilla/5.0 Ubuntu/10.04 Chromium/52.0.813.0", false),
		Entry("does not support v66", "Mozilla/5.0 Ubuntu/10.04 Chromium/66.0.813.0", false),
		Entry("supports v67", "Mozilla/5.0 Ubuntu/10.04 Chromium/67.0.813.0", true),
		Entry("supports v68", "Mozilla/5.0 Ubuntu/10.04 Chromium/68.0.813.0", true),
	)

	DescribeTable("for other user agents", userAgentChecker,
		Entry("supports Firefox 71.0 on Windows 10", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0", true),
		Entry("supports Firefox 71.0 on Linux", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:71.0) Gecko/20100101 Firefox/71.0", true),
		Entry("supports Firefox 71.0 on MacOSX 10.15", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:71.0) Gecko/20100101 Firefox/71.0", true),
		Entry("supports Firefox 68.0 on Windows 10", "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0", true),
		Entry("supports Firefox 71.0 on Windows 7", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0", true),
		Entry("supports Firefox on MacOSX 12.1", "Mozilla/5.0 (Macintosh; CPU Intel Mac OS X 12_1) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", true),
		Entry("supports Embedded Webkit on Windows 7", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko)", true),
		Entry("supports Firefox on MacOSX v10.14", "(Macintosh; Intel Mac OS X 10.14; rv:58.0) Gecko/20100101 Firefox/58.0", true),
	)
})
