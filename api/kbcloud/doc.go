// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

// List of APIs:
//   - [URLCheckerApi.batchCheckURLConnectivity]
//   - [accountApi.createAccount]
//   - [accountApi.deleteAccount]
//   - [accountApi.getDSN]
//   - [accountApi.listAccounts]
//   - [accountApi.updateAccount]
//   - [accountApi.updateAccountPrivileges]
//   - [alertConfigApi.getAlertConfig]
//   - [alertConfigApi.setAlertConfig]
//   - [alertInhibitApi.createAlertInhibit]
//   - [alertInhibitApi.deleteAlertInhibit]
//   - [alertInhibitApi.getAlertInhibit]
//   - [alertInhibitApi.listAlertInhibits]
//   - [alertInhibitApi.patchAlertInhibit]
//   - [alertMetricsApi.listAlertMetrics]
//   - [alertObjectApi.listAlertObjects]
//   - [alertObjectApi.setAlertObjectStatus]
//   - [alertObjectApi.setAlertObjectsStatus]
//   - [alertReceiverApi.createAlertReceiver]
//   - [alertReceiverApi.deleteAlertReceiver]
//   - [alertReceiverApi.getAlertReceiver]
//   - [alertReceiverApi.listAlertReceivers]
//   - [alertReceiverApi.patchAlertReceiver]
//   - [alertRuleApi.createAlertRule]
//   - [alertRuleApi.deleteAlertRule]
//   - [alertRuleApi.getAlertRule]
//   - [alertRuleApi.getRuleConfig]
//   - [alertRuleApi.listAlertRules]
//   - [alertRuleApi.updateAlertRule]
//   - [alertStrategyApi.createAlertStrategy]
//   - [alertStrategyApi.deleteAlertStrategy]
//   - [alertStrategyApi.listAlertStrategies]
//   - [alertStrategyApi.patchAlertStrategy]
//   - [alertStrategyApi.updateAlertStrategy]
//   - [analyzeApi.analyzeBackup]
//   - [analyzeApi.analyzeClusterParam]
//   - [analyzeApi.analyzeClusterRestore]
//   - [analyzeApi.analyzeLogs]
//   - [analyzeApi.analyzeOps]
//   - [analyzeApi.analyzeParam]
//   - [analyzeApi.analyzeService]
//   - [analyzeApi.analyzeSlowLogs]
//   - [analyzeApi.analyzeView]
//   - [autohealingApi.getAutohealing]
//   - [backupApi.createClusterBackup]
//   - [backupApi.deleteBackup]
//   - [backupApi.downloadBackup]
//   - [backupApi.downloadMutipleBackups]
//   - [backupApi.getBackup]
//   - [backupApi.getBackupLog]
//   - [backupApi.getBackupStats]
//   - [backupApi.getClusterBackupPolicy]
//   - [backupApi.listBackups]
//   - [backupApi.patchBackupPolicy]
//   - [backupApi.viewBackup]
//   - [backupRepoApi.deleteBackupRepo]
//   - [backupRepoApi.getBackupRepo]
//   - [backupRepoApi.listBackupRepos]
//   - [backupRepoApi.updateBackupRepo]
//   - [benchmarkApi.createPgbench]
//   - [benchmarkApi.createSysbench]
//   - [benchmarkApi.createTpcc]
//   - [benchmarkApi.createYcsb]
//   - [benchmarkApi.deleteBenchmark]
//   - [benchmarkApi.getBenchmark]
//   - [benchmarkApi.listBenchmark]
//   - [classApi.createClass]
//   - [classApi.deleteClass]
//   - [classApi.listClasses]
//   - [classApi.patchClass]
//   - [clusterApi.createCluster]
//   - [clusterApi.deleteCluster]
//   - [clusterApi.describeClusterHaHistory]
//   - [clusterApi.getCluster]
//   - [clusterApi.getClusterByID]
//   - [clusterApi.getClusterInstanceLog]
//   - [clusterApi.getClusterManifest]
//   - [clusterApi.getInstacesMetrics]
//   - [clusterApi.listCluster]
//   - [clusterApi.listEndpoints]
//   - [clusterApi.listInstance]
//   - [clusterApi.patchCluster]
//   - [clusterAlertSwitchApi.getClusterAlertDisabled]
//   - [clusterAlertSwitchApi.setClusterAlertDisabled]
//   - [clusterLogApi.queryAuditLogs]
//   - [clusterLogApi.queryErrorLogs]
//   - [clusterLogApi.queryPodLogs]
//   - [clusterLogApi.queryRunningLogs]
//   - [clusterLogApi.querySlowLogs]
//   - [databaseApi.createDatabase]
//   - [databaseApi.deleteDatabase]
//   - [databaseApi.listDatabases]
//   - [disasterRecoveryApi.createDisasterRecovery]
//   - [disasterRecoveryApi.deleteDisasterRecovery]
//   - [disasterRecoveryApi.getDisasterRecoveryHistory]
//   - [disasterRecoveryApi.getDisasterRecoveryStatus]
//   - [disasterRecoveryApi.listDisasterRecovery]
//   - [disasterRecoveryApi.promoteDisasterRecovery]
//   - [dmsApi.AlterVolumes]
//   - [dmsApi.CreateVolumes]
//   - [dmsApi.DataExport]
//   - [dmsApi.DataImport]
//   - [dmsApi.DropVolumes]
//   - [dmsApi.GetObjectInfo]
//   - [dmsApi.GetTaskList]
//   - [dmsApi.GetTaskProgress]
//   - [dmsApi.ListObjectNamesByType]
//   - [dmsApi.ListObjectTypesInSchema]
//   - [dmsApi.ListVolumes]
//   - [dmsApi.SetDefaultVolumes]
//   - [dmsApi.alterParameter]
//   - [dmsApi.closeSessions]
//   - [dmsApi.createDataSourceV2]
//   - [dmsApi.deleteDataSourceV2]
//   - [dmsApi.generateDDL]
//   - [dmsApi.getDataSourceV2]
//   - [dmsApi.getSchemaList]
//   - [dmsApi.listDataSourceV2]
//   - [dmsApi.listParameters]
//   - [dmsApi.listQueryHistory]
//   - [dmsApi.listSessions]
//   - [dmsApi.query]
//   - [dmsApi.showData]
//   - [dmsApi.sqlExplain]
//   - [dmsApi.tenantParameterHistory]
//   - [dmsApi.testDataSourceV2]
//   - [dmsApi.updateDataSourceV2]
//   - [engineApi.engineActionInOrg]
//   - [engineApi.listEnginesInEnv]
//   - [engineApi.listEnginesInOrg]
//   - [engineLicenseApi.listEngineLicenses]
//   - [engineOptionApi.ListUpgradeableServiceVersion]
//   - [engineOptionApi.getEngineOption]
//   - [engineOptionApi.listEngineOptions]
//   - [environmentApi.getEnvironment]
//   - [environmentApi.listEnvNodeZone]
//   - [environmentApi.listEnvironment]
//   - [environmentApi.listNodeGroup]
//   - [eventApi.queryClusterEvents]
//   - [featureApi.listFeature]
//   - [featureApi.readFeature]
//   - [inspectionApi.createAutoInspection]
//   - [inspectionApi.createInspectionScript]
//   - [inspectionApi.deleteInspectionScript]
//   - [inspectionApi.listAutoInspection]
//   - [inspectionApi.listInspectionScripts]
//   - [inspectionApi.listInspections]
//   - [inspectionApi.updateAutoInspection]
//   - [inspectionApi.updateInspection]
//   - [inspectionApi.updateInspectionScript]
//   - [instanceTypesApi.createInstanceType]
//   - [instanceTypesApi.deleteInstanceType]
//   - [instanceTypesApi.getInstanceTypeById]
//   - [instanceTypesApi.getInstanceTypes]
//   - [instanceTypesApi.updateInstanceType]
//   - [invitationApi.acceptInvitation]
//   - [invitationApi.createInvitation]
//   - [invitationApi.deleteInvitation]
//   - [invitationApi.listInvitation]
//   - [invitationApi.readInvitation]
//   - [invitationApi.rejectInvitation]
//   - [invitationApi.resendInvitation]
//   - [ipWhitelistApi.createIPWhitelist]
//   - [ipWhitelistApi.listIPWhitelist]
//   - [ipWhitelistApi.updateIPWhitelist]
//   - [loadBalancerApi.getLoadBalancer]
//   - [markClusterApi.markClusterRestoreCompleted]
//   - [memberApi.addOrgMember]
//   - [memberApi.deleteOrgMember]
//   - [memberApi.listOrgMember]
//   - [memberApi.listOrgMemberPermission]
//   - [memberApi.patchOrgMember]
//   - [memberApi.readOrgMember]
//   - [metricsApi.queryClusterMetrics]
//   - [oceanbaseApi.getTenant]
//   - [oceanbaseApi.listTenants]
//   - [opsrequestApi.cancelOps]
//   - [opsrequestApi.clusterVolumeExpand]
//   - [opsrequestApi.customOps]
//   - [opsrequestApi.exposeCluster]
//   - [opsrequestApi.getOpsRequestStatus]
//   - [opsrequestApi.horizontalScaleCluster]
//   - [opsrequestApi.promoteCluster]
//   - [opsrequestApi.rebuildInstance]
//   - [opsrequestApi.reconfigureCluster]
//   - [opsrequestApi.restartCluster]
//   - [opsrequestApi.startCluster]
//   - [opsrequestApi.stopCluster]
//   - [opsrequestApi.updateClusterLicense]
//   - [opsrequestApi.upgradeCluster]
//   - [opsrequestApi.verticalScaleCluster]
//   - [organizationApi.batchUpdateOrgParameters]
//   - [organizationApi.createOrg]
//   - [organizationApi.freezeMember]
//   - [organizationApi.getOrgParameter]
//   - [organizationApi.listOrg]
//   - [organizationApi.listOrgParameters]
//   - [organizationApi.patchOrg]
//   - [organizationApi.patchOrgParameter]
//   - [organizationApi.readOrg]
//   - [organizationApi.unfreezeMember]
//   - [paramTplApi.createParamTpl]
//   - [paramTplApi.createParamTplFromCluster]
//   - [paramTplApi.deleteParamTpl]
//   - [paramTplApi.getClusterParamTpls]
//   - [paramTplApi.listParamTpl]
//   - [paramTplApi.patchParamTpl]
//   - [paramTplApi.readParamTpl]
//   - [parameterApi.listConfigurations]
//   - [parameterApi.listParameterSpecs]
//   - [parameterApi.listParametersHistory]
//   - [projectApi.listProjects]
//   - [providerApi.getCloudProvider]
//   - [providerApi.listCloudProviders]
//   - [recycleBinClusterApi.deleteRecycleBinCluster]
//   - [recycleBinClusterApi.getRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinCluster]
//   - [recycleBinClusterApi.restoreRecycleBinCluster]
//   - [regionApi.getRegion]
//   - [regionApi.listRegions]
//   - [relationApi.createRelation]
//   - [relationApi.deleteRelation]
//   - [relationApi.listAvailableClustersForRelation]
//   - [relationApi.listRelatedClusters]
//   - [restoreApi.GetRestoreLog]
//   - [restoreApi.deleteRestoreObject]
//   - [restoreApi.doRestore]
//   - [restoreApi.getRestoreTimeRange]
//   - [restoreApi.listClusterRestore]
//   - [restoreApi.listRestores]
//   - [restoreApi.restoreCluster]
//   - [roleApi.batchAddRolePermissions]
//   - [roleApi.batchRemoveRolePermissions]
//   - [roleApi.createRole]
//   - [roleApi.deleteRoleByName]
//   - [roleApi.getRoleByName]
//   - [roleApi.listPermissions]
//   - [roleApi.listRolePermissions]
//   - [roleApi.listRoles]
//   - [roleApi.updateRoleByName]
//   - [serviceVersionApi.ListServiceVersion]
//   - [sqlEditorApi.runSQLOnCluster]
//   - [statisticApi.alertStatistic]
//   - [storageClassApi.getStorageClassStats]
//   - [tagApi.createTag]
//   - [tagApi.deleteTags]
//   - [tagApi.getTags]
//   - [tagApi.listOrgTags]
//   - [tagApi.updateTag]
//   - [taskApi.queryClusterTaskDetail]
//   - [tlsApi.getTLSCertificate]
//   - [tlsApi.tlsSwitcher]
//   - [userApi.createUserApikey]
//   - [userApi.deleteApikey]
//   - [userApi.patchAPIkey]
//   - [userApi.patchUser]
//   - [userApi.phoneVerification]
//   - [userApi.readUser]
//   - [userApi.readUserApikeys]
//   - [userApi.sendVerificationEmail]
//   - [userApi.updateUserPassword]
//   - [viewApi.getTreeView]
//   - [viewApi.getViewByCluster]
//   - [whitelistApi.deleteIPWhiteList]
//   - [zoneApi.getZones]
//   - [zoneApi.listZones]
package kbcloud
