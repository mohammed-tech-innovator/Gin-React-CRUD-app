@echo off
setlocal enabledelayedexpansion

REM Define the JSON data as variables
set "Password=greatman"
set "Newpassword1=great man"
set "Newpassword2=great man"
set "Email=mohammed.yah.yousif@gmail.com"

REM Create a JSON string using the defined variables
set "jsonData={\"Email\":\"!Email!\",\"Password\":\"!Password!\",\"Newpassword1\":\"!Newpassword1!\",\"Newpassword2\":\"!Newpassword2!\"}"

REM Send the PUT request with JSON data
curl -X PUT -H "Content-Type: application/json" -d "%jsonData%" http://localhost:8000/change-password/

REM Pause to keep the window open 
pause

endlocal
