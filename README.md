# Price Checker CLI

Este programa en Go permite obtener el precio actual de una acción usando la API de Yahoo Finance.

## Requisitos

- Go 1.23 o superior

## Instalación

1. Clona este repositorio:

    ```sh
    git clone https://github.com/francoJalil/cli-stocks-price-viewer
    ```

2. Navega al directorio del proyecto:

    ```sh
    
    cd cli-stocks-price-viewer
    ```

3. Compila el programa:

    ```sh

    go build -o price ./cmd/main.go
    ```

   Esto generará un ejecutable llamado `price`.

## Uso

Para ejecutar el programa, usa el siguiente comando:

```sh
price <TICKER> [x]
```

# Ejemplos:

```sh
$ price AAPL
>> AAPL ARS $1000.00 -2%
```

```sh
$ price AAPL x
>> AAPL USD $100.00 -2%
```

El programa obtendrá el precio actual de la acción AAPL y lo mostrará en la consola.
Si se agrega el parámetro `x`, el programa mostrará el precio en USD.

