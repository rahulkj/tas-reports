package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/rahulkj/app-info/cmd"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type TASReport struct {
	EnvironmentName string `json:"envName"`
	Orgs            []Org  `json:"orgs"`
}

type Org struct {
	OrgName string  `json:"orgName"`
	Spaces  []Space `json:"spaces"`
}

type Space struct {
	SpaceName string                  `json:"spaceName"`
	Apps      []cmd.AppSearchResource `json:"apps"`
}

type Buildpack struct {
	BuildpackName string `json:"name"`
	Apps          []App  `json:"children"`
}

type App struct {
	AppName   string `json:"name"`
	Instances int    `json:"value"`
}

func createOrgSpaceAppReport(reportsDirectory string, env string) {
	var tasReport TASReport = TASReport{EnvironmentName: env}

	if _, err := os.Stat(reportsDirectory); err == nil {
		fmt.Println("Path to the directory is exists")
	} else if os.IsNotExist(err) {
		fmt.Println("Path to the directory is not accessible or does not exist")
		os.Exit(1)
	}

	orgDirs, err1 := os.ReadDir(reportsDirectory)

	if err1 != nil {
		log.Fatal(err1)
	}

	for _, od := range orgDirs {
		if od.IsDir() {
			fmt.Println(od.Name())
			org := Org{OrgName: od.Name()}

			spacesDirs, _ := os.ReadDir(reportsDirectory + "/" + od.Name())
			for _, sd := range spacesDirs {
				if sd.IsDir() {
					fmt.Println(sd.Name())

					space := Space{SpaceName: sd.Name()}

					appManifests, _ := os.ReadDir(reportsDirectory + "/" + od.Name() + "/" + sd.Name())

					for _, am := range appManifests {
						if strings.Contains(am.Name(), ".yml") {
							fmt.Println(am.Name())
							appManifestPath := reportsDirectory + "/" + od.Name() + "/" + sd.Name() + "/" + am.Name()
							data := GetRaw(appManifestPath)
							var appResource cmd.AppSearchResource

							err := yaml.Unmarshal(data, &appResource)
							if err != nil {
								fmt.Println("Error unmarshalling YAML:", err)
								return
							}

							fmt.Println(appResource)

							space.Apps = append(space.Apps, appResource)
						}
					}
					org.Spaces = append(org.Spaces, space)
				}
			}
			tasReport.Orgs = append(tasReport.Orgs, org)
		}
	}

	fileName := "org-space-app.js"
	tasReportData, _ := json.Marshal(tasReport)

	stringOutput := string(tasReportData)

	stringOutput = "let data = " + stringOutput

	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)

	reportsDir := currentDir + "/reports-for-analysis"

	WriteReport(reportsDir, fileName, []byte(stringOutput))
}
