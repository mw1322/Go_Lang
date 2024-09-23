package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
	myDefer()
}
//hello 43210 world
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    // Deferred transaction handling
    defer func() {
        if err == nil {
            err = tx.Commit()
        }
        if err != nil {
            tx.Rollback()
        }
    }()

    // First insert
    _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values ($1)", value1)
    if err != nil {
        return err // No need to manually rollback; defer handles it
    }

    // Second insert (hypothetical)
    _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values ($1)", value2)
    if err != nil {
        return err // Defer will handle rollback
    }

    return err // The deferred function will commit or rollback based on this
}
