package main

import (
	"fmt"
	"context"

        "github.com/ledgerwatch/erigon/ethdb/kv"
        "github.com/ledgerwatch/erigon/common/dbutils"
)

/* 

from erigon/common/dbutils/bucket.go
----
//key - contract code hash
//value - contract code
CodeBucket = "Code"
----
*/

// dump all unique (codehash, bytecode)
func dumpCode(ctx context.Context) error {
	dbpath := "/nvme-disk/erigon-data/erigon/chaindata"
        db, err := kv.Open(dbpath, true)
        if err != nil {
                return err
        }
        defer db.Close()

        tx, err := db.BeginRo(ctx)
        if err != nil {
                return err
        }
        defer tx.Rollback()

	return tx.ForEach(dbutils.CodeBucket, []byte{}, func(k, v []byte) error {
		fmt.Println(fmt.Sprintf("%x, %x", k, v))
		return nil
	})
}

func main() {
	if err := dumpCode(context.Background()); err != nil {
		panic(err)
	}
}
