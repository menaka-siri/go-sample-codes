package clidemo

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func CliDemo() {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./functions/clidemo/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Print(line)
		}
	}
}