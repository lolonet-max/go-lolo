// Copyright 2015 The go-lolo Authors
// This file is part of the go-lolo library.
//
// The go-lolo library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-lolo library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-lolo library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/lolo/go-lolo/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	//"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",   // bootnode-aws-ap-southeast-1-001
	//"enode://cc1c02dfc2c2d616f40cd888118269196e16c1bcd2de8f3193fc198f6d0f4dd4fbcd407bdae1b8ded517f1f85aae15f51ff58209edf42e67d014476042453270@103.72.167.16:30005",
	//"enode://fa48bc5c69139aec7c70f2ac0c4dd45c1055f9af081a4dc7b262c4ee434e86f99aa3479dc8d3d578e83bf7782baa49e90dd07c877ac5b73f71257000a0b69ada@160.19.51.70:30005",
	
	"enode://9bf283202c6ca3e19748a241b2f39e92e550b7e6a63b2fe4d1f2e3840909c7ac3203f1e935f635eac2178af68e249228401dff59e9be96e966c84777a72159c7@39.109.12.220:30005",
	"enode://418c835cb06a79084256f48e808d5a3dd2f83fe9fe1e5cd8f71e767c405bd9ed4d4623af68c75b19c07c8b93d41cb766791b5e0c77f3cd62f90f3b1fca23766b@39.109.12.149:30005",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{

}

// YoloV1Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv1 ephemeral test network.
var YoloV1Bootnodes = []string{
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/lolo/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
