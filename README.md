##Installation instructions
In order to run the app you need to have one of mysql or mongodb installed 

Two type of database option available `mongodb` `sql`
By default app work with mongodb if you are not specify database when running app 

To use `sql` you should create a database and run sql file located in extra folder

You can change the `port` and database `name` in config

Then just need to run code by following command
```sh
go run main.go
```

To specify database add `db_type` flag like below
```sh
go run main.go -db_type=mongodb
```
or

```sh
go run main.go -db_type=sql
```

##Technologies

library I have used for this app are in the below
```sh
github.com/dgrijalva/jwt-go v3.2.0+incompatible
github.com/gin-contrib/cors v1.3.1
github.com/gin-gonic/gin v1.6.2
github.com/go-sql-driver/mysql v1.5.0
github.com/google/uuid v1.1.1
github.com/jmoiron/sqlx v1.2.0
github.com/pkg/errors v0.9.1
go.mongodb.org/mongo-driver v1.3.2
golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
```

##Requirement

We should create 3 api as following

    > register user
    > login user
    > get user info by token

All covered!

##Test
I haven't got time to write test against api yet 
to test them you can find postman profile in extra folder which you can import and test the api

