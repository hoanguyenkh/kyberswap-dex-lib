package madmex

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PancakePair) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	if cap(z.Reserves) >= int(zb0002) {
		z.Reserves = (z.Reserves)[:zb0002]
	} else {
		z.Reserves = make([]*big.Int, zb0002)
	}
	for za0001 := range z.Reserves {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "Reserves", za0001)
				return
			}
			z.Reserves[za0001] = nil
		} else {
			{
				var zb0003 []byte
				zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Reserves[za0001]))
				if err != nil {
					err = msgp.WrapError(err, "Reserves", za0001)
					return
				}
				z.Reserves[za0001] = msgpencode.DecodeInt(zb0003)
			}
		}
	}
	z.TimestampLast, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "TimestampLast")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PancakePair) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Reserves)))
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	for za0001 := range z.Reserves {
		if z.Reserves[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(z.Reserves[za0001]))
			if err != nil {
				err = msgp.WrapError(err, "Reserves", za0001)
				return
			}
		}
	}
	err = en.WriteUint32(z.TimestampLast)
	if err != nil {
		err = msgp.WrapError(err, "TimestampLast")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PancakePair) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Reserves)))
	for za0001 := range z.Reserves {
		if z.Reserves[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Reserves[za0001]))
		}
	}
	o = msgp.AppendUint32(o, z.TimestampLast)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PancakePair) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	if cap(z.Reserves) >= int(zb0002) {
		z.Reserves = (z.Reserves)[:zb0002]
	} else {
		z.Reserves = make([]*big.Int, zb0002)
	}
	for za0001 := range z.Reserves {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.Reserves[za0001] = nil
		} else {
			{
				var zb0003 []byte
				zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Reserves[za0001]))
				if err != nil {
					err = msgp.WrapError(err, "Reserves", za0001)
					return
				}
				z.Reserves[za0001] = msgpencode.DecodeInt(zb0003)
			}
		}
	}
	z.TimestampLast, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TimestampLast")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PancakePair) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize
	for za0001 := range z.Reserves {
		if z.Reserves[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Reserves[za0001]))
		}
	}
	s += msgp.Uint32Size
	return
}
