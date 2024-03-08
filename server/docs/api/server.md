# Endpoints

## Authentification

### 1. Authentification d'un utilisateur

- Requête:
  - Methode: POST
  - path: /auth/client
- body:
  ```json
  {
      "email": str,
      "password": str
  }
  ```
- reponse:
  - Si authentification réussie
    - status: OK (200)
    - body:
      ```json
      {
          "token": str
      }
      ```
  - Sinon: code 401|400, reponse vide

### 2. Authentification d'un restaurant

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
- reponse:
  - Si authentification réussie
    - status: OK (200)
    - body:
      ```json
      {
          "token": str
      }
      ```
  - Sinon: code 401|400, reponse vide

## Order

### Gestion du cycle de vie d'une commande

#### 1. Create

#### 2. Pending to accepted

#### 3. Pending to declined

#### 4. Accepted to In progress

#### 5. In progress to Completed

###
