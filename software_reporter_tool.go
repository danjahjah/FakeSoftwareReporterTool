package main

import (
	"fmt" 
	"os"
	"time"
	"strings"
)

// this is a comment

func main() {
    fmt.Println("Hello! I am fake software_reporter_tool.exe ")
    currentTime := time.Now()

    f, err := os.OpenFile("swReporterTool.log",	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(currentTime.String()+ ":: called me with args:: "+strings.Join(os.Args, ", ")+" \n"); err != nil {
			fmt.Println(err)
		}
}