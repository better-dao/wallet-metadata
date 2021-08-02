package main

import (
	"os"
	"strings"

	"github.com/better-go/pkg/log"
	"github.com/urfave/cli/v2"

	"metadata/app/internal/task"
)

func main() {

	// do run:
	Runner()

}

// run server:
func Runner() {
	// parse cmd args:
	r := &cli.App{
		Name:    "token contract generator tool",
		Version: "v0.0.1",

		//
		// 参数解析:
		//
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "chainType",
				Aliases: []string{"type"},
				Usage:   "chainType: [data/from.json]",
				EnvVars: []string{"CHAIN_TYPE", "chainType"},
			},

			&cli.StringFlag{
				Name:    "fromFile",
				Aliases: []string{"from"},
				Usage:   "from json file: [data/from.json]",
				EnvVars: []string{"FROM_FILE", "fromFile"},
			},
			&cli.StringFlag{
				Name:    "toFile",
				Aliases: []string{"to"},
				Usage:   "generate result file: [dist/to.json]",
				EnvVars: []string{"TO_FILE", "toFile"},
				//FilePath: "configs/configs.yaml", // 会自动解析文件内容
			},
		},


		//
		// 执行动作:
		//
		Action: func(ctx *cli.Context) error {
			// 服务类型:(小写转换)
			chainType := strings.ToLower(ctx.String("chainType"))

			// from file:
			from := ctx.String("fromFile")

			// to file:
			to := ctx.String("toFile")

			log.Infof("run action: chainType=%v,  fromFile=%v, toFile=%v", chainType, from, to)

			// dispatch:
			switch chainType {
			case "eth":
				_ = task.GenEthMainNetContractMeta(from, to)
			case "bsc":
				if to == "" {
					to = "dist/bsc_contract_map.json"
				}
				_ = task.GenBscMainNetTokenMeta(to)
			}
			return nil
		},

		Before:   nil,
		After:    nil,
		Commands: nil,
	}

	// do run:
	if err := r.Run(os.Args); err != nil {
		log.Errorf("start run failed, error: %v", err)
	}
}
