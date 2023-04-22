# Sample Server
Sampleのサーバー実装です。
- Go
- OpenAPI
- ent
- dig


## パッケージ構成
パッケージ構成は以下のようになっています
- main
  - sampleのメインサーバー
- package
  - 共通コンポーネント
- Makefile
  - セットアップを行うためのmakeコマンドが書かれたファイル
- README.md
  - 説明文（今読んでいるファイル）

## セットアップ
#### 1.Goのインストール
homebrewでGoの最新バージョンをインストールしてください。Goは後方互換があるので最新バージョンをインストールしてもらって問題ありません。
```sh
$ brew install go
```

#### 2.Docker Desktop for Macのインストール
ローカルの開発はDockerを使います。
https://docs.docker.jp/docker-for-mac/install.html
からDocker Desctop for Macをインストールしてください。

#### 3. ローカルのコンテナを起動する
ローカルで開発に使うコンテナを起動します。
```
docker compose -f main/docker-compose.yml up -d
```
mysqlが立ち上がるのに少し時間がかかることがあります。

#### 4. Makeコマンドを叩く
ルートディレクトリ配下にあるmakeファイルを叩く。
```sh
$ make setup
```

## 設計思想
|                   |
|:-----------------:|
|    controller     |
|      usecase      |
|  domain/service   |
| domain/repository |

必ず下位層のみに依存するようにしてください。
例えばusecase層ではdomain/service層のメソッドしか呼び出せません。
上位のcontroller層や同位のusecase層、２つ下層のdomain/repository層のメソッドは呼び出すことはできません。
上記ルールを守らないと依存に循環ができ、複雑性が大幅に増してリファクタリングが難しくなることからこちらのルールを守ってください。

## ディレクトリ構成
```
main
├── cmd             // バッチなどのコマンド処理を実装。
│   └── migration   // マイグレーションを実装。
├── domain          // ドメイン層の実装。
│   ├── ent         // Entityの実装。Facebookのentを使用。
│   ├── repository  // リポジトリの実装。
│   └── service     // ドメインサービスの実装。リポジトリ層の前に置く層。
├── environment     // 環境変数により実装を出し分ける。
├── infra           // インフラ層。controllerから下層に向かって依存注入する。
│   ├── db          // DBに関する実装。
│   ├── di          // DIの実装。uberのdigを使用。
│   └── service     // インフラ層の機能をまとめて環境ごとに実装を出し分ける。
├── interface       // フロントとサーバーの境界の実装。
│   ├── converter   // リクエストやレスポンスとEntityのデータ構造を変換する。
│   └── openapi     // OpenAPIの実装。oapi-codegenを使って自動生成する。
├── server          // サーバー機能に関する実装。
│   ├── controller  // コントローラーの実装。
│   ├── errorcode   // エラーコードの実装。定義したcsvからコードを自動生成する。
│   └── middleware  // middlewareの実装。
├── usecase         // ユースケース層の実装
└── main.go         // エントリーポイント。実行すればサーバーが起動する。
```