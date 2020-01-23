package main

// IsSameSiteCookieSupported returns true if through the user agent we have
// detected a browser that supports SameSite=None behaviour. It returns
// false otherwise.
func IsSameSiteCookieSupported(userAgent string) bool {
	parser := NewParser(userAgent)
	return !parser.dropsUnrecognizedSameSiteCookies() && !parser.hasWebKitSameSiteBug()
}
