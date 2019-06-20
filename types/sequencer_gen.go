package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BlsSigSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PublicKey, err = dc.ReadBytes(z.PublicKey)
	if err != nil {
		return
	}
	z.BlsSignature, err = dc.ReadBytes(z.BlsSignature)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BlsSigSet) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PublicKey)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.BlsSignature)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BlsSigSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendBytes(o, z.PublicKey)
	o = msgp.AppendBytes(o, z.BlsSignature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlsSigSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PublicKey, bts, err = msgp.ReadBytesBytes(bts, z.PublicKey)
	if err != nil {
		return
	}
	z.BlsSignature, bts, err = msgp.ReadBytesBytes(bts, z.BlsSignature)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BlsSigSet) Msgsize() (s int) {
	s = 1 + msgp.BytesPrefixSize + len(z.PublicKey) + msgp.BytesPrefixSize + len(z.BlsSignature)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sequencer) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.Issuer.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.BlsJointSig.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Sequencer) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.Issuer.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.BlsJointSig.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Sequencer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.Issuer.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.BlsJointSig.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.BlsJointPubKey.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sequencer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.Issuer.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.BlsJointSig.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.BlsJointPubKey.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Sequencer) Msgsize() (s int) {
	s = 1 + z.TxBase.Msgsize() + z.Issuer.Msgsize() + z.BlsJointSig.Msgsize() + z.BlsJointPubKey.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SequencerJson) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBaseJson":
			err = z.TxBaseJson.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Issuer":
			err = z.Issuer.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "BlsJointSig":
			err = z.BlsJointSig.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "BlsJointPubKey":
			err = z.BlsJointPubKey.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SequencerJson) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "TxBaseJson"
	err = en.Append(0x84, 0xaa, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = z.TxBaseJson.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Issuer"
	err = en.Append(0xa6, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.Issuer.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "BlsJointSig"
	err = en.Append(0xab, 0x42, 0x6c, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x67)
	if err != nil {
		return
	}
	err = z.BlsJointSig.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "BlsJointPubKey"
	err = en.Append(0xae, 0x42, 0x6c, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SequencerJson) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "TxBaseJson"
	o = append(o, 0x84, 0xaa, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e)
	o, err = z.TxBaseJson.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Issuer"
	o = append(o, 0xa6, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72)
	o, err = z.Issuer.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "BlsJointSig"
	o = append(o, 0xab, 0x42, 0x6c, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x67)
	o, err = z.BlsJointSig.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "BlsJointPubKey"
	o = append(o, 0xae, 0x42, 0x6c, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79)
	o, err = z.BlsJointPubKey.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SequencerJson) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBaseJson":
			bts, err = z.TxBaseJson.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Issuer":
			bts, err = z.Issuer.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "BlsJointSig":
			bts, err = z.BlsJointSig.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "BlsJointPubKey":
			bts, err = z.BlsJointPubKey.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SequencerJson) Msgsize() (s int) {
	s = 1 + 11 + z.TxBaseJson.Msgsize() + 7 + z.Issuer.Msgsize() + 12 + z.BlsJointSig.Msgsize() + 15 + z.BlsJointPubKey.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sequencers) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
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
				(*z)[zb0001] = new(Sequencer)
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
func (z Sequencers) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z Sequencers) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *Sequencers) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
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
				(*z)[zb0001] = new(Sequencer)
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
func (z Sequencers) Msgsize() (s int) {
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
