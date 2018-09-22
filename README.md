## About Crawler Detector

Crawler Detector is a Golang package for detecting bots/crawlers/spiders via the user agent.

### Installation

Run `go run github.com/langaner/crawlerdetector`.

### Usage

```
detector := crawlerdetector.New()
isCrawler := detector.isCrawler(r.Header.Get("User-Agent"))

if isCrawler {
    // Do something
}

```

### Contributing

If you find a bot/spider/crawler user agent that Crawler Detector not detect, please submit a pull request with the regex pattern.