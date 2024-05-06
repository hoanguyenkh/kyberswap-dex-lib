package fraxswap

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PoolSimulator) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
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
			err = msgp.WrapError(err, "Fee")
			return
		}
		z.Fee = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Fee))
			if err != nil {
				err = msgp.WrapError(err, "Fee")
				return
			}
			z.Fee = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Reserve0")
			return
		}
		z.Reserve0 = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Reserve0))
			if err != nil {
				err = msgp.WrapError(err, "Reserve0")
				return
			}
			z.Reserve0 = msgpencode.DecodeInt(zb0003)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Reserve1")
			return
		}
		z.Reserve1 = nil
	} else {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(z.Reserve1))
			if err != nil {
				err = msgp.WrapError(err, "Reserve1")
				return
			}
			z.Reserve1 = msgpencode.DecodeInt(zb0004)
		}
	}
	err = z.gas.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolSimulator) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = z.Pool.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if z.Fee == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Fee))
		if err != nil {
			err = msgp.WrapError(err, "Fee")
			return
		}
	}
	if z.Reserve0 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Reserve0))
		if err != nil {
			err = msgp.WrapError(err, "Reserve0")
			return
		}
	}
	if z.Reserve1 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Reserve1))
		if err != nil {
			err = msgp.WrapError(err, "Reserve1")
			return
		}
	}
	err = z.gas.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolSimulator) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o, err = z.Pool.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	if z.Fee == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Fee))
	}
	if z.Reserve0 == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Reserve0))
	}
	if z.Reserve1 == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Reserve1))
	}
	o, err = z.gas.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
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
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
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
		z.Fee = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Fee))
			if err != nil {
				err = msgp.WrapError(err, "Fee")
				return
			}
			z.Fee = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Reserve0 = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Reserve0))
			if err != nil {
				err = msgp.WrapError(err, "Reserve0")
				return
			}
			z.Reserve0 = msgpencode.DecodeInt(zb0003)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Reserve1 = nil
	} else {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Reserve1))
			if err != nil {
				err = msgp.WrapError(err, "Reserve1")
				return
			}
			z.Reserve1 = msgpencode.DecodeInt(zb0004)
		}
	}
	bts, err = z.gas.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolSimulator) Msgsize() (s int) {
	s = 1 + z.Pool.Msgsize()
	if z.Fee == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Fee))
	}
	if z.Reserve0 == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Reserve0))
	}
	if z.Reserve1 == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Reserve1))
	}
	s += z.gas.Msgsize()
	return
}
