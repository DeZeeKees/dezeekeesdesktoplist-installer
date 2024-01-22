# build.ps1
$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
$packageName = "./" # Replace with your package name if needed
$outputPath = "./bin/"
$outputName = "dezeekeesdesktoplist-installer.exe" # Replace with your program name

# Create the output directory if it doesn't exist
if (!(Test-Path -Path $outputPath)) {
    New-Item -ItemType Directory -Force -Path $outputPath
}

# Build the Go program
go build -o ($outputPath + $outputName) $packageName