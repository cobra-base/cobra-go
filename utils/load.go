package utils

import (
	"encoding/json"
	"io"
	"os"
)

func LoadJsonConf(configFile string, v any) error {
	jf, err := os.Open(configFile)
	if err != nil {
		return err
	}

	defer jf.Close()

	ba, err := io.ReadAll(jf)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(ba, v); err != nil {
		return err
	}

	return nil
}
