/*
prj needs to build a git code base with different build script,
also we need to check the log when it's fail.

eg:
$>source envsetup lotus cfg=lotus.mk
$>make sdk
$>make sdk_install
$>make

should check the screen output and find whether the compile success.
then get the result.

this program do previous build work by :
1.get build config file and build
2.check each steps result
3.save the log if build fails.

by cmd input:
$>./gitbuild lotusbuild

lotusbuild file is:


*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var buildconfigname string
var cmderr error
var cmdret []byte

const tmplogfolder  = "templog"

func main() {
	curargs := os.Args
	if len(curargs) < 2 {
		help()
		return
	}
	//patch folder name only contains first 5 bytes in commit id.
	if len(curargs) == 2 {
		buildconfigname = os.Args[1]
	}

	fmt.Println("buildconfig:", buildconfigname)

	checkerr := checkenv()
	if checkerr != nil {
		return
	}

	fmt.Println("Warning! Due to save build log files, this program will create a temp folder in current pwd")

	checkerr = preparefolder(tmplogfolder)
	if checkerr != nil {
		return
	}

	//build check result

	//remove folder at last
	checkerr = removefolder(tmplogfolder)
	if cmderr != nil {
		return
	}

	return
}


func preparefolder(name string) error {
	cmdret, cmderr = exec.Command("mkdir", "-p", name).Output()

	if cmderr != nil {
		fmt.Println("mkdir -p " + name + " failed")
		return cmderr
	}

	return cmderr
}

func removefolder(name string) error {
	cmdret, cmderr = exec.Command("rm", "-r", name).Output()

	if cmderr != nil {
		fmt.Println("rm -r " + name + " failed")
		return cmderr
	}

	return cmderr
}



func help() {
	fmt.Println("Wrong Args! Check Usage:")
	fmt.Println("gitbuild lotusbuild")
	fmt.Println("gitbuild will save output with each build step in temp folder")
	fmt.Println("if build success, temp folder will be delete")
	fmt.Println()
	return
}

func checkenv() error {
	var cmdret []byte
	var cmderr error

	cmdret, cmderr = exec.Command("make", "--version").Output()
	buf := string(cmdret[:])
	if cmderr != nil {
		fmt.Println("Check make version Failed:" + buf)
		return cmderr
	}

	_, cmderr = os.Stat(buildconfigname);
	if cmderr != nil {
		fmt.Println("Check buildconfig file Failed:" + buildconfigname)
		return cmderr
	}

	return cmderr
}
