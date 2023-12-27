# urlfusion
junta urls com parâmetros 


Instalação

go install github.com/erickfernandox/urlfusion@latest

Como usar

echo "https://example.com" | urlfusion -p "?id=XSS"

Resultado:

example.com?id=XSS

echo "https://example.com" | urlfusion -p "?id=XSS" -http

Resultado:

https://example.com?id=XSS
