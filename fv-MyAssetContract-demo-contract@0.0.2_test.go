/**
 * SPDX-License-Identifier: Apache-2.0
 */

/**
 * Use this file for functional testing of your smart contract.
 * Fill out the arguments and return values for a function and
 * use the CodeLens links above the transaction blocks to
 * invoke/submit transactions.
 * All transactions defined in your smart contract are used here
 * to generate tests, including those functions that would
 * normally only be used on instantiate and upgrade operations.
 * This basic test file can also be used as the basis for building
 * further functional tests to run as part of a continuous
 * integration pipeline, or for debugging locally deployed smart
 * contracts by invoking/submitting individual transactions.
 *
 * Generating this test file will also run 'go mod vendor'.
 */

package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	gw "github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/stretchr/testify/assert"
)

var homedir = os.Getenv("HOME")
var walletPath = filepath.Join(homedir, ".fabric-vscode", "v2", "environments", "1 Org Local Fabric", "wallets", "Org1")
var connectionProfilePath = filepath.Join(homedir, ".fabric-vscode", "v2", "environments", "1 Org Local Fabric", "gateways", "conn.json")
var isLocalhostURL, _ = HasLocalhostURLs(connectionProfilePath)

// const identityName = "Org1 Admin"

func setup(t *testing.T) (*gw.Contract, func(t *testing.T)) {
	SetDiscoverAsLocalHost(isLocalhostURL)

	log.Println("******************************************")
	log.Println(homedir)
	log.Println("walletPath: "+walletPath)
	log.Println("connectionProfilePath: " + connectionProfilePath)
	log.Println(isLocalhostURL)
	log.Println("******************************************")


	fabricWallet, err := gw.NewFileSystemWallet(walletPath)
	if err != nil {
		t.Fatalf("Failed to create wallet: %s\n", err)
	} else if !fabricWallet.Exists("Org1 Admin") {
		t.Fatalf("Identity %s\n not present in wallet", "Org1 Admin")
	}

	gateway, err := gw.Connect(
		gw.WithConfig(config.FromFile(filepath.Clean(connectionProfilePath))),
		gw.WithIdentity(fabricWallet, "Org1 Admin"),
	)
	if err != nil {
		t.Fatalf("Failed to connect to gateway: %s\n", err)
	}

	network, err := gateway.GetNetwork("mychannel")
	if err != nil {
		t.Fatalf("Failed to get network: %s\n", err)
	}

	contract := network.GetContractWithName("demo-contract", "MyAssetContract")

	return contract, func(t *testing.T) {
		gateway.Close()
	}
}

func TestCreateMyAsset_fv(t *testing.T) {
    t.Run("SubmitCreateMyAssetTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        // TODO: populate transaction parameters
        param0 := "EXAMPLE"
        param1 := "EXAMPLE"
        transaction, err := contract.CreateTransaction("CreateMyAsset")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit(param0,param1)
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

        // TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
    })
}

func TestDeleteMyAsset_fv(t *testing.T) {
    t.Run("SubmitDeleteMyAssetTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        // TODO: populate transaction parameters
        param0 := "EXAMPLE"
        transaction, err := contract.CreateTransaction("DeleteMyAsset")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit(param0)
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

        // TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
    })
}

func TestMyAssetExists_fv(t *testing.T) {
    t.Run("SubmitMyAssetExistsTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        // TODO: populate transaction parameters
        param0 := "EXAMPLE"
        transaction, err := contract.CreateTransaction("MyAssetExists")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit(param0)
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

		// TODO: Update with return value of transaction
		assert.EqualValues(t, false,result)
    })
}

func TestQueryAllAssets_fv(t *testing.T) {
    t.Run("SubmitQueryAllAssetsTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        transaction, err := contract.CreateTransaction("QueryAllAssets")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit()
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

        // TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
    })
}

func TestReadMyAsset_fv(t *testing.T) {
    t.Run("SubmitReadMyAssetTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        // TODO: populate transaction parameters
        param0 := "EXAMPLE"
        transaction, err := contract.CreateTransaction("ReadMyAsset")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit(param0)
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

        // TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
    })
}

func TestUpdateMyAsset_fv(t *testing.T) {
    t.Run("SubmitUpdateMyAssetTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
        // TODO: populate transaction parameters
        param0 := "EXAMPLE"
        param1 := "EXAMPLE"
        transaction, err := contract.CreateTransaction("UpdateMyAsset")
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
        result, err := transaction.Submit(param0,param1)
        
		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
        }

        // TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
    })
}
