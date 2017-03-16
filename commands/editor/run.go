package editor

import (
    "github.com/codegangsta/cli"
    "github.com/labstack/echo"

    "github.com/zeuxisoo/go-contix/editor"

    static "github.com/Code-Hex/echo-static"
    assetfs "github.com/elazarl/go-bindata-assetfs"
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
    e := echo.New()
    e.Use(static.ServeRoot("/", &assetfs.AssetFS{
        Asset: editor.Asset,
        AssetDir: editor.AssetDir,
        AssetInfo: editor.AssetInfo,
        Prefix: "",
    }))
    e.Logger.Fatal(e.Start(":8312"))

    return nil
}
