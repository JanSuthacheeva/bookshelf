package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os/exec"
	"path/filepath"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "test_web:pass@/test_bookshelf?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal("Unable to ping the database:", err)
	}
	projectRoot, err := filepath.Abs("../../")
	if err != nil {
		db.Close()
		t.Fatal(err)
	}
	cmd := exec.Command("make", "test_up")
	cmd.Dir = projectRoot
	err = cmd.Run()
	if err != nil {
		db.Close()
		t.Fatal(err)
	}
	cmd = exec.Command("make", "test_run_seeder")
	cmd.Dir = projectRoot
	err = cmd.Run()
	if err != nil {
		db.Close()
		t.Fatal(err)
	}

	t.Cleanup(func() {
		defer db.Close()

		cmd := exec.Command("make", "test_reset")
		cmd.Dir = projectRoot

		err = cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	return db
}
