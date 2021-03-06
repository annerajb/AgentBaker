// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package agent

import (
	"github.com/Azure/aks-engine/pkg/api"
)

func getBootstrappingCSE(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	if profile.IsWindows() {
		return "[concat('echo %DATE%,%TIME%,%COMPUTERNAME% && powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',parameters('kubernetesEndpoint'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -TargetEnvironment ',parameters('targetEnvironment'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('singleQuote'),variables('singleQuote'),base64(variables('servicePrincipalClientSecret')),variables('singleQuote'),variables('singleQuote'),' -NetworkAPIVersion ',variables('apiVersionNetwork'),' ',variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1 ; exit $LASTEXITCODE')]"
	} else {
		return ""
	}
}

func generateUserAssignedIdentityClientIDParameter(cs *api.ContainerService) string {
	if cs.Properties.OrchestratorProfile != nil &&
		cs.Properties.OrchestratorProfile.KubernetesConfig != nil &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedIDEnabled() {
		return "' USER_ASSIGNED_IDENTITY_ID=',reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID')), '2018-11-30').clientId, ' '"
	}
	return "' USER_ASSIGNED_IDENTITY_ID=',' '"
}
