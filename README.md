

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

### 顧客
<details>
  <summary>顧客一覧取得</summary>

- **URL:** `/v1/customers/customers`
- **メソッド:** GET
- **説明:** 顧客情報を一覧取得
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
            "customerUUId": "0038eae7-56ee-4ff3-bba5-380de72fb3ba",
            "customerName": "室谷",
            "methodType": 1,
            "tellAddress": "54-0854",
            "tellType": 1,
            "note": "",
            "csvId": 162
          },]
      }
      ```

</details>

### 定期
<details>
  <summary>雑誌を主キーに定期一覧取得</summary>

- **URL:** `/v1/regulars/regulars/magazines`
- **メソッド:** GET
- **説明:** 雑誌情報を主キーに定期情報を一覧取得
- **リクエスト:**
  - ヘッダー: application/json

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": 200,
        "srvResData":[
          {
            "magazine": {
              "magazineCode": "\ufeff00010",
              "magazineName": "WITH HARLEY(ヤングマシン増)",
              "takerUUID": "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
              "takerName": "やづ"
            },
            "regulars": null
          },
          {
            "magazine": {
              "magazineCode": "01017",
              "magazineName": "てれびくん",
              "takerUUID": "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
              "takerName": "やづ"
            },
            "regulars":[
              {
                "regularUUID": "2243990e-0a0e-4eab-be6b-8bb5a6a8aca5",
                "quantity": 1,
                "Customer": {
                  "customerUUId": "0327c618-1402-43e5-ba9a-baf8dee4ecec",
                  "customerName": "やづこ",
                  "methodType": 1,
                  "tellAddress": "",
                  "tellType": 0,
                  "note": "「てれびくん」必要な時だけ連絡予定。ない場合は不要。",
                  "csvId": 204
                }
              },]
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

