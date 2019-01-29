> ### 创建tls证书

####  ```openssl genrsa -out -server.key 2048```

<br>

####  ```openssl req -new -x509	-key server.key -out server.crt -days 365```


<br>

####  ```openssl x509 -in server.crt -noout -text```
