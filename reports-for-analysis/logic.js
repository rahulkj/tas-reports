function getBuildpackChartData() {
    let buildpackRelationData = {"name": data.envName, "children": []}

    for (const org of data.orgs) {

        let orgData = {name: org.orgName, children: []}

        for (const space of org.spaces) {
            let spaceData = {"name": space.spaceName, "children": []}

            for (const spaceApp of space.apps) {
                let buildpackName = spaceApp.entity.detected_buildpack

                let appName = spaceApp.entity.name

                let appData = {name: appName, value: spaceApp.entity.instances}

                let foundAndUpdatedBuildpack = false

                for (const buildpack of spaceData.children) {
                    if (buildpack.name === buildpackName) {
                        buildpack.children.push(appData)
                        foundAndUpdatedBuildpack = true
                    }
                }

                if (!foundAndUpdatedBuildpack) {
                    let buildpackData = {"name": buildpackName, children: []}
                    buildpackData.children.push(appData)
                    spaceData.children.push(buildpackData)
                }
            }

            orgData.children.push(spaceData)
        }

        buildpackRelationData.children.push(orgData)

    }

    return buildpackRelationData
}

function getAppResourceData() {
    let renderData = {apps: [], instances: [], memoryUsage: [], diskQuota: []}

    for (const org of data.orgs) {

        let orgData = {name: org.orgName, children: []}

        for (const space of org.spaces) {
            let spaceData = {"name": space.spaceName, "children": []}

            for (const spaceApp of space.apps) {
                renderData.apps.push(spaceApp.entity.name)
                renderData.instances.push(spaceApp.entity.instances)
                renderData.memoryUsage.push(spaceApp.entity.memory/1024)
                renderData.diskQuota.push(spaceApp.entity.disk_quota/1024)
            }
        }

    }

    return renderData
}