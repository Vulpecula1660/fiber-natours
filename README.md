# fiber-natours

## 專案介紹
fiber-natours 是一個使用 [Go](https://golang.org) 語言及 [Fiber 框架](https://gofiber.io) 建置的後端 API 專案。專案提供會員註冊、登入、資訊查詢、修改及刪除等功能，並使用 PostgreSQL 作為資料庫、Redis 作為快取機制。

## 專案架構
- **api/**
  包含前台與後台 API 控制器，相關程式碼請參考：
  - 後台會員管理：[api/controller/back/member](api/controller/back/member)
  - 前台會員功能：[api/controller/front/member](api/controller/front/member)

- **model/**
  - **dao/**：資料庫存取層，如 [`model/dao/user/ins.go`](model/dao/user/ins.go) 新增用戶、[`model/dao/user/get.go`](model/dao/user/get.go) 取用戶資料等。
  - **dto/**：資料傳輸物件定義，請參照 [`model/dto/User.go`](model/dto/User.go)。
  - **postgresql/**：PostgreSQL 連線建立，參考 [`model/postgresql/create_conn.go`](model/postgresql/create_conn.go)。
  - **redis/**：Redis 連線及字串操作，參考 [`model/redis/create_conn.go`](model/redis/create_conn.go) 與 [`model/redis/string.go`](model/redis/string.go)。

- **service/**
  包含會員及 token 相關的商業邏輯，如：
  - 會員註冊與登入：[service/member/register.go](service/member/register.go)、[service/member/login.go](service/member/login.go)
  - token 產生：[service/token/create.go](service/token/create.go)

- **util/**
  密碼加密工具，請參考 [`util/password.go`](util/password.go)。

- **enum/**
  定義專案錯誤代碼與自訂錯誤處理，請參照 [`enum/error_code.go`](enum/error_code.go)。

## 環境變數設定
請確認設定下列環境變數：
- `DATABASE_Host`、`DATABASE_Port`、`DATABASE_User`、`DATABASE_Password`、`DATABASE_Name`：PostgreSQL 資料庫連線資訊。
- `REDIS_URL`：Redis 連線 URL。

## 安裝與執行
1. 下載專案：
    ```sh
    git clone https://github.com/Vulpecula1660/fiber-natours.git
    cd fiber-natours
    ```
2. 下載相依套件：
    ```sh
    go mod tidy
    ```
3. 設定環境變數，然後啟動專案：
    ```sh
    go run main.go
    ```

## API 路由
主要 API 路由定義於 [`api/router.go`](api/router.go)：
- **前台未登入 API：**
  - `POST /v1/front/member/register`：會員註冊 ([參見](api/controller/front/member/register.go))
  - `POST /v1/front/member/login`：會員登入 ([參見](api/controller/front/member/login.go))

（後台 API 可依需求擴充）

## 錯誤處理
錯誤代碼與訊息定義於 [`enum/error_code.go`](enum/error_code.go)，錯誤處理邏輯則由 [`api/middleware/error_hanlder.go`](api/middleware/error_hanlder.go) 完成。

## 貢獻
歡迎提交 PR 以及 Issues 討論功能改進與錯誤修正。

## 授權
請參考專案的 [LICENSE](LICENSE) 說明。
