package v1

import "encoding/json"

//Server version representing the server version
type ServerVersion struct {
	Version string `json:"version"`
}

var versionJson []byte

func init() {
	var err error
	versionJson, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}
}
