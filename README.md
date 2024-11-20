## QUICK START

1. Install Docker and Docker Compose
2. Clone this repository
3. Diagram:
   
   ![](images/diagram.png)
5. Run Flask app:

    ```bash
    make run-python
    ```
   ![](images/python.png)
6. Run Golang app:

    ```bash
    make run-go
    ```
   ![](images/go.png)
7. Run Flask and Golang app with Docker Compose:

    ```bash
    make docker-run
    ```
8. Run Kong Gateway:

    ```bash
    make run-kong
    ```
   ![](images/kong.png)
9. Http request:

    - Test Golang app:
        ```bash
        curl -X GET http://localhost:8000/
        ```
      ![](images/test-go.png)
    - Test Login:
        ```bash
        curl -X POST http://localhost:8000/login -H "Content-Type: application/json" -d '{"username": "ngdangkietswe"}'
        ```
      ![](images/test-login.png)
    - Test Flask app with Token:
        ```bash
        curl -X GET http://localhost:8000/api/health -H "Authorization: Bearer {token}" 
        ```
      ![](images/test-python-health.png)
        ```bash
        curl -X POST http://localhost:8000/api/check-authorize -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"username": "ngdangkietswe"}' 
        ```
      ![](images/test-python-authorize.png)
<<<<<<< Updated upstream
=======
9. Documentation:
    - https://docs.konghq.com/
>>>>>>> Stashed changes
