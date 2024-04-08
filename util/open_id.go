package util

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

type Env int

const (
	EnvSdev0 Env = iota
	EnvSdev
	EnvSdev2
	EnvDev
	EnvDev2
	EnvDev3
	EnvStg
	EnvLive
)

type IdType int

const (
	IdTypeOpenId = iota
	IdTypeUnionId
)

func GenerateOpenId(env Env, appId uint32, memberNo uint64) string {

	var b [16]byte
	binary.BigEndian.PutUint16(b[0:], uint16(env))
	binary.BigEndian.PutUint16(b[2:], uint16(IdTypeOpenId))
	binary.BigEndian.PutUint32(b[4:], appId)
	binary.BigEndian.PutUint64(b[8:], memberNo)

	var buf [32]byte
	hex.Encode(buf[:], b[:])

	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum(buf[:])))
}

func GenerateUnionId(env Env, appId uint32, ownerMemberNo uint64) string {

	var b [16]byte
	binary.BigEndian.PutUint16(b[0:], uint16(env))
	binary.BigEndian.PutUint16(b[2:], uint16(IdTypeUnionId))
	binary.BigEndian.PutUint32(b[4:], appId)
	binary.BigEndian.PutUint64(b[8:], ownerMemberNo)

	var buf [32]byte
	hex.Encode(buf[:], b[:])

	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum(buf[:])))
}
