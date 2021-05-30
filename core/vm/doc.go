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

/*
Package vm implements the Ethereum Virtual Machine.

The vm package implements one EVM, a byte code VM. The BC (Byte Code) VM loops
over a set of bytes and executes them according to the set of rules defined
in the Ethereum yellow paper.
*/
package vm
