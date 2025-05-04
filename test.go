package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const target string = "http://localhost:8000/"

func main() {
	// Create client with cookie jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Parse target URL
	targetUrl, _ := url.Parse(target)

	// Make request
	resp, err := client.Get(targetUrl.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	body, _ := io.ReadAll(resp.Body)

	// 1. Check cookies in the JAR (this is where they're stored)
	fmt.Println("\nCookies in jar:")
	for _, cookie := range jar.Cookies(targetUrl) {
		fmt.Printf("→ %s=%s\n", cookie.Name, cookie.Value)
	}

	// 2. Check response cookies (should be empty)
	fmt.Println("\nResponse cookies:")
	for _, cookie := range resp.Cookies() {
		fmt.Printf("→ %s=%s\n", cookie.Name, cookie.Value)
	}

	// 3. Check response headers (no Set-Cookie)
	fmt.Println("\nResponse headers:")
	for k, v := range resp.Header {
		fmt.Printf("→ %s: %v\n", k, v)
	}

	fmt.Println("\nResponse body:")
	fmt.Println(string(body))
}
