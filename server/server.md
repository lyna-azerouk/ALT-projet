# Endpoints
## Restaurants
### Authentification
- Requête:
  - Methode: POST
  - path: /auth/restaurant
- body: 
  ```json
  {
      "id": int,
      "password": str
  }
  ```
- response: 
    ```json
    {
        "token": str
    }
    ```
## Order
### Life cycle
#### 1. Create
#### 2. Pending to accepted
#### 3. Pending to declined
#### 4. Accepted to In progress
#### 5. In progress to Completed

### 