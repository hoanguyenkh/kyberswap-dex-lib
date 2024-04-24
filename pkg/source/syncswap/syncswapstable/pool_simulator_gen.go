package syncswapstable

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

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
	z.vaultAddress, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "vaultAddress")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "swapFees")
		return
	}
	if cap(z.swapFees) >= int(zb0002) {
		z.swapFees = (z.swapFees)[:zb0002]
	} else {
		z.swapFees = make([]*big.Int, zb0002)
	}
	for za0001 := range z.swapFees {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "swapFees", za0001)
				return
			}
			z.swapFees[za0001] = nil
		} else {
			{
				var zb0003 []byte
				zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.swapFees[za0001]))
				if err != nil {
					err = msgp.WrapError(err, "swapFees", za0001)
					return
				}
				z.swapFees[za0001] = msgpencode.DecodeInt(zb0003)
			}
		}
	}
	var zb0004 uint32
	zb0004, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "tokenPrecisionMultipliers")
		return
	}
	if cap(z.tokenPrecisionMultipliers) >= int(zb0004) {
		z.tokenPrecisionMultipliers = (z.tokenPrecisionMultipliers)[:zb0004]
	} else {
		z.tokenPrecisionMultipliers = make([]*big.Int, zb0004)
	}
	for za0002 := range z.tokenPrecisionMultipliers {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "tokenPrecisionMultipliers", za0002)
				return
			}
			z.tokenPrecisionMultipliers[za0002] = nil
		} else {
			{
				var zb0005 []byte
				zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.tokenPrecisionMultipliers[za0002]))
				if err != nil {
					err = msgp.WrapError(err, "tokenPrecisionMultipliers", za0002)
					return
				}
				z.tokenPrecisionMultipliers[za0002] = msgpencode.DecodeInt(zb0005)
			}
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
	err = en.WriteString(z.vaultAddress)
	if err != nil {
		err = msgp.WrapError(err, "vaultAddress")
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.swapFees)))
	if err != nil {
		err = msgp.WrapError(err, "swapFees")
		return
	}
	for za0001 := range z.swapFees {
		if z.swapFees[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(z.swapFees[za0001]))
			if err != nil {
				err = msgp.WrapError(err, "swapFees", za0001)
				return
			}
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.tokenPrecisionMultipliers)))
	if err != nil {
		err = msgp.WrapError(err, "tokenPrecisionMultipliers")
		return
	}
	for za0002 := range z.tokenPrecisionMultipliers {
		if z.tokenPrecisionMultipliers[za0002] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(z.tokenPrecisionMultipliers[za0002]))
			if err != nil {
				err = msgp.WrapError(err, "tokenPrecisionMultipliers", za0002)
				return
			}
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
	o = msgp.AppendString(o, z.vaultAddress)
	o = msgp.AppendArrayHeader(o, uint32(len(z.swapFees)))
	for za0001 := range z.swapFees {
		if z.swapFees[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.swapFees[za0001]))
		}
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.tokenPrecisionMultipliers)))
	for za0002 := range z.tokenPrecisionMultipliers {
		if z.tokenPrecisionMultipliers[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.tokenPrecisionMultipliers[za0002]))
		}
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
	z.vaultAddress, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "vaultAddress")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "swapFees")
		return
	}
	if cap(z.swapFees) >= int(zb0002) {
		z.swapFees = (z.swapFees)[:zb0002]
	} else {
		z.swapFees = make([]*big.Int, zb0002)
	}
	for za0001 := range z.swapFees {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.swapFees[za0001] = nil
		} else {
			{
				var zb0003 []byte
				zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.swapFees[za0001]))
				if err != nil {
					err = msgp.WrapError(err, "swapFees", za0001)
					return
				}
				z.swapFees[za0001] = msgpencode.DecodeInt(zb0003)
			}
		}
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "tokenPrecisionMultipliers")
		return
	}
	if cap(z.tokenPrecisionMultipliers) >= int(zb0004) {
		z.tokenPrecisionMultipliers = (z.tokenPrecisionMultipliers)[:zb0004]
	} else {
		z.tokenPrecisionMultipliers = make([]*big.Int, zb0004)
	}
	for za0002 := range z.tokenPrecisionMultipliers {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.tokenPrecisionMultipliers[za0002] = nil
		} else {
			{
				var zb0005 []byte
				zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.tokenPrecisionMultipliers[za0002]))
				if err != nil {
					err = msgp.WrapError(err, "tokenPrecisionMultipliers", za0002)
					return
				}
				z.tokenPrecisionMultipliers[za0002] = msgpencode.DecodeInt(zb0005)
			}
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
	s = 1 + z.Pool.Msgsize() + msgp.StringPrefixSize + len(z.vaultAddress) + msgp.ArrayHeaderSize
	for za0001 := range z.swapFees {
		if z.swapFees[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.swapFees[za0001]))
		}
	}
	s += msgp.ArrayHeaderSize
	for za0002 := range z.tokenPrecisionMultipliers {
		if z.tokenPrecisionMultipliers[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.tokenPrecisionMultipliers[za0002]))
		}
	}
	s += z.gas.Msgsize()
	return
}