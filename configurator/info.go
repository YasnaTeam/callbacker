package configurator

import (
	"fmt"
	"time"
)

var (
	AppVersion string
	BuildTime  string
	GitCommit  string
)

const cb = ` ▄████████ ▀█████████▄  
███    ███   ███    ███ 
███    █▀    ███    ███ 
███         ▄███▄▄▄██▀  
███        ▀▀███▀▀▀██▄  
███    █▄    ███    ██▄ 
███    ███   ███    ███ 
████████▀  ▄█████████▀ `

func PrintBuildInformation() {
	buildDate, _ := time.Parse(time.RFC3339, BuildTime)

	fmt.Println(cb)
	fmt.Printf(
		"Version:\t\t%s\n"+
			"Build Commit:\t\t%s\n"+
			"Build Date:\t\t%s\n",
		AppVersion,
		GitCommit,
		buildDate.UTC(),
	)
}
