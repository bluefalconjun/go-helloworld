package main

import (
    "fmt"
    "strings"
    "os/exec"
)

//
var curpwd []byte
var folders []string


func getfolder() []string {
    cmd := exec.Command("pwd")
    
    curpwd, _ := cmd.Output()
    
    buf := string(curpwd[:])
    
    folders := strings.Split(buf, "/")
    
    fmt.Println(folders)
    
    return folders
}