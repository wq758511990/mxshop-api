package utils

import (
	"github.com/goccy/go-json"
	"io"
)

func JSONDecode(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return nil
}
