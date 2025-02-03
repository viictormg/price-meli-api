### Ejecucion price-api-meli

para la api “price-api-meli” los pasos son mas pocos, que la mayoria de dependencias están definidas en el docker-compose de product

1. entrar la carpeta donde está el proyecto
    
    ```bash
    cd price-api-meli
    ```
    
2. despues de estar en el directorio del repositorio ejecutamos el comando para instalar las dependencias 
    
    ```bash
    go mod tidy 
    ```
    
3. tambien debes crear un archivo .env en la raiz del proyecto y pegar el siguiente codigo
    
    ```bash
    DB_HOST=localhost
    DB_USER=postgresql
    DB_PASS=root
    DB_NAME=products_db
    DB_PORT=5432
    KAFKA_BROKERS=localhost:9092
    KAFKA_TOPIC=price
    ```
    
4. ahora solo falta ejecutar nuestra api
    
    ```bash
    go run cmd/main
    ```
