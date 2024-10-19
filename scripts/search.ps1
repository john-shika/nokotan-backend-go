#!pwsh

$CurrentWorkDir = Get-Location
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Set-Location $ScriptDir -ErrorAction Stop
Set-Location ..

# not working well, need to fix it, im boring, help me out :)
Get-ChildItem -Path "app" -Recurse -Filter "*.go" | ForEach-Object {
    $file = $_
    Write-Output $file.FullName
    Get-Content $file.FullName | Select-String -Pattern $args[0] -CaseSensitive
}

Set-Location $CurrentWorkDir
