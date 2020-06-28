package main

import(
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const(
  host = "localhost"
  port = 5432
  user = "postgres"
  password = "postgres"
  dbname = "Persons_DB"
)

var db *sql.DB

type Person struct {
    Id   int
    FirstName  string
    LastName string
    Bdate  string
    Gender string
    Email string
    Address string
}
func connectDB() *sql.DB{
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }else{
    fmt.Println("DB is opened")
  }
  return db
}

func insertQuery(fname string, lname string, bdate string, gender string, email string, address string){
  fmt.Println(bdate)
  db = connectDB()
  rows, err := db.Query("INSERT INTO persons (f_name, l_name, b_date, gender, email, address) VALUES ('"+fname+"', '"+lname+"','"+bdate+"','"+gender+"','"+email+"','"+address+"')")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  defer db.Close()
}

func deleteQuery(id string){
  db = connectDB()
  rows, err := db.Query("DELETE FROM persons WHERE id="+id+"")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  defer db.Close()
}

func selectQuery() []Person{
  var id int
  var firstName string
  var lastName string
  var birthDate string
  var gender string
  var email string
  var address string
  var persons []Person
  db = connectDB()
  rows, err := db.Query("SELECT * FROM persons")
  if err != nil {
    // handle this error better than this
    panic(err)
  }

    //loop through the db
  for rows.Next() {
    err = rows.Scan(&id, &firstName, &lastName, &birthDate, &gender, &email, &address)
    if err != nil {
      panic(err)
    }
    persons = append(persons, Person{Id: id, FirstName: firstName, LastName: lastName, Bdate: birthDate, Gender: gender, Email: email, Address: address})
  }

  defer rows.Close()
  defer db.Close()

  return persons
}

func updateQuery(id string, fname string, lname string, email string, address string){
  db = connectDB()
  rows, err := db.Query("UPDATE persons SET f_name='"+fname+"', l_name='"+lname+"', email='"+email+"', address='"+address+"' WHERE id="+id+"")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  defer db.Close()
}

func searchQuery(fname string, lname string) []Person{
  var id int
  var firstName string
  var lastName string
  var birthDate string
  var gender string
  var email string
  var address string
  var persons []Person
  db = connectDB()
  rows, err := db.Query("SELECT * FROM persons WHERE f_name LIKE '%"+fname+"%' AND l_name LIKE '%"+lname+"%'")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
    //loop through the db
  for rows.Next() {
    err = rows.Scan(&id, &firstName, &lastName, &birthDate, &gender, &email, &address)
    if err != nil {
      panic(err)
    }
    persons = append(persons, Person{Id: id, FirstName: firstName, LastName: lastName, Bdate: birthDate, Gender: gender, Email: email, Address: address})
  }
  defer rows.Close()
  defer db.Close()
  return persons
}
