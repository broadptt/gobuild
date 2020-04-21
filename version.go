package buildversion

import (
	"fmt"
)

type versionInfo struct {
	BuildVersion string
	GitCommit    string
	BuildTime    string
	GoVersion    string
}

var version = new(versionInfo)

func SettingVersionInfo(buildVersion, gitCommit, buildTime, goVersion string) {
	version.BuildVersion = buildVersion
	version.GitCommit = gitCommit
	version.BuildTime = buildTime
	version.GoVersion = goVersion
}
func PrintBuildingVerionInfo() {
	fmt.Println("        ------  build infomation ! ------ ")
	fmt.Println("|--------------------------------------------------|")
	fmt.Println("   version :|", version.BuildVersion)
	fmt.Println("|--------------------------------------------------|")
	fmt.Println(" git commit:|", version.GitCommit)
	fmt.Println("|--------------------------------------------------|")
	fmt.Println(" build time:|", version.BuildTime)
	fmt.Println("|--------------------------------------------------|")
	fmt.Println(" go version:|", version.GoVersion)
	fmt.Println("|--------------------------------------------------|")
}

func createVersion() {

}
