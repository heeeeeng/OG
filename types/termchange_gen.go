package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *SigSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	err = z.PublicKey.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.Signature.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SigSet) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = z.PublicKey.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.Signature.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SigSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o, err = z.PublicKey.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.Signature.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SigSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	bts, err = z.PublicKey.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.Signature.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SigSet) Msgsize() (s int) {
	s = 1 + z.PublicKey.Msgsize() + z.Signature.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TermChange) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	err = z.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	z.TermID, err = dc.ReadUint64()
	if err != nil {
		return
	}
	err = z.PkBls.DecodeMsg(dc)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.SigSet) >= int(zb0002) {
		z.SigSet = (z.SigSet)[:zb0002]
	} else {
		z.SigSet = make([]*SigSet, zb0002)
	}
	for za0001 := range z.SigSet {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			z.SigSet[za0001] = nil
		} else {
			if z.SigSet[za0001] == nil {
				z.SigSet[za0001] = new(SigSet)
			}
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if zb0003 != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zb0003}
				return
			}
			err = z.SigSet[za0001].PublicKey.DecodeMsg(dc)
			if err != nil {
				return
			}
			err = z.SigSet[za0001].Signature.DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	err = z.Issuer.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TermChange) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TermID)
	if err != nil {
		return
	}
	err = z.PkBls.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.SigSet)))
	if err != nil {
		return
	}
	for za0001 := range z.SigSet {
		if z.SigSet[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// array header, size 2
			err = en.Append(0x92)
			if err != nil {
				return
			}
			err = z.SigSet[za0001].PublicKey.EncodeMsg(en)
			if err != nil {
				return
			}
			err = z.SigSet[za0001].Signature.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	err = z.Issuer.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TermChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendUint64(o, z.TermID)
	o, err = z.PkBls.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.SigSet)))
	for za0001 := range z.SigSet {
		if z.SigSet[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// array header, size 2
			o = append(o, 0x92)
			o, err = z.SigSet[za0001].PublicKey.MarshalMsg(o)
			if err != nil {
				return
			}
			o, err = z.SigSet[za0001].Signature.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	o, err = z.Issuer.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TermChange) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	bts, err = z.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	z.TermID, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	bts, err = z.PkBls.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.SigSet) >= int(zb0002) {
		z.SigSet = (z.SigSet)[:zb0002]
	} else {
		z.SigSet = make([]*SigSet, zb0002)
	}
	for za0001 := range z.SigSet {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.SigSet[za0001] = nil
		} else {
			if z.SigSet[za0001] == nil {
				z.SigSet[za0001] = new(SigSet)
			}
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if zb0003 != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zb0003}
				return
			}
			bts, err = z.SigSet[za0001].PublicKey.UnmarshalMsg(bts)
			if err != nil {
				return
			}
			bts, err = z.SigSet[za0001].Signature.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	bts, err = z.Issuer.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TermChange) Msgsize() (s int) {
	s = 1 + z.TxBase.Msgsize() + msgp.Uint64Size + z.PkBls.Msgsize() + msgp.ArrayHeaderSize
	for za0001 := range z.SigSet {
		if z.SigSet[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + z.SigSet[za0001].PublicKey.Msgsize() + z.SigSet[za0001].Signature.Msgsize()
		}
	}
	s += z.Issuer.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TermChanges) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TermChanges, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(TermChange)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TermChanges) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TermChanges) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TermChanges) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TermChanges, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(TermChange)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TermChanges) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
