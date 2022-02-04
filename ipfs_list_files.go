package main

import (
        "context"
        "fmt"

        "github.com/alexflint/go-arg"
        ipfs "github.com/ipfs/go-ipfs-api"
)

var (
        sh *ipfs.Shell

        args struct {
                APIServer string `arg:"-a, --api" help:"the API endpoint URL for your IPFS" default:"127.1:5001"`
                Path      string `arg:"-p, --path" help:"path to show listing for" default:"/"`
        }
)

func byteCountFromDecimal(b uint64) string {
        const unit = 1000
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func main() {
        arg.MustParse(&args)

        sh = ipfs.NewShell(args.APIServer)

        opts := func(rb *ipfs.RequestBuilder) error {
                rb.Option("long", true)
                return nil
        }

        // list files
        files, err := sh.FilesLs(context.Background(), args.Path, opts)
        if err != nil {
                fmt.Println(err)
                return
        }

        for _, file := range files {
                var file_type string
                if file.Type == 1 {
                        file_type = "D"
                } else {
                        file_type = "F"
                }
                fmtSize := byteCountFromDecimal(file.Size)
                fmt.Printf("[%s] %s:%s, size: %s\n", file_type, file.Name, file.Hash, fmtSize)
        }
}
