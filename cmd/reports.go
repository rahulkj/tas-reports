package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/rahulkj/app-info/cmd"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type TASReport struct {
	EnvironmentName string `json:"name"`
	Orgs            []Org  `json:"children"`
}

type Org struct {
	OrgName string  `json:"name"`
	Spaces  []Space `json:"children"`
}

type Space struct {
	SpaceName  string      `json:"name"`
	Buildpacks []Buildpack `json:"children"`
}

type Buildpack struct {
	BuildpackName string `json:"name"`
	Apps          []App  `json:"children"`
}

type App struct {
	AppName   string `json:"name"`
	AppGUID   string
	Instances int `json:"value"`
}

func ParseManifests(reportsDirectory string, env string) {

	if reportsDirectory == "" || env == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

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

							space = locateBuildpackAndAddApplication(space, appResource)
						}
					}
					org.Spaces = append(org.Spaces, space)
				}
			}
			tasReport.Orgs = append(tasReport.Orgs, org)
		}
	}

	fileName := "org-space-buildpack.json"
	tasReportData, _ := json.Marshal(tasReport)
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)

	reportsDir := currentDir + "/reports-for-analysis"
	
	WriteReport(reportsDir, fileName, tasReportData)
}

func locateBuildpackAndAddApplication(space Space, appSearchResource cmd.AppSearchResource) Space {
	buildpacksArray := space.Buildpacks

	var buildpack Buildpack
	buildpackName := appSearchResource.Entity.Buildpack

	app := App{AppName: appSearchResource.Entity.Name, AppGUID: appSearchResource.Metadata.AppGUID, Instances: appSearchResource.Entity.Instances}

	var foundAndUpdatedBuildpack bool
	var buildpackIndex int = -1

	for i, bp := range buildpacksArray {
		if bp.BuildpackName == buildpackName {
			buildpack = bp
			buildpack.Apps = append(buildpack.Apps, app)
			foundAndUpdatedBuildpack = true
			buildpackIndex = i
		}
	}

	if !foundAndUpdatedBuildpack {
		buildpack.BuildpackName = buildpackName
		buildpack.Apps = append(buildpack.Apps, app)
	}

	if buildpackIndex == -1 {
		space.Buildpacks = append(space.Buildpacks, buildpack)
	} else {
		space.Buildpacks[buildpackIndex] = buildpack
	}

	return space
}
