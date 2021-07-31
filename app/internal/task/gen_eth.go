package task

import (
	"github.com/better-go/pkg/log"

	"metadata/app/pkg"
)

///
func GenEthContractMeta(from string, to string) error {
	dist := make(map[string]interface{})

	// fix &
	err := pkg.ReadJsonFile(from, &dist)

	log.Infof("read: $+v, err: $v", dist, err)
	if err != nil {
		return err
	}

	/// convert:
	output, _ := pkg.ConvertToJsonBytes(dist)

	// write:
	return pkg.WriteJsonFile(to, output)
}
