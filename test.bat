@echo off
for /l %%i in (1, 1, 1000) do (
    curl http://localhost:8000/%%i
)
