# If-Else (শর্ত বা Condition) - Go Language

প্রোগ্রামিংয়ে সিদ্ধান্ত নেওয়ার জন্য `if-else` ব্যবহার করা হয়। "যদি এটা হয়, তবে এটা করো, নাহলে অন্যটা করো" - এটাই হলো `if-else` এর মূল কথা।

### সাধারণ If-Else:
```go
package main

import "fmt"

func main() {
    age := 18

    // যদি বয়স 18 বা তার বেশি হয়
    if age >= 18 {
        fmt.Println("আপনি ভোটার হতে পারবেন।")
    } else {
        // যদি উপরের শর্ত সত্যি না হয়
        fmt.Println("আপনি এখনও ভোটার হতে পারবেন না।")
    }
}
```

### If-Else If (একাধিক শর্ত):
যখন একাধিক শর্ত চেক করতে হয়, তখন `else if` ব্যবহার করা হয়।
```go
package main

import "fmt"

func main() {
    marks := 85

    if marks >= 80 {
        fmt.Println("আপনি A+ পেয়েছেন।")
    } else if marks >= 70 {
        fmt.Println("আপনি A পেয়েছেন।")
    } else {
        fmt.Println("আপনাকে আরও পড়তে হবে।")
    }
}
```

### If-এর ভেতরে Variable ডিক্লেয়ার করা:
Go-তে `if` কন্ডিশনের ভেতরেই ছোট করে ভ্যারিয়েবল ডিক্লেয়ার করা যায়। এটি বেশ কাজের!
```go
package main

import "fmt"

func main() {
    // num := 10 এখানে ডিক্লেয়ার করে সাথে সাথেই চেক করা হচ্ছে
    if num := 10; num%2 == 0 {
        fmt.Println(num, "হলো জোড় সংখ্যা")
    } else {
        fmt.Println(num, "হলো বিজোড় সংখ্যা")
    }
}
```
