> ### 创建tls证书

####  ```openssl genrsa -out -server.key 2048```

<br>

####  ```openssl req -new -x509	-key server.key -out server.crt -days 365```


<br>

####  ```openssl x509 -in server.crt -noout -text```

<br>

#### 将证书由.crt转为.pem
```openssl x509 -in mycert.crt -out mycert.pem -outform PEM```
==按照上面步骤生成的.crt文件转成.pem后，内容是没有区别的==
