## Let's start algorithm


This repository implements search algorithm written in [Let's start algorithm](https://book.impress.co.jp/books/3201) by Golang


### Install

1. Download this repository
2. Execute `go build` command as following

```
go build -o search search.go
```

3. Execute built commands

```
./search <Algorithm> <Searched Value> <Array(Comma Separated)>
```

### Implemented algorithm

#### Linear Search

##### Usage

```
./search linear 10 1,7,10,12,15,20
// expected results
// 3番目の文字と一致しました 
```

#### Binary Search

##### Usage

```
./search binary 4 1,2,4,8,16,32,64
// expected results
// 4番目の値を検索します
// 2番目の値を検索します
// 3番目の値を検索します
// 3番目の値と一致しました
```

Caution: Array must have been sorted

#### Hash Search

##### Usage

```
./search hash 10 12,14,9,1,7,10,5,2
// expected results
// 10番目の値は10です
// 10番目の値と一致しました
```