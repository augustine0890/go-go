package main

import (
	"database/sql"
	"flag"
	"log"

	"snippetbox/pkg/models/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port = flag.String("port", ":4000", "Port to listen")
	dsn  = flag.String("dsn", "admin:admin@/snippetbox?parseTime=true", "MariaDB data source name")
)

func main() {
	// Parse flags
	flag.Parse()

	// Connection pool: many connections
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize html tamplate engine
	engine := html.New("./ui/html", ".html")
	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	app := fiber.New(fiber.Config{
		AppName: "SnippetBox",
		Views:   engine,
	})

	// Log middleware config
	app.Use(logger.New())

	// Static files
	app.Static("/", "./ui/static")

	repository := mysql.NewRepo(db)
	service := NewService(repository)
	Routes(app, service)

	log.Printf("Starting server on %s", *port)
	log.Fatal(app.Listen(*port))
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
