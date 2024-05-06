package makerpsm

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/ethereum/go-ethereum/common"
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
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.BuyGem, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "BuyGem")
		return
	}
	z.SellGem, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "SellGem")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Gas) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BuyGem)
	if err != nil {
		err = msgp.WrapError(err, "BuyGem")
		return
	}
	err = en.WriteInt64(z.SellGem)
	if err != nil {
		err = msgp.WrapError(err, "SellGem")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Gas) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt64(o, z.BuyGem)
	o = msgp.AppendInt64(o, z.SellGem)
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
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.BuyGem, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BuyGem")
		return
	}
	z.SellGem, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "SellGem")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Gas) Msgsize() (s int) {
	s = 1 + msgp.Int64Size + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ILK) DecodeMsg(dc *msgp.Reader) (err error) {
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
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Art")
			return
		}
		z.Art = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Art))
			if err != nil {
				err = msgp.WrapError(err, "Art")
				return
			}
			z.Art = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Rate")
			return
		}
		z.Rate = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Rate))
			if err != nil {
				err = msgp.WrapError(err, "Rate")
				return
			}
			z.Rate = msgpencode.DecodeInt(zb0003)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Line")
			return
		}
		z.Line = nil
	} else {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(z.Line))
			if err != nil {
				err = msgp.WrapError(err, "Line")
				return
			}
			z.Line = msgpencode.DecodeInt(zb0004)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Spot")
			return
		}
		z.Spot = nil
	} else {
		{
			var zb0005 []byte
			zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.Spot))
			if err != nil {
				err = msgp.WrapError(err, "Spot")
				return
			}
			z.Spot = msgpencode.DecodeInt(zb0005)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Dust")
			return
		}
		z.Dust = nil
	} else {
		{
			var zb0006 []byte
			zb0006, err = dc.ReadBytes(msgpencode.EncodeInt(z.Dust))
			if err != nil {
				err = msgp.WrapError(err, "Dust")
				return
			}
			z.Dust = msgpencode.DecodeInt(zb0006)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ILK) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	if z.Art == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Art))
		if err != nil {
			err = msgp.WrapError(err, "Art")
			return
		}
	}
	if z.Rate == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Rate))
		if err != nil {
			err = msgp.WrapError(err, "Rate")
			return
		}
	}
	if z.Line == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Line))
		if err != nil {
			err = msgp.WrapError(err, "Line")
			return
		}
	}
	if z.Spot == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Spot))
		if err != nil {
			err = msgp.WrapError(err, "Spot")
			return
		}
	}
	if z.Dust == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Dust))
		if err != nil {
			err = msgp.WrapError(err, "Dust")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ILK) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	if z.Art == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Art))
	}
	if z.Rate == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Rate))
	}
	if z.Line == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Line))
	}
	if z.Spot == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Spot))
	}
	if z.Dust == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Dust))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ILK) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Art = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Art))
			if err != nil {
				err = msgp.WrapError(err, "Art")
				return
			}
			z.Art = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Rate = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Rate))
			if err != nil {
				err = msgp.WrapError(err, "Rate")
				return
			}
			z.Rate = msgpencode.DecodeInt(zb0003)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Line = nil
	} else {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Line))
			if err != nil {
				err = msgp.WrapError(err, "Line")
				return
			}
			z.Line = msgpencode.DecodeInt(zb0004)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Spot = nil
	} else {
		{
			var zb0005 []byte
			zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Spot))
			if err != nil {
				err = msgp.WrapError(err, "Spot")
				return
			}
			z.Spot = msgpencode.DecodeInt(zb0005)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Dust = nil
	} else {
		{
			var zb0006 []byte
			zb0006, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Dust))
			if err != nil {
				err = msgp.WrapError(err, "Dust")
				return
			}
			z.Dust = msgpencode.DecodeInt(zb0006)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ILK) Msgsize() (s int) {
	s = 1
	if z.Art == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Art))
	}
	if z.Rate == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Rate))
	}
	if z.Line == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Line))
	}
	if z.Spot == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Spot))
	}
	if z.Dust == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Dust))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PSM) DecodeMsg(dc *msgp.Reader) (err error) {
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
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "TIn")
			return
		}
		z.TIn = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.TIn))
			if err != nil {
				err = msgp.WrapError(err, "TIn")
				return
			}
			z.TIn = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "TOut")
			return
		}
		z.TOut = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.TOut))
			if err != nil {
				err = msgp.WrapError(err, "TOut")
				return
			}
			z.TOut = msgpencode.DecodeInt(zb0003)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Vat")
			return
		}
		z.Vat = nil
	} else {
		if z.Vat == nil {
			z.Vat = new(Vat)
		}
		err = z.Vat.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Vat")
			return
		}
	}
	{
		var zb0004 []byte
		zb0004, err = dc.ReadBytes((common.Address).Bytes(z.VatAddress))
		if err != nil {
			err = msgp.WrapError(err, "VatAddress")
			return
		}
		z.VatAddress = common.BytesToAddress(zb0004)
	}
	err = dc.ReadExactBytes((z.ILK)[:])
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "To18ConversionFactor")
			return
		}
		z.To18ConversionFactor = nil
	} else {
		{
			var zb0005 []byte
			zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.To18ConversionFactor))
			if err != nil {
				err = msgp.WrapError(err, "To18ConversionFactor")
				return
			}
			z.To18ConversionFactor = msgpencode.DecodeInt(zb0005)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PSM) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 6
	err = en.Append(0x96)
	if err != nil {
		return
	}
	if z.TIn == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.TIn))
		if err != nil {
			err = msgp.WrapError(err, "TIn")
			return
		}
	}
	if z.TOut == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.TOut))
		if err != nil {
			err = msgp.WrapError(err, "TOut")
			return
		}
	}
	if z.Vat == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Vat.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Vat")
			return
		}
	}
	err = en.WriteBytes((common.Address).Bytes(z.VatAddress))
	if err != nil {
		err = msgp.WrapError(err, "VatAddress")
		return
	}
	err = en.WriteBytes((z.ILK)[:])
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if z.To18ConversionFactor == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.To18ConversionFactor))
		if err != nil {
			err = msgp.WrapError(err, "To18ConversionFactor")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PSM) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 6
	o = append(o, 0x96)
	if z.TIn == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.TIn))
	}
	if z.TOut == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.TOut))
	}
	if z.Vat == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Vat.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Vat")
			return
		}
	}
	o = msgp.AppendBytes(o, (common.Address).Bytes(z.VatAddress))
	o = msgp.AppendBytes(o, (z.ILK)[:])
	if z.To18ConversionFactor == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.To18ConversionFactor))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PSM) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.TIn = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.TIn))
			if err != nil {
				err = msgp.WrapError(err, "TIn")
				return
			}
			z.TIn = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.TOut = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.TOut))
			if err != nil {
				err = msgp.WrapError(err, "TOut")
				return
			}
			z.TOut = msgpencode.DecodeInt(zb0003)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Vat = nil
	} else {
		if z.Vat == nil {
			z.Vat = new(Vat)
		}
		bts, err = z.Vat.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Vat")
			return
		}
	}
	{
		var zb0004 []byte
		zb0004, bts, err = msgp.ReadBytesBytes(bts, (common.Address).Bytes(z.VatAddress))
		if err != nil {
			err = msgp.WrapError(err, "VatAddress")
			return
		}
		z.VatAddress = common.BytesToAddress(zb0004)
	}
	bts, err = msgp.ReadExactBytes(bts, (z.ILK)[:])
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.To18ConversionFactor = nil
	} else {
		{
			var zb0005 []byte
			zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.To18ConversionFactor))
			if err != nil {
				err = msgp.WrapError(err, "To18ConversionFactor")
				return
			}
			z.To18ConversionFactor = msgpencode.DecodeInt(zb0005)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PSM) Msgsize() (s int) {
	s = 1
	if z.TIn == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.TIn))
	}
	if z.TOut == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.TOut))
	}
	if z.Vat == nil {
		s += msgp.NilSize
	} else {
		s += z.Vat.Msgsize()
	}
	s += msgp.BytesPrefixSize + len((common.Address).Bytes(z.VatAddress)) + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize))
	if z.To18ConversionFactor == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.To18ConversionFactor))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Vat) DecodeMsg(dc *msgp.Reader) (err error) {
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
	err = z.ILK.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Debt")
			return
		}
		z.Debt = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Debt))
			if err != nil {
				err = msgp.WrapError(err, "Debt")
				return
			}
			z.Debt = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Line")
			return
		}
		z.Line = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Line))
			if err != nil {
				err = msgp.WrapError(err, "Line")
				return
			}
			z.Line = msgpencode.DecodeInt(zb0003)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Vat) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = z.ILK.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if z.Debt == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Debt))
		if err != nil {
			err = msgp.WrapError(err, "Debt")
			return
		}
	}
	if z.Line == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Line))
		if err != nil {
			err = msgp.WrapError(err, "Line")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Vat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o, err = z.ILK.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if z.Debt == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Debt))
	}
	if z.Line == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Line))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Vat) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	bts, err = z.ILK.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "ILK")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Debt = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Debt))
			if err != nil {
				err = msgp.WrapError(err, "Debt")
				return
			}
			z.Debt = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Line = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Line))
			if err != nil {
				err = msgp.WrapError(err, "Line")
				return
			}
			z.Line = msgpencode.DecodeInt(zb0003)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Vat) Msgsize() (s int) {
	s = 1 + z.ILK.Msgsize()
	if z.Debt == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Debt))
	}
	if z.Line == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Line))
	}
	return
}
