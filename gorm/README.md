# GORM

Demonstrate basic CRUD operation using a store based on MySQL and GORM

[GORM](https://github.com/jinzhu/gorm) is simple ORM library for Go.
In this example the same struct is used both as the GORM model and as the JSON model.

curl demo:
```
curl -i -H 'Content-Type: application/json' \
    -d '{"Message":"this is a test"}' http://127.0.0.1:8080/reminders
curl -i http://127.0.0.1:8080/reminders/1
curl -i http://127.0.0.1:8080/reminders
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Message":"is updated"}' http://127.0.0.1:8080/reminders/1
curl -i -X DELETE http://127.0.0.1:8080/reminders/1
```
