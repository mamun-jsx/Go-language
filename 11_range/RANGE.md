# Range with For Loop (রেঞ্জ এবং ফর লুপ) - Go Language

Go-তে **range** কি-ওয়ার্ডটি `for` লুপের সাথে ব্যবহার করা হয়। এটি দিয়ে আমরা খুব সহজেই Array, Slice, এবং Map-এর ভেতরের প্রতিটি ডাটা একে একে পড়তে পারি। 

### Slice এর সাথে range-এর ব্যবহার:
Slice-এ range ব্যবহার করলে আমরা দুইটি জিনিস পাই: `index` (কত নম্বর পজিশন) এবং `value` (পজিশনে থাকা ডাটা)।

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40}

    // range দিয়ে Slice এর ডাটা প্রিন্ট করা
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### Map এর সাথে range-এর ব্যবহার:
Map-এর তো কোনো index থাকে না। তাই Map-এ range ব্যবহার করলে আমরা পাই: `key` এবং `value`।

```go
package main

import "fmt"

func main() {
    ages := map[string]int{
        "Mamun": 25,
        "Rahim": 30,
    }

    // range দিয়ে Map এর ডাটা প্রিন্ট করা
    for key, value := range ages {
        fmt.Printf("নাম (Key): %s, বয়স (Value): %d\n", key, value)
    }
}
```

### শুধু ভ্যালু (Value) দরকার হলে:
মাঝে মাঝে আমাদের index বা key-র দরকার হয় না, শুধু value দরকার হয়। তখন আমরা আন্ডারস্কোর `_` (blank identifier) ব্যবহার করে index বা key-কে ইগনোর করতে পারি।

```go
package main

import "fmt"

func main() {
    names := []string{"A", "B", "C"}

    for _, name := range names {
        fmt.Println("নাম:", name) // এখানে index ইগনোর করা হয়েছে
    }
}
```
