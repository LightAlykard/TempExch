# аутентификация по логину и паролю в JSON, ответ: JWT в cookies

rm -rf /tmp/cookie.txt && \
curl \
    --verbose \
    --request POST \
    --header "Content-Type: application/json" \
    --data '{"username":"user1","password":"123"}' \
    --cookie-jar /tmp/cookie.txt \
    http://localhost:8080/login
