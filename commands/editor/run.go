package editor

import (
    "fmt"
    "time"

    "github.com/codegangsta/cli"
    "github.com/labstack/echo"
    "github.com/skratchdot/open-golang/open"

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
        cli.StringFlag{
            Name:  "address, a",
            Usage: "Custom address for server",
            Value: "127.0.0.1",
        },
        cli.IntFlag{
            Name : "port, p",
            Usage: "Custom port for server",
            Value: 8312,
        },
        cli.BoolTFlag{
            Name: "browser, b",
            Usage: "Auto open editor browser (default: true)",
        },
    },
}

func editorRun(ctx *cli.Context) error {
    serverAddress := fmt.Sprintf("%s:%d", ctx.String("address"), ctx.Int("port"))

    if ctx.Bool("browser") == true {
        time.AfterFunc(2 * time.Second, func() {
            fmt.Println("⇛ openning editor in default browser ...")

            open.Run(fmt.Sprintf("http://%s/", serverAddress))
        })
    }

    e := echo.New()
    e.Use(static.ServeRoot("/", &assetfs.AssetFS{
        Asset: editor.Asset,
        AssetDir: editor.AssetDir,
        AssetInfo: editor.AssetInfo,
        Prefix: "",
    }))
    e.Logger.Fatal(e.Start(serverAddress))

    return nil
}
