package kv

import (
	"bytes"
	"encoding/json"

	"github.com/hashicorp/consul/api"
)

// KV is used to represent a single K/V entry.
type KV struct {
	*api.KVPair
	// Value overrides embedded KVPair's Value
	// to represent value as json object instead of raw string.
	Value interface{}
}

// Decode overrides given value representation for each kv entry as an object
// and returns all kv as json array.
func Decode(kvs ...*api.KVPair) ([]byte, error) {
	decoded := make([]KV, len(kvs))
	for i, kv := range kvs {
		decoded[i] = KV{KVPair: kv, Value: make(map[string]interface{})}
		if len(kv.Value) == 0 {
			continue
		}
		if err := json.Unmarshal(kv.Value, &decoded[i].Value); err != nil {
			return nil, err
		}
	}
	return indent(decoded)
}

// Encode accepts json array of kv entries to transform each
// value as []byte from json object.
func Encode(data []byte) (encoded api.KVPairs, err error) {
	var kvs []KV
	if err := json.Unmarshal(data, &kvs); err != nil {
		return nil, err
	}
	encoded = make(api.KVPairs, len(kvs))
	for i, kv := range kvs {
		encoded[i] = kv.KVPair
		encoded[i].Value, err = indent(kv.Value)
		if err != nil {
			return nil, err
		}
	}
	return encoded, nil
}

func indent(v interface{}) ([]byte, error) {
	var w bytes.Buffer
	e := json.NewEncoder(&w)
	e.SetEscapeHTML(false)
	e.SetIndent("", "	")
	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
