// Copyright 2021 The go-lolo Authors
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

package lolo

import (
	"os"

	"github.com/lolo/go-lolo/log"
)

// SetVerbosity sets the global verbosity level (between 0 and 6 - see logger/verbosity.go).
func SetVerbosity(level int) {
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(level), log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}