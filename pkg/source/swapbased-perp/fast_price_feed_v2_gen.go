package swapbasedperp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *FastPriceFeedV2) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 12 {
		err = msgp.ArrayError{Wanted: 12, Got: zb0001}
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
			err = msgp.WrapError(err, "MaxPriceUpdateDelay")
			return
		}
		z.MaxPriceUpdateDelay = nil
	} else {
		{
			var zb0007 []byte
			zb0007, err = dc.ReadBytes(msgpencode.EncodeInt(z.MaxPriceUpdateDelay))
			if err != nil {
				err = msgp.WrapError(err, "MaxPriceUpdateDelay")
				return
			}
			z.MaxPriceUpdateDelay = msgpencode.DecodeInt(zb0007)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "SpreadBasisPointsIfChainError")
			return
		}
		z.SpreadBasisPointsIfChainError = nil
	} else {
		{
			var zb0008 []byte
			zb0008, err = dc.ReadBytes(msgpencode.EncodeInt(z.SpreadBasisPointsIfChainError))
			if err != nil {
				err = msgp.WrapError(err, "SpreadBasisPointsIfChainError")
				return
			}
			z.SpreadBasisPointsIfChainError = msgpencode.DecodeInt(zb0008)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "SpreadBasisPointsIfInactive")
			return
		}
		z.SpreadBasisPointsIfInactive = nil
	} else {
		{
			var zb0009 []byte
			zb0009, err = dc.ReadBytes(msgpencode.EncodeInt(z.SpreadBasisPointsIfInactive))
			if err != nil {
				err = msgp.WrapError(err, "SpreadBasisPointsIfInactive")
				return
			}
			z.SpreadBasisPointsIfInactive = msgpencode.DecodeInt(zb0009)
		}
	}
	var zb0010 uint32
	zb0010, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Prices")
		return
	}
	if z.Prices == nil {
		z.Prices = make(map[string]*big.Int, zb0010)
	} else if len(z.Prices) > 0 {
		for key := range z.Prices {
			delete(z.Prices, key)
		}
	}
	for zb0010 > 0 {
		zb0010--
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
				var zb0011 []byte
				zb0011, err = dc.ReadBytes(msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Prices", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0011)
			}
		}
		z.Prices[za0001] = za0002
	}
	var zb0012 uint32
	zb0012, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "PriceData")
		return
	}
	if z.PriceData == nil {
		z.PriceData = make(map[string]PriceDataItem, zb0012)
	} else if len(z.PriceData) > 0 {
		for key := range z.PriceData {
			delete(z.PriceData, key)
		}
	}
	for zb0012 > 0 {
		zb0012--
		var za0003 string
		var za0004 PriceDataItem
		za0003, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "PriceData")
			return
		}
		err = za0004.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "PriceData", za0003)
			return
		}
		z.PriceData[za0003] = za0004
	}
	var zb0013 uint32
	zb0013, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
		return
	}
	if z.MaxCumulativeDeltaDiffs == nil {
		z.MaxCumulativeDeltaDiffs = make(map[string]*big.Int, zb0013)
	} else if len(z.MaxCumulativeDeltaDiffs) > 0 {
		for key := range z.MaxCumulativeDeltaDiffs {
			delete(z.MaxCumulativeDeltaDiffs, key)
		}
	}
	for zb0013 > 0 {
		zb0013--
		var za0005 string
		var za0006 *big.Int
		za0005, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs", za0005)
				return
			}
			za0006 = nil
		} else {
			{
				var zb0014 []byte
				zb0014, err = dc.ReadBytes(msgpencode.EncodeInt(za0006))
				if err != nil {
					err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs", za0005)
					return
				}
				za0006 = msgpencode.DecodeInt(zb0014)
			}
		}
		z.MaxCumulativeDeltaDiffs[za0005] = za0006
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FastPriceFeedV2) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 12
	err = en.Append(0x9c)
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
	if z.MaxPriceUpdateDelay == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.MaxPriceUpdateDelay))
		if err != nil {
			err = msgp.WrapError(err, "MaxPriceUpdateDelay")
			return
		}
	}
	if z.SpreadBasisPointsIfChainError == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.SpreadBasisPointsIfChainError))
		if err != nil {
			err = msgp.WrapError(err, "SpreadBasisPointsIfChainError")
			return
		}
	}
	if z.SpreadBasisPointsIfInactive == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.SpreadBasisPointsIfInactive))
		if err != nil {
			err = msgp.WrapError(err, "SpreadBasisPointsIfInactive")
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
	err = en.WriteMapHeader(uint32(len(z.PriceData)))
	if err != nil {
		err = msgp.WrapError(err, "PriceData")
		return
	}
	for za0003, za0004 := range z.PriceData {
		err = en.WriteString(za0003)
		if err != nil {
			err = msgp.WrapError(err, "PriceData")
			return
		}
		err = za0004.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "PriceData", za0003)
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.MaxCumulativeDeltaDiffs)))
	if err != nil {
		err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
		return
	}
	for za0005, za0006 := range z.MaxCumulativeDeltaDiffs {
		err = en.WriteString(za0005)
		if err != nil {
			err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
			return
		}
		if za0006 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(za0006))
			if err != nil {
				err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs", za0005)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FastPriceFeedV2) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 12
	o = append(o, 0x9c)
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
	if z.MaxPriceUpdateDelay == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.MaxPriceUpdateDelay))
	}
	if z.SpreadBasisPointsIfChainError == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.SpreadBasisPointsIfChainError))
	}
	if z.SpreadBasisPointsIfInactive == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.SpreadBasisPointsIfInactive))
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
	o = msgp.AppendMapHeader(o, uint32(len(z.PriceData)))
	for za0003, za0004 := range z.PriceData {
		o = msgp.AppendString(o, za0003)
		o, err = za0004.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "PriceData", za0003)
			return
		}
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.MaxCumulativeDeltaDiffs)))
	for za0005, za0006 := range z.MaxCumulativeDeltaDiffs {
		o = msgp.AppendString(o, za0005)
		if za0006 == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(za0006))
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FastPriceFeedV2) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 12 {
		err = msgp.ArrayError{Wanted: 12, Got: zb0001}
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
		z.MaxPriceUpdateDelay = nil
	} else {
		{
			var zb0007 []byte
			zb0007, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.MaxPriceUpdateDelay))
			if err != nil {
				err = msgp.WrapError(err, "MaxPriceUpdateDelay")
				return
			}
			z.MaxPriceUpdateDelay = msgpencode.DecodeInt(zb0007)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.SpreadBasisPointsIfChainError = nil
	} else {
		{
			var zb0008 []byte
			zb0008, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.SpreadBasisPointsIfChainError))
			if err != nil {
				err = msgp.WrapError(err, "SpreadBasisPointsIfChainError")
				return
			}
			z.SpreadBasisPointsIfChainError = msgpencode.DecodeInt(zb0008)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.SpreadBasisPointsIfInactive = nil
	} else {
		{
			var zb0009 []byte
			zb0009, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.SpreadBasisPointsIfInactive))
			if err != nil {
				err = msgp.WrapError(err, "SpreadBasisPointsIfInactive")
				return
			}
			z.SpreadBasisPointsIfInactive = msgpencode.DecodeInt(zb0009)
		}
	}
	var zb0010 uint32
	zb0010, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Prices")
		return
	}
	if z.Prices == nil {
		z.Prices = make(map[string]*big.Int, zb0010)
	} else if len(z.Prices) > 0 {
		for key := range z.Prices {
			delete(z.Prices, key)
		}
	}
	for zb0010 > 0 {
		var za0001 string
		var za0002 *big.Int
		zb0010--
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
				var zb0011 []byte
				zb0011, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Prices", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0011)
			}
		}
		z.Prices[za0001] = za0002
	}
	var zb0012 uint32
	zb0012, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "PriceData")
		return
	}
	if z.PriceData == nil {
		z.PriceData = make(map[string]PriceDataItem, zb0012)
	} else if len(z.PriceData) > 0 {
		for key := range z.PriceData {
			delete(z.PriceData, key)
		}
	}
	for zb0012 > 0 {
		var za0003 string
		var za0004 PriceDataItem
		zb0012--
		za0003, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "PriceData")
			return
		}
		bts, err = za0004.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "PriceData", za0003)
			return
		}
		z.PriceData[za0003] = za0004
	}
	var zb0013 uint32
	zb0013, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
		return
	}
	if z.MaxCumulativeDeltaDiffs == nil {
		z.MaxCumulativeDeltaDiffs = make(map[string]*big.Int, zb0013)
	} else if len(z.MaxCumulativeDeltaDiffs) > 0 {
		for key := range z.MaxCumulativeDeltaDiffs {
			delete(z.MaxCumulativeDeltaDiffs, key)
		}
	}
	for zb0013 > 0 {
		var za0005 string
		var za0006 *big.Int
		zb0013--
		za0005, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs")
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			za0006 = nil
		} else {
			{
				var zb0014 []byte
				zb0014, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0006))
				if err != nil {
					err = msgp.WrapError(err, "MaxCumulativeDeltaDiffs", za0005)
					return
				}
				za0006 = msgpencode.DecodeInt(zb0014)
			}
		}
		z.MaxCumulativeDeltaDiffs[za0005] = za0006
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FastPriceFeedV2) Msgsize() (s int) {
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
	if z.MaxPriceUpdateDelay == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.MaxPriceUpdateDelay))
	}
	if z.SpreadBasisPointsIfChainError == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.SpreadBasisPointsIfChainError))
	}
	if z.SpreadBasisPointsIfInactive == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.SpreadBasisPointsIfInactive))
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
	s += msgp.MapHeaderSize
	if z.PriceData != nil {
		for za0003, za0004 := range z.PriceData {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + za0004.Msgsize()
		}
	}
	s += msgp.MapHeaderSize
	if z.MaxCumulativeDeltaDiffs != nil {
		for za0005, za0006 := range z.MaxCumulativeDeltaDiffs {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005)
			if za0006 == nil {
				s += msgp.NilSize
			} else {
				s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(za0006))
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PriceDataItem) DecodeMsg(dc *msgp.Reader) (err error) {
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
			err = msgp.WrapError(err, "RefPrice")
			return
		}
		z.RefPrice = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.RefPrice))
			if err != nil {
				err = msgp.WrapError(err, "RefPrice")
				return
			}
			z.RefPrice = msgpencode.DecodeInt(zb0002)
		}
	}
	z.RefTime, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "RefTime")
		return
	}
	z.CumulativeRefDelta, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "CumulativeRefDelta")
		return
	}
	z.CumulativeFastDelta, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "CumulativeFastDelta")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PriceDataItem) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	if z.RefPrice == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.RefPrice))
		if err != nil {
			err = msgp.WrapError(err, "RefPrice")
			return
		}
	}
	err = en.WriteUint64(z.RefTime)
	if err != nil {
		err = msgp.WrapError(err, "RefTime")
		return
	}
	err = en.WriteUint64(z.CumulativeRefDelta)
	if err != nil {
		err = msgp.WrapError(err, "CumulativeRefDelta")
		return
	}
	err = en.WriteUint64(z.CumulativeFastDelta)
	if err != nil {
		err = msgp.WrapError(err, "CumulativeFastDelta")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PriceDataItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	if z.RefPrice == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.RefPrice))
	}
	o = msgp.AppendUint64(o, z.RefTime)
	o = msgp.AppendUint64(o, z.CumulativeRefDelta)
	o = msgp.AppendUint64(o, z.CumulativeFastDelta)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PriceDataItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		z.RefPrice = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.RefPrice))
			if err != nil {
				err = msgp.WrapError(err, "RefPrice")
				return
			}
			z.RefPrice = msgpencode.DecodeInt(zb0002)
		}
	}
	z.RefTime, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "RefTime")
		return
	}
	z.CumulativeRefDelta, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CumulativeRefDelta")
		return
	}
	z.CumulativeFastDelta, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CumulativeFastDelta")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PriceDataItem) Msgsize() (s int) {
	s = 1
	if z.RefPrice == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.RefPrice))
	}
	s += msgp.Uint64Size + msgp.Uint64Size + msgp.Uint64Size
	return
}
