param(
  [Parameter(Mandatory=$false,
  HelpMessage="The friendly name of the app registration")]
  [String]
  $AppName = "Kiota ToDoItem Sample API Go Client",

  [Parameter(Mandatory=$false,
  HelpMessage="Your Azure Active Directory tenant ID")]
  [String]
  $TenantId,

  [Parameter(Mandatory=$false)]
  [Switch]
  $StayConnected = $false
)

$AppSettingsFile = ".env.local"
$AppSettingsPath = "./src/" + $AppSettingsFile
$ApiScopePlaceHolder = "api://YOUR_TODO_API_CLIENT_ID/ToDoItem.ReadWrite"

# Requires an admin
if ($TenantId)
{
  Connect-MgGraph -Scopes "Application.ReadWrite.All User.Read" -TenantId $TenantId -ErrorAction Stop
}
else
{
  Connect-MgGraph -Scopes "Application.ReadWrite.All User.Read" -ErrorAction Stop
}

# Get context for access to tenant ID
$context = Get-MgContext -ErrorAction Stop

# Create app registration
$appRegistration = New-MgApplication -DisplayName $AppName -SignInAudience "AzureADMyOrg" `
 -IsFallbackPublicClient -ErrorAction Stop
Write-Host -ForegroundColor Cyan "App registration created with app ID" $appRegistration.AppId

# Create corresponding service principal
New-MgServicePrincipal -AppId $appRegistration.AppId | Out-Null -ErrorAction SilentlyContinue -ErrorVariable SPError
if ($SPError)
{
  WriteHost -ForegroundColor Red "A service principal for the app could not be created."
  WriteHost -ForegroundColor Red $SPError
  Exit
}

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
  Write-Host -ForegroundColor Yellow "Renamed existing " + $AppSettingsFile + " to " + $newFileName
}

$clientIdLine = "CLIENT_ID=" + $appRegistration.AppId
$tenantIdLine = "TENANT_ID=" + $context.TenantId
$scopesLine = "API_SCOPE=" + $ApiScopePlaceHolder

Add-Content -Path $AppSettingsPath -Value $clientIdLine
Add-Content -Path $AppSettingsPath -Value $tenantIdLine
Add-Content -Path $AppSettingsPath -Value $scopesLine

Write-Host -ForegroundColor Cyan "App registration details saved in " + $AppSettingsPath
Write-Host -ForegroundColor Cyan "Update the API_SCOPE value with the API scope of your sample API"
Write-Host

if ($StayConnected -eq $false)
{
  Disconnect-MgGraph | Out-Null
  Write-Host "Disconnected from Microsoft Graph"
}
else
{
  Write-Warning "The connection to Microsoft Graph is still active. To disconnect, use Disconnect-MgGraph"
}
