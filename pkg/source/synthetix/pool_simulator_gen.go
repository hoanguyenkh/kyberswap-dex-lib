package synthetix

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
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
	z.ExchangeAtomically, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "ExchangeAtomically")
		return
	}
	z.Exchange, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
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
	err = en.WriteInt64(z.ExchangeAtomically)
	if err != nil {
		err = msgp.WrapError(err, "ExchangeAtomically")
		return
	}
	err = en.WriteInt64(z.Exchange)
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Gas) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt64(o, z.ExchangeAtomically)
	o = msgp.AppendInt64(o, z.Exchange)
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
	z.ExchangeAtomically, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ExchangeAtomically")
		return
	}
	z.Exchange, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
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
func (z *PoolSimulator) DecodeMsg(dc *msgp.Reader) (err error) {
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
	err = z.Pool.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	{
		var zb0002 uint
		zb0002, err = dc.ReadUint()
		if err != nil {
			err = msgp.WrapError(err, "poolStateVersion")
			return
		}
		z.poolStateVersion = PoolStateVersion(zb0002)
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "poolState")
			return
		}
		z.poolState = nil
	} else {
		if z.poolState == nil {
			z.poolState = new(PoolState)
		}
		err = z.poolState.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "poolState")
			return
		}
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0003 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0003}
		return
	}
	z.gas.ExchangeAtomically, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "gas", "ExchangeAtomically")
		return
	}
	z.gas.Exchange, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "gas", "Exchange")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolSimulator) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.Pool.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	err = en.WriteUint(uint(z.poolStateVersion))
	if err != nil {
		err = msgp.WrapError(err, "poolStateVersion")
		return
	}
	if z.poolState == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.poolState.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "poolState")
			return
		}
	}
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.gas.ExchangeAtomically)
	if err != nil {
		err = msgp.WrapError(err, "gas", "ExchangeAtomically")
		return
	}
	err = en.WriteInt64(z.gas.Exchange)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Exchange")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolSimulator) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.Pool.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	o = msgp.AppendUint(o, uint(z.poolStateVersion))
	if z.poolState == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.poolState.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "poolState")
			return
		}
	}
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt64(o, z.gas.ExchangeAtomically)
	o = msgp.AppendInt64(o, z.gas.Exchange)
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
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.Pool.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	{
		var zb0002 uint
		zb0002, bts, err = msgp.ReadUintBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "poolStateVersion")
			return
		}
		z.poolStateVersion = PoolStateVersion(zb0002)
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.poolState = nil
	} else {
		if z.poolState == nil {
			z.poolState = new(PoolState)
		}
		bts, err = z.poolState.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "poolState")
			return
		}
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0003 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0003}
		return
	}
	z.gas.ExchangeAtomically, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas", "ExchangeAtomically")
		return
	}
	z.gas.Exchange, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Exchange")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolSimulator) Msgsize() (s int) {
	s = 1 + z.Pool.Msgsize() + msgp.UintSize
	if z.poolState == nil {
		s += msgp.NilSize
	} else {
		s += z.poolState.Msgsize()
	}
	s += 1 + msgp.Int64Size + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolStateVersion) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint
		zb0001, err = dc.ReadUint()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = PoolStateVersion(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z PoolStateVersion) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint(uint(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z PoolStateVersion) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint(o, uint(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolStateVersion) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint
		zb0001, bts, err = msgp.ReadUintBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = PoolStateVersion(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z PoolStateVersion) Msgsize() (s int) {
	s = msgp.UintSize
	return
}
