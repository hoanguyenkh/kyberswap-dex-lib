package gyro3clp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PoolTokenInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
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
			zb0002, err = dc.ReadBytes(msgpencode.EncodeUint256(z.Cash))
			if err != nil {
				err = msgp.WrapError(err, "Cash")
				return
			}
			z.Cash = msgpencode.DecodeUint256(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Managed")
			return
		}
		z.Managed = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeUint256(z.Managed))
			if err != nil {
				err = msgp.WrapError(err, "Managed")
				return
			}
			z.Managed = msgpencode.DecodeUint256(zb0003)
		}
	}
	z.LastChangeBlock, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "LastChangeBlock")
		return
	}
	z.AssetManager, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "AssetManager")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolTokenInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	if z.Cash == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeUint256(z.Cash))
		if err != nil {
			err = msgp.WrapError(err, "Cash")
			return
		}
	}
	if z.Managed == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeUint256(z.Managed))
		if err != nil {
			err = msgp.WrapError(err, "Managed")
			return
		}
	}
	err = en.WriteUint64(z.LastChangeBlock)
	if err != nil {
		err = msgp.WrapError(err, "LastChangeBlock")
		return
	}
	err = en.WriteString(z.AssetManager)
	if err != nil {
		err = msgp.WrapError(err, "AssetManager")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolTokenInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	if z.Cash == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeUint256(z.Cash))
	}
	if z.Managed == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeUint256(z.Managed))
	}
	o = msgp.AppendUint64(o, z.LastChangeBlock)
	o = msgp.AppendString(o, z.AssetManager)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolTokenInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
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
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeUint256(z.Cash))
			if err != nil {
				err = msgp.WrapError(err, "Cash")
				return
			}
			z.Cash = msgpencode.DecodeUint256(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Managed = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeUint256(z.Managed))
			if err != nil {
				err = msgp.WrapError(err, "Managed")
				return
			}
			z.Managed = msgpencode.DecodeUint256(zb0003)
		}
	}
	z.LastChangeBlock, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastChangeBlock")
		return
	}
	z.AssetManager, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AssetManager")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolTokenInfo) Msgsize() (s int) {
	s = 1
	if z.Cash == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeUint256(z.Cash))
	}
	if z.Managed == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeUint256(z.Managed))
	}
	s += msgp.Uint64Size + msgp.StringPrefixSize + len(z.AssetManager)
	return
}
