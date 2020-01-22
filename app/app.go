package app

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/unqnown/consulkv/pkg/check"
	"github.com/unqnown/consulkv/pkg/kv"
	"github.com/urfave/cli"
)

// Run is the entry point to the cli app.
func Run() error {
	app := cli.NewApp()

	app.Name = "consulkv"
	app.Version = "v0.1.0"
	app.Usage = "Consul kv export/import tool."

	app.Commands = []cli.Command{
		_export(),
		_import(),
	}

	return app.Run(os.Args)
}

func verbosef(v bool, format string, a ...interface{}) {
	if !v {
		return
	}
	log.Printf(format, a...)
}

func _export() cli.Command {
	action := func(ctx *cli.Context) {
		v := ctx.Bool("verbose")

		conf := api.DefaultConfig()
		conf.Address = ctx.String("address")
		consul, err := api.NewClient(conf)
		check.Fatal(err, "init consul client")

		kvs, _, err := consul.KV().List(ctx.String("prefix"), nil)
		check.Fatal(err, "get kv")
		verbosef(v, "kv list retrieved")

		decoded, err := kv.Decode(kvs...)
		check.Fatal(err, "decode kv")
		verbosef(v, "kv decoded")

		check.Fatal(save(decoded, ctx.String("file")), "save kv")
		verbosef(v, "kv successfully exported")
	}

	return cli.Command{
		Name:        "export",
		Aliases:     []string{"e"},
		Usage:       "exports kv from consul server.",
		Description: ``,
		Category:    "kv",
		Action:      action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "file, f",
				Required: true,
				Usage:    "Path to kv `FILE`",
			},
			cli.StringFlag{
				Name:  "address, a",
				Usage: "Consul server address",
			},
			cli.StringFlag{
				Name:  "prefix, p",
				Usage: "KV prefix",
			},
			cli.BoolFlag{
				Name:  "verbose, v",
				Usage: "Prints progress information",
			},
		},
	}
}

func _import() cli.Command {
	action := func(ctx *cli.Context) {
		v := ctx.Bool("verbose")

		conf := api.DefaultConfig()
		conf.Address = ctx.String("address")
		consul, err := api.NewClient(conf)
		check.Fatal(err, "init consul client")

		data, err := load(ctx.String("file"))
		check.Fatal(err, "load kv")
		verbosef(v, "kv loaded")

		kvs, err := kv.Encode(data)
		check.Fatal(err, "encode kv")
		verbosef(v, "kv encoded")

		for _, kv := range kvs {
			_, err := consul.KV().Put(kv, nil)
			check.Fatal(err, "put kv %s", kv.Key)
			verbosef(v, "put %s", kv.Key)
		}

		verbosef(v, "kv successfully imported")
	}

	return cli.Command{
		Name:        "import",
		Aliases:     []string{"i"},
		Usage:       "imports kv to consul server.",
		Description: ``,
		Category:    "kv",
		Action:      action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "file, f",
				Required: true,
				Usage:    "Path to kv `FILE`",
			},
			cli.StringFlag{
				Name:  "address, a",
				Usage: "Consul server address",
			},
			cli.BoolFlag{
				Name:  "verbose, v",
				Usage: "Prints progress information",
			},
		},
	}
}

func json(filename string) string {
	if strings.HasSuffix(filename, ".json") {
		return filename
	}
	return filename + ".json"
}

func save(data []byte, filename string) error {
	f, err := os.Create(json(filename))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func load(filename string) ([]byte, error) {
	f, err := os.Open(json(filename))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
