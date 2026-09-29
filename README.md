-- バックエンド全体設計
～フロントはしらん
レイヤードでいく

discordのボタン駆動でfetch＆get

別パッケージでFETCH
     googleNewsでrss2.0かAtom
     取得したらhandler層に渡す

--------------------------------------

HANDLER層
    httpメソッドなど
    ここはdiscord仕様に

MODEL層
    fetchしたやつの型とか整える？
    getでも型整えたり
    ここは軽く

DB層
sqliteで
    保存（fetchのやつ）
    dbから取得
    削除

    fetch & CREATE
        goのライブラリ調べる、コピペ
        dbに登録

        ※関数は分けよう


    GET
        書く

    DELETE
        自動で48時間より前削除