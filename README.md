# To-do List
(Work in Progress)

Todo list web application. Made with Svelte and Go

You'll need MariaDB to run the aplication


On the server folder, create a file named '.env'. It's content should follow the following format:

```
DBURL = myuser:mypass@tcp(127.0.0.1:3306)/mydbname?charset=utf8mb4&parseTime=True&loc=Local
```

Inside the server folder, run the following commands:

```
go mod init server-app
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware@v4.10.2
go get github.com/joho/godotenv
go run main.go
```

Then you can launch the apllication with:
```
cd ..
npm install
npm run dev -- --open
```
