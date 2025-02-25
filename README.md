# 🚀 Meme Coin API

## 📌 專案簡介
這是一個用於管理 Meme Coin 的 API 服務，使用 **Golang + Gin**，並整合 **Swaggo** 來生成 API 文件。專案支援本機開發與 Docker 容器化部署。

---

## 🛠️ 安裝與運行方式

### 本機運行
1. **安裝 Go 1.22.5**
2. Clone:
    ```sh
    git clone https://github.com/jacobintern/MemeCoin.git
    cd your-repo-name
    ```
3. 安裝依賴： `go mod tidy`
4. 產生環境變數檔案： `make gen-env`
5. 產生 or 更新 swagger 文件： `make swag`
6. 啟動應用程式：
    - go run：
        ```sh
        go run cmd/api/main.go
        ```
    - vscode debug mode：
        - 建置 launch.json
        - 指定 `program` 啟動位置 `"${workspaceFolder}/cmd/api/main.go"`
        - 指定 `cwd` 目前根目錄 `"${workspaceFolder}"`
        - 指定環境變數檔案位置 `"${workspaceFolder}/config/.env"`

### Docker 運行
1. 安裝 Docker 與 Docker Compose
2. 使用 Docker Compose 啟動：
    ```sh
    # 啟動
    docker compose up -d

    # 關閉
    docker compose down
    # 刪除相關 image
    docker compose down --rmi all
    # 刪除資料 volume
    docker compose down -v
    ```
3. 確保服務已運行：
    ```sh
    docker ps
    ```

---

## 📂 專案結構
```bash
meme-coin-api/
│── cmd/                # 應用程式入口
│   ├── api/            # API 進入點
│     ├── main.go       # 主程式
│   ├── injection/      # 依賴註冊
│── controller/         # 業務邏輯（處理 HTTP 請求）
│── service/            # 商業邏輯（應用核心邏輯）
│── repository/         # 資料庫存取層
│── config/             # 環境變數與設定
│── docs/               # Swaggo 自動生成的 Swagger 文件
│── scripts/            # mysql table schema
│── pkg/                # 共用工具
│   ├── conf/           # 讀取環境變數的工具
│   ├── errors/         # 自定義錯誤格式
│   ├── mysqlx/         # 自定義 MySQL client
│   ├── util/           # 擴充工具函式
│── dockerfile          # Docker 映像檔設定
│── docker-compose.yml  # Docker Compose 設定
```

- 以上是簡化 DDD 架構來簡單處理 meme coin API 的程式架構

---

## ⚙️ 設定與配置
- 環境變數設定
    應用程式使用 `.env` 來管理環境變數，請使用 `make swag` 並修改內容：
    ```sh
    make swag
    ```

    .env 內容範例：
    ```
    # server setting
    SERVER_PORT=8080

    # log level
    LOG_LEVEL=ERROR

    HTTP_TIMEOUTSEC=10

    DATABASE_MYSQL_DATABASE=mydatabase
    DATABASE_MYSQL_USER=user
    DATABASE_MYSQL_PASSWORD=passw0rd
    DATABASE_MYSQL_HOST=localhost
    DATABASE_MYSQL_PORT=3306

    CORS_ALLOW_ALLORIGINS=true
    CORS_ALLOW_CREDENTIALS=false
    CORS_ALLOW_METHODS=GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS
    CORS_ALLOW_HEADERS=Origin,Content-Length,Content-Type,  Authorization
    CORS_MAXAGE=43200
    ```

---

## 📜 API 文件
此專案使用 Swaggo 來產生 Swagger API 文件，啟動應用後可透過以下方式存取：

- Swagger UI：http://localhost:8080/swagger/index.html