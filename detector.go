package crawlerdetector

import (
	"regexp"
	"strings"
)

// CrawlerDetector is crawler detector structure
type CrawlerDetector struct {
	Crawlers   []string
	Exclusions []string
	Matched    []string
}

// New returns a new initialized CrawlerDetector
func New() *CrawlerDetector {
	return &CrawlerDetector{CrawlersList(), ExclusionsList(), []string{}}
}

// IsCrawler is detect crawlers/spiders/bots by user agent
func (cd *CrawlerDetector) IsCrawler(userAgent string) bool {
	if cd.IsExclusion(userAgent) {
		return false
	}

	cReg := regexp.MustCompile(cd.CombineRegexp(cd.Crawlers))
	cd.Matched = cReg.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		return true
	}

	return false
}

// IsExclusion is detect exclusion from user agent
func (cd *CrawlerDetector) IsExclusion(userAgent string) bool {
	eReg := regexp.MustCompile(cd.CombineRegexp(cd.Exclusions))
	isExclusion := eReg.ReplaceAllString(userAgent, "")

	if len(isExclusion) == 0 {
		return true
	}

	return false
}

// CombineRegexp is build regex from givement patterns list
func (cd *CrawlerDetector) CombineRegexp(patterns []string) string {
	return "(" + strings.Join(patterns, "|") + ")"
}

// SetCrawlers is setter for custom crawlers list
func (cd *CrawlerDetector) SetCrawlers(list []string) *CrawlerDetector {
	cd.Crawlers = list
	return cd
}

// SetExclusions is setter for custom exclusions list
func (cd *CrawlerDetector) SetExclusions(list []string) *CrawlerDetector {
	cd.Exclusions = list
	return cd
}

// GetMatched is getter of matched result
func (cd *CrawlerDetector) GetMatched() []string {
	return cd.Matched
}
