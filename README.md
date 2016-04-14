# ToDoList API

## Installation
	1.you need to installaion MYSQL
	2.Change the db.go file database name also

## Import packages
	go get github.com/op/go-logging
	go get github.com/gin-gonic/gin
	go get github.com/go-sql-driver/mysql
	go get github.com/jinzhu/gorm

## To run the server
	
	go run server.go	

## Routs

	### users

		1. View all the data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users

		2. View by ID
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users/{id}

		3. Add data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"ID":"1", "Login":"admin", "Email":"admin@localhost","Passwords":"86a65acd94b33daa87c1c6a2d1408593",  "TofactorSecret":"", "Ful1lname":"", "Description":"", "Enabled":"1", "Verified":"", "Role":"1", "CreateAt":"2013-11-06 09:24:00", "UpdateAt":"2013-11-06 09:25:07"}' http://localhost:8000/users

		4. Update data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"ID":"1", "Login":"blowmedia", "Email":"sergey.page@blowmedianow.com", "Passwords":"86a65acd94b33daa87c1c6a2d1408593", "Fullname":"", "Description":"", "Enabled":"1", "Verified":"1", "Role":"1", "CreateAt":"2013-11-06 09:24:00", "UpdateAt":"2013-11-06 09:25:07"}' http://localhost:8000/users/{id}

		5. Delete data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X DELETE -d '{"ID":"1", "Login":"blowmedia", "Email":"sergey.page@blowmedianow.com", "Passwords":"86a65acd94b33daa87c1c6a2d1408593", "Fu1llname":"", "Description":"", "Enabled":"0", "Verified":"", "Role":"1", "CreateAt":"2013-11-06 09:24:00", "UpdateAt":"2013-11-06 09:25:07"}' http://localhost:8000/users/{id}
	
	### todolist

		1. View all the data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/todolist 

		2. View by ID
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/todolist/{id}

		3. Add data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{ "ID":"1", "UserID":"1", "Item":"", "Description":"", "StartDate":"2016-04-03 22:57:27", "EndDate":"2016-04-06 22:57:27","Enabled":"1", "CreatedAt":"null", "UpdatedAt":"null"}' http://localhost:8000/todolist

		4. Update data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{ "ID":"1", "UserID":"1", "Item":"", "Description":"test description", "StartDate":"2016-04-03 22:57:27", "EndDate":"2016-04-07 22:57:27","Enabled":"1", "CreatedAt":"2014-09-03 22:57:27", "UpdatedAt":"null"}' http://localhost:8000/todolist/{id}

		5. Delete data
			curl -v -H "Accept: application/json" -H "Content-type: application/json" -X DELETE -d '{"ID":"1", "UserID":"1", "Item":"", "Description":"test description", "StartDate":"2016-04-03 22:57:27", "EndDate":"2016-04-07 22:57:27","Enabled":"0", "CreatedAt":"2014-09-03 22:57:27", "UpdatedAt":"null"}' http://localhost:8000/todolist/{id}

	
