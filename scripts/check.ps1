# Paikallinen tarkistus: gofmt, vet, build ja testit.
$ErrorActionPreference = 'Stop'
Set-Location (Split-Path -Parent $PSScriptRoot)
$unformatted = gofmt -l .
if ($unformatted) { throw "gofmt: muotoile tiedostot:`n$unformatted" }
go vet ./...
if ($LASTEXITCODE) { throw 'go vet epäonnistui' }
go build -trimpath ./...
if ($LASTEXITCODE) { throw 'go build epäonnistui' }
go test ./...
if ($LASTEXITCODE) { throw 'go test epäonnistui' }
Write-Host 'OK'
