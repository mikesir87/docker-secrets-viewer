package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    files, _ := ioutil.ReadDir("/run/secrets")
    for _, f := range files {
        dat, err := ioutil.ReadFile("/run/secrets/" + f.Name())
        check(err)
        fmt.Print(f.Name() + " : " + string(dat))
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
