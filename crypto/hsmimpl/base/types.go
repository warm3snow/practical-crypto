/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package base

import "C"

type DeviceInfo struct {
	IssuerName      string
	DeviceName      string
	DeviceSerial    string
	DeviceVersion   uint
	StandardVersion uint
	AsymAlgAbility  [2]uint
	SymAlgAbility   uint
	HashAlgAbility  uint
	BufferSize      uint
}

type DeviceRunStatus struct {
	Onboot      uint
	Service     uint
	Concurrency uint
	Memtotal    uint
	Memfree     uint
	Cpu         uint
	Reserve1    uint
	Reserve2    uint
}

type ECCrefPublicKey struct {
	Bits uint
	X    string
	Y    string
}

type ECCrefPrivateKey struct {
	Bits uint
	K    string
}

type ECCCipher struct {
	X string
	Y string
	M string
	L uint
	C string
}

type ECCSignature struct {
	R string
	S string
}
