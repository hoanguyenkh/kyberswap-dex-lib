package quickperps

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *FastPriceFeedV1) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "DisableFastPriceVoteCount")
			return
		}
		z.DisableFastPriceVoteCount = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.DisableFastPriceVoteCount))
			if err != nil {
				err = msgp.WrapError(err, "DisableFastPriceVoteCount")
				return
			}
			z.DisableFastPriceVoteCount = msgpencode.DecodeInt(zb0002)
		}
	}
	z.IsSpreadEnabled, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsSpreadEnabled")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "LastUpdatedAt")
			return
		}
		z.LastUpdatedAt = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.LastUpdatedAt))
			if err != nil {
				err = msgp.WrapError(err, "LastUpdatedAt")
				return
			}
			z.LastUpdatedAt = msgpencode.DecodeInt(zb0003)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "MaxDeviationBasisPoints")
			return
		}
		z.MaxDeviationBasisPoints = nil
	} else {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(z.MaxDeviationBasisPoints))
			if err != nil {
				err = msgp.WrapError(err, "MaxDeviationBasisPoints")
				return
			}
			z.MaxDeviationBasisPoints = msgpencode.DecodeInt(zb0004)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "MinAuthorizations")
			return
		}
		z.MinAuthorizations = nil
	} else {
		{
			var zb0005 []byte
			zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.MinAuthorizations))
			if err != nil {
				err = msgp.WrapError(err, "MinAuthorizations")
				return
			}
			z.MinAuthorizations = msgpencode.DecodeInt(zb0005)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "PriceDuration")
			return
		}
		z.PriceDuration = nil
	} else {
		{
			var zb0006 []byte
			zb0006, err = dc.ReadBytes(msgpencode.EncodeInt(z.PriceDuration))
			if err != nil {
				err = msgp.WrapError(err, "PriceDuration")
				return
			}
			z.PriceDuration = msgpencode.DecodeInt(zb0006)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "VolBasisPoints")
			return
		}
		z.VolBasisPoints = nil
	} else {
		{
			var zb0007 []byte
			zb0007, err = dc.ReadBytes(msgpencode.EncodeInt(z.VolBasisPoints))
			if err != nil {
				err = msgp.WrapError(err, "VolBasisPoints")
				return
			}
			z.VolBasisPoints = msgpencode.DecodeInt(zb0007)
		}
	}
	var zb0008 uint32
	zb0008, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Prices")
		return
	}
	if z.Prices == nil {
		z.Prices = make(map[string]*big.Int, zb0008)
	} else if len(z.Prices) > 0 {
		for key := range z.Prices {
			delete(z.Prices, key)
		}
	}
	for zb0008 > 0 {
		zb0008--
		var za0001 string
		var za0002 *big.Int
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Prices")
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "Prices", za0001)
				return
			}
			za0002 = nil
		} else {
			{
				var zb0009 []byte
				zb0009, err = dc.ReadBytes(msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Prices", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0009)
			}
		}
		z.Prices[za0001] = za0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FastPriceFeedV1) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 8
	err = en.Append(0x98)
	if err != nil {
		return
	}
	if z.DisableFastPriceVoteCount == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.DisableFastPriceVoteCount))
		if err != nil {
			err = msgp.WrapError(err, "DisableFastPriceVoteCount")
			return
		}
	}
	err = en.WriteBool(z.IsSpreadEnabled)
	if err != nil {
		err = msgp.WrapError(err, "IsSpreadEnabled")
		return
	}
	if z.LastUpdatedAt == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.LastUpdatedAt))
		if err != nil {
			err = msgp.WrapError(err, "LastUpdatedAt")
			return
		}
	}
	if z.MaxDeviationBasisPoints == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.MaxDeviationBasisPoints))
		if err != nil {
			err = msgp.WrapError(err, "MaxDeviationBasisPoints")
			return
		}
	}
	if z.MinAuthorizations == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.MinAuthorizations))
		if err != nil {
			err = msgp.WrapError(err, "MinAuthorizations")
			return
		}
	}
	if z.PriceDuration == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.PriceDuration))
		if err != nil {
			err = msgp.WrapError(err, "PriceDuration")
			return
		}
	}
	if z.VolBasisPoints == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.VolBasisPoints))
		if err != nil {
			err = msgp.WrapError(err, "VolBasisPoints")
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.Prices)))
	if err != nil {
		err = msgp.WrapError(err, "Prices")
		return
	}
	for za0001, za0002 := range z.Prices {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Prices")
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
				err = msgp.WrapError(err, "Prices", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FastPriceFeedV1) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 8
	o = append(o, 0x98)
	if z.DisableFastPriceVoteCount == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.DisableFastPriceVoteCount))
	}
	o = msgp.AppendBool(o, z.IsSpreadEnabled)
	if z.LastUpdatedAt == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.LastUpdatedAt))
	}
	if z.MaxDeviationBasisPoints == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.MaxDeviationBasisPoints))
	}
	if z.MinAuthorizations == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.MinAuthorizations))
	}
	if z.PriceDuration == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.PriceDuration))
	}
	if z.VolBasisPoints == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.VolBasisPoints))
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.Prices)))
	for za0001, za0002 := range z.Prices {
		o = msgp.AppendString(o, za0001)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(za0002))
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FastPriceFeedV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.DisableFastPriceVoteCount = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.DisableFastPriceVoteCount))
			if err != nil {
				err = msgp.WrapError(err, "DisableFastPriceVoteCount")
				return
			}
			z.DisableFastPriceVoteCount = msgpencode.DecodeInt(zb0002)
		}
	}
	z.IsSpreadEnabled, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsSpreadEnabled")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.LastUpdatedAt = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.LastUpdatedAt))
			if err != nil {
				err = msgp.WrapError(err, "LastUpdatedAt")
				return
			}
			z.LastUpdatedAt = msgpencode.DecodeInt(zb0003)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.MaxDeviationBasisPoints = nil
	} else {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.MaxDeviationBasisPoints))
			if err != nil {
				err = msgp.WrapError(err, "MaxDeviationBasisPoints")
				return
			}
			z.MaxDeviationBasisPoints = msgpencode.DecodeInt(zb0004)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.MinAuthorizations = nil
	} else {
		{
			var zb0005 []byte
			zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.MinAuthorizations))
			if err != nil {
				err = msgp.WrapError(err, "MinAuthorizations")
				return
			}
			z.MinAuthorizations = msgpencode.DecodeInt(zb0005)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.PriceDuration = nil
	} else {
		{
			var zb0006 []byte
			zb0006, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.PriceDuration))
			if err != nil {
				err = msgp.WrapError(err, "PriceDuration")
				return
			}
			z.PriceDuration = msgpencode.DecodeInt(zb0006)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.VolBasisPoints = nil
	} else {
		{
			var zb0007 []byte
			zb0007, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.VolBasisPoints))
			if err != nil {
				err = msgp.WrapError(err, "VolBasisPoints")
				return
			}
			z.VolBasisPoints = msgpencode.DecodeInt(zb0007)
		}
	}
	var zb0008 uint32
	zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Prices")
		return
	}
	if z.Prices == nil {
		z.Prices = make(map[string]*big.Int, zb0008)
	} else if len(z.Prices) > 0 {
		for key := range z.Prices {
			delete(z.Prices, key)
		}
	}
	for zb0008 > 0 {
		var za0001 string
		var za0002 *big.Int
		zb0008--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Prices")
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
				var zb0009 []byte
				zb0009, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Prices", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0009)
			}
		}
		z.Prices[za0001] = za0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FastPriceFeedV1) Msgsize() (s int) {
	s = 1
	if z.DisableFastPriceVoteCount == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.DisableFastPriceVoteCount))
	}
	s += msgp.BoolSize
	if z.LastUpdatedAt == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.LastUpdatedAt))
	}
	if z.MaxDeviationBasisPoints == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.MaxDeviationBasisPoints))
	}
	if z.MinAuthorizations == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.MinAuthorizations))
	}
	if z.PriceDuration == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.PriceDuration))
	}
	if z.VolBasisPoints == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.VolBasisPoints))
	}
	s += msgp.MapHeaderSize
	if z.Prices != nil {
		for za0001, za0002 := range z.Prices {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += msgp.NilSize
			} else {
				s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(za0002))
			}
		}
	}
	return
}
