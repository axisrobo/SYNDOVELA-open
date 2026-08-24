# Runs the checks required before publishing the current open release.
$ErrorActionPreference = 'Stop'

git diff --check
if ($LASTEXITCODE -ne 0) {
  exit $LASTEXITCODE
}

Push-Location (Join-Path $PSScriptRoot '..\sdk\go')
try {
  go test ./...
  if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
  }
} finally {
  Pop-Location
}

& (Join-Path $PSScriptRoot 'verify-release.ps1') -ReleaseDir (Join-Path $PSScriptRoot '..\releases\v1.3.0\windows-amd64')
exit $LASTEXITCODE
