//prj needs to generate a git patch format file into a  patch file list.
//eg:
/*
$>git lg
* 808369a - (HEAD, origin/master, origin/HEAD) remove x (6 days ago) <falcon.xu>
* 49ffbc1 - update path (6 days ago) <juxu>
$>git diff 808369a 49ffbc1 --stat
 gostudy/effective.go.md  |   0
 gostudy/goslice_show.png | Bin
 gostudy/gostudy.go       |   0
 3 files changed, 0 insertions(+), 0 deletions(-)

$>mkdir -p gostudy_new
$>cp gostudy/effective.go.md gostudy_new
$>cp gostudy/goslice_show.png gostudy_new
$>cp gostudy/gostudy.go gostudy_new
$>tar czvf gostudy_new.tgz gostduy_new
$>rm -fr gostduy_new

this program do previews works
by cmd input:
$>./gitpatch . 808369a 49ffbc1
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var gitdiffbuf string
var patchfoldername string
var cmderr error
var cmdret []byte

func main() {
	curargs := os.Args
	if len(curargs) != 4 {
		help()
		return
	}

	//patch folder name only contains first 5 bytes in commit id.
	patchfoldername = os.Args[2][:5] + "_" + os.Args[3][:5]
	fmt.Println("create folder:", patchfoldername)

	checkerr := checkenv()
	if checkerr != nil {
		return
	}

	checkerr = checkgitlog(&gitdiffbuf)
	if checkerr != nil {
		return
	}
	//fmt.Println("git diff result:")
	//fmt.Println(buf)
	checkerr = preparefolder(patchfoldername)
	if checkerr != nil {
		return
	}

	difffiles := strings.Split(gitdiffbuf, "\n")
	for _, char := range difffiles {
		if strings.Contains(char, "|") {
			var firstblankidx, fisrtdelimiter int
			firstblankidx = strings.Index(char, " ")
			fisrtdelimiter = strings.Index(char, "|")
			//fmt.Println("*" + char[firstblankidx+1:fisrtdelimiter] + "*")
			copysinglefile(char[firstblankidx+1 : fisrtdelimiter])
		}
	}

	checkerr = tarfiles(patchfoldername)
	if checkerr != nil {
		return
	}

	//remove folder at last
	cmdret, cmderr = exec.Command("rm", "-fr", patchfoldername).Output()
	if cmderr != nil {
		fmt.Println("rm " + patchfoldername + " failed")
		return
	}

	return
}

func tarfiles(foldername string) error {
	cmderr = nil

	cmdret, cmderr = exec.Command("tar", "czvf", foldername+".tgz", foldername).Output()
	if cmderr != nil {
		fmt.Println("tar " + foldername + " failed")
		return cmderr
	}
	fmt.Println("tar " + foldername + ".tgz success")
	return cmderr
}

func copysinglefile(singlefile string) error {
	cmderr = nil
	var singlefilepath string
	var singlefilename string
	var lastslashidx, firstblankidx int

	singlefilename = ""

	//fmt.Println("*" + singlefile + "*")

	if strings.Contains(singlefile, "/") {
		lastslashidx = strings.LastIndex(singlefile, "/")
		firstblankidx = strings.Index(singlefile, " ")
		singlefilepath = singlefile[:lastslashidx]
		singlefilename = singlefile[lastslashidx+1 : firstblankidx]

		cmdret, cmderr = exec.Command("mkdir", "-p", patchfoldername+"/"+singlefilepath).Output()
		if cmderr != nil {
			fmt.Println("mkdir -p " + singlefilepath + " failed")
			return cmderr
		}
	} else {
		singlefilepath = ""
		firstblankidx = strings.Index(singlefile, " ")
		singlefilename = singlefile[:firstblankidx]
	}

	if len(singlefilepath) != 0 {
		cmdret, cmderr = exec.Command("cp", "-d", singlefilepath+"/"+singlefilename, patchfoldername+"/"+singlefilepath+"/").Output()
	} else {
		cmdret, cmderr = exec.Command("cp", "-d", singlefilename, patchfoldername+"/").Output()
	}

	if cmderr != nil {
		fmt.Println("cp " + singlefilename + " failed")
		return cmderr
	}
	fmt.Println("cp " + singlefilename + " success")
	return cmderr
}

func preparefolder(name string) error {
	cmdret, cmderr = exec.Command("mkdir", "-p", name).Output()

	if cmderr != nil {
		fmt.Println("mkdir -p " + name + " failed")
		return cmderr
	}

	return cmderr
}

func checkgitlog(buf *string) error {

	//git diff $commit1 $commit2 and need stat-name-width more for no ... .
	cmdret, cmderr = exec.Command("git", "diff", os.Args[2], os.Args[3], "--stat-width=512").Output()
	*buf = string(cmdret[:])

	if cmderr != nil {
		fmt.Println("git diff "+os.Args[2]+" "+os.Args[3], "failed")
		return cmderr
	}
	fmt.Println(*buf)

	return cmderr
}

func help() {
	fmt.Println("Wrong Args! Check Usage:")
	fmt.Println("gitpatch . 808369a 49ffbc1:")
	fmt.Println("gitpatch <$curpath> <$commit1> <$commit2>:")
	fmt.Println("gitpatch only works on rootdir of git repository, commitid should more than 6 bytes.")
	fmt.Println()
	return
}

func checkenv() error {
	var cmdret []byte
	var cmderr error
	cmdret, cmderr = exec.Command("git", "--version").Output()
	buf := string(cmdret[:])
	if cmderr != nil {
		fmt.Println("Check git version Failed:" + buf)
		return cmderr
	}

	cmdret, cmderr = exec.Command("tar", "--version").Output()
	i := strings.Index(buf, "\n")
	buf += string(cmdret[:i])
	if cmderr != nil {
		fmt.Println("Check tar version Failed:" + buf)
		return cmderr
	}

	//fmt.Println("Check result:")
	//fmt.Println(buf)
	//fmt.Println()

	return cmderr
}
