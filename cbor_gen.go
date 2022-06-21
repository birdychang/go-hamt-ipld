// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package hamt

import (
	"fmt"
	"io"
	"math/big"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// NOTE: This is a generated file, but it has been modified to encode the
// bitfield big.Int as a byte array. The bitfield is only a big.Int because
// thats a convenient type for the operations we need to perform on it, but it
// is fundamentally an array of bytes (bits)

var _ = xerrors.Errorf

var lengthBufNode = []byte{130}

func (t *Node) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufNode); err != nil {
		return err
	}

	cw := cbg.NewCborWriter(w)

	// t.Bitfield (big.Int) (struct)
	{
		var b []byte
		if t.Bitfield != nil {
			b = t.Bitfield.Bytes()
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(b))); err != nil {
			return err
		}
		if _, err := w.Write(b); err != nil {
			return err
		}
	}

	// t.Pointers ([]*hamt.Pointer) (slice)
	if len(t.Pointers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Pointers was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Pointers))); err != nil {
		return err
	}
	for _, v := range t.Pointers {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *Node) UnmarshalCBOR(r io.Reader) error {
	*t = Node{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bitfield (big.Int) (struct)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("big ints should be tagged cbor byte strings")
	}

	if extra > 256 {
		return fmt.Errorf("t.Bitfield: cbor bignum was too large")
	}

	if extra > 0 {
		buf := make([]byte, extra)
		if _, err := io.ReadFull(cr, buf); err != nil {
			return err
		}
		t.Bitfield = big.NewInt(0).SetBytes(buf)
	} else {
		t.Bitfield = big.NewInt(0)
	}
	// t.Pointers ([]*hamt.Pointer) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Pointers: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Pointers = make([]*Pointer, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v Pointer
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Pointers[i] = &v
	}

	return nil
}

var lengthBufKV = []byte{130}

func (t *KV) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufKV); err != nil {
		return err
	}

	cw := cbg.NewCborWriter(w)

	// t.Key ([]uint8) (slice)
	if len(t.Key) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Key was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err
	}

	// t.Value (typegen.Deferred) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *KV) UnmarshalCBOR(r io.Reader) error {
	*t = KV{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Key ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Key: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Key = make([]byte, extra)
	}

	if _, err := io.ReadFull(cr, t.Key[:]); err != nil {
		return err
	}
	// t.Value (typegen.Deferred) (struct)

	{

		t.Value = new(cbg.Deferred)

		if err := t.Value.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("failed to read deferred field: %w", err)
		}
	}
	return nil
}
