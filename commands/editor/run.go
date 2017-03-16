package editor

import (
    "fmt"

    "github.com/codegangsta/cli"
)

var CmdEditorRun = cli.Command{
    Name: "run",
    Usage: "Run the editor ui",
    Description: "The tools provide you to edit / create / update the configure file",
    Action: editorRun,
    Flags: []cli.Flag{
    },
}

func editorRun(ctx *cli.Context) error {
    fmt.Println("Editor run")
    return nil
}
