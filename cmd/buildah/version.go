package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	ispecs "github.com/opencontainers/image-spec/specs-go"
	rspecs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/projectatomic/buildah"
	"github.com/urfave/cli"
)

//Overwritten at build time
var (
	gitCommit string
	buildInfo string
)

//Function to get and print info for version command
func versionCmd(c *cli.Context) error {

	//converting unix time from string to int64
	buildTime, err := strconv.ParseInt(buildInfo, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println("Version:      ", buildah.Version)
	fmt.Println("Go Version:   ", runtime.Version())
	fmt.Println("Image Spec:   ", ispecs.Version)
	fmt.Println("Runtime Spec: ", rspecs.Version)
	fmt.Println("Git Commit:   ", gitCommit)

	//Prints out the build time in readable format
	fmt.Println("Built:        ", time.Unix(buildTime, 0).Format(time.ANSIC))
	fmt.Println("OS/Arch:      ", runtime.GOOS+"/"+runtime.GOARCH)

	return nil
}

//cli command to print out the version info of buildah
var versionCommand = cli.Command{
	Name:           "version",
	Usage:          "Display the Buildah Version Information",
	Action:         versionCmd,
	SkipArgReorder: true,
}
