# GoBlogClean
This is Go Blog API server using clean architecture

## Directory structure
### cmd
- 全てのファイルのルートとなるパッケージ
### common
- どのパッケージからも共通で使われる処理を書くパッケージ
### domain
- クリーンアーキテクチャのドメインを定義するパッケージ
### injector
- 後日記載予定
### pkg
- APIやHandler、repository, usecaseなどを記載する
- アプリケーション単位(auth, blogなど)でさらに分割している