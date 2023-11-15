@echo off
setlocal enabledelayedexpansion

REM Define the JSON data as variables
set "Title=Reem note"
set "OwnerID='654fbfec6aef15bbafd568c1'"
set "Content=Hello world"

REM Create a JSON string using the defined variables
set "jsonData={\"Title\":\"!Title!\",\"OwnerID\":\"!OwnerID!\",\"Content\":\"!Content!\"}"

REM Send the POST request with JSON data
curl -X POST -H "Content-Type: application/json" -d "%jsonData%" http://localhost:8000/note/

REM Pause to keep the window open 
pause

endlocal
