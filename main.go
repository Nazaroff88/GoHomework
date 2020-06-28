package main

import(
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
  "strconv"
  "io/ioutil"
  "regexp"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
  http.HandleFunc("/", mainPage)
  http.HandleFunc("/creat", createPage)
  http.HandleFunc("/read", readPage)
  http.HandleFunc("/update", updatePage)
  http.HandleFunc("/createprocess", createprocessor)
  http.HandleFunc("/updateprocess", updateprocessor)
  http.HandleFunc("/deleteprocess", deleteprocessor)
  http.HandleFunc("/email-validation.js", SendEmailValidation)
  http.HandleFunc("/sort-table.js", SendPageSort)
  http.ListenAndServe(":8080", nil)
}

func SendEmailValidation(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("templates/js/email-validation.js")
    if err != nil {
        http.Error(w, "Couldn't read file", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
    w.Write(data)
}

func SendPageSort(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("templates/js/sort-table.js")
    if err != nil {
        http.Error(w, "Couldn't read file", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
    w.Write(data)
}

func mainPage(w http.ResponseWriter, r *http.Request){
  tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func createPage(w http.ResponseWriter, r *http.Request){
  tpl.ExecuteTemplate(w, "creat.gohtml", nil)
}

func readPage(w http.ResponseWriter, r *http.Request){
  var persons []Person
  if r.Method != "POST"{
    persons := selectQuery()
    tpl.ExecuteTemplate(w, "read.gohtml", persons)
  }else{
    searchFname := r.FormValue("FirstName")
    searchLname := r.FormValue("LastName")
    reg, err := regexp.Compile("[^a-zA]")
    if err != nil {
      panic(err)
    }
    searchFname = reg.ReplaceAllString(searchFname, "")
    searchLname = reg.ReplaceAllString(searchLname, "")
    persons = searchQuery(searchFname, searchLname)
    tpl.ExecuteTemplate(w, "read.gohtml", persons)
    }
}

func updatePage(w http.ResponseWriter, r *http.Request){
  id, _ := strconv.Atoi(r.URL.Query().Get("id"))
  fname := r.URL.Query().Get("firstName")
  lname := r.URL.Query().Get("lastName")
  bdate := r.URL.Query().Get("bdate")
  email := r.URL.Query().Get("email")
  address := r.URL.Query().Get("address")
    person := Person{
      Id: id,
      FirstName: fname,
      LastName: lname,
      Bdate: bdate,
      Gender: "ff",
      Email: email,
      Address: address,
    }
    tpl.ExecuteTemplate(w, "update.gohtml", person)
}
