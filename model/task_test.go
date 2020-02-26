package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
	"testing"
)

const TestConnectionString = "user=postgres password=admin dbname=todo_test host=localhost port=5432 sslmode=disable"

var dbtest *gorm.DB

func TestMain(m *testing.M) {
	db, err := gorm.Open("postgres", TestConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	dbtest = db

	InitService(dbtest)
	os.Exit(m.Run())
}

func prepareTestDatabase(name string) {
	fixtures, err := testfixtures.NewFolder(dbtest.DB(), &testfixtures.PostgreSQL{}, name)
	if err != nil {
		panic(err)
	}
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}

func TestGetTask(t *testing.T) {
	prepareTestDatabase("../fixtures")
	task, err := GetTask(1)
	if err != nil {
		t.Fatalf("did not expect error")
	}
	expected := "First task"

	if task.Title != expected {
		t.Fatalf("expected task title %s, got %s", expected, task.Title)
	}
}