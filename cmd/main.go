package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"sync"
)

type Conf struct {
	PgUser string `envconfig:"PG_USER" required:"true"`
	PgPass string `envconfig:"PG_PASS" required:"true"`
	PgDb   string `envconfig:"PG_DB" required:"true"`
	PgPort int16  `envconfig:"PG_PORT" required:"true"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	urlList := [...]string{
		"https://ya.ru",
		"https://vk.com",
		"https://asdgsdadwhadsh.dsdshg",
		"https://hh.ru",
		"https://hh.ru",
		"https://hh.ru",
		"htts://hh.ru",
		"https://hh.ru",
		"https://hh.ru",
		"htt://hh.ru",
		"https://hh.ru",
		"https://hh.ru",
		"http://hh.ru",
		"https://hh.ru",
		"https://hh.ru",
		"https://x5.ru",
		"https://avito.ru",
	}

	var wg sync.WaitGroup

	wg.Add(len(urlList))
	errCh := make(chan error, len(urlList))
	for _, url := range urlList {
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Println("no valid http request", err)
				errCh <- err
			} else {
				log.Println(res.StatusCode)
				_ = res.Body.Close()
			}
		}(url)
	}
	wg.Wait()

	close(errCh)
	for v := range errCh {
		log.Println(v)
	}
	return
	//gracefulShutdown := make(chan os.Signal)
	//signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
	//<-gracefulShutdown

	var conf Conf
	err := envconfig.Process("app_config", &conf)
	if err != nil {
		log.Fatal(err)
	}

	dbh := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s", conf.PgUser, conf.PgPass, conf.PgPort, conf.PgDb)
	conn, err := pgx.Connect(context.Background(), dbh)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Printf("Ping err: %v", err)
	}
	defer func() { _ = conn.Close(context.Background()) }()
	type testTable struct {
		name string
		age  int
	}

	sql := "select name, age from test_table"
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		log.Printf("Err from query: %v", err)
	}

	var tt []testTable

	for rows.Next() {
		var t testTable
		err = rows.Scan(&t.name, &t.age)
		if err != nil {
			fmt.Println(err)
		} else {
			tt = append(tt, t)
		}
	}
	if rows.Err() != nil {
		fmt.Println(err)
	}

	for _, u := range tt {
		fmt.Println(u.age, u.name)
	}

	//server.Start()
}
