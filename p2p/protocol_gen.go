package p2p

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Cap) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Version":
			z.Version, err = dc.ReadUint()
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
func (z Cap) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Cap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Cap) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadUintBytes(bts)
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
func (z Cap) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.UintSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ProtoHandshake) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "version":
			z.Version, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "caps":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Caps) >= int(zb0002) {
				z.Caps = (z.Caps)[:zb0002]
			} else {
				z.Caps = make([]Cap, zb0002)
			}
			for za0001 := range z.Caps {
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Name":
						z.Caps[za0001].Name, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Version":
						z.Caps[za0001].Version, err = dc.ReadUint()
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
			}
		case "listen_port":
			z.ListenPort, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "id":
			err = dc.ReadExactBytes((z.ID)[:])
			if err != nil {
				return
			}
		case "tail":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rest) >= int(zb0004) {
				z.Rest = (z.Rest)[:zb0004]
			} else {
				z.Rest = make([][]byte, zb0004)
			}
			for za0003 := range z.Rest {
				z.Rest[za0003], err = dc.ReadBytes(z.Rest[za0003])
				if err != nil {
					return
				}
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
func (z *ProtoHandshake) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "version"
	err = en.Append(0x86, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Version)
	if err != nil {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "caps"
	err = en.Append(0xa4, 0x63, 0x61, 0x70, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Caps)))
	if err != nil {
		return
	}
	for za0001 := range z.Caps {
		// map header, size 2
		// write "Name"
		err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return
		}
		err = en.WriteString(z.Caps[za0001].Name)
		if err != nil {
			return
		}
		// write "Version"
		err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
		if err != nil {
			return
		}
		err = en.WriteUint(z.Caps[za0001].Version)
		if err != nil {
			return
		}
	}
	// write "listen_port"
	err = en.Append(0xab, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ListenPort)
	if err != nil {
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.ID)[:])
	if err != nil {
		return
	}
	// write "tail"
	err = en.Append(0xa4, 0x74, 0x61, 0x69, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Rest)))
	if err != nil {
		return
	}
	for za0003 := range z.Rest {
		err = en.WriteBytes(z.Rest[za0003])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ProtoHandshake) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "version"
	o = append(o, 0x86, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint64(o, z.Version)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "caps"
	o = append(o, 0xa4, 0x63, 0x61, 0x70, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Caps)))
	for za0001 := range z.Caps {
		// map header, size 2
		// string "Name"
		o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Caps[za0001].Name)
		// string "Version"
		o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
		o = msgp.AppendUint(o, z.Caps[za0001].Version)
	}
	// string "listen_port"
	o = append(o, 0xab, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74)
	o = msgp.AppendUint64(o, z.ListenPort)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendBytes(o, (z.ID)[:])
	// string "tail"
	o = append(o, 0xa4, 0x74, 0x61, 0x69, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rest)))
	for za0003 := range z.Rest {
		o = msgp.AppendBytes(o, z.Rest[za0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ProtoHandshake) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "version":
			z.Version, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "caps":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Caps) >= int(zb0002) {
				z.Caps = (z.Caps)[:zb0002]
			} else {
				z.Caps = make([]Cap, zb0002)
			}
			for za0001 := range z.Caps {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Name":
						z.Caps[za0001].Name, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Version":
						z.Caps[za0001].Version, bts, err = msgp.ReadUintBytes(bts)
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
			}
		case "listen_port":
			z.ListenPort, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "id":
			bts, err = msgp.ReadExactBytes(bts, (z.ID)[:])
			if err != nil {
				return
			}
		case "tail":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rest) >= int(zb0004) {
				z.Rest = (z.Rest)[:zb0004]
			} else {
				z.Rest = make([][]byte, zb0004)
			}
			for za0003 := range z.Rest {
				z.Rest[za0003], bts, err = msgp.ReadBytesBytes(bts, z.Rest[za0003])
				if err != nil {
					return
				}
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
func (z *ProtoHandshake) Msgsize() (s int) {
	s = 1 + 8 + msgp.Uint64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Caps {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Caps[za0001].Name) + 8 + msgp.UintSize
	}
	s += 12 + msgp.Uint64Size + 3 + msgp.ArrayHeaderSize + (64 * (msgp.ByteSize)) + 5 + msgp.ArrayHeaderSize
	for za0003 := range z.Rest {
		s += msgp.BytesPrefixSize + len(z.Rest[za0003])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Protocol) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Version":
			z.Version, err = dc.ReadUint()
			if err != nil {
				return
			}
		case "Length":
			z.Length, err = dc.ReadUint64()
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
func (z Protocol) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Name"
	err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint(z.Version)
	if err != nil {
		return
	}
	// write "Length"
	err = en.Append(0xa6, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Length)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Protocol) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint(o, z.Version)
	// string "Length"
	o = append(o, 0xa6, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint64(o, z.Length)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Protocol) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadUintBytes(bts)
			if err != nil {
				return
			}
		case "Length":
			z.Length, bts, err = msgp.ReadUint64Bytes(bts)
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
func (z Protocol) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.UintSize + 7 + msgp.Uint64Size
	return
}
