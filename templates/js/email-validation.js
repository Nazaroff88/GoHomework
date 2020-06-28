function ValidateEmail(inputText)
{
  var mailformat = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
  if(document.getElementById("mail").value.match(mailformat))
  {
    //document.form1.text1.focus();
    return true;
  }
  else
  {
    alert("You have entered an invalid email address!");
    document.form.tableEmail.focus();
    return false;
  }
}
