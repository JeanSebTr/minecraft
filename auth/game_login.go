package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	// "encoding/hex"
	"fmt"
	"github.com/NetherrackDev/netherrack/protocol"
	"github.com/NetherrackDev/netherrack/protocol/auth"
)

const Version = 47

var (
	publicKeyBytes []byte
	privateKey     *rsa.PrivateKey
)

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	privateKey.Precompute()

	publicKeyBytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
}

//Auths the user and returns their username.
//Uses infomation from http://wiki.vg/Protocol_Encryption
func Login(conn *protocol.Conn, handshake protocol.Handshake, authenticator auth.Authenticator) (username string, uuid string, err error) {
	if handshake.ProtocolVersion != Version {
		if handshake.ProtocolVersion < Version {
			return "", uuid, protocol.ErrorOutOfDateClient
		} else {
			return "", uuid, protocol.ErrorOutOfDateServer
		}
	}

	conn.State = protocol.Login

	packet, err := conn.ReadPacket()
	if err != nil {
		return
	}
	lStart, ok := packet.(protocol.LoginStart)
	if !ok {
		err = fmt.Errorf("Unexpected packet")
		return
	}
	username = lStart.Username

	verifyToken := make([]byte, 16) //Used by the server to check encryption is working correctly
	rand.Read(verifyToken)

	var serverID = ""
	// if authenticator != nil {
	// serverBytes := make([]byte, 10)
	// rand.Read(serverBytes)
	// serverID = hex.EncodeToString(serverBytes)
	// } else {
	// 	if len(username) > 16 {
	// 		username = username[:16]
	// 	}
	// }
	encReq := protocol.EncryptionKeyRequest{
		ServerID:    serverID,
		PublicKey:   publicKeyBytes,
		VerifyToken: verifyToken,
	}
	fmt.Printf("EncReq %+v \n", encReq)
	conn.WritePacket(encReq)

	packet, err = conn.ReadPacket()
	if err != nil {
		return
	}
	encryptionResponse, ok := packet.(protocol.EncryptionKeyResponse)
	if !ok {
		err = fmt.Errorf("Unexpected packet")
		return
	}

	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptionResponse.SharedSecret)
	if err != nil {
		return
	}

	verifyTokenResponse, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptionResponse.VerifyToken)
	if err != nil {
		return
	}
	if !bytes.Equal(verifyToken, verifyTokenResponse) {
		return
	}

	if uuid, err = authenticator.Authenticate(username, serverID, sharedSecret, publicKeyBytes); err != nil {
		return
	}

	aesCipher, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return
	}

	conn.In = cipher.StreamReader{
		R: conn.In,
		S: newCFB8Decrypt(aesCipher, sharedSecret),
	}
	conn.Out = cipher.StreamWriter{
		W: conn.Out,
		S: newCFB8Encrypt(aesCipher, sharedSecret),
	}

	// fix uuid format
	uuid = fmt.Sprintf("%s-%s-%s-%s-%s", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:32])

	conn.WritePacket(protocol.LoginSuccess{uuid, username})
	conn.State = protocol.Play
	return
}

func ClientLogin(conn *protocol.Conn) (err error) {

	packet, err := conn.ReadPacket()
	if err != nil {
		return
	}
	encReq, ok := packet.(protocol.EncryptionKeyRequest)
	if !ok {
		return fmt.Errorf("Packet isn't EncryptionKeyRequest %T %+v\n", packet, packet)
	}

	sharedSecret := make([]byte, 16)
	_, _ = rand.Read(sharedSecret)

	pubKeyIntf, err := x509.ParsePKIXPublicKey(encReq.PublicKey)
	if err != nil {
		return
	}

	pubKey, ok := pubKeyIntf.(rsa.PublicKey)
	if !ok {
		return fmt.Errorf("Not a RSA public key. %T", pubKeyIntf)
	}

	encSharedSecret, err := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, sharedSecret)
	if err != nil {
		return
	}

	encVerifyToken, err := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, encReq.VerifyToken)
	if err != nil {
		return
	}

	conn.WritePacket(protocol.EncryptionKeyResponse{
		SharedSecret: encSharedSecret,
		VerifyToken:  encVerifyToken,
	})

	aesCipher, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return
	}

	conn.In = cipher.StreamReader{
		R: conn.In,
		S: newCFB8Decrypt(aesCipher, sharedSecret),
	}
	conn.Out = cipher.StreamWriter{
		W: conn.Out,
		S: newCFB8Encrypt(aesCipher, sharedSecret),
	}

	packet, err = conn.ReadPacket()
	if err != nil {
		return
	}

	if _, ok := packet.(protocol.LoginSuccess); !ok {
		return fmt.Errorf("Packet isn't LoginSuccess %T %+v\n", packet, packet)
	}
	conn.State = protocol.Play

	return
}
