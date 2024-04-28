package metavault

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PriceFeed) DecodeMsg(dc *msgp.Reader) (err error) {
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
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "RoundID")
			return
		}
		z.RoundID = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.RoundID))
			if err != nil {
				err = msgp.WrapError(err, "RoundID")
				return
			}
			z.RoundID = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Answer")
			return
		}
		z.Answer = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Answer))
			if err != nil {
				err = msgp.WrapError(err, "Answer")
				return
			}
			z.Answer = msgpencode.DecodeInt(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Answers")
		return
	}
	if z.Answers == nil {
		z.Answers = make(map[string]*big.Int, zb0004)
	} else if len(z.Answers) > 0 {
		for key := range z.Answers {
			delete(z.Answers, key)
		}
	}
	var field []byte
	_ = field
	for zb0004 > 0 {
		zb0004--
		var za0001 string
		var za0002 *big.Int
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Answers")
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "Answers", za0001)
				return
			}
			za0002 = nil
		} else {
			{
				var zb0005 []byte
				zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Answers", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0005)
			}
		}
		z.Answers[za0001] = za0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PriceFeed) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	if z.RoundID == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.RoundID))
		if err != nil {
			err = msgp.WrapError(err, "RoundID")
			return
		}
	}
	if z.Answer == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Answer))
		if err != nil {
			err = msgp.WrapError(err, "Answer")
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.Answers)))
	if err != nil {
		err = msgp.WrapError(err, "Answers")
		return
	}
	for za0001, za0002 := range z.Answers {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Answers")
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
				err = msgp.WrapError(err, "Answers", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PriceFeed) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	if z.RoundID == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.RoundID))
	}
	if z.Answer == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Answer))
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.Answers)))
	for za0001, za0002 := range z.Answers {
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
func (z *PriceFeed) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.RoundID = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.RoundID))
			if err != nil {
				err = msgp.WrapError(err, "RoundID")
				return
			}
			z.RoundID = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Answer = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Answer))
			if err != nil {
				err = msgp.WrapError(err, "Answer")
				return
			}
			z.Answer = msgpencode.DecodeInt(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Answers")
		return
	}
	if z.Answers == nil {
		z.Answers = make(map[string]*big.Int, zb0004)
	} else if len(z.Answers) > 0 {
		for key := range z.Answers {
			delete(z.Answers, key)
		}
	}
	var field []byte
	_ = field
	for zb0004 > 0 {
		var za0001 string
		var za0002 *big.Int
		zb0004--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Answers")
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
				var zb0005 []byte
				zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(za0002))
				if err != nil {
					err = msgp.WrapError(err, "Answers", za0001)
					return
				}
				za0002 = msgpencode.DecodeInt(zb0005)
			}
		}
		z.Answers[za0001] = za0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PriceFeed) Msgsize() (s int) {
	s = 1
	if z.RoundID == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.RoundID))
	}
	if z.Answer == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Answer))
	}
	s += msgp.MapHeaderSize
	if z.Answers != nil {
		for za0001, za0002 := range z.Answers {
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
