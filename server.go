package main

import (
	"encoding/json"
	"reflect"
	// "encoding/json"
	// authOv "github.com/JeanSebTr/minecraft/auth"
	"github.com/JeanSebTr/minecraft/protocol"
	// "github.com/NetherrackDev/netherrack/message"
	// "github.com/NetherrackDev/netherrack/protocol/auth"
	"log"
	"net"
)

func main() {
	log.Println("Starting server")

	srv, err := net.Listen("tcp", ":25565")

	if err != nil {
		log.Panicf("Error listening on 25564: %v \n", err)
	}

	for {
		conn, err := srv.Accept()

		if err != nil {
			log.Panicf("Error accepting: %v \n", err)
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Connection from %s", conn.RemoteAddr().String())

	mcConn := &protocol.Conn{
		In:             conn,
		Out:            conn,
		ReadDirection:  protocol.Serverbound,
		WriteDirection: protocol.Clientbound,
		State:          protocol.Handshaking,
	}

	packet, err := mcConn.ReadPacket()
	if err != nil {
		//The client either had a early connection issue or
		//it isn't a minecraft client
		return
	}
	handshake, ok := packet.(*protocol.Handshake)
	if !ok {
		//Client sent the wrong packet. This shouldn't
		//happen because in the handshaking protocol (default state)
		//only has the handshake packet as a valid packet
		return
	}

	log.Printf("handshake %+v", handshake)

	// log.Printf("Client v%d\n", handshake.ProtocolVersion)

	//Status ping
	if handshake.State == 1 {
		mcConn.State = protocol.Status
		packet, err := mcConn.ReadPacket()
		if _, ok := packet.(*protocol.StatusGet); !ok || err != nil {
			return
		}

		by, err := json.Marshal(Ping{
			Version:     Version{Name: "1.8.0", Protocol: int(handshake.ProtocolVersion)},
			Description: "Wow mine that!",
		})
		if err != nil {
			panic(err)
		}

		pkt := protocol.Packet(&protocol.StatusResponse{string(by)})

		log.Printf("%+v => %+v", reflect.TypeOf(pkt).Elem(), reflect.TypeOf((*protocol.StatusResponse)(nil)).Elem())

		mcConn.WritePacket(pkt)
		packet, err = mcConn.ReadPacket()
		if err != nil {
			panic(err)
			return
		}
		cPing, ok := packet.(*protocol.ClientStatusPing)
		log.Printf("handshake %+v", cPing)
		if !ok {
			return
		}
		mcConn.WritePacket(&protocol.StatusPing{Time: cPing.Time})
		return
	}

	// if handshake.State != 2 {
	// 	return
	// }

	// defer log.Printf("Killed %s", conn.RemoteAddr())
	// log.Printf("Connection %s", conn.RemoteAddr())

	// //handshake.ProtocolVersion = 4
	// username, uuid, err := authOv.Login(&mcConn, handshake, auth.Instance)
	// if err != nil {
	// 	log.Printf("Player %s(%s) login error: %s", uuid, username, err)
	// 	mcConn.WritePacket(protocol.LoginDisconnect{(&message.Message{Text: err.Error(), Color: message.Red}).JSONString()})
	// 	return
	// }

	// sConn, err := net.Dial("tcp", "127.0.0.1:25565")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sConn.Close()

	// msConn := &protocol.Conn{
	// 	In:             conn,
	// 	Out:            conn,
	// 	Deadliner:      conn,
	// 	ReadDirection:  protocol.Clientbound,
	// 	WriteDirection: protocol.Serverbound,
	// }

	// log.Printf("Auth %s(%s)\n", uuid, username)

	// msConn.WritePacket(protocol.Handshake{
	// 	Address:         handshake.Address,
	// 	Port:            25565,
	// 	ProtocolVersion: handshake.ProtocolVersion,
	// 	State:           handshake.State,
	// })

	// msConn.WritePacket(protocol.LoginStart{username})

	// authOv.ClientLogin(msConn)

	// mcConn.WritePacket(protocol.JoinGame{
	// 	EntityID:     45,
	// 	Gamemode:     0,
	// 	Dimension:    int8(0),
	// 	Difficulty:   2,
	// 	MaxPlayers:   20,
	// 	LevelType:    "default",
	// 	ReducedDebug: true,
	// })

	// mcConn.WritePacket(protocol.SpawnPosition{
	// 	protocol.Position{
	// 		X: 20,
	// 		Y: 20,
	// 		Z: 20,
	// 	},
	// })

	// mcConn.WritePacket(protocol.PlayerPositionLook{
	// 	X:        20,
	// 	Y:        20,
	// 	Z:        20,
	// 	Yaw:      20,
	// 	Pitch:    20,
	// 	OnGround: true,
	// })

	// for {
	// 	packet, err := mcConn.ReadPacket()
	// 	if err != nil {
	// 		log.Panicln(err)
	// 	}
	// 	log.Printf("Packet %T %+v\n", packet, packet)
	// }
}

type Ping struct {
	Version     Version `json:"version"`
	Description string  `json:"description"`
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}
