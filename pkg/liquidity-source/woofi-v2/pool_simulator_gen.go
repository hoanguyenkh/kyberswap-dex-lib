package woofiv2

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
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Swap, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Gas) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Swap)
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Gas) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt64(o, z.Swap)
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
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Swap, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Swap")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Gas) Msgsize() (s int) {
	s = 1 + msgp.Int64Size
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
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
		return
	}
	err = z.Pool.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	z.quoteToken, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "quoteToken")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "tokenInfos")
		return
	}
	if z.tokenInfos == nil {
		z.tokenInfos = make(map[string]TokenInfo, zb0002)
	} else if len(z.tokenInfos) > 0 {
		for key := range z.tokenInfos {
			delete(z.tokenInfos, key)
		}
	}
	var field []byte
	_ = field
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 TokenInfo
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos")
			return
		}
		err = za0002.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos", za0001)
			return
		}
		z.tokenInfos[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	if z.decimals == nil {
		z.decimals = make(map[string]uint8, zb0003)
	} else if len(z.decimals) > 0 {
		for key := range z.decimals {
			delete(z.decimals, key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var za0003 string
		var za0004 uint8
		za0003, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "decimals")
			return
		}
		za0004, err = dc.ReadUint8()
		if err != nil {
			err = msgp.WrapError(err, "decimals", za0003)
			return
		}
		z.decimals[za0003] = za0004
	}
	err = z.wooracle.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "wooracle")
		return
	}
	var zb0004 uint32
	zb0004, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "cloracle")
		return
	}
	if z.cloracle == nil {
		z.cloracle = make(map[string]Cloracle, zb0004)
	} else if len(z.cloracle) > 0 {
		for key := range z.cloracle {
			delete(z.cloracle, key)
		}
	}
	for zb0004 > 0 {
		zb0004--
		var za0005 string
		var za0006 Cloracle
		za0005, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "cloracle")
			return
		}
		err = za0006.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "cloracle", za0005)
			return
		}
		z.cloracle[za0005] = za0006
	}
	var zb0005 uint32
	zb0005, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0005 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0005}
		return
	}
	z.gas.Swap, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolSimulator) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 7
	err = en.Append(0x97)
	if err != nil {
		return
	}
	err = z.Pool.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	err = en.WriteString(z.quoteToken)
	if err != nil {
		err = msgp.WrapError(err, "quoteToken")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.tokenInfos)))
	if err != nil {
		err = msgp.WrapError(err, "tokenInfos")
		return
	}
	for za0001, za0002 := range z.tokenInfos {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos")
			return
		}
		err = za0002.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos", za0001)
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.decimals)))
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	for za0003, za0004 := range z.decimals {
		err = en.WriteString(za0003)
		if err != nil {
			err = msgp.WrapError(err, "decimals")
			return
		}
		err = en.WriteUint8(za0004)
		if err != nil {
			err = msgp.WrapError(err, "decimals", za0003)
			return
		}
	}
	err = z.wooracle.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "wooracle")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.cloracle)))
	if err != nil {
		err = msgp.WrapError(err, "cloracle")
		return
	}
	for za0005, za0006 := range z.cloracle {
		err = en.WriteString(za0005)
		if err != nil {
			err = msgp.WrapError(err, "cloracle")
			return
		}
		err = za0006.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "cloracle", za0005)
			return
		}
	}
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.gas.Swap)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolSimulator) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 7
	o = append(o, 0x97)
	o, err = z.Pool.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	o = msgp.AppendString(o, z.quoteToken)
	o = msgp.AppendMapHeader(o, uint32(len(z.tokenInfos)))
	for za0001, za0002 := range z.tokenInfos {
		o = msgp.AppendString(o, za0001)
		o, err = za0002.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos", za0001)
			return
		}
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.decimals)))
	for za0003, za0004 := range z.decimals {
		o = msgp.AppendString(o, za0003)
		o = msgp.AppendUint8(o, za0004)
	}
	o, err = z.wooracle.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "wooracle")
		return
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.cloracle)))
	for za0005, za0006 := range z.cloracle {
		o = msgp.AppendString(o, za0005)
		o, err = za0006.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "cloracle", za0005)
			return
		}
	}
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt64(o, z.gas.Swap)
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
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
		return
	}
	bts, err = z.Pool.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Pool")
		return
	}
	z.quoteToken, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "quoteToken")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "tokenInfos")
		return
	}
	if z.tokenInfos == nil {
		z.tokenInfos = make(map[string]TokenInfo, zb0002)
	} else if len(z.tokenInfos) > 0 {
		for key := range z.tokenInfos {
			delete(z.tokenInfos, key)
		}
	}
	var field []byte
	_ = field
	for zb0002 > 0 {
		var za0001 string
		var za0002 TokenInfo
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos")
			return
		}
		bts, err = za0002.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "tokenInfos", za0001)
			return
		}
		z.tokenInfos[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	if z.decimals == nil {
		z.decimals = make(map[string]uint8, zb0003)
	} else if len(z.decimals) > 0 {
		for key := range z.decimals {
			delete(z.decimals, key)
		}
	}
	for zb0003 > 0 {
		var za0003 string
		var za0004 uint8
		zb0003--
		za0003, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "decimals")
			return
		}
		za0004, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "decimals", za0003)
			return
		}
		z.decimals[za0003] = za0004
	}
	bts, err = z.wooracle.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "wooracle")
		return
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "cloracle")
		return
	}
	if z.cloracle == nil {
		z.cloracle = make(map[string]Cloracle, zb0004)
	} else if len(z.cloracle) > 0 {
		for key := range z.cloracle {
			delete(z.cloracle, key)
		}
	}
	for zb0004 > 0 {
		var za0005 string
		var za0006 Cloracle
		zb0004--
		za0005, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "cloracle")
			return
		}
		bts, err = za0006.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "cloracle", za0005)
			return
		}
		z.cloracle[za0005] = za0006
	}
	var zb0005 uint32
	zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas")
		return
	}
	if zb0005 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0005}
		return
	}
	z.gas.Swap, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "gas", "Swap")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolSimulator) Msgsize() (s int) {
	s = 1 + z.Pool.Msgsize() + msgp.StringPrefixSize + len(z.quoteToken) + msgp.MapHeaderSize
	if z.tokenInfos != nil {
		for za0001, za0002 := range z.tokenInfos {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + za0002.Msgsize()
		}
	}
	s += msgp.MapHeaderSize
	if z.decimals != nil {
		for za0003, za0004 := range z.decimals {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.Uint8Size
		}
	}
	s += z.wooracle.Msgsize() + msgp.MapHeaderSize
	if z.cloracle != nil {
		for za0005, za0006 := range z.cloracle {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005) + za0006.Msgsize()
		}
	}
	s += 1 + msgp.Int64Size
	return
}
