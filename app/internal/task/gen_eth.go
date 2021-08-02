package task

import (
	"github.com/better-go/pkg/log"

	"metadata/app/pkg"
)

/*


{
	// mainnet:
	'chainId': {
		{
			'tokenName': {
				'address': '',
				''
			}
		}
	}
}


*/

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

///
type HandlerFunc func(in map[string]interface{}) (out interface{}, err error)

func genEthContractMeta(from string, to string, taskFn HandlerFunc) error {
	input := make(map[string]interface{})

	// fix &
	err := pkg.ReadJsonFile(from, &input)

	log.Infof("read: $+v, err: $v", input, err)
	if err != nil {
		return err
	}

	// do task:
	data, err := taskFn(input)
	if err != nil {
		return err
	}

	/// convert:
	output, _ := pkg.ConvertToJsonBytes(data)

	// write:
	return pkg.WriteJsonFile(to, output)
}

/// ETH 正式链:
func GenEthMainNetContractMeta(from string, to string) error {
	return genEthContractMeta(from, to, func(in map[string]interface{}) (out interface{}, err error) {
		result := make(map[string]interface{})
		contracts := make(map[string]interface{})

		// mainNet:
		key := "1" // chainID = 1, = mainnet

		// fmt:
		for k, v := range in {

			// assert type:
			vv, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			// fmt key:
			symbol, _ := vv["symbol"].(string)
			name, _ := vv["name"].(string)

			// fix key name empty:
			keyTokenName := symbol
			if keyTokenName == "" {
				keyTokenName = name
			}

			log.Infof("contract item: %+v, key=%v", vv, keyTokenName)

			value := vv
			value["address"] = k

			// add to map:
			contracts[keyTokenName] = value

		}

		// fmt: mainNet
		result[key] = contracts
		// kovan:
		result["42"] = make(map[string]interface{})
		return result, nil
	})
}
