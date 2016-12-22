package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// variables
var destinations []map[string]string

var sources = []string{}
var excludePatterns = []string{}

// Create wait group to make sure we can track when all is done
var wg sync.WaitGroup

func main() {
	homeDir, err := homedir.Dir()
	//
	// Add exclusions from file
	//
	exclusionsFile, err := os.Open(homeDir + "/.sync-remote-configs/exclusions")
	if err != nil {
		fmt.Println("Please create a exclusions file. See README")
	}
	defer exclusionsFile.Close()

	exclusionsScanner := bufio.NewScanner(exclusionsFile)
	for exclusionsScanner.Scan() {
		// For each line that isn't a comment
		if strings.HasPrefix(exclusionsScanner.Text(), "#") == false && len(strings.Trim(exclusionsScanner.Text(), " ")) != 0 {
			excludePatterns = append(excludePatterns, exclusionsScanner.Text())
		}
	}

	if err := exclusionsScanner.Err(); err != nil {
		fmt.Println("There is an error with your exclusions file")
		fmt.Println(err)
	}

	//
	// Add sources from file
	//
	sourcesFile, err := os.Open(homeDir + "/.sync-remote-configs/sources")
	if err != nil {
		fmt.Println("Please create a sources file. See README")
	}
	defer sourcesFile.Close()

	sourcesScanner := bufio.NewScanner(sourcesFile)
	for sourcesScanner.Scan() {
		// For each line that isn't a comment
		if strings.HasPrefix(sourcesScanner.Text(), "#") == false && len(strings.Trim(sourcesScanner.Text(), " ")) != 0 {
			sources = append(sources, sourcesScanner.Text())
		}
	}

	if err := sourcesScanner.Err(); err != nil {
		fmt.Println("There is an error with your sources file")
		fmt.Println(err)
	}

	//
	// Add destinations from file
	//
	destinationsFile, err := os.Open(homeDir + "/.sync-remote-configs/destinations")
	if err != nil {
		fmt.Println("Please create a destinations file. See README")
	}
	defer destinationsFile.Close()

	destinationsScanner := bufio.NewScanner(destinationsFile)
	fmt.Println("Destinations")
	fmt.Println("-------------------------------")
	for destinationsScanner.Scan() {
		// For each line that isn't a comment and contains the @ symbol, add it as a destination
		if strings.HasPrefix(destinationsScanner.Text(), "#") == false && strings.Contains(destinationsScanner.Text(), "@") {
			s := strings.Split(destinationsScanner.Text(), "@")
			add_dest(s[0], s[1])
			fmt.Println(destinationsScanner.Text())
		}
	}

	if err := destinationsScanner.Err(); err != nil {
		fmt.Println("There is an error with your .destinations file")
		fmt.Println(err)
	}

	fmt.Println("\r\nStarting sync")
	fmt.Println("-------------------------------")

	// Need to make this use go routines
	for _, dest := range destinations {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go run_rsync(dest)
	}

	//Show output as running
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			fmt.Printf(".")
		}
	}()

	// Wait for all Rsyncs to complete
	wg.Wait()
	ticker.Stop()
	fmt.Println("\n\nAll servers updated")
}

func add_dest(username string, remote_address string) {
	var m = make(map[string]string)
	m["username"] = username
	m["remote_address"] = remote_address

	destinations = append(destinations, m)
}

func run_rsync(dest map[string]string) {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	directory := "/" + dest["username"] + "/"
	// for username that is not root, prepend the home dir
	if dest["username"] != "root" {
		directory = "/home" + directory
	}
	var updatedLine = "\n\nupdated " + dest["username"] + "@" + dest["remote_address"] + "...\n"
	var ssh = dest["username"] + "@" + dest["remote_address"] + ":" + directory
	// Need to do rsync command here. How do I add multiple commands?

	cmdArgs := append([]string{"-rv", "--delete", "--update", "--progress", "--copy-links", "--delete-excluded"}, sources...)
	for _, excludePattern := range excludePatterns {
		cmdArgs = append(cmdArgs, "--exclude")
		cmdArgs = append(cmdArgs, excludePattern)
	}
	cmdArgs = append(cmdArgs, ssh)

	cmd := exec.Command("rsync", cmdArgs...)

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	out, outErr := cmd.Output()

	if outErr != nil {
		fmt.Printf(outErr.Error())
	}

	fmt.Printf(updatedLine + string(out))
}
