package main


import (
    "fmt"
    //"os/exec"
)


func main(){
    
    fmt.Println("Sync Start")
    
    execsync(getfolder())
    
    fmt.Println("Sync End")
}