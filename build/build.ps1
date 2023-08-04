# Name of the Go file to build
$GOFILE = "../main.go"

# Use the first command-line argument as the output name, or default to "saas-squash"
$OUTNAME = If ($args.Length -eq 0) { "saas-squash" } Else { $args[0] }

# Check if the "small" argument is passed
$LDFLAGS = ""
if ($args.Length -gt 1 -and $args[1] -eq "small") {
  $LDFLAGS = '-ldflags "-s -w"'
}

# Iterate through OS and ARCH
foreach ($GOOS in "darwin", "linux", "windows") {
  foreach ($GOARCH in "amd64") {
    # Set the environment variables
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH

    # Name of the output binary file
    $OUTFILE = "$OUTNAME-$GOOS-$GOARCH"

    # If it's a Windows build, add the .exe extension
    if ($GOOS -eq "windows") {
      $OUTFILE = "$OUTFILE.exe"
    }

    # Run the build command
    Write-Host "Building $OUTFILE"
    Invoke-Expression "go build $LDFLAGS -o $OUTFILE $GOFILE"
  }
}

Write-Host "Build process completed"