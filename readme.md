Glossika Backend Interview

How to use:

若尚未安裝 docker-compose, 請先安裝

安裝完成後，進入專案，執行: docker-compose up -d
主程式會聽取 port 8080

APIs:
1. POST /user/create -> 建立用戶
body: 
{
    "email": "string",
    "password": "string"
}

2. POST /user/login -> 登入
body:
{
    "email": "string",
    "password": "string"
}
response: jwt token string -> 在 get recommendations 需要使用到

3. POST /emailVerify/send -> 寄出驗證 email
body: 
{
    "userId": "string" -> find in database
}
response: verify code -> 由於並未真的寄出郵件，所以 verify code 會在這裡回傳

4. POST /emailVerify/verify -> 驗證 email
body:
{
    "code": "string" -> 從 API3 取得
}

5. GET /recommendation/all -> 取得 recommendations
Header: UserToken: "jwt token string" -> 從 API2 的 response 中取得


為了方便驗證服務的正確性，我 expose mysql 以及 redis 的 port，所以啟動後可以直接用 localhost 連接查看
連線資訊：
Mysql:
    username: user
    password: password
    port: 3306
Redis:
    port: 6379