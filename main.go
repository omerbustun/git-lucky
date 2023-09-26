package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func getSnapConfig(key string) (string, error) {
	socketPath := os.Getenv("SNAPD_SOCKET")
	if socketPath == "" {
		socketPath = "/run/snapd.socket"
	}

	url := fmt.Sprintf("http://unix/v2/snaps/git-lucky/conf?keys=%s", key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Snapd-Socket-Path", socketPath)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	configs, ok := result["result"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response structure")
	}

	value, ok := configs[key].(string)
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return value, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var language string
	flag.StringVar(&language, "lang", "", "Specify the programming language (e.g. Go, Python, JavaScript). If not specified, a random language will be chosen.")
	help := flag.Bool("h", false, "Display the help text.")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	// Attempt to retrieve the API token from the snap configuration
	apiToken, _ := getSnapConfig("api-token")

	if apiToken == "" {
		fmt.Println("Warning: You're making unauthenticated requests to the GitHub API. Consider adding an API token to avoid rate limit issues.")
	}

	languages := []string{"Python", "JavaScript", "Ruby", "Go", "Java", "C++", "TypeScript", "PHP", "C#", "Swift", "Kotlin", "Rust", "R", "Scala", "Perl", "Objective-C", "Lua", "Shell", "Haskell", "Dart"}

	if language != "" {
		validLanguage := false
		for _, validLang := range languages {
			if validLang == language {
				validLanguage = true
				break
			}
		}

		if !validLanguage {
			fmt.Printf("Error: Invalid language specified. Please choose from: %v\n", languages)
			return
		}
	} else {
		language = languages[rand.Intn(len(languages))]
	}

	randomPage := rand.Intn(10) + 1

	apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=updated&order=desc&per_page=100&page=%d", language, randomPage)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if apiToken != "" {
		req.Header.Set("Authorization", "token "+apiToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	type Repository struct {
		Name    string `json:"name"`
		HTMLURL string `json:"html_url"`
	}

	var result struct {
		Items []Repository `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(result.Items) == 0 {
		fmt.Println("No repositories found!")
		return
	}

	randomRepo := result.Items[rand.Intn(len(result.Items))]

	fmt.Println(randomRepo.HTMLURL)

}
