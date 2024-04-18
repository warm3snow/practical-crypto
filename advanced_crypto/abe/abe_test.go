/**
 * @Author: xueyanghan
 * @File: abe_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/18 18:20
 */

package abe

import (
	"encoding/json"
	"github.com/fentec-project/gofe/abe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFAME(t *testing.T) {
	// create a new FAME struct with the universe of attributes
	// denoted by integer
	a := abe.NewFAME()

	// generate a public key and a secret key for the scheme
	pubKey, secKey, err := a.GenerateMasterKeys()
	if err != nil {
		t.Fatalf("Failed to generate master keys: %v", err)
	}

	// create a message to be encrypted
	msg := "Attack at dawn!"

	// create a msp struct out of a boolean expression representing the
	// policy specifying which attributes are needed to decrypt the ciphertext;
	// the boolean expression is a string of attributes joined by AND and OR
	// hence the names of the attributes should not include "AND" or "OR"
	// as a substring and '(' or ')' as a character

	// note that safety of the encryption is only proved if the mapping
	// msp.RowToAttrib from the rows of msp.Mat to attributes is injective, i.e.
	// only boolean expressions in which each attribute appears at most once
	// are allowed - if expressions with multiple appearances of an attribute
	// are needed, then this attribute can be split into more sub-attributes
	msp, err := abe.BooleanToMSP("((0 AND 1) OR (2 AND 3)) AND 5", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	// encrypt the message msg with the decryption policy specified by the
	// msp structure
	cipher, err := a.Encrypt(msg, msp, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gamma := []string{"0", "2", "3", "5"}

	// generate keys for decryption for an entity with
	// attributes gamma
	keys, err := a.GenerateAttribKeys(gamma, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheck, err := a.Decrypt(cipher, keys, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheck)

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaInsuff := []string{"1", "3", "5"}

	// generate keys for decryption for an entity with
	// attributes gammaInsuff
	keysInsuff, err := a.GenerateAttribKeys(gammaInsuff, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipher, keysInsuff, pubKey)
	assert.Error(t, err)

	mspSingleCondition, err := abe.BooleanToMSP("0", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	// encrypt the message msg with the decryption policy specified by the
	// msp structure
	cipherSingleCondition, err := a.Encrypt(msg, mspSingleCondition, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	msgCheckSingleCondition, err := a.Decrypt(cipherSingleCondition, keys, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckSingleCondition)

	_, err = a.Decrypt(cipherSingleCondition, keysInsuff, pubKey)
	assert.Error(t, err)

	// test with Single UUID
	mspSingleUUID, err := abe.BooleanToMSP("123e4567-e89b-12d3-a456-426655440000", false)

	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	cipherSingleUUID, err := a.Encrypt(msg, mspSingleUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaUUID := []string{"123e4567-e89b-12d3-a456-426655440000", "123e4567-e89b-12d3-a456-4266554400001"}

	// generate keys for decryption for an entity with
	// attributes gamma
	keysUUID, err := a.GenerateAttribKeys(gammaUUID, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheckSingleUUID, err := a.Decrypt(cipherSingleUUID, keysUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckSingleUUID)

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaInsuffUUID := []string{"123e4567-e89b-12d3-a456-426655440099"}

	// generate keys for decryption for an entity with
	// attributes gammaInsuff
	keysInsuffUUID, err := a.GenerateAttribKeys(gammaInsuffUUID, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipherSingleUUID, keysInsuffUUID, pubKey)
	assert.Error(t, err)

	//
	// test with Multi UUID
	mspMultiUUID, err := abe.BooleanToMSP("123e4567-e89b-12d3-a456-426655440000 OR 123e4567-e89b-12d3-a456-426655440001", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	cipherMultiUUID, err := a.Encrypt(msg, mspMultiUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheckMultiUUID, err := a.Decrypt(cipherMultiUUID, keysUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckMultiUUID)

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipherMultiUUID, keysInsuffUUID, pubKey)
	assert.Error(t, err)
}

func TestABEMarshal(t *testing.T) {

	// create a new FAME struct with the universe of attributes
	// denoted by integer
	a := abe.NewFAME()

	// generate a public key and a secret key for the scheme
	pubKey, secKey, err := a.GenerateMasterKeys()
	if err != nil {
		t.Fatalf("Failed to generate master keys: %v", err)
	}

	// create a message to be encrypted
	msg := "Attack at dawn!"

	// create a msp struct out of a boolean expression representing the
	// policy specifying which attributes are needed to decrypt the ciphertext;
	// the boolean expression is a string of attributes joined by AND and OR
	// hence the names of the attributes should not include "AND" or "OR"
	// as a substring and '(' or ')' as a character

	// note that safety of the encryption is only proved if the mapping
	// msp.RowToAttrib from the rows of msp.Mat to attributes is injective, i.e.
	// only boolean expressions in which each attribute appears at most once
	// are allowed - if expressions with multiple appearances of an attribute
	// are needed, then this attribute can be split into more sub-attributes
	msp, err := abe.BooleanToMSP("((0 AND 1) OR (2 AND 3)) AND 5", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	// encrypt the message msg with the decryption policy specified by the
	// msp structure
	cipher, err := a.Encrypt(msg, msp, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	cipherBytes, err := json.Marshal(cipher)
	if err != nil {
		t.Fatalf("Failed to marshal cipher: %v", err)
	}

	t.Logf("len of cipherBytes: %d", len(cipherBytes))

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gamma := []string{"0", "2", "3", "5"}

	// generate keys for decryption for an entity with
	// attributes gamma
	keys, err := a.GenerateAttribKeys(gamma, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes

	var cipher2 abe.FAMECipher
	err = json.Unmarshal(cipherBytes, &cipher2)
	if err != nil {
		t.Fatalf("Failed to unmarshal cipher: %v", err)
	}

	msgCheck, err := a.Decrypt(&cipher2, keys, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheck)

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaInsuff := []string{"1", "3", "5"}

	// generate keys for decryption for an entity with
	// attributes gammaInsuff
	keysInsuff, err := a.GenerateAttribKeys(gammaInsuff, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipher, keysInsuff, pubKey)
	assert.Error(t, err)

	mspSingleCondition, err := abe.BooleanToMSP("0", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	// encrypt the message msg with the decryption policy specified by the
	// msp structure
	cipherSingleCondition, err := a.Encrypt(msg, mspSingleCondition, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	msgCheckSingleCondition, err := a.Decrypt(cipherSingleCondition, keys, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckSingleCondition)

	_, err = a.Decrypt(cipherSingleCondition, keysInsuff, pubKey)
	assert.Error(t, err)

	// test with Single UUID
	mspSingleUUID, err := abe.BooleanToMSP("123e4567-e89b-12d3-a456-426655440000", false)

	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	cipherSingleUUID, err := a.Encrypt(msg, mspSingleUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaUUID := []string{"123e4567-e89b-12d3-a456-426655440000", "123e4567-e89b-12d3-a456-4266554400001"}

	// generate keys for decryption for an entity with
	// attributes gamma
	keysUUID, err := a.GenerateAttribKeys(gammaUUID, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheckSingleUUID, err := a.Decrypt(cipherSingleUUID, keysUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckSingleUUID)

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gammaInsuffUUID := []string{"123e4567-e89b-12d3-a456-426655440099"}

	// generate keys for decryption for an entity with
	// attributes gammaInsuff
	keysInsuffUUID, err := a.GenerateAttribKeys(gammaInsuffUUID, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipherSingleUUID, keysInsuffUUID, pubKey)
	assert.Error(t, err)

	//
	// test with Multi UUID
	mspMultiUUID, err := abe.BooleanToMSP("123e4567-e89b-12d3-a456-426655440000 OR 123e4567-e89b-12d3-a456-426655440001", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	cipherMultiUUID, err := a.Encrypt(msg, mspMultiUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheckMultiUUID, err := a.Decrypt(cipherMultiUUID, keysUUID, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheckMultiUUID)

	// try to decrypt the ciphertext with the keys of an entity
	// that has insufficient attributes
	_, err = a.Decrypt(cipherMultiUUID, keysInsuffUUID, pubKey)
	assert.Error(t, err)
}
