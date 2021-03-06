package wire

import (
	"bytes"
	"crypto/rand"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/utils"
)

// ComposeGQUICVersionNegotiation composes a Version Negotiation Packet for gQUIC
func ComposeGQUICVersionNegotiation(connID protocol.ConnectionID, versions []protocol.VersionNumber) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1+8+len(versions)*4))
	buf.Write([]byte{0x1 | 0x8}) // type byte
	utils.BigEndian.WriteUint64(buf, uint64(connID))
	for _, v := range versions {
		utils.BigEndian.WriteUint32(buf, uint32(v))
	}
	return buf.Bytes()
}

// ComposeVersionNegotiation composes a Version Negotiation according to the IETF draft
func ComposeVersionNegotiation(
	connID protocol.ConnectionID,
	versions []protocol.VersionNumber,
) []byte {
	greasedVersions := protocol.GetGreasedVersions(versions)
	buf := bytes.NewBuffer(make([]byte, 0, 1+8+4+len(greasedVersions)*4))
	r := make([]byte, 1)
	_, _ = rand.Read(r) // ignore the error here. It is not critical to have perfect random here.
	buf.WriteByte(r[0] | 0x80)
	utils.BigEndian.WriteUint64(buf, uint64(connID))
	utils.BigEndian.WriteUint32(buf, 0) // version 0
	for _, v := range greasedVersions {
		utils.BigEndian.WriteUint32(buf, uint32(v))
	}
	return buf.Bytes()
}
