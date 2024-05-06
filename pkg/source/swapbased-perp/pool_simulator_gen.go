package swapbasedperp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Gas) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Swap, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Gas) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Swap)
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Gas) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt64(o, z.Swap)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Gas) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Swap, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Gas) Msgsize() (s int) {
	s = 1 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolSimulator) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	err = z.Pool.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "vault")
			return
		}
		z.vault = nil
	} else {
		if z.vault == nil {
			z.vault = new(Vault)
		}
		err = z.vault.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "vault")
			return
		}
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.gas.Swap, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolSimulator) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = z.Pool.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if z.vault == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.vault.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "vault")
			return
		}
	}
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.gas.Swap)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolSimulator) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o, err = z.Pool.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if z.vault == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.vault.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "vault")
			return
		}
	}
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt64(o, z.gas.Swap)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolSimulator) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	bts, err = z.Pool.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.vault = nil
	} else {
		if z.vault == nil {
			z.vault = new(Vault)
		}
		bts, err = z.vault.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "vault")
			return
		}
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.gas.Swap, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolSimulator) Msgsize() (s int) {
	s = 1 + z.Pool.Msgsize()
	if z.vault == nil {
		s += msgp.NilSize
	} else {
		s += z.vault.Msgsize()
	}
	s += 1 + msgp.Int64Size
	return
}
