## QUICK START

1. Install Docker and Docker Compose
2. Clone this repository
3. Run Flask app:

    ```bash
    make run-python
    ```
   ![](images/python.png)
4. Run Golang app:

    ```bash
    make run-go
    ```
   ![](images/go.png)
5. Run Flask and Golang app with Docker Compose:

    ```bash
    make docker-run
    ```
   ![](images/compose.png)
6. Run Kong Gateway:

    ```bash
    make run-kong
    ```
   ![](images/kong.png)
7. Http request:

    - Test Golang app:
        ```bash
        curl -X GET http://localhost:8005/
        ```
      ![](images/test-go.png)
    - Test Login:
        ```bash
        curl -X POST http://localhost:8005/login -H "Content-Type: application/json" -d '{"username": "ngdangkietswe"}'
        ```
      ![](images/test-login.png)
    - Test Flask app with Token:
        ```bash
        curl -X GET http://localhost:8005/api/health -H "Authorization: Bearer {token}" 
        ```
      ![](images/test-python-health.png)
        ```bash
        curl -X POST http://localhost:8005/api/check-authorize -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"username": "ngdangkietswe"}' 
        ```
      ![](images/test-python-authorize.png)