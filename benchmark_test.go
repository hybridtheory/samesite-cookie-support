package main

import (
	"github.com/montanaflynn/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type data struct {
	message   string
	userAgent string
	expected  bool
}

var _ = Describe("performance", func() {

	const (
		times = 100
	)

	var (
		testResults = map[string][]int64{}
		testData    = []data{
			data{message: "chrome", userAgent: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Chrome/51.0.154.53 Safari/525.19", expected: false},
			data{message: "ucbrowser", userAgent: "Mozilla/5.0 UCBrowser/12.12.5426.1034 Safari/537.36", expected: false},
			data{message: "iphone", userAgent: "iPhone; CPU iPhone OS 13_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/8.1.2 Mobile/15E148 Safari/605.1.15", expected: true},
			data{message: "safari", userAgent: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6; it-it) Version/4.0 Safari/528.16", expected: true},
			data{message: "ubuntu", userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:71.0) Gecko/20100101 Firefox/71.0", expected: true},
			data{message: "embedded browser", userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) AppleWebKit/605.1.15 (KHTML, like Gecko)", expected: false},
			data{message: "chrome version", userAgent: "Mozilla/5.0 Ubuntu/10.04 Chromium/67.0.813.0", expected: true},
		}
	)

	AfterSuite(func() {
		for testName, results := range testResults {
			p99, _ := stats.Percentile(stats.LoadRawData(results), 99)
			Expect(p99).Should(BeNumerically("<", 1000), testName+" shouldn't take too long.")
		}
	})

	var run = func(userAgent string, result bool) func(b Benchmarker) {
		return func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				output := IsSameSiteCookieSupported(userAgent)
				Expect(output).To(Equal(result))
			})
			testName := CurrentGinkgoTestDescription().TestText
			testResults[testName] = append(testResults[testName], runtime.Microseconds())
			b.RecordValue("spent in microseconds", float64(runtime.Microseconds()))
		}
	}

	Context("custom parser", func() {

		var runCustom = func(userAgent string, result bool) func(b Benchmarker) {
			return run(userAgent, result)
		}

		for _, data := range testData {
			testResults[data.message] = []int64{}
			Measure(data.message, runCustom(data.userAgent, data.expected), times)
		}
	})
})
