/*
   Copyright 2013 Matthew Collins (purggames@gmail.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package protocol

import (
	"reflect"
)

type State int

const (
	Handshaking State = iota
	Play
	Login
	Status
)

type Direction int

const (
	Serverbound Direction = iota
	Clientbound
)

var (
	packets = [4][2][]reflect.Type{
		Handshaking: [2][]reflect.Type{
			Clientbound: []reflect.Type{},
			Serverbound: []reflect.Type{
				reflect.TypeOf((*Handshake)(nil)).Elem(),
			},
		},
		Play: [2][]reflect.Type{
			Clientbound: []reflect.Type{
				reflect.TypeOf((*KeepAlive)(nil)).Elem(),
				reflect.TypeOf((*JoinGame)(nil)).Elem(),
				reflect.TypeOf((*ServerMessage)(nil)).Elem(),
				reflect.TypeOf((*TimeUpdate)(nil)).Elem(),
				reflect.TypeOf((*EntityEquipment)(nil)).Elem(),
				reflect.TypeOf((*SpawnPosition)(nil)).Elem(),
				reflect.TypeOf((*UpdateHealth)(nil)).Elem(),
				reflect.TypeOf((*Respawn)(nil)).Elem(),
				reflect.TypeOf((*PlayerPositionLook)(nil)).Elem(),
				reflect.TypeOf((*HeldItemChange)(nil)).Elem(),
				reflect.TypeOf((*UseBed)(nil)).Elem(),
				reflect.TypeOf((*Animation)(nil)).Elem(),
				reflect.TypeOf((*SpawnPlayer)(nil)).Elem(),
				reflect.TypeOf((*CollectItem)(nil)).Elem(),
				reflect.TypeOf((*SpawnObject)(nil)).Elem(),
				reflect.TypeOf((*SpawnMob)(nil)).Elem(),
				reflect.TypeOf((*SpawnPainting)(nil)).Elem(),
				reflect.TypeOf((*SpawnExperienceOrb)(nil)).Elem(),
				reflect.TypeOf((*EntityVelocity)(nil)).Elem(),
				reflect.TypeOf((*EntityDestroy)(nil)).Elem(),
				reflect.TypeOf((*Entity)(nil)).Elem(),
				reflect.TypeOf((*EntityMove)(nil)).Elem(),
				reflect.TypeOf((*EntityLook)(nil)).Elem(),
				reflect.TypeOf((*EntityLookMove)(nil)).Elem(),
				reflect.TypeOf((*EntityTeleport)(nil)).Elem(),
				reflect.TypeOf((*EntityHeadLook)(nil)).Elem(),
				reflect.TypeOf((*EntityStatus)(nil)).Elem(),
				reflect.TypeOf((*EntityAttach)(nil)).Elem(),
				reflect.TypeOf((*EntityMetadata)(nil)).Elem(),
				reflect.TypeOf((*EntityEffect)(nil)).Elem(),
				reflect.TypeOf((*EntityEffectRemove)(nil)).Elem(),
				reflect.TypeOf((*SetExperience)(nil)).Elem(),
				reflect.TypeOf((*EntityProperties)(nil)).Elem(),
				reflect.TypeOf((*ChunkData)(nil)).Elem(),
				reflect.TypeOf((*MultiBlockChange)(nil)).Elem(),
				reflect.TypeOf((*BlockChange)(nil)).Elem(),
				reflect.TypeOf((*BlockAction)(nil)).Elem(),
				reflect.TypeOf((*BlockBreakAnimation)(nil)).Elem(),
				reflect.TypeOf((*MapChunkBulk)(nil)).Elem(),
				reflect.TypeOf((*Explosion)(nil)).Elem(),
				reflect.TypeOf((*Effect)(nil)).Elem(),
				reflect.TypeOf((*SoundEffect)(nil)).Elem(),
				reflect.TypeOf((*Particle)(nil)).Elem(),
				reflect.TypeOf((*GameState)(nil)).Elem(),
				reflect.TypeOf((*SpawnGlobalEntity)(nil)).Elem(),
				reflect.TypeOf((*WindowOpen)(nil)).Elem(),
				reflect.TypeOf((*WindowClose)(nil)).Elem(),
				reflect.TypeOf((*WindowSetSlot)(nil)).Elem(),
				reflect.TypeOf((*WindowItems)(nil)).Elem(),
				reflect.TypeOf((*WindowUpdateProperty)(nil)).Elem(),
				reflect.TypeOf((*WindowTransactionConfirm)(nil)).Elem(),
				reflect.TypeOf((*UpdateSign)(nil)).Elem(),
				reflect.TypeOf((*Maps)(nil)).Elem(),
				reflect.TypeOf((*UpdateBlockEntity)(nil)).Elem(),
				reflect.TypeOf((*SignEditorOpen)(nil)).Elem(),
				reflect.TypeOf((*Statistics)(nil)).Elem(),
				reflect.TypeOf((*PlayerListItem)(nil)).Elem(),
				reflect.TypeOf((*PlayerAbilities)(nil)).Elem(),
				reflect.TypeOf((*TabComplete)(nil)).Elem(),
				reflect.TypeOf((*ScoreboardObjective)(nil)).Elem(),
				reflect.TypeOf((*UpdateScore)(nil)).Elem(),
				reflect.TypeOf((*DisplayScoreboard)(nil)).Elem(),
				reflect.TypeOf((*Teams)(nil)).Elem(),
				reflect.TypeOf((*PluginMessage)(nil)).Elem(),
				reflect.TypeOf((*Disconnect)(nil)).Elem(),
			},
			Serverbound: []reflect.Type{
				reflect.TypeOf((*ClientKeepAlive)(nil)).Elem(),
				reflect.TypeOf((*ChatMessage)(nil)).Elem(),
				reflect.TypeOf((*UseEntity)(nil)).Elem(),
				reflect.TypeOf((*ClientPlayer)(nil)).Elem(),
				reflect.TypeOf((*ClientPlayerPosition)(nil)).Elem(),
				reflect.TypeOf((*ClientPlayerLook)(nil)).Elem(),
				reflect.TypeOf((*ClientPlayerPositionLook)(nil)).Elem(),
				reflect.TypeOf((*PlayerDigging)(nil)).Elem(),
				reflect.TypeOf((*PlayerBlockPlacement)(nil)).Elem(),
				reflect.TypeOf((*ClientHeldItemChange)(nil)).Elem(),
				reflect.TypeOf((*ClientAnimation)(nil)).Elem(),
				reflect.TypeOf((*EntityAction)(nil)).Elem(),
				reflect.TypeOf((*SteerVehicle)(nil)).Elem(),
				reflect.TypeOf((*ClientWindowClose)(nil)).Elem(),
				reflect.TypeOf((*WindowClick)(nil)).Elem(),
				reflect.TypeOf((*ClientWindowTransactionConfirm)(nil)).Elem(),
				reflect.TypeOf((*CreativeInventoryAction)(nil)).Elem(),
				reflect.TypeOf((*EnchantItem)(nil)).Elem(),
				reflect.TypeOf((*ClientUpdateSign)(nil)).Elem(),
				reflect.TypeOf((*ClientPlayerAbilities)(nil)).Elem(),
				reflect.TypeOf((*ClientTabComplete)(nil)).Elem(),
				reflect.TypeOf((*ClientSettings)(nil)).Elem(),
				reflect.TypeOf((*ClientStatuses)(nil)).Elem(),
				reflect.TypeOf((*ClientPluginMessage)(nil)).Elem(),
			},
		},
		Login: [2][]reflect.Type{
			Clientbound: []reflect.Type{
				reflect.TypeOf((*LoginDisconnect)(nil)).Elem(),
				reflect.TypeOf((*EncryptionKeyRequest)(nil)).Elem(),
				reflect.TypeOf((*LoginSuccess)(nil)).Elem(),
			},
			Serverbound: []reflect.Type{
				reflect.TypeOf((*LoginStart)(nil)).Elem(),
				reflect.TypeOf((*EncryptionKeyResponse)(nil)).Elem(),
			},
		},
		Status: [2][]reflect.Type{
			Clientbound: []reflect.Type{
				reflect.TypeOf((*StatusResponse)(nil)).Elem(),
				reflect.TypeOf((*StatusPing)(nil)).Elem(),
			},
			Serverbound: []reflect.Type{
				reflect.TypeOf((*StatusGet)(nil)).Elem(),
				reflect.TypeOf((*ClientStatusPing)(nil)).Elem(),
			},
		},
	}
	packetsToID = [2]map[reflect.Type]int{
		Clientbound: map[reflect.Type]int{},
		Serverbound: map[reflect.Type]int{},
	}
	//Needed because the packet passer struggles with this packet
	_MapChunkBulkID int
)

func init() {
	for _, st := range packets {
		for d, dir := range st {
			for i, p := range dir {
				if _, ok := packetsToID[d][p]; ok {
					panic("Duplicate packet " + p.Name())
				}
				packetsToID[d][p] = i
			}
		}
	}
	_MapChunkBulkID = packetsToID[Clientbound][reflect.TypeOf((*MapChunkBulk)(nil)).Elem()]
}
