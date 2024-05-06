package metavault

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PriceFeedEnum) DecodeMsg(dc *msgp.Reader) (err error) {
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
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "V1")
			return
		}
		z.V1 = nil
	} else {
		if z.V1 == nil {
			z.V1 = new(FastPriceFeedV1)
		}
		err = z.V1.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "V1")
			return
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "V2")
			return
		}
		z.V2 = nil
	} else {
		if z.V2 == nil {
			z.V2 = new(FastPriceFeedV2)
		}
		err = z.V2.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "V2")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PriceFeedEnum) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	if z.V1 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.V1.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "V1")
			return
		}
	}
	if z.V2 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.V2.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "V2")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PriceFeedEnum) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	if z.V1 == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.V1.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "V1")
			return
		}
	}
	if z.V2 == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.V2.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "V2")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PriceFeedEnum) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.V1 = nil
	} else {
		if z.V1 == nil {
			z.V1 = new(FastPriceFeedV1)
		}
		bts, err = z.V1.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "V1")
			return
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.V2 = nil
	} else {
		if z.V2 == nil {
			z.V2 = new(FastPriceFeedV2)
		}
		bts, err = z.V2.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "V2")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PriceFeedEnum) Msgsize() (s int) {
	s = 1
	if z.V1 == nil {
		s += msgp.NilSize
	} else {
		s += z.V1.Msgsize()
	}
	if z.V2 == nil {
		s += msgp.NilSize
	} else {
		s += z.V2.Msgsize()
	}
	return
}
