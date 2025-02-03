### Ejecución price-api-meli


    
1. ejecutar el comando para instalar dependencias
    
    ```bash
    go mod tidy 
    ```
    
2. crear un archivo .env en la raiz del proyecto y pegar el siguiente codigo
    
    ```bash
    DB_HOST=localhost
    DB_USER=postgresql
    DB_PASS=root
    DB_NAME=products_db
    DB_PORT=5432
    KAFKA_BROKERS=localhost:9092
    KAFKA_TOPIC=price
    ```
    
3. ahora solo falta ejecutar nuestra api
    
    ```bash
    go run cmd/main
    ```
