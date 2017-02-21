package main

import (
    "os"

    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Run(os.Args)
}
