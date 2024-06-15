

docker-compose ps
docker exec -it safety_c_sv-db-1 bash
mysql -u root -proot
go mod tidy

## api仕様書

### 雑誌
- **URL:** `/api/v1/auth/users/list`
- **メソッド:** `GET`
- **説明:** パラメーターで"id"を指定、そのユーザが所属する会社のユーザ一覧を返す。 
- **リクエスト:**
  - パラメーター:
    - `id`: (int)ID。トークンと合わせて本人のものか確認、所属している会社を特定する。
  - ヘッダー:
    - `Authorization`: (string) 認証トークン

#### Response
```json
[
    {
        "id": 1,
        "name": "Item 1",
        "description": "Description of Item 1"
    },
    {
        "id": 2,
        "name": "Item 2",
        "description": "Description of Item 2"
    }
]

