@echo off
for /l %%i in (1, 1, 100) do (
    curl http://localhost:8000/%%i
)

for /l %%i in (1, 1, 100) do (
    curl http://localhost:8000/owners/
)

for /l %%i in (1, 1, 100) do (
    curl http://localhost:8000/estate/
)