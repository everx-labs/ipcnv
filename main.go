/*
	ipcnv – Simple IP address conversion tool

	Copyright (C) 2023 EverX

	This file is part of ipcnv.

	ipcnv is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	ipcnv is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with ipcnv.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"unsafe"
)

var IS_BIG_ENDIAN bool

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(1)
}

func checkEndianness() {
	var i int = 0x0100
	ptr := unsafe.Pointer(&i)

	if *(*byte)(ptr) == 0x01 {
		IS_BIG_ENDIAN = true
	} else if *(*byte)(ptr) == 0x00 {
		IS_BIG_ENDIAN = false
	} else {
		fatal(errors.New("can not check endianness"))
	}
}

func ipv4ToInt32(input string) (string, error) {
	ip := net.ParseIP(input)
	if ip == nil {
		return "", errors.New("invalid ipv4 address")
	}

	v4 := ip.To4()
	if v4 == nil || len(v4) != 4 {
		return "", errors.New("invalid ipv4 address")
	}

	if !IS_BIG_ENDIAN {
		v4le := make([]byte, 4)
		for i := 0; i < 4; i++ {
			v4le[i] = v4[4-i-1]
		}
		v4 = v4le
	}

	i := (*(*int32)(unsafe.Pointer(&v4[0])))
	return strconv.Itoa(int(i)), nil
}

func int32ToIpv4(input string) (string, error) {
	i, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return "", err
	}

	v4 := make([]byte, 4)
	v4[0] = byte(i >> 24)
	v4[1] = byte(i >> 16)
	v4[2] = byte(i >> 8)
	v4[3] = byte(i)

	return net.IP(v4).To4().String(), nil
}

func main() {
	checkEndianness()

	var input string
	var mode int

	flag.StringVar(&input, "i", "", "input ip address")
	flag.IntVar(&mode, "m", -1, "0 - ipv4 to int32\n1 - int32 to ipv4")

	flag.Parse()

	if mode < 0 || mode > 1 {
		fatal(errors.New("mode must be >= 0 and <= 1"))
	}

	if input == "" {
		fatal(errors.New("-i flag must not be empty"))
	}

	var result string
	var err error

	switch mode {
	case 0:
		result, err = ipv4ToInt32(input)
	case 1:
		result, err = int32ToIpv4(input)
	}

	if err != nil {
		fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", result)
}
