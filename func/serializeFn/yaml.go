package serializeFn

import (
	"gopkg.in/yaml.v2"
	"io"
)

func EncodeYaml(w io.Writer, v any) error {
	return yaml.NewEncoder(w).Encode(v)
}

func DecodeYaml(w io.Reader, v any) error {
	return yaml.NewDecoder(w).Decode(v)
}

func MarshalYaml(v any) (string, error) {
	byteBody, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(byteBody), nil
}

func UnmarshalYaml(s string, v any) error {
	return yaml.Unmarshal([]byte(s), v)
}
