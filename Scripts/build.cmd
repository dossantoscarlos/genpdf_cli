@echo off 

echo Iniciando build de projeto...

go build cmd/main.go

if %errorlevel% neq 0 (
    echo Erro ao compilar o projeto.
    exit /b %errorlevel%
)

echo Build concluído com sucesso.
echo renomeando arquivo de saída...
rename main.exe cli_op_pdf.exe
if %errorlevel% neq 0 (
    echo Erro ao renomear o arquivo de saída.
    exit /b %errorlevel%
)

echo Arquivo de saída renomeado com sucesso.

echo movendo arquivo de saída para pasta bin...

move cli_op_pdf.exe build\cli_op_pdf.exe
if %errorlevel% neq 0 (
    echo Erro ao mover o arquivo de saída.
    exit /b %errorlevel%
)

echo Arquivo de saída movido com sucesso.

echo finalizando build de projeto...
