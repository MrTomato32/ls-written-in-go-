package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    path := "."

    for _, arg := range os.Args[1:] {
        if !strings.HasPrefix(arg, "-") {
            path = arg
            continue
        }
        for _, char := range arg[1:] {
            switch char {
            case 'a':
                // idk
            default:
                fmt.Fprintf(os.Stderr, "usage: %s [-a] [path]\n, os.Args[0]")
                os.Exit(1)
        }
    }
}
    
   entries, error := os.ReadDir(path)
    if error != nil {
        fmt.Println("Directory not found!")
        os.Exit(1)
    }
    for _, entry:= range entries {
        if strings.HasPrefix(entry.Name(), ".") {
            continue
        }
        fmt.Println(entry.Name())
    }
}
