package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	version = "0.1.0"

	baseURL         = "https://courses.edx.org"
	loginEndpoint   = baseURL + "/api/user/v1/account/login_session/"
	outlineEndpoint = baseURL + "/api/course_home/v1/outline/"
)

// loop over an array of cookies looking for the specific `crsftoken` name
func retrieveCSRFToken(cookies []*http.Cookie) (string, error) {
	for _, c := range cookies {
		if c.Name == "csrftoken" {
			return c.Value, nil
		}
	}
	return "", fmt.Errorf("csrftoken cookie not found")
}

func main() {
	// step -1: parse command line arguments
	email := flag.String("email", "", "edX account email (*required)")
	password := flag.String("password", "", "edX account password (*required)")
	course_key := flag.String("course", "", "edX course key (*required)")

	flag.Parse()

	if *email == "" || *password == "" || *course_key == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(1)
	}

	// step 0: setup HTTP client
	jar, _ := cookiejar.New(&cookiejar.Options{publicsuffix.List})
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       15 * time.Second,
	}

	// step 1: retrieve CSRF token
	req, _ := http.NewRequest("GET", loginEndpoint, nil)
	req.Header.Add("User-Agent", "edxurl/"+version)

	resp, _ := client.Do(req)

	u, _ := url.Parse(baseURL)
	CSRFToken, err := retrieveCSRFToken(jar.Cookies(u))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("csrftoken: %v\n", CSRFToken)

	// step 2: authenticate
	payload := url.Values{
		"email":    {*email},
		"password": {*password},
	}
	req, _ = http.NewRequest("POST", loginEndpoint, strings.NewReader(payload.Encode()))
	req.Header.Add("User-Agent", "edxurl/"+version)
	req.Header.Add("X-CSRFToken", CSRFToken)
	req.Header.Add("Referer", baseURL+"/login")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, _ = client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var result struct {
		Success     bool   `json:"success"`
		RedirectURL string `json:"redirect_url"`
		Value       string `json:"value"`
		Email       string `json:"email"`
	}
	_ = json.Unmarshal(body, &result)

	if !result.Success {
		log.Println("could not authenticate")
		log.Fatalln(result)
	}
	log.Println("authentication successful")

	// step 3: get course outline
	req, _ = http.NewRequest("GET", outlineEndpoint+(*course_key), nil)
	req.Header.Add("User-Agent", "edxurl/"+version)

	resp, _ = client.Do(req)
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("%s\n", body)

	// Alas, this is where the journey ends...
	// apparently edX has multiple block types - from the Open edX API docs:
	// > type: (str) The type of block. Possible values the names of any
	// >     XBlock type in the system, including custom blocks. Examples are
	// >     course, chapter, sequential, vertical, html, problem, video, and
	// >     discussion.
	//
	// each of this type of block would require a different scraping algorithm, and even
	// then I could still miss some of the course's videos due to these "custom blocks".
}
