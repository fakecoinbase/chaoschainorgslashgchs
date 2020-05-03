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

	"enode://06f4536f3956646c49e32d82ac31721669012324c4da28210795865d2cd821ffc88737c442c5458788e0ef9252df4c549226d69092df0d035e2d5abdac7723e5@49.235.119.123:30303",

	"enode://464ecec2a80488b93c83592da368b38338dc530dff0d3f37e0d60f02ec41f8e29a1c09d592ad6352cf2f256ed2efe73e72e65401d253217056801b5fd708d3ee@[2408:8642:4ff:8:1::db7]:10101",
	"enode://464ecec2a80488b93c83592da368b38338dc530dff0d3f37e0d60f02ec41f8e29a1c09d592ad6352cf2f256ed2efe73e72e65401d253217056801b5fd708d3ee@221.12.12.211:10101",

	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@[2408:8642:4ff:8:1::fb10]:10102",
	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@221.12.12.211:10102",

	"enode://e10302c0f953edb02dc1f46f5489a9ac97dea4c0db76305763239df63541917457d5c087571e53842f7c61feb5cf5e4a0424d913fd8882738c787166f91a6ad8@[2408:8642:4ff:8:1::80f]:10103",
	"enode://e10302c0f953edb02dc1f46f5489a9ac97dea4c0db76305763239df63541917457d5c087571e53842f7c61feb5cf5e4a0424d913fd8882738c787166f91a6ad8@221.12.12.211:10103",

	"enode://2c3791581ee083cb55eb6f5afcd8bc2e46962cbab0bd95ad1aac961ab941eb3cf45480b98801fc49787cbff42a3fea8871964464ae95d29cd2ca911a20d270e0@[2408:8642:4ff:8:1::e0a0]:10104",
	"enode://2c3791581ee083cb55eb6f5afcd8bc2e46962cbab0bd95ad1aac961ab941eb3cf45480b98801fc49787cbff42a3fea8871964464ae95d29cd2ca911a20d270e0@221.12.12.211:10104",

	"enode://03ae723ec5dd6dfdc926d7d73cf03ea8c9da0a0a2546a82b348af45cfc52235cc60a78c5699075644a7e98fd2a57232ccf0b034d070945baa1fbde9ba1a700c9@221.12.12.211:30301",

	"enode://7f44aa20a629f43661e260b4653b623f0c32620d02b1321a6762c09707930a4ec22ec7529859bd4dfc1317e55bae6635a2a71fbc4ead0df3ec2e4a4625eac3a2@gchs.p2sp.net:30303",
	"enode://b2a13b6fb4b969354b264de1ba7e9d8c054b0c2330b4efb4df7024e615c070f782cb88568f4fb45f7783920b2826be0b761501bcc802d608f2a138df0c5c3b4c@gchs.microcai.org:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06f4536f3956646c49e32d82ac31721669012324c4da28210795865d2cd821ffc88737c442c5458788e0ef9252df4c549226d69092df0d035e2d5abdac7723e5@49.235.119.123:30303",

	"enode://464ecec2a80488b93c83592da368b38338dc530dff0d3f37e0d60f02ec41f8e29a1c09d592ad6352cf2f256ed2efe73e72e65401d253217056801b5fd708d3ee@[2408:8642:4ff:8:1::db7]:10101",
	"enode://464ecec2a80488b93c83592da368b38338dc530dff0d3f37e0d60f02ec41f8e29a1c09d592ad6352cf2f256ed2efe73e72e65401d253217056801b5fd708d3ee@221.12.12.211:10101",

	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@[2408:8642:4ff:8:1::fb10]:10102",
	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@221.12.12.211:10102",

	"enode://e10302c0f953edb02dc1f46f5489a9ac97dea4c0db76305763239df63541917457d5c087571e53842f7c61feb5cf5e4a0424d913fd8882738c787166f91a6ad8@[2408:8642:4ff:8:1::80f]:10103",
	"enode://e10302c0f953edb02dc1f46f5489a9ac97dea4c0db76305763239df63541917457d5c087571e53842f7c61feb5cf5e4a0424d913fd8882738c787166f91a6ad8@221.12.12.211:10103",

	"enode://2c3791581ee083cb55eb6f5afcd8bc2e46962cbab0bd95ad1aac961ab941eb3cf45480b98801fc49787cbff42a3fea8871964464ae95d29cd2ca911a20d270e0@[2408:8642:4ff:8:1::e0a0]:10104",
	"enode://2c3791581ee083cb55eb6f5afcd8bc2e46962cbab0bd95ad1aac961ab941eb3cf45480b98801fc49787cbff42a3fea8871964464ae95d29cd2ca911a20d270e0@221.12.12.211:10104",

	"enode://03ae723ec5dd6dfdc926d7d73cf03ea8c9da0a0a2546a82b348af45cfc52235cc60a78c5699075644a7e98fd2a57232ccf0b034d070945baa1fbde9ba1a700c9@221.12.12.211:30301",
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
