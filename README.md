# GoBlogClean
This is Go Blog API server using clean architecture

## Directory structure
### cmd
- 全てのファイルのルートとなるパッケージ
### config
- databaseの設定などを記載するパッケージ
### domain
- クリーンアーキテクチャのドメインを定義するパッケージ
### pkg
- APIやHandler、repository, usecaseなどを記載する
- アプリケーション単位(auth, blogなど)でさらに分割している