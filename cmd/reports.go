package cmd

import (
	"flag"
	"os"
)

func ParseManifests(reportsDirectory string, env string) {

	if reportsDirectory == "" || env == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	createOrgSpaceAppReport(reportsDirectory, env)
}
