package shims

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

type Serializer[T any] interface {
	ToFile(fn string, data T)
	ToWriter(w io.Writer, data T) error
	FromFile(fn string) (T, error)
	FromReader(r io.Reader) (T, error)
}

type JsonSerializer[T any] struct {
}

// ToFile serializes an arbitrary data type in JSON to the
// specified file
func (s JsonSerializer[T]) ToFile(fn string, data T) error {
	out, err := json.Marshal(data)
	//log.Printf("TRACE: DATA = %s", string(out))
	if err != nil {
		return err
	}
	return os.WriteFile(fn, out, 0644)
}

// ToWriter serializes an arbitrary data type in JSON to the
// specified io.Writer
func (s JsonSerializer[T]) ToWriter(w io.Writer, data T) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	return err
}

// FromFile deserializes a JSON representation from a
// specified file into an arbitrary object
func (s JsonSerializer[T]) FromFile(fn string) (T, error) {
	var out T
	data, err := os.ReadFile(fn)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)
	return out, err
}

// FromWriter deserializes a JSON representation from a
// specified io.Reader into an arbitrary object
func (s JsonSerializer[T]) FromReader(r io.Reader) (T, error) {
	var out T
	data, err := io.ReadAll(r)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)
	return out, err
}

// SerializeJsonToFile serializes an arbitrary data type in JSON to the
// specified file
func SerializeJsonToFile[T any](fn string, data T) error {
	s := JsonSerializer[T]{}
	return s.ToFile(fn, data)
}

// SerializeJsonToWriter serializes an arbitrary data type in JSON to the
// specified io.Writer
func SerializeJsonToWriter[T any](w io.Writer, data T) error {
	s := JsonSerializer[T]{}
	return s.ToWriter(w, data)
}

// UnserializeJsonFromFile deserializes a JSON representation from a
// specified file into an arbitrary object
func UnserializeJsonFromFile[T any](fn string) (T, error) {
	s := JsonSerializer[T]{}
	return s.FromFile(fn)
}

// UnserializeJsonFromReader deserializes a JSON representation from a
// specified io.Reader into an arbitrary object
func UnserializeJsonFromReader[T any](r io.Reader) (T, error) {
	s := JsonSerializer[T]{}
	return s.FromReader(r)
}

type GOB64Serializer[T any] struct {
}

// ToFile serializes an arbitrary data type in GOB64 to the
// specified file
func (s GOB64Serializer[T]) ToFile(fn string, data T) error {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(data)
	if err != nil {
		log.Printf("failed gob Encode: %s", err.Error())
		return err
	}
	fp, err := os.OpenFile(fn, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	enc := base64.NewEncoder(base64.StdEncoding, fp)
	_, err = enc.Write(b.Bytes())
	return err
}

// ToWriter serializes an arbitrary data type in GOB64 to the
// specified io.Writer
func (s GOB64Serializer[T]) ToWriter(w io.Writer, data T) error {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(data)
	if err != nil {
		log.Printf("failed gob Encode: %s", err.Error())
		return err
	}
	enc := base64.NewEncoder(base64.StdEncoding, w)
	_, err = enc.Write(b.Bytes())
	return err
}

// FromFile deserializes a GOB64 representation from a
// specified file into an arbitrary object
func (s GOB64Serializer[T]) FromFile(fn string) (T, error) {
	return *new(T), errors.ErrUnsupported
}

// FromWriter deserializes a GOB64 representation from a
// specified io.Reader into an arbitrary object
func (s GOB64Serializer[T]) FromReader(r io.Reader) (T, error) {
	return *new(T), errors.ErrUnsupported
}

// ToGOB64 encodes an arbotrary object into a base64 wrapped GOB
// encoded string
func ToGOB64[T any](m T) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Printf("failed gob Encode: %s", err.Error())
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
