package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/go-openapi/loads"
)

func main() {
    swaggerDoc, err := loads.Spec(os.Args[1])
    if err != nil {
        panic(err)
    }

    for path, item := range swaggerDoc.Spec().Paths.Paths {
        for method, op := range item.Operations() {
            fmt.Printf("%s %s %s\n", path, method, op.ID)
        }
    }
}
