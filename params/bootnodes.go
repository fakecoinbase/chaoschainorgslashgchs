// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// CHS Foundation Go Bootnodes
	"enode://7c62ac5fd9dbb73387fb37410baf78e96b767f746f11e1c19d88ab4a345bcb7a54211efd4764988c1b6ddfb95d76a3db52bbb696cbd6d6b0c40927da29a6a6c3@[2408:8642:4ff:8:1::2ea2]:40400",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@[2408:8642:4ff:8:1::80f]:30303",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@221.12.12.211:30303",
	// "enode://374a44efd8d58c0223d1db080aa42ff28268ad5e823b76dec552878c096fa316193d9a1136d115802fed4a01ff943f1b187354a58740ebf26bc7dbb1bfc2106b@gchs.microcai.org:30303",
	// "enode://e750149c15c96bc55f4f29d94960872a8ad699fec91e8e8ef8eb2209c9330093835a830f5f4f725ac2a81b5544f6af53ef3e9ff71f3e38858fb83111dc5c43f9@gchs02.microcai.org:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	// CHS Foundation Go Bootnodes
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@[2408:8642:4ff:8:1::80f]:30303",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@221.12.12.211:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	// CHS Foundation Go Bootnodes
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@[2408:8642:4ff:8:1::80f]:30303",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@221.12.12.211:30303",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	// "enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	// "enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	// "enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	// "enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	// "enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// CHS Foundation Go Bootnodes
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@[2408:8642:4ff:8:1::80f]:30303",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@221.12.12.211:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	// "enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	// "enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	// "enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	// "enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",\

	// CHS Foundation Go Bootnodes
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@[2408:8642:4ff:8:1::80f]:30303",
	"enode://434a2b30d484449565ab0c51fcd800a11199826ed12ee1c857d9b4cd4fec46d90e44552165609d7576a775a4089746aa382794bdd3bbe31dd4923f4429ce70c8@221.12.12.211:30303",
	// "enode://374a44efd8d58c0223d1db080aa42ff28268ad5e823b76dec552878c096fa316193d9a1136d115802fed4a01ff943f1b187354a58740ebf26bc7dbb1bfc2106b@gchs.microcai.org:30303",
	// "enode://e750149c15c96bc55f4f29d94960872a8ad699fec91e8e8ef8eb2209c9330093835a830f5f4f725ac2a81b5544f6af53ef3e9ff71f3e38858fb83111dc5c43f9@gchs02.microcai.org:30303",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// These DNS names provide bootstrap connectivity for public testnets and the mainnet.
// See https://github.com/ethereum/discv4-dns-lists for more information.
var KnownDNSNetworks = map[common.Hash]string{
	MainnetGenesisHash: dnsPrefix + "mainnet.nodes.superpool.io",
	RopstenGenesisHash: dnsPrefix + "ropsten.nodes.superpool.io",
	RinkebyGenesisHash: dnsPrefix + "rinkeby.nodes.superpool.io",
	GoerliGenesisHash:  dnsPrefix + "goerli.nodes.superpool.io",
}
