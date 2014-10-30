# minecraft tools & libs written in golang

Here's some code about Minecraft… big WIP

* Mostly about the main game protocol
* Some, originaly from [NetherrackDev/netherrack][netherrack] (packets, packets' table & auth)

The MasterPlan would be to make a server supporting multiple client's versions on a distributed back-end.
Storing terrain chunks in redis or something like that.
Or sharding the generated chunks between running instances with custom pubsub ans raft. Because. Fun.

## build/build.go

Generate protocol read/write methods from packet's struct definitions using golang AST

run `go run build/build.go < protocol/packets.go > protocol/somefile.go`

## protocol/

WIP reimplementation of [netherrack][netherrack]'s mc protocol handler without the reflection based parser
Make huge use of golang type alias so every packet's field type has a Read/Write method

The rational behind it is that the original implementation :
 * can't support multiple protocol versions at once
 * make heavy use of reflexion to read/write packets
 * did a clever use of struct's tag for buffer's and array's length, but I think it's better solved by type alias

And the more complete implementation [mc-server/MCServer][mcserver] (support 1.8)
 * duplicate the parser methods to handle changes in protocol
 * manually implement every packet's reading/writing
 * is in C++

## server.go

allow to login from an 1.8 client with some patch to [netherrack][netherrack] and then login to an other server
should add packets proxying to catch broken packet parsing and record session for automated tests
mostly experiments residues

## Food for thought

* http://wiki.vg/
* be more performant and maintainable than [MCServer][mcserver]
* api to control game entities
* of course [a map like that](http://mc.westeroscraft.com/)
* allow unlimited client with virtual entity ids
* coordinate translation to allow wrapping around the map
* dispatcher patern + golang chan's sweetness : be aware of blocking channels & slow clients


[netherrack]: https://github.com/NetherrackDev/netherrack
[mcserver]: https://github.com/mc-server/MCServer
