package main

import (
    "fmt"
    "strings"
    "os/exec"
    "os"
)

//
const KEYFOLDER string = "workspace"
const GITFOLDER string = ".git"

var cmdret []byte
var folders []string


func getfolder() []string {
    //get current pwd    
    cmd := exec.Command("pwd")
    cmdret, cmderr := cmd.Output()

    buf := string(cmdret[:])
    i := strings.Index(buf, KEYFOLDER)
    buf = buf[:i+len(KEYFOLDER)]
    fmt.Println("curpwd:" + buf)
    
    //check folder list
    cmderr = os.Chdir(buf)
    if (cmderr!=nil) {
        fmt.Println(cmderr)
    }
    
    cmd = exec.Command("pwd")
    cmdret, cmderr = cmd.Output()
    buf = string(cmdret[:])
    fmt.Println("keypwd:"+buf)
    
    cmd = exec.Command("find",".", "-name", GITFOLDER)
    cmdret, cmderr = cmd.Output()
    buf = string(cmdret[:])
    fmt.Println(buf)

    folders = strings.Fields(buf)
    fmt.Println(folders)

    
    return folders
}