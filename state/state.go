// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package state

import (
	"bytes"
	"encoding/gob"
	"sort"

	"github.com/pkg/errors"
)

var (
	// ErrStateSerialization is the error that the state marshaling is failed
	ErrStateSerialization = errors.New("failed to marshal state")

	// ErrStateDeserialization is the error that the state un-marshaling is failed
	ErrStateDeserialization = errors.New("failed to unmarshal state")

	// ErrStateNotExist is the error that the state does not exist
	ErrStateNotExist = errors.New("state does not exist")
)

// State is the interface, which defines the common methods for state struct to be handled by state factory
type State interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}

// Serializer has Serialize method to serialize struct to binary data.
type Serializer interface {
	Serialize() ([]byte, error)
}

// Deserializer has Deserialize method to deserialize binary data to struct.
type Deserializer interface {
	Deserialize(data []byte) error
}

// Serialize check if input is Serializer, if it is, use the input's Serialize method, otherwise use Gob.
func Serialize(d interface{}) ([]byte, error) {
	if s, ok := d.(Serializer); ok {
		return s.Serialize()
	}
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)
	if err := e.Encode(d); err != nil {
		return nil, errors.Wrapf(err, "error when serializing %T state", d)
	}
	return buf.Bytes(), nil
}

// Deserialize check if input is Deserializer, if it is, use the input's Deserialize method, otherwise use Gob.
func Deserialize(x interface{}, data []byte) error {
	if s, ok := x.(Deserializer); ok {
		return s.Deserialize(data)
	}
	buf := bytes.NewBuffer(data)
	d := gob.NewDecoder(buf)
	if err := d.Decode(x); err != nil {
		return errors.Wrapf(err, "error when deserializing %v state to %T", data, x)
	}
	return nil
}

// GobBasedSerialize serializes a state into bytes via gob
func GobBasedSerialize(state State) ([]byte, error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)
	if err := e.Encode(state); err != nil {
		return nil, errors.Wrapf(err, "error when serializing %T state", state)
	}
	return buf.Bytes(), nil
}

// GobBasedDeserialize deserialize a state from bytes via gob
func GobBasedDeserialize(state State, data []byte) error {
	buf := bytes.NewBuffer(data)
	e := gob.NewDecoder(buf)
	if err := e.Decode(state); err != nil {
		return errors.Wrapf(err, "error when deserializing %T state", state)
	}
	return nil
}

// SortedSlice represents the state slice in the state factory, which is sorted by the function:
//
//   func(i interface{}, j interface{}) int
//
// The function is expected to output 3 type of values. 0 means i and j are equal; negative integer means i is smaller
// i; and positive integer means i is bigger than j.
//
// SortedSlice will be ser/des as a whole.
type SortedSlice []interface{}

// Serialize serializes the state slice into bytes
func (slice *SortedSlice) Serialize() ([]byte, error) {
	return GobBasedSerialize(slice)
}

// Deserialize deserializes bytes into the state slice
func (slice *SortedSlice) Deserialize(data []byte) error {
	return GobBasedDeserialize(slice, data)
}

// index returns the smallest index of state with value e
func (slice SortedSlice) index(e interface{}, f func(interface{}, interface{}) int) int {
	return sort.Search(len(slice), func(i int) bool {
		return f(slice[i], e) >= 0
	})
}

// Get check if a state exists in the slice
func (slice SortedSlice) Get(e interface{}, f func(interface{}, interface{}) int) (interface{}, bool) {
	idx := slice.index(e, f)
	if idx < len(slice) && f(slice[idx], e) == 0 {
		return slice[idx], true
	}
	return nil, false
}

// Append appends a state into the state slice
func (slice SortedSlice) Append(e interface{}, f func(interface{}, interface{}) int) SortedSlice {
	s := append(slice, e)
	sort.Slice(s, func(i, j int) bool {
		return f(s[i], s[j]) < 0
	})
	return s
}

// Delete deletes a state from the state slice
func (slice SortedSlice) Delete(e interface{}, f func(interface{}, interface{}) int) (SortedSlice, int) {
	idx := slice.index(e, f)
	if idx >= len(slice) || f(slice[idx], e) != 0 {
		return slice, 0
	}
	last := idx + 1
	for last < len(slice) && f(slice[last], e) == 0 {
		last++
	}

	return append(slice[:idx], slice[last:]...), last - idx
}
