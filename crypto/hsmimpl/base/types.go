/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package base

// DeviceInfo sdf device info
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

// DeviceRunStatus sdf device status
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

// ECCrefPublicKey ecc public key
type ECCrefPublicKey struct {
	Bits uint
	X    string
	Y    string
}

// ECCrefPrivateKey ecc private key
type ECCrefPrivateKey struct {
	Bits uint
	K    string
}

// ECCCipher ecc cipher
type ECCCipher struct {
	X string
	Y string
	M string
	L uint
	C string
}

// ECCSignature ecc signature
type ECCSignature struct {
	R string
	S string
}
