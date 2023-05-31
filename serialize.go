package shims

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// SerializeJsonToFile serializes an arbitrary data type in JSON to the
// specified file
func SerializeJsonToFile[T any](fn string, data T) error {
	out, err := json.Marshal(data)
	//log.Printf("TRACE: DATA = %s", string(out))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fn, out, 0644)
}

// SerializeJsonToWriter serializes an arbitrary data type in JSON to the
// specified io.Writer
func SerializeJsonToWriter[T any](w io.Writer, data T) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	return err
}

// UnserializeJsonFromFile deserializes a JSON representation from a
// specified file into an arbitrary object
func UnserializeJsonFromFile[T any](fn string) (T, error) {
	var out T
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)
	return out, nil
}

// UnserializeJsonFromReader deserializes a JSON representation from a
// specified io.Reader into an arbitrary object
func UnserializeJsonFromReader[T any](r io.Reader) (T, error) {
	var out T
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)
	return out, nil
}

// ToGOB64 encodes an arbotrary object into a base64 wrapped GOB
// encoded string
func ToGOB64[T any](m T) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Printf(`failed gob Encode`, err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// FromGOB64 decodes a base64 wrapped GOB encoded string into an
// arbitrary object
func FromGOB64[T any](str string) (T, error) {
	m := new(T)
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Printf("ERR: failed base64 Decode: %s", err.Error())
		return *m, err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		log.Printf("ERR: failed gob Decode: %s", err.Error())
		return *m, err
	}
	return *m, nil
}
