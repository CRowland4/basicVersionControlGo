package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	help = `These are SVCS commands:
config     Get and set a username.
add        Add a file to the index.
log        Show commit logs.
commit     Save changes.
checkout   Restore a file.`
	log      = "Show commit logs."
	commit   = "Save changes."
	checkout = "Restore a file."
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(help)
		return
	}

	createFiles()
	writeDefaultValues()
	// TODO try reading form input rather than from CLI
	command := os.Args[1]
	switch command {
	case "--help":
		fmt.Println(help)
	case "config":
		configCommand(os.Args)
	case "add":
		addCommand(os.Args)
	case "log":
		fmt.Println(log)
	case "commit":
		fmt.Println(commit)
	case "checkout":
		fmt.Println(checkout)
	default:
		fmt.Printf("'%s' is not a SVCS command.", command)
	}
}

func addCommand(args []string) {
	trackedFiles := getTrackedFiles()
	if len(args) == 2 && len(trackedFiles) == 0 { // add, with no currently tracked files
		fmt.Println("Add a file to the index.")
	} else if len(args) == 2 { // add, with currently tracked files
		printTrackedFiles(trackedFiles)
	} else if len(args) == 3 { // add, new file to be tracked
		addFileToIndex(args[2])
	}
}

func printTrackedFiles(files []string) {
	fmt.Println("Tracked files:")
	for _, file := range files {
		fmt.Println(file)
	}

	return
}

func getTrackedFiles() (trackedFiles []string) {
	file, _ := os.Open(".vcs/index.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "Tracked files:" {
			trackedFiles = append(trackedFiles, scanner.Text())
		}
	}

	return trackedFiles
}

func readConfigName() (name string) {
	file, _ := os.Open(".vcs/config.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if strings.Contains(fields[0], "name") || fields[1] != "" {
			name = fields[1]
			break
		}
	}

	return name
}

func addFileToIndex(fileName string) {
	file, err := os.OpenFile("./vcs/index.txt", os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Printf("Can't find '%s'.\n", fileName)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, fileName)
	fmt.Printf("The file '%s' is tracked\n", fileName)
	return
}

func writeDefaultValues() {
	writeDefaultConfigValues()
	writeDefaultIndexValues()
	return
}

func writeDefaultIndexValues() {
	file, _ := os.OpenFile("./vcs/index.txt", os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	fmt.Fprintln(file, "Tracked files:")
	return
}

func writeDefaultConfigValues() {
	file, _ := os.OpenFile("./vcs/config.txt", os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	fmt.Fprintln(file, "name: ")
	return
}

func configCommand(args []string) {
	name := readConfigName()
	if len(args) == 2 && strings.TrimSpace(name) == "" { // config, no new name passed
		fmt.Println("Please, tell me who you are.")
	} else if len(args) == 2 { // config, no new name passed, with a name already set in config.txt
		fmt.Println("The username is,", name)
	} else if len(args) == 3 { // config, new name passed
		updateConfigName(name)
	}

	return
}

func updateConfigName(name string) {
	file, _ := os.Create("./vcs/config.txt")
	defer file.Close()
	fmt.Fprintln(file, "name:", name)
	return
}

func createFiles() {
	os.Mkdir("./vcs", os.ModePerm)
	os.Create("./vcs/config.txt")
	os.Create("./vcs/index.txt")
	return
}
