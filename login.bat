@echo off
setlocal enabledelayedexpansion

REM Define the JSON data as variables
set "Password=greatman"
set "Email=mohammed.yah.yousif@gmail.com"

REM Create a JSON string using the defined variables
set "jsonData={\"Password\":\"!Password!\",\"Email\":\"!Email!\"}"

REM Send the POST request with JSON data
curl -X POST -H "Content-Type: application/json" -d "%jsonData%" http://localhost:8000/login/

REM Pause to keep the window open (optional)
pause

endlocal
