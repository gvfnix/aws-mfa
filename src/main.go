package main

import (
	"fmt"
	"flag"
	"time"
	// "log"
)


var (
	fromProfile = flag.String("from-profile", "", "AWS profile that is granted permissions to get temporary credentials")
	toProfile = flag.String("to-profile", "", "AWS profile that will be configured with temporary credentials")
	duration = flag.Duration("duration", 12*time.Hour, "Duration of the session token")
)

func main() {
	flag.Parse()
	fmt.Println(*fromProfile)
}