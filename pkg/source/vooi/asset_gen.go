package vooi

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/ethereum/go-ethereum/common"
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
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
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
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "MaxSupply")
			return
		}
		z.MaxSupply = nil
	} else {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(z.MaxSupply))
			if err != nil {
				err = msgp.WrapError(err, "MaxSupply")
				return
			}
			z.MaxSupply = msgpencode.DecodeInt(zb0004)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "TotalSupply")
			return
		}
		z.TotalSupply = nil
	} else {
		{
			var zb0005 []byte
			zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.TotalSupply))
			if err != nil {
				err = msgp.WrapError(err, "TotalSupply")
				return
			}
			z.TotalSupply = msgpencode.DecodeInt(zb0005)
		}
	}
	z.Decimals, err = dc.ReadUint8()
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	{
		var zb0006 []byte
		zb0006, err = dc.ReadBytes((common.Address).Bytes(z.Token))
		if err != nil {
			err = msgp.WrapError(err, "Token")
			return
		}
		z.Token = common.BytesToAddress(zb0006)
	}
	z.Active, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "Active")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Asset) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 7
	err = en.Append(0x97)
	if err != nil {
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
	if z.MaxSupply == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.MaxSupply))
		if err != nil {
			err = msgp.WrapError(err, "MaxSupply")
			return
		}
	}
	if z.TotalSupply == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.TotalSupply))
		if err != nil {
			err = msgp.WrapError(err, "TotalSupply")
			return
		}
	}
	err = en.WriteUint8(z.Decimals)
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	err = en.WriteBytes((common.Address).Bytes(z.Token))
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	err = en.WriteBool(z.Active)
	if err != nil {
		err = msgp.WrapError(err, "Active")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Asset) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 7
	o = append(o, 0x97)
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
	if z.MaxSupply == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.MaxSupply))
	}
	if z.TotalSupply == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.TotalSupply))
	}
	o = msgp.AppendUint8(o, z.Decimals)
	o = msgp.AppendBytes(o, (common.Address).Bytes(z.Token))
	o = msgp.AppendBool(o, z.Active)
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
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
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
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.MaxSupply = nil
	} else {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.MaxSupply))
			if err != nil {
				err = msgp.WrapError(err, "MaxSupply")
				return
			}
			z.MaxSupply = msgpencode.DecodeInt(zb0004)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.TotalSupply = nil
	} else {
		{
			var zb0005 []byte
			zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.TotalSupply))
			if err != nil {
				err = msgp.WrapError(err, "TotalSupply")
				return
			}
			z.TotalSupply = msgpencode.DecodeInt(zb0005)
		}
	}
	z.Decimals, bts, err = msgp.ReadUint8Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	{
		var zb0006 []byte
		zb0006, bts, err = msgp.ReadBytesBytes(bts, (common.Address).Bytes(z.Token))
		if err != nil {
			err = msgp.WrapError(err, "Token")
			return
		}
		z.Token = common.BytesToAddress(zb0006)
	}
	z.Active, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Active")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Asset) Msgsize() (s int) {
	s = 1
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
	if z.MaxSupply == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.MaxSupply))
	}
	if z.TotalSupply == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.TotalSupply))
	}
	s += msgp.Uint8Size + msgp.BytesPrefixSize + len((common.Address).Bytes(z.Token)) + msgp.BoolSize
	return
}
