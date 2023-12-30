package main

import (
	"flag"
	"tas-reports/cmd"
)

func main() {

	environment := flag.String("e", "", "TAS environment name, like dev, non-prod, qa, prod, etc")
	reportsDirectory := flag.String("d", "", "absolute path to directory where the manifests from app-info cf plugin are available")
	flag.Parse()

	//if args[1] == "--reports-dir" || args[1] == "-d" {
	cmd.ParseManifests(*reportsDirectory, *environment)
	//} else {
	//	fmt.Printf("Invalid flags, please run help to see the valid options")
	//}
}
