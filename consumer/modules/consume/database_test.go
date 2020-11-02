package consume

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDbConn(t *testing.T) {
	t.Log("Connecting To SQL Server...")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		t.Fatal("Connection To SQL Failed: ", err)
	}
	t.Log("Connected To SQL Successfully with connection details:", db)
}
