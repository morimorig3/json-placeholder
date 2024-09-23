# json-placeholder

echoを使用して[json-placeholder](https://jsonplaceholder.typicode.com/)的なものを実装してみる練習

# Usage

※デプロイ先はRender無料枠のためインスタンス起床までに30秒必要

## タスク1件

```
https://json-placeholder-8s87.onrender.com/todo/1
```

```
[
  {
    "id": 0,
    "task": "タスク No.0",
    "isCompleted": false
  }
]
```

## タスク5件

```
https://json-placeholder-8s87.onrender.com/todo/5
```

```
[
  {
    "id": 0,
    "task": "タスク No.0",
    "isCompleted": false
  },
  {
    "id": 1,
    "task": "タスク No.1",
    "isCompleted": false
  },
  {
    "id": 2,
    "task": "タスク No.2",
    "isCompleted": false
  },
  {
    "id": 3,
    "task": "タスク No.3",
    "isCompleted": false
  },
  {
    "id": 4,
    "task": "タスク No.4",
    "isCompleted": false
  }
]
```

# Update

## 2024-09-23

Goらしい書きっぷりに四苦八苦。ある処理を単体の関数として実装するのか構造体に関数を生やして実装するべきなのかの判断ができていない。
在野のGoのコードを見ていると構造体から生やしているものが多く。ファイル構成も`構造体A.go`で、ロジックの処理も1つのファイルにまとめているように思える。ここら辺はGoの達人コードを見て学んでいくしかない。