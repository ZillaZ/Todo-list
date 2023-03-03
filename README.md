# To-do List
(Work in Progress)

Todo list web application. Made with Svelte and Go

You'll need SQLite to run the aplication


On the server folder, initialize a database called "todo.db":

```
sqlite3 todo.db
```

Inside the server folder, run the following commands:

```
go mod init server-app
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware@v4.10.2
go run server/main.go
```

Then you can launch the apllication with:
```
npm install
npm run dev -- --open
```
