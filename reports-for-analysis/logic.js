function getBuildpackChartData() {
  let buildpackRelationData = { name: data.envName, children: [] };

  for (const org of data.orgs) {
    let orgData = { name: org.orgName, children: [] };

    for (const space of org.spaces) {
      let spaceData = { name: space.spaceName, children: [] };

      for (const spaceApp of space.apps) {
        let buildpackName = spaceApp.entity.detected_buildpack;

        let appName = spaceApp.entity.name;

        let appData = { name: appName, value: spaceApp.entity.instances };

        let foundAndUpdatedBuildpack = false;

        for (const buildpack of spaceData.children) {
          if (buildpack.name === buildpackName) {
            buildpack.children.push(appData);
            foundAndUpdatedBuildpack = true;
          }
        }

        if (!foundAndUpdatedBuildpack) {
          let buildpackData = { name: buildpackName, children: [] };
          buildpackData.children.push(appData);
          spaceData.children.push(buildpackData);
        }
      }

      orgData.children.push(spaceData);
    }

    buildpackRelationData.children.push(orgData);
  }

  return buildpackRelationData;
}

function getOrgSpaceAppsChartData() {
  let buildpackRelationData = { name: data.envName, children: [] };

  for (const org of data.orgs) {
    let orgData = { name: org.orgName, children: [] };

    for (const space of org.spaces) {
      let spaceData = { name: space.spaceName, children: [] };

      for (const spaceApp of space.apps) {
        let appName = spaceApp.entity.name;

        let appData = { name: appName, value: spaceApp.entity.instances };

        spaceData.children.push(appData);
      }

      orgData.children.push(spaceData);
    }

    buildpackRelationData.children.push(orgData);
  }

  return buildpackRelationData;
}

function getAppResourceData() {
  let renderData = { apps: [], instances: [], memoryUsage: [], diskQuota: [] };

  for (const org of data.orgs) {
    let orgData = { name: org.orgName, children: [] };

    for (const space of org.spaces) {
      let spaceData = { name: space.spaceName, children: [] };

      for (const spaceApp of space.apps) {
        renderData.apps.push(spaceApp.entity.name);
        renderData.instances.push(spaceApp.entity.instances);
        renderData.memoryUsage.push(spaceApp.entity.memory / 1024);
        renderData.diskQuota.push(spaceApp.entity.disk_quota / 1024);
      }
    }
  }

  return renderData;
}

function getStackAppData() {
  let stackAppData = [];

  const stackAppMap = new Map();

  for (const org of data.orgs) {
    for (const space of org.spaces) {
      for (const spaceApp of space.apps) {
        if (stackAppMap.has(spaceApp.entity.stack)) {
          let count = stackAppMap.get(spaceApp.entity.stack) + 1;
          stackAppMap.set(spaceApp.entity.stack, count);
        } else {
          stackAppMap.set(spaceApp.entity.stack, 1);
        }
      }
    }
  }

  stackAppMap.forEach(function (value, key) {
    let stackApp = {name: key, value: value}
    stackAppData.push(stackApp)
  });

  return stackAppData;
}

function getAppsPowerStateData() {
  let appsData = [];

  const powerStateMap = new Map();

  for (const org of data.orgs) {
    for (const space of org.spaces) {
      for (const spaceApp of space.apps) {
        if (powerStateMap.has(spaceApp.entity.state)) {
          let count = powerStateMap.get(spaceApp.entity.state) + 1;
          powerStateMap.set(spaceApp.entity.state, count);
        } else {
          powerStateMap.set(spaceApp.entity.state, 1);
        }
      }
    }
  }

  powerStateMap.forEach(function (value, key) {
    let appPowerStatus = {name: key, value: value}
    appsData.push(appPowerStatus)
  });

  return appsData;
}

function getOrgSpaceAppTableData() {
  let buildpackRelationData = [];

  for (const org of data.orgs) {
    for (const space of org.spaces) {
      for (const spaceApp of space.apps) {
        let data = []

        data.push(org.orgName);
        data.push(space.spaceName);
        data.push(spaceApp.entity.name);
        data.push(spaceApp.entity.instances);
        data.push(spaceApp.entity.state);
        data.push(spaceApp.entity.memory);
        data.push(spaceApp.entity.disk_quota);
        data.push(spaceApp.entity.detected_buildpack_filename);
        buildpackRelationData.push(data);
      }

    }
  }

  return buildpackRelationData;
}
