# Performance Tests

1. Install K6
   ```bash
   choco install k6
   ```
2. Install postman-to-k6
   ```bash
   npm install -g @apideck/postman-to-k6
   ```
3. Execute tests:
   ```bash
   k6 run ./k6-script.js
   ```