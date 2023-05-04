package generates

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ModuleNameGet() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	if len(fields) != 2 || fields[0] != "module" {
		return "", errors.New("invalid go.mod file")
	}
	moduleName := fields[1]

	return moduleName, nil
}

func PackageIntall(link string) {
	cmd := exec.Command("go", "get", "-u", link)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(link, "Installed")
}

func PackageIntallList(links []string) {
	for _, v := range links {
		PackageIntall(v)
	}
}
