
## Recommendation before use

# 🔥🔥 Go Version Tested 1.23.1 🔥🔥

## Features

|       Feature       | Supported |
| :-----------------: | :-------: |
|   Multithreading    |    ✅     |
|   Use Query Data    |    ✅     |
|   Auto Claim Task   |    ✅     |
|  Auto Daily Streak  |    ✅     |
| Auto Connect Wallet |    ✅     |

## [Settings](https://github.com/ehhramaaa/UnionExBot/blob/main/config.yml)

|       Settings       |                  Description                  |
| :------------------: | :-------------------------------------------: |
|     **API-URL**      |                 BASE API URL                  |
|    **REFER-URL**     |                 BASE BOT URL                  |
| **AUTO-BIND-WALLET** |    Auto Bind Wallet If Not Ready Connected    |
|    **MAX-THREAD**    |        Max Thread Worker Run Parallel         |
|   **RANDOM-SLEEP**   | Delay before the next lap (e.g. [1800, 3600]) |

## Prerequisites 📚

Before you begin, make sure you have the following installed:

- [Golang](https://go.dev/doc/install) **version > 1.22**


```shell
git clone https://github.com/ehhramaaa/UnionExBot.git
cd UnionExBot
go mod tidy
```

```shell
cp query.txt.example query.txt && cp wallet.txt.example wallet.txt
rm -rf query.txt.example wallet.txt.example
```

```shell
go build -o UnionExBot
chmod +x UnionExBot
./UnionExBot
```
