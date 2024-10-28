package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var similarFlag bool

func init() {
	// Define the -s or --similar flag for similar URL deduplication
	flag.BoolVar(&similarFlag, "s", false, "Enable similar URL deduplication")
	flag.BoolVar(&similarFlag, "similar", false, "Enable similar URL deduplication")
}

func main() {
	flag.Parse()

	uniqueURLs := make(map[string]bool)
	// Refine the list to exclude specific media file extensions
	mediaExtensions := regexp.MustCompile(`(?i)\.(jpg|jpeg|png|gif|bmp|webp|mp4|mp3|avi|mov|svg|pdf|doc|xls)$`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputURL := scanner.Text()
		if inputURL == "" || mediaExtensions.MatchString(inputURL) {
			continue
		}

		parsedURL, err := url.Parse(inputURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid URL: %s\n", inputURL)
			continue
		}

		uniqueKey := generateKey(parsedURL)
		if _, exists := uniqueURLs[uniqueKey]; !exists {
			uniqueURLs[uniqueKey] = true
			fmt.Println(inputURL)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

// generateKey generates a unique key for each URL based on the deduplication mode
func generateKey(parsedURL *url.URL) string {
	if similarFlag {
		return similarPathKey(parsedURL)
	} else {
		return queryInsensitiveKey(parsedURL)
	}
}

// queryInsensitiveKey creates a key based on path and query parameters ignoring values
func queryInsensitiveKey(parsedURL *url.URL) string {
	queryKeys := []string{}
	for key := range parsedURL.Query() {
		queryKeys = append(queryKeys, key)
	}
	return parsedURL.Path + "?" + strings.Join(queryKeys, "&")
}

// similarPathKey creates a key using regex to identify similar paths (e.g., removing IDs)
func similarPathKey(parsedURL *url.URL) string {
	re := regexp.MustCompile(`/\d+`)
	sanitizedPath := re.ReplaceAllString(parsedURL.Path, "/id")
	return sanitizedPath
}

