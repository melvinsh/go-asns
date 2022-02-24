package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	var matches []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		out := find(scanner.Text())

		if out != "" {
			matches = append(matches, out)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for _, match := range unique(matches) {
		fmt.Println(match)
	}
}

func find(ip string) string {
	var re = regexp.MustCompile(`(?m)(?:\n)(\d*)`)
	var str = whois(ip)

	out := re.FindStringSubmatch(str)[1]
	return out
}

func whois(ip string) string {
	command := fmt.Sprintf("whois -h v4.whois.cymru.com \"  -v %s\"", ip)

	out, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		log.Println(err)
	}

	return string(out[:])
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
