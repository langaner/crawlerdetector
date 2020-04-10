package crawlerdetector

import (
	"regexp"
	"strings"
)

var (
	crawlerReg   = regexp.MustCompile(CombineRegexp(CrawlersList()))
	exclusionReg = regexp.MustCompile(CombineRegexp(ExclusionsList()))
)

// CrawlerDetector is crawler detector structure
type CrawlerDetector struct {
	Matched []string
}

// New returns a new initialized CrawlerDetector
func New() *CrawlerDetector {
	return &CrawlerDetector{Matched: []string{}}
}

// IsCrawler is detect crawlers/spiders/bots by user agent
func (cd *CrawlerDetector) IsCrawler(userAgent string) bool {
	if cd.IsExclusion(userAgent) {
		return false
	}

	cd.Matched = crawlerReg.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		return true
	}

	return false
}

// IsExclusion is detect exclusion from user agent
func (cd *CrawlerDetector) IsExclusion(userAgent string) bool {
	isExclusion := exclusionReg.ReplaceAllString(userAgent, "")

	if len(isExclusion) == 0 {
		return true
	}

	return false
}

// CombineRegexp is build regex from givement patterns list
func CombineRegexp(patterns []string) string {
	return "(" + strings.Join(patterns, "|") + ")"
}

// SetCrawlers is setter for custom crawlers list
func (cd *CrawlerDetector) SetCrawlers(list []string) *CrawlerDetector {
	crawlerReg = regexp.MustCompile(CombineRegexp(list))
	return cd
}

// SetExclusions is setter for custom exclusions list
func (cd *CrawlerDetector) SetExclusions(list []string) *CrawlerDetector {
	exclusionReg = regexp.MustCompile(CombineRegexp(list))
	return cd
}

// GetMatched is getter of matched result
func (cd *CrawlerDetector) GetMatched() []string {
	return cd.Matched
}
