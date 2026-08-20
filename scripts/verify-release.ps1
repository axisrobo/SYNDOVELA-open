#verify-release.ps1 — verify a SYNDOVELA Open release directory.
#
#   ./scripts/verify-release.ps1 -Version v1.0.0 -Platform windows-amd64
#
# Verifies, for every binary listed in SHA256SUMS:
#   1. the file exists next to the checksum file, and
#   2. its SHA-256 matches the published checksum.
#
# If SHA256SUMS.asc contains a real (non-placeholder) OpenPGP signature,
# gpg is additionally used to verify it against the bundled release key.
# As of v1.0.0 the .asc files are placeholders ("no GPG key available");
# the script warns when that is the case instead of failing.
#
# Exit code is 0 only when every checksum verified.

param(
    [Parameter(Mandatory = $true)]
    [string]$Version,

    [Parameter(Mandatory = $true)]
    [string]$Platform,

    [string]$Root
)

$ErrorActionPreference = "Stop"
$scriptDir = if ($PSScriptRoot) { $PSScriptRoot } else { Split-Path -Parent $MyInvocation.MyCommand.Path }
if (-not $Root) {
    $Root = [System.IO.Path]::GetFullPath((Join-Path $scriptDir "..\releases"))
}
$dir = [System.IO.Path]::GetFullPath((Join-Path (Join-Path $Root $Version) $Platform))
if (-not (Test-Path $dir)) {
    Write-Error "release directory not found: $dir"
    exit 1
}

$sumsPath = Join-Path $dir "SHA256SUMS"
if (-not (Test-Path $sumsPath)) {
    Write-Error "missing $sumsPath"
    exit 1
}

$failures = 0
$checked = 0
foreach ($line in Get-Content $sumsPath) {
    $line = $line.Trim()
    if ($line -eq "" -or $line.StartsWith("#")) { continue }
    $parts = $line -split "\s+", 2
    if ($parts.Length -ne 2) {
        Write-Warning "malformed checksum line: $line"
        continue
    }
    $want = $parts[0]
    $name = $parts[1].Trim()
    $file = Join-Path $dir $name
    if (-not (Test-Path $file)) {
        Write-Error "missing binary: $name"
        $failures++
        continue
    }
    $got = (Get-FileHash $file -Algorithm SHA256).Hash.ToLower()
    if ($got -ne $want) {
        Write-Error "CHECKSUM MISMATCH: $name"
        Write-Error "  want: $want"
        Write-Error "  got:  $got"
        $failures++
    } else {
        Write-Host "OK  $name"
    }
    $checked++
}

$ascPath = Join-Path $dir "SHA256SUMS.asc"
if (Test-Path $ascPath) {
    $asc = (Get-Content $ascPath -Raw).Trim()
    if ($asc -match "placeholder") {
        Write-Warning "SHA256SUMS.asc is a placeholder (no signing key); skipping GPG verification"
    } else {
        $key = Get-ChildItem $dir -Filter "*.asc" | Where-Object { $_.Name -ne "SHA256SUMS.asc" } | Select-Object -First 1
        if (-not $key) { $key = Get-ChildItem (Join-Path $dir "..") -Filter "*.asc" | Select-Object -First 1 }
        if ($key) {
            gpg --import (Resolve-Path $key.FullName) | Out-Null
            gpg --verify $ascPath $sumsPath
            if ($LASTEXITCODE -ne 0) {
                Write-Error "GPG signature verification failed"
                $failures++
            }
        } else {
            Write-Warning "signature present but no public key found to verify against"
        }
    }
}

Write-Host ""
Write-Host "checked $checked binaries, $failures failure(s)"
if ($failures -gt 0) { exit 1 }
Write-Host "verified $Version/$Platform OK"
