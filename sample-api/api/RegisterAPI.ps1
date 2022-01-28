param(
  [Parameter(Mandatory=$false,
  HelpMessage="The friendly name of the app registration")]
  [String]
  $AppName = "Kiota ToDoItem Sample API",

  [Parameter(Mandatory=$false,
  HelpMessage="Your Azure Active Directory tenant ID")]
  [String]
  $TenantId,

  [Parameter(Mandatory=$false)]
  [Switch]
  $StayConnected = $false
)

# API Scope descriptions
$AdminConsentDisplayName = "Full access to users' ToDoItems"
$AdminConsentDescription = "Allows the app to read and write users' ToDoItems"
$UserConsentDisplayName = "Full access to your ToDoItems"
$UserConsentDescription = "Allows the app to read and write your ToDoItems"
$ScopeName = "ToDoItem.ReadWrite"
$AppSettingsFile = "appsettings.Development.json"
$AppSettingsPath = "./src/" + $AppSettingsFile

# Requires an admin
if ($TenantId)
{
  Connect-MgGraph -Scopes "Application.ReadWrite.All User.Read" -TenantId $TenantId
}
else
{
  Connect-MgGraph -Scopes "Application.ReadWrite.All User.Read"
}

# Get context for access to tenant ID
$context = Get-MgContext

$scopeId = New-Guid
# Create app registration
$appRegistration = New-MgApplication -DisplayName $AppName -SignInAudience "AzureADMyOrg"
Update-MgApplication -ApplicationId $appRegistration.Id -IdentifierUris @("api://" + $appRegistration.AppId) `
 -Api @{ Oauth2PermissionScopes=@(@{ Id=$scopeId; AdminConsentDescription=$AdminConsentDescription; `
 AdminConsentDisplayName=$AdminConsentDisplayName; UserConsentDescription=$UserConsentDescription; `
 UserConsentDisplayName=$UserConsentDisplayName; Value=$ScopeName; IsEnabled=$true; Type="User" })} `

 Write-Host -ForegroundColor Cyan "App registration created with app ID" $appRegistration.AppId

# Create corresponding service principal
New-MgServicePrincipal -AppId $appRegistration.AppId | Out-Null
Write-Host -ForegroundColor Cyan "Service principal created"
Write-Host
Write-Host -ForegroundColor Green "Success"
Write-Host

# Generate appsettings.Development.json
# If file already exists, back it up
if (Test-Path $AppSettingsPath)
{
  $timestamp = Get-Date -UFormat %s
  $newFileName = $AppSettingsFile + ".backup." + $timestamp
  Rename-Item -Path $AppSettingsPath -NewName $newFileName
  Write-Host -ForegroundColor Yellow "Renamed existing appsettings.Development.json to " + $newFileName
}

$appSettings = @{ AzureAd=@{ ClientId=$appRegistration.AppId; TenantId=$context.TenantId } }
$appSettings | ConvertTo-Json | Add-Content -Path $AppSettingsPath
$apiScope = "api://" + $appRegistration.AppId + "/" + $ScopeName
Write-Host -ForegroundColor Cyan "App registration details saved in " + $AppSettingsPath
Write-Host -ForegroundColor Cyan "Use the following API scope to configure your client application:"
Write-Host -ForegroundColor Yellow $apiScope

if ($StayConnected -eq $false)
{
  Disconnect-MgGraph
  Write-Host "Disconnected from Microsoft Graph"
}
else
{
  Write-Host
  Write-Host -ForegroundColor Yellow "The connection to Microsoft Graph is still active. To disconnect, use Disconnect-MgGraph"
}
