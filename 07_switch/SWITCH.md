# Switch Statement (সুইচ) - Go Language

`switch` অনেকটা `if-else` এর মতোই কাজ করে। যখন অনেকগুলো শর্ত (condition) চেক করার প্রয়োজন হয়, তখন অনেকগুলো `else if` লেখার চেয়ে `switch` ব্যবহার করা দেখতে সুন্দর এবং সহজ।

Go-তে `switch` এর একটি বিশেষ সুবিধা হলো, এর প্রতিটি কেসের (case) শেষে `break` লিখতে হয় না (অন্য অনেক ল্যাঙ্গুয়েজে লিখতে হয়)।

### উদাহরণ:
```go
package main

import "fmt"

func main() {
    day := "Friday"

    switch day {
    case "Friday":
        fmt.Println("আজ ছুটির দিন!")
    case "Saturday":
        fmt.Println("আজকেও ছুটির দিন!")
    case "Sunday":
        fmt.Println("কাজের দিন শুরু।")
    default:
        // উপরের কোনোটির সাথেই না মিললে এটি কাজ করবে
        fmt.Println("এটি সাধারণ একটি কাজের দিন।")
    }
}
```

### একসাথে একাধিক শর্ত:
আমরা চাইলে এক কেসেই কমা দিয়ে একাধিক ভ্যালু চেক করতে পারি।
```go
package main

import "fmt"

func main() {
    day := "Saturday"

    switch day {
    case "Friday", "Saturday":
        fmt.Println("উইকেন্ড (Weekend) বা ছুটির দিন!")
    default:
        fmt.Println("কাজের দিন!")
    }
}
```

### কন্ডিশন ছাড়া Switch (If-Else এর মতো):
`switch` এর পাশে কোনো ভ্যারিয়েবল না দিয়ে সরাসরি `case`-এ কন্ডিশন লেখা যায়।
```go
package main

import "fmt"

func main() {
    time := 10

    switch {
    case time < 12:
        fmt.Println("শুভ সকাল")
    case time < 18:
        fmt.Println("শুভ বিকাল")
    default:
        fmt.Println("শুভ রাত্রি")
    }
}
```
