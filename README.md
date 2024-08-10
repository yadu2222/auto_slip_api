

docker-compose ps
docker exec -it safety_c_sv-db-1 bash
mysql -u root -proot
go mod tidy

## api仕様書

### 雑誌

<details>
  <summary>雑誌一覧取得</summary>

- **URL:** `/v1/magazines/magazines`
- **メソッド:** GET
- **説明:** 説明
- **リクエスト:**
  - ヘッダー: application/json

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
      "srvResCode": 200,
      "srvResData": [
        {
          "magazineCode": "\ufeff00010",
          "magazineName": "WITH HARLEY(ヤングマシン増)",
          "takerUUID": "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
          "takerName": "やづ"
        },]
      }
      ```

</details>

<details>
  <summary>数取り</summary>

- **URL:** `/v1/csv/counting`
- **メソッド:** POST
- **説明:** CSVファイルを投げて数を取る
- **リクエスト:**
  - ヘッダー:
    ```
    Content-type : application/octet-stream
    ```
  - ボディ:
    ```
     {
        file : aaaaaa.csv
     }
    ```

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```
      {
        "srvResCode": 200,
        "srvResData":[
          {
            "Agency": {
              "countingUUId": "095cb246-c988-4038-bc16-08ae88fcbd5d",
              "magazineName": "SPA!",
              "magazineCode": "23451",
              "number": "07",
              "quantity": 0
            },
            "RegularAgencys":[
              {
                "regularUUID": "1386dea6-2c09-4679-b5bf-51744d0cc671",
                "customerUUID": "1386dea6-2c09-4679-b5bf-51744d0cc673",
                "customerName": "てすと書店",
                "quantity": 1,
                "methodType": "配達"
              }],
            "CountFlag": true
          },]
      }
      ```

</details>




### テンプレート
<details>
  <summary>タイトル</summary>

- **URL:** `url`
- **メソッド:** GET
- **説明:** 説明
- **リクエスト:**
  - ヘッダー:
  - ボディ:

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": "OK",
        "srvResData": {
          "message": "hello go server!"
        }
      }      
      ```

</details>

