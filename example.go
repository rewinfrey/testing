package main

func main() {
  type T1 string
  type T2 T1
  type T3 []T1

  type s3 struct {
    f, g int
  }
}
