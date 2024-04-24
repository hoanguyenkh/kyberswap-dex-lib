package fxdx

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *FeeUtilsV2) DecodeMsg(dc *msgp.Reader) (err error) {
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
	z.Address, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.IsInitialized, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsInitialized")
		return
	}
	z.IsActive, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsActive")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "FeeMultiplierIfInactive")
			return
		}
		z.FeeMultiplierIfInactive = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.FeeMultiplierIfInactive))
			if err != nil {
				err = msgp.WrapError(err, "FeeMultiplierIfInactive")
				return
			}
			z.FeeMultiplierIfInactive = msgpencode.DecodeInt(zb0002)
		}
	}
	z.HasDynamicFees, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "HasDynamicFees")
		return
	}
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "TaxBasisPoints")
		return
	}
	if z.TaxBasisPoints == nil {
		z.TaxBasisPoints = make(map[string]*big.Int, zb0003)
	} else if len(z.TaxBasisPoints) > 0 {
		for key := range z.TaxBasisPoints {
			delete(z.TaxBasisPoints, key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var za0001 string
		var za0002 *big.Int
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "TaxBasisPoints")
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "TaxBasisPoints", za0001)
				return
			}
			za0002 = nil
		} else {
			{
				var zb0004 []byte
				zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "TaxBasisPoints", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0004)
			}
		}
		z.TaxBasisPoints[za0001] = za0002
	}
	var zb0005 uint32
	zb0005, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "SwapFeeBasisPoints")
		return
	}
	if z.SwapFeeBasisPoints == nil {
		z.SwapFeeBasisPoints = make(map[string]*big.Int, zb0005)
	} else if len(z.SwapFeeBasisPoints) > 0 {
		for key := range z.SwapFeeBasisPoints {
			delete(z.SwapFeeBasisPoints, key)
		}
	}
	for zb0005 > 0 {
		zb0005--
		var za0003 string
		var za0004 *big.Int
		za0003, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "SwapFeeBasisPoints")
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "SwapFeeBasisPoints", za0003)
				return
			}
			za0004 = nil
		} else {
			{
				var zb0006 []byte
				zb0006, err = dc.ReadBytes(msgpencode.EncodeInt(za0004))
				if err != nil {
					err = msgp.WrapError(err, "SwapFeeBasisPoints", za0003)
					return
				}
				za0004 = msgpencode.DecodeInt(zb0006)
			}
		}
		z.SwapFeeBasisPoints[za0003] = za0004
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FeeUtilsV2) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 7
	err = en.Append(0x97)
	if err != nil {
		return
	}
	err = en.WriteString(z.Address)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	err = en.WriteBool(z.IsInitialized)
	if err != nil {
		err = msgp.WrapError(err, "IsInitialized")
		return
	}
	err = en.WriteBool(z.IsActive)
	if err != nil {
		err = msgp.WrapError(err, "IsActive")
		return
	}
	if z.FeeMultiplierIfInactive == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.FeeMultiplierIfInactive))
		if err != nil {
			err = msgp.WrapError(err, "FeeMultiplierIfInactive")
			return
		}
	}
	err = en.WriteBool(z.HasDynamicFees)
	if err != nil {
		err = msgp.WrapError(err, "HasDynamicFees")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.TaxBasisPoints)))
	if err != nil {
		err = msgp.WrapError(err, "TaxBasisPoints")
		return
	}
	for za0001, za0002 := range z.TaxBasisPoints {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "TaxBasisPoints")
			return
		}
		if za0002 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(za0002))
			if err != nil {
				err = msgp.WrapError(err, "TaxBasisPoints", za0001)
				return
			}
		}
	}
	err = en.WriteMapHeader(uint32(len(z.SwapFeeBasisPoints)))
	if err != nil {
		err = msgp.WrapError(err, "SwapFeeBasisPoints")
		return
	}
	for za0003, za0004 := range z.SwapFeeBasisPoints {
		err = en.WriteString(za0003)
		if err != nil {
			err = msgp.WrapError(err, "SwapFeeBasisPoints")
			return
		}
		if za0004 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(za0004))
			if err != nil {
				err = msgp.WrapError(err, "SwapFeeBasisPoints", za0003)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FeeUtilsV2) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 7
	o = append(o, 0x97)
	o = msgp.AppendString(o, z.Address)
	o = msgp.AppendBool(o, z.IsInitialized)
	o = msgp.AppendBool(o, z.IsActive)
	if z.FeeMultiplierIfInactive == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.FeeMultiplierIfInactive))
	}
	o = msgp.AppendBool(o, z.HasDynamicFees)
	o = msgp.AppendMapHeader(o, uint32(len(z.TaxBasisPoints)))
	for za0001, za0002 := range z.TaxBasisPoints {
		o = msgp.AppendString(o, za0001)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(za0002))
		}
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.SwapFeeBasisPoints)))
	for za0003, za0004 := range z.SwapFeeBasisPoints {
		o = msgp.AppendString(o, za0003)
		if za0004 == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(za0004))
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FeeUtilsV2) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	z.Address, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.IsInitialized, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsInitialized")
		return
	}
	z.IsActive, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsActive")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.FeeMultiplierIfInactive = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.FeeMultiplierIfInactive))
			if err != nil {
				err = msgp.WrapError(err, "FeeMultiplierIfInactive")
				return
			}
			z.FeeMultiplierIfInactive = msgpencode.DecodeInt(zb0002)
		}
	}
	z.HasDynamicFees, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "HasDynamicFees")
		return
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TaxBasisPoints")
		return
	}
	if z.TaxBasisPoints == nil {
		z.TaxBasisPoints = make(map[string]*big.Int, zb0003)
	} else if len(z.TaxBasisPoints) > 0 {
		for key := range z.TaxBasisPoints {
			delete(z.TaxBasisPoints, key)
		}
	}
	for zb0003 > 0 {
		var za0001 string
		var za0002 *big.Int
		zb0003--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "TaxBasisPoints")
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			za0002 = nil
		} else {
			{
				var zb0004 []byte
				zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "TaxBasisPoints", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0004)
			}
		}
		z.TaxBasisPoints[za0001] = za0002
	}
	var zb0005 uint32
	zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "SwapFeeBasisPoints")
		return
	}
	if z.SwapFeeBasisPoints == nil {
		z.SwapFeeBasisPoints = make(map[string]*big.Int, zb0005)
	} else if len(z.SwapFeeBasisPoints) > 0 {
		for key := range z.SwapFeeBasisPoints {
			delete(z.SwapFeeBasisPoints, key)
		}
	}
	for zb0005 > 0 {
		var za0003 string
		var za0004 *big.Int
		zb0005--
		za0003, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "SwapFeeBasisPoints")
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			za0004 = nil
		} else {
			{
				var zb0006 []byte
				zb0006, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0004))
				if err != nil {
					err = msgp.WrapError(err, "SwapFeeBasisPoints", za0003)
					return
				}
				za0004 = msgpencode.DecodeInt(zb0006)
			}
		}
		z.SwapFeeBasisPoints[za0003] = za0004
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FeeUtilsV2) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Address) + msgp.BoolSize + msgp.BoolSize
	if z.FeeMultiplierIfInactive == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.FeeMultiplierIfInactive))
	}
	s += msgp.BoolSize + msgp.MapHeaderSize
	if z.TaxBasisPoints != nil {
		for za0001, za0002 := range z.TaxBasisPoints {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += msgp.NilSize
			} else {
				s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(za0002))
			}
		}
	}
	s += msgp.MapHeaderSize
	if z.SwapFeeBasisPoints != nil {
		for za0003, za0004 := range z.SwapFeeBasisPoints {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003)
			if za0004 == nil {
				s += msgp.NilSize
			} else {
				s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(za0004))
			}
		}
	}
	return
}
