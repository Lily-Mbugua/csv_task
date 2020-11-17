package main

import (
	"csv_task/menu"
	"csv_task/sylabs"
	"flag"
	"fmt"
)


var (
	version = "1.0.0" // version of tool
	singularityBuildsCSV string
	singularityBuilds    sylabs.SingularityBuilds
)

func init() {
	fmt.Printf("Sylabs Cloud Remote Builder Service v%s\n\n", version)
	flag.StringVar(&singularityBuildsCSV, "file", "stats.csv", "SingularityBuilds CSV file")
	flag.Parse()
	singularityBuilds = sylabs.Load(singularityBuildsCSV)
}
func main() {
	menu.Start(&singularityBuilds)
}
