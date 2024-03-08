# Endpoints

## Inscription

### 1. Utilisateur

- Requête:
  - Methode: POST
  - path: /signup/client
- body:
  ```json
  {
      "email": str,
      "password": str
  }
  ```
- Response:
  - Si cet email est deja inscrit: 409
  - Si tout va bien: 200. user peut par la suite s'authentifier normalement.

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
