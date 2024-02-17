# メモ

## ディレクトリと説明

### custom_backend
カスタムバッグエンドのGoのサーバー

### src
webpackがbuildするソースコード

### dist
webpackがbuildした成果物  
公開用ディレクトリ

## コマンドと説明

### Firebase
```
npm run build && firebase serve
```

### カスタムバックエンド
```
ilts@irumarudev:~/custom_backend/develop$ docker compose up
```
```
godev go run main.go
```

### 

## Hostingの構築
[プロジェクトの作成](https://firebase.google.com/docs/web/setup?hl=ja&_gl=1*qel1us*_up*MQ..*_ga*OTIwMTgyNTExLjE3MDgxODE3MzA.*_ga_CW55HF8NVT*MTcwODE4MTczMC4xLjAuMTcwODE4MTczMC4wLjAuMA..)

[webpackの構築](https://firebase.google.com/docs/web/module-bundling?hl=ja&_gl=1*qel1us*_up*MQ..*_ga*OTIwMTgyNTExLjE3MDgxODE3MzA.*_ga_CW55HF8NVT*MTcwODE4MTczMC4xLjAuMTcwODE4MTczMC4wLjAuMA..)

## カスタムバッグエンド

### カスタムバッグエンド側の認証(Go AdminSDKの初期化時に使用)

[カスタムバッグエンドDocs](https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja)

[プロジェクト設定&サンプルコードURL](https://console.firebase.google.com/project/intro-8f3f7/settings/serviceaccounts/adminsdk?hl=ja)

プロジェクト設定/サービスアカウント/Firebase Admin SDK/Go  

以下を使用  
- コード
- 新しい秘密鍵を生成
