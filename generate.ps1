<#
.SYNOPSIS
    Regenerate generated/gatewayapi from the LocalGateway management OpenAPI document.

.DESCRIPTION
    The generated Go files are committed, so `go build` does not need Java,
    the generator jar, or the .cache directory. Run this script only when the
    management contract changes.

    Paths default relative to the repository root. Override -Jar / -Spec /
    -Output to use another generator build or spec. In a LocalGateway monorepo
    checkout, -ExportSpec runs `cargo run -p gateway-admin --example dump-openapi`
    first; the standalone localgateway-go repository has no Rust workspace, so
    pass -Spec explicitly there.

.EXAMPLE
    powershell -NoProfile -ExecutionPolicy Bypass -File sdk/go/generate.ps1 -ExportSpec

.EXAMPLE
    powershell -NoProfile -ExecutionPolicy Bypass -File generate.ps1 -Jar C:\tools\openapi-generator-cli.jar -Spec C:\path\openapi.json

.EXAMPLE
    powershell -NoProfile -ExecutionPolicy Bypass -File sdk/go/generate.ps1 -JavaHome D:\JDKs\jdk-27 -Spec target/openapi.json
#>
[CmdletBinding()]
param(
    [string]$Jar = "",
    [string]$Spec = "",
    [string]$Output = "",
    [string]$JavaHome = "",
    [string]$Java = "",
    [string]$PackageName = "gatewayapi",
    [switch]$ExportSpec
)

$ErrorActionPreference = "Stop"

# Locate the repository root. In a LocalGateway monorepo checkout this script
# lives at sdk/go/generate.ps1, so the monorepo root is two levels up. In the
# standalone localgateway-go repository this script sits at the repository
# root, so use $PSScriptRoot.
$monorepoCandidate = Join-Path $PSScriptRoot "..\.."
if (Test-Path (Join-Path $monorepoCandidate "crates\gateway-admin")) {
    $repoRoot = (Resolve-Path $monorepoCandidate).Path
} else {
    $repoRoot = (Resolve-Path $PSScriptRoot).Path
}

if (-not $Jar) {
    $Jar = Join-Path $repoRoot ".cache\openapi-generator\7.26.0\openapi-generator-cli.jar"
}
if (-not $Spec) {
    $Spec = Join-Path $repoRoot "target\openapi.json"
}
if (-not $Output) {
    $Output = Join-Path $PSScriptRoot "generated/gatewayapi"
}

if (-not $Java) {
    $javaCandidates = @()
    if ($JavaHome) {
        $javaCandidates += (Join-Path $JavaHome "bin\java.exe")
    }
    if ($env:JAVA_HOME) {
        $javaCandidates += (Join-Path $env:JAVA_HOME "bin\java.exe")
    }
    $resolved = $javaCandidates | Where-Object { Test-Path $_ } | Select-Object -First 1
    if ($resolved) {
        $Java = $resolved
    } else {
        $Java = "java"
    }
}

if ($ExportSpec) {
    if (-not (Test-Path (Join-Path $repoRoot "crates\gateway-admin"))) {
        throw "-ExportSpec requires the LocalGateway monorepo checkout (crates/gateway-admin not found under $repoRoot). Pass -Spec <openapi.json> instead."
    }
    Push-Location $repoRoot
    try {
        Write-Host "Exporting OpenAPI document to $Spec"
        & cargo run -q -p gateway-admin --example dump-openapi -- target/openapi.json
        if ($LASTEXITCODE -ne 0) {
            throw "cargo dump-openapi failed with exit code $LASTEXITCODE"
        }
    } finally {
        Pop-Location
    }
}

if (-not (Test-Path $Jar)) {
    throw "OpenAPI Generator jar not found: $Jar. Pass -Jar <path> or place the jar under $repoRoot\.cache\openapi-generator\7.26.0\."
}
if (-not (Test-Path $Spec)) {
    throw "OpenAPI spec not found: $Spec. Pass -Spec <path> (or run with -ExportSpec inside the monorepo checkout)."
}

if (Test-Path $Output) {
    Remove-Item -Recurse -Force $Output
}

$generatorArgs = @(
    "-jar", $Jar,
    "generate",
    "-i", $Spec,
    "-g", "go",
    "-o", $Output,
    "--global-property=apiTests=false,modelTests=false",
    "--additional-properties=packageName=$PackageName,isGoSubmodule=true,withGoMod=false,generateInterfaces=true"
)

Write-Host "Running OpenAPI Generator with $Java"
& $Java @generatorArgs
if ($LASTEXITCODE -ne 0) {
    throw "openapi-generator failed with exit code $LASTEXITCODE"
}

# These files are generator boilerplate that this repository does not use.
foreach ($noise in @(".travis.yml", "git_push.sh", ".gitignore")) {
    $path = Join-Path $Output $noise
    if (Test-Path $path) {
        Remove-Item -Force $path
    }
}

# openapi-generator emits Go code that is not gofmt-formatted unless
# GO_POST_PROCESS_FILE is set. Format the output so regeneration is
# reproducible, `gofmt -l` stays clean, and scripts/verify-generated.ps1 can
# compare the committed files byte for byte.
$gofmt = Get-Command gofmt -ErrorAction SilentlyContinue
if ($gofmt) {
    Get-ChildItem -LiteralPath $Output -Recurse -File -Filter "*.go" | ForEach-Object {
        & $gofmt.Source -w $_.FullName
        if ($LASTEXITCODE -ne 0) {
            throw "gofmt failed for $($_.FullName)"
        }
    }
} else {
    Write-Warning "gofmt not found; generated Go files are not formatted. Install the Go toolchain to reproduce the committed output."
}

# The repository stores LF via .gitattributes, but the Java generator writes
# CRLF and the Go docs template leaves trailing spaces in Markdown tables.
# Normalize both so regenerating leaves a clean working tree and
# scripts/verify-generated.ps1 can compare content reproducibly.
Get-ChildItem -LiteralPath $Output -Recurse -File | ForEach-Object {
    $text = [System.IO.File]::ReadAllText($_.FullName, [System.Text.Encoding]::UTF8)
    $normalized = $text.Replace("`r`n", "`n").Replace("`r", "`n")
    if ($_.Extension -eq ".md") {
        $normalized = (($normalized -split "`n") | ForEach-Object { $_.TrimEnd() }) -join "`n"
        $normalized = $normalized.TrimEnd([char[]]"`n") + "`n"
    }
    if ($normalized -ne $text) {
        [System.IO.File]::WriteAllText($_.FullName, $normalized, (New-Object System.Text.UTF8Encoding($false)))
    }
}

$versionFile = Join-Path $Output ".openapi-generator\VERSION"
if (Test-Path $versionFile) {
    $generatorVersion = (Get-Content $versionFile -Raw).Trim()
    Write-Host "Generated $Output with OpenAPI Generator $generatorVersion"
} else {
    Write-Host "Generated $Output"
}
