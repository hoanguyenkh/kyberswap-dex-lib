package platypus

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Asset) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	z.Address, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.Decimals, err = dc.ReadUint8()
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Cash")
			return
		}
		z.Cash = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Cash))
			if err != nil {
				err = msgp.WrapError(err, "Cash")
				return
			}
			z.Cash = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Liability")
			return
		}
		z.Liability = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Liability))
			if err != nil {
				err = msgp.WrapError(err, "Liability")
				return
			}
			z.Liability = msgpencode.DecodeInt(zb0003)
		}
	}
	z.UnderlyingToken, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "UnderlyingToken")
		return
	}
	z.AggregateAccount, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "AggregateAccount")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Asset) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 6
	err = en.Append(0x96)
	if err != nil {
		return
	}
	err = en.WriteString(z.Address)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	err = en.WriteUint8(z.Decimals)
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	if z.Cash == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Cash))
		if err != nil {
			err = msgp.WrapError(err, "Cash")
			return
		}
	}
	if z.Liability == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Liability))
		if err != nil {
			err = msgp.WrapError(err, "Liability")
			return
		}
	}
	err = en.WriteString(z.UnderlyingToken)
	if err != nil {
		err = msgp.WrapError(err, "UnderlyingToken")
		return
	}
	err = en.WriteString(z.AggregateAccount)
	if err != nil {
		err = msgp.WrapError(err, "AggregateAccount")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Asset) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 6
	o = append(o, 0x96)
	o = msgp.AppendString(o, z.Address)
	o = msgp.AppendUint8(o, z.Decimals)
	if z.Cash == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Cash))
	}
	if z.Liability == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Liability))
	}
	o = msgp.AppendString(o, z.UnderlyingToken)
	o = msgp.AppendString(o, z.AggregateAccount)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Asset) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	z.Address, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.Decimals, bts, err = msgp.ReadUint8Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Cash = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Cash))
			if err != nil {
				err = msgp.WrapError(err, "Cash")
				return
			}
			z.Cash = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Liability = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Liability))
			if err != nil {
				err = msgp.WrapError(err, "Liability")
				return
			}
			z.Liability = msgpencode.DecodeInt(zb0003)
		}
	}
	z.UnderlyingToken, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "UnderlyingToken")
		return
	}
	z.AggregateAccount, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AggregateAccount")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Asset) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Address) + msgp.Uint8Size
	if z.Cash == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Cash))
	}
	if z.Liability == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Liability))
	}
	s += msgp.StringPrefixSize + len(z.UnderlyingToken) + msgp.StringPrefixSize + len(z.AggregateAccount)
	return
}

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