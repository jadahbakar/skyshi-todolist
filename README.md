<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h2 align="center">Skyshi - TodoList - Backend</h2> <br />
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
     <li>
      <a href="#testing">Testing</a>
      <ul>
        <li><a href="#testing-daftar">Testing on Daftar</a></li>
        <li><a href="#testing-tabung">Testing on Tabung</a></li>
        <li><a href="#testing-tarik">Testing on Tarik</a></li>
        <li><a href="#testing-saldo">Testing on Saldo</a></li>
        <li><a href="#testing-mutasi">Testing on Mutasi</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

This system is using for Skyshi - TodoList - Challenge.


<!-- BUILD WITH -->
#### Built With

* [Golang](https://golang.org)
* [PostgreSQL](https://www.postgresql.org/)
* [Fiber](https://docs.gofiber.io/)
* [Viper](https://github.com/spf13/viper)
* [PgxPool](https://github.com/jackc/pgx)


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.


<!-- PREREQUISITES -->
#### Prerequisites

* I use macOS Catalina 10.15.7
  ```sh
  ❯ uname -a
  Darwin Kernel Version 19.6.0
   ```
* And then you need to install [Golang](https://golang.org/doc/install)

* After that check the installation and Golang version it must be above than 1.11 because we need the [Golang Modules](https://blog.golang.org/using-go-modules)
  ```sh
  > go version
  go version go1.21.0 darwin/amd64
   ```


<!-- INSTALLATION -->
#### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/jadahbakar/skyshi-todolist.git
   ```
2. Install Modules packages
   ```sh
   go mod tidy
   ```
   or
   ```sh
   go mod download
   ```
3. Create PostgreSQL Database
   ```sh
   make postgres
   ```
4. Create PostgreSQL Database
   ```sh
   make createdb
   ```
5. Create PostgreSQL Table
   ```sh
   make migrateup
   ```
6. Run program
   ```sh
   make server
   ```
  ![running](assets/images/command-running.png)

#### OR
7. Run from Docker
   ```sh
   docker compose up
   ```
   ![docker](assets/images/docker-run.png)

7. Test using Postman, with address and port
  ```sh
  localhost:8000

  ```

<!-- Testing Testing -->
## Testing
#### testing-daftar
1. Testing daftar

  ![daftar](assets/images/postman-daftar.png)

#### testing-tabung
2. Testing tabung

  ![daftar](assets/images/postman-tabung.png)

#### testing-tarik
3. Testing tarik

  ![daftar](assets/images/postman-tarik.png)

#### testing-saldo
4. Testing saldo

  ![daftar](assets/images/postman-saldo.png)

#### testing-mutasi
5. Testing mutasi

  ![daftar](assets/images/postman-mutasi.png)


<!-- Project structure -->
## Project Structure

```sh
.
├── Dockerfile
├── Makefile
├── README.md
├── assets
│   └── images
│       ├── command-running.png
│       ├── postman-daftar.png
│       ├── postman-mutasi.png
│       ├── postman-saldo.png
│       ├── postman-tabung.png
│       └── postman-tarik.png
├── config.env
├── db
│   └── migration
│       ├── 000001_init_schema.down.sql
│       └── 000001_init_schema.up.sql
├── docker-compose.yml
├── go.mod
├── go.sum
├── log
│   ├── fiber.log
│   └── skyshi-todolist
├── main.go
├── repository
│   └── postgres
│       └── repository.go
├── script
│   ├── start.sh
│   └── wait-for.sh
├── tabungan
│   ├── handler.go
│   ├── model.go
│   ├── repository.go
│   └── service.go
└── util
    ├── config
    │   └── config.go
    ├── engine
    │   └── engine.go
    ├── generate
    │   └── generate.go
    ├── greet
    │   └── greet.go
    ├── logger
    │   ├── appLogger.go
    │   ├── fiberLogger.go
    │   └── fiberLogger_test.go
    ├── response
    │   ├── error.go
    │   └── success.go
    └── validator
        └── validator.go

18 directories, 35 files
```

