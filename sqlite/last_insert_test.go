package sqlite

import "testing"

const (
    nInserts = 5
)

var cCreateTestTable = `
    CREATE TABLE IF NOT EXISTS
    test (
        str TEXT
    )
`

var cTestInsert = `
    INSERT INTO test (
        str
    ) VALUES (
        ?
    )
`

func TestSimpleLastInsertRowId(t *testing.T) {

    db, err := OpenV2(":memory:", OpenFullmutex|OpenReadwrite)

    if err != nil {
        t.Fatalf("Error opening database: %s\n", err);
    }

    // create a test table
    err = db.Exec(cCreateTestTable)

    if err != nil {
        db.Close()
        t.Fatalf("Error executing test table creation query: %s\n", err)
    }   

    for i := 1; i <= nInserts; i++ {

        err = db.Exec(cTestInsert, "test")

        if err != nil {
            db.Close()
            t.Fatalf("Error executing test insert query: %s\n", err)
        }

        rowid := db.LastInsertRowid()

        if rowid != int64(i) {
            db.Close()
            t.Fatalf("incorrect row id: expected %d, got %d\n", i, rowid)
        }
    }

    db.Close()
}
