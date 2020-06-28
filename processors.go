package main

import(
  "net/http"
  _ "github.com/lib/pq"
  "regexp"
)

func createprocessor(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  fname := r.FormValue("tableFirstName")
  lname := r.FormValue("tableLastName")
  bdate := r.FormValue("tableBDate")
  gender := r.FormValue("tableGender")
  email := r.FormValue("tableEmail")
  address := r.FormValue("tableAddress")
  reg, err := regexp.Compile("[^a-z-A-Z]")
  if err != nil {
    panic(err)
  }
  fname = reg.ReplaceAllString(fname, "")
  lname = reg.ReplaceAllString(lname, "")
  insertQuery(fname, lname, bdate, gender, email, address)

  http.Redirect(w, r, "/read", http.StatusSeeOther)
}

func updateprocessor(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  id := r.FormValue("tableId")
  fname := r.FormValue("tableFirstName")
  lname := r.FormValue("tableLastName")
  email := r.FormValue("tableEmail")
  address := r.FormValue("tableAddress")
  reg, err := regexp.Compile("[^a-z-A-Z]")
  if err != nil {
    panic(err)
  }
  fname = reg.ReplaceAllString(fname, "")
  lname = reg.ReplaceAllString(lname, "")
  updateQuery(id, fname, lname, email, address)

  http.Redirect(w, r, "/read", http.StatusSeeOther)
}

func deleteprocessor(w http.ResponseWriter, r *http.Request){
  id := r.URL.Query().Get("id")
  deleteQuery(id)
  http.Redirect(w, r, "/read", http.StatusSeeOther)
}
