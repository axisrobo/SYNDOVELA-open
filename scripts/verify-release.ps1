# verify-release.ps1 — verify a SYNDOVELA Open release directory.
#
# Usage:  .\scripts\verify-release.ps1 -ReleaseDir releases\v1.0.0\windows-amd64
#
# Checks, in order:
#   1. every binary named in SHA256SUMS exists and its SHA-256 matches;
#   2. if SHA256SUMS.asc is a real OpenPGP signature, verifies it with gpg
#      (a "placeholder" signature is reported as a warning, not an error).
#
# Exits non-zero if any checksum check fails. Exit code 2 means the
# signature check was skipped or degraded.

param(
  [Parameter(Mandatory = $true)]
  [string]$ReleaseDir
)

if (-not (Test-Path (Join-Path $ReleaseDir 'SHA256SUMS'))) {
  Write-Error "no SHA256SUMS in $ReleaseDir"
  exit 2
}

$fail = $false

# 1. Checksum verification.
Get-Content (Join-Path $ReleaseDir 'SHA256SUMS') | ForEach-Object {
  $line = $_ -split '\s+', 2
  if ($line.Length -lt 2 -or -not $line[0]) { return }
  $digest = $line[0].ToLowerInvariant()
  $file = $line[1]
  $path = Join-Path $ReleaseDir $file
  if (-not (Test-Path $path)) {
    Write-Error "FAIL  missing: $file"
    $fail = $true
    return
  }
  $actual = (Get-FileHash $path -Algorithm SHA256).Hash.ToLowerInvariant()
  if ($actual -ne $digest) {
    Write-Error "FAIL  digest mismatch: $file (got $actual)"
    $fail = $true
  } else {
    Write-Output "ok    $file"
  }
}

# 2. Signature verification (real OpenPGP only).
$asc = Join-Path $ReleaseDir 'SHA256SUMS.asc'
if (Test-Path $asc) {
  $content = Get-Content $asc -Raw
  if ($content -match '^placeholder') {
    Write-Warning "SHA256SUMS.asc is a placeholder; checksums above are authoritative"
    exit $(if ($fail) { 1 } else { 2 })
  }
  if (Get-Command gpg -ErrorAction SilentlyContinue) {
    # Import the published release key if it is not already in the keyring.
    $keyFile = @(
      (Join-Path $ReleaseDir 'SYNDOVELA-RELEASE-SIGNING-KEY.asc'),
      (Join-Path $ReleaseDir '..\SYNDOVELA-RELEASE-SIGNING-KEY.asc'),
      (Join-Path $ReleaseDir '..\..\SYNDOVELA-RELEASE-SIGNING-KEY.asc')
    ) | Where-Object { Test-Path $_ } | Select-Object -First 1
    $hasKey = & gpg --list-keys releases@syndovela.dev 2>$null
    if ($keyFile -and $LASTEXITCODE -ne 0) {
      & gpg --batch --quiet --import $keyFile 2>$null
    }
    $sums = Join-Path $ReleaseDir 'SHA256SUMS'
    & gpg --verify $asc $sums 2>$null
    if ($LASTEXITCODE -eq 0) {
      Write-Output "ok    OpenPGP signature verified"
    } else {
      Write-Error "FAIL  OpenPGP signature did not verify"
      $fail = $true
    }
  } else {
    Write-Warning "gpg not installed; skipped signature verification"
    exit $(if ($fail) { 1 } else { 2 })
  }
} else {
  Write-Warning "no SHA256SUMS.asc present"
}

exit $(if ($fail) { 1 } else { 0 })
