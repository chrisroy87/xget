package main

import (
	flag "github.com/spf13/pflag"
)

var (
	flagVersion        bool
	flagHelp           bool
	flagDebug          bool
	flagVerbose        bool
	flagQuiet          bool
	flagShowProgress   bool
	flagSpider         bool
	flagDownload       bool
	flagFollowExternal bool
	flagRecursive      bool

	flagTries   int
	flagTimeout int
	flagWait    int

	flagUser        string
	flagPass        string
	flagAskUserPass bool
	flagAskPass     bool

	flagLimitRate string
	flagLogFile   string
	flagOutputDir string
	flagUserAgent string
)

func init() {
	flag.BoolVarP(&flagVersion, "version", "V", false, "Show program version and exit")
	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help message and exit")

	flag.BoolVarP(&flagDebug, "debug", "D", false, "Enable debug output")
	flag.BoolVarP(&flagVerbose, "verbose", "v", false, "Enable verbose output")
	flag.BoolVarP(&flagQuiet, "quiet", "q", false, "Suppress all output")

	flag.BoolVarP(&flagShowProgress, "progress", "", false, "Show progress bar")
	flag.BoolVarP(&flagSpider, "spider", "s", false, "Crawl URLs but do not download (spider mode)")
	flag.BoolVarP(&flagDownload, "download", "d", false, "Download files")
	flag.BoolVarP(&flagFollowExternal, "follow-external", "E", false, "Follow external links")
	flag.BoolVarP(&flagRecursive, "recursive", "r", false, "Enable recursive crawling")

	flag.IntVarP(&flagTries, "tries", "T", 10, "Maximum number of connection retries")
	flag.IntVarP(&flagTimeout, "timeout", "O", 10, "Download timeout in seconds")
	flag.IntVarP(&flagWait, "wait", "w", 0, "Seconds to wait between requests")

	flag.StringVarP(&flagUser, "user", "U", "", "Username for HTTP authentication")
	flag.StringVarP(&flagPass, "pass", "P", "", "Password for HTTP authentication")
	flag.BoolVarP(&flagAskUserPass, "ask-user-pass", "", false, "Prompt for username and password")
	flag.BoolVarP(&flagAskPass, "ask-pass", "", false, "Prompt for password")

	flag.StringVarP(&flagLimitRate, "rate", "L", "", "Limit download rate (e.g., 500K, 2M)")
	flag.StringVarP(&flagLogFile, "log", "g", "", "Log file path")
	flag.StringVarP(&flagOutputDir, "output", "o", "", "Directory to save downloads")
	flag.StringVarP(&flagUserAgent, "user-agent", "a", "", "Set custom User-Agent header")
}

func main() {
	flag.Parse()
}
