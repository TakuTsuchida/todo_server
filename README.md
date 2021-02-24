# 概要

golang echoにおけるクリーンアーキテクチャーを意識したテスト実装

# 技術スタック

go version: 1.15.7
packageに関しては、go.modを参照すること

- github.com/labstack/echo
- gorm.io/gorm

# ディレクトリ構造

基本ルールとしては、以下の円の外から内への依存関係とする。

![Clean Architecture](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.amazonaws.com%2F0%2F44142%2Fa7643c53-8cc0-b079-0745-a20f06f23372.jpeg?ixlib=rb-1.2.2&auto=format&gif-q=60&q=75&w=1400&fit=max&s=cedd83cad1f4ce805e663ffd96ec335c)

## 全体

``` text
├── domains
│   └── todo.go
├── go.mod
├── go.sum
├── infrastructures
│   ├── router.go
│   └── sql_handler.go
├── interfaces
│   ├── controllers
│   │   └── todo_controller.go
│   └── databases
│       ├── sql_handler.go
│       └── todo_repository.go
├── server.go
└── usecases
    ├── interactors
    │   └── todo_interactor.go
    └── repositories
        └── todo_repository.go
```

## 詳細説明

### domains

ビジネスルールの為のデータ構造、もしくはメソッドを持ったオブジェクト。
ビジネスドメインモデルを定義する場所としている。

### usecases

アプリケーション固有のビジネスルールを定義する層である。
ドメインとのデータの流れを組み立てる。

- interactors
    `interfaces/controllers`へのGatewayの役割を定義する

- repositories
    `interfaces/databases`からのInput Portの役割を定義する

### interfaces

外部から、ユースケースとドメインで使われる内部形式にデータを変換、または、内部から外部の機能に最も便利な形式に、データを変換するアダプター的役割を果たす層である。

- controllers
    APIの内容を表すコントローラーを定義する。
    Middleware, CustomContextはここで定義する

- databases
    具体的なデータをやり取りする処理を定義する

### infrastructures

フレームワークやツールを定義する層である。
デバイスやフレームワークを介した通信を定義するコードを記載する。

- router.go
    echoを介したroutingを定義する

- sql_handler.go
    mysqlへのクエリーの発行等の接続を定義する

# 与件

- mysql configuration fileの定義、共通化
