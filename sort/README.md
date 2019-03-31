## Let's start algorithm


This repository implements sort algorithm written in [Let's start algorithm](https://book.impress.co.jp/books/3201) by Golang


### Install

1. Download this repository
2. Execute `go build` command as following

```
go build -o sort sort.go
```

3. Execute built commands

```
./sort <Algorithm> <Array(Comma Separated)>
```

### Implemented algorithm

#### Selection Sort

##### Usage

```
./sort selection 10,2,7,15,8
// expected results
//1番目までソート済み：[]string{"2", "10", "7", "15", "8"}
//2番目までソート済み：[]string{"2", "7", "10", "15", "8"}
//3番目までソート済み：[]string{"2", "7", "8", "15", "10"}
//4番目までソート済み：[]string{"2", "7", "8", "10", "15"}
//[2 7 8 10 15]
```

#### Bubble Sort

##### Usage

```
./sort bubble 10,2,7,15,8
// expected results
// 1番目までソート済み：[]string{"2", "10", "7", "8", "15"}
// 2番目までソート済み：[]string{"2", "7", "10", "8", "15"}
// 3番目までソート済み：[]string{"2", "7", "8", "10", "15"}
// 4番目までソート済み：[]string{"2", "7", "8", "10", "15"}
// 5番目までソート済み：[]string{"2", "7", "8", "10", "15"}
// [2 7 8 10 15]
```


#### Insert Sort

##### Usage

```
./sort insert 10,2,7,15,8
// expected results
//2番目までソート済み：[]string{"2", "10", "7", "15", "8"}
//3番目までソート済み：[]string{"2", "7", "10", "15", "8"}
//4番目までソート済み：[]string{"2", "7", "10", "15", "8"}
//5番目までソート済み：[]string{"2", "7", "8", "10", "15"}
//[2 7 8 10 15]
```

#### Quick Sort

##### Usage

```
./sort quick 10,2,7,15,8
// expected results
//[8 2 7 10 15]
//left is 0, right is 2
//left side sort
//[7 2 8 10 15]
//left is 0, right is 1
//left side sort
//[2 7 8 10 15]
```
