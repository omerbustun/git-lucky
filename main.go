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

type Config struct {
	APIToken string `json:"api_token"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var language string
  var apiToken string
	configPath := flag.String("config", "config.json", "Path to the configuration file")
  flag.StringVar(&apiToken, "api","","API token (optional)")
	flag.StringVar(&language, "lang", "", "Specify the programming language (e.g. Go, Python, JavaScript). If not specified, a random language will be chosen.")
	help := flag.Bool("h", false, "Display the help text.")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

  var config Config
  var err error

  if apiToken != "" {
    config.APIToken = apiToken
  } else {
    config, err = LoadConfig(*configPath)
  }

	if err != nil {
		fmt.Printf("Error loading configuration: %s\n", err)
		os.Exit(1)
	}

	if config.APIToken == "" {
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

	if config.APIToken != "" {
		req.Header.Set("Authorization", "token "+config.APIToken)
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
