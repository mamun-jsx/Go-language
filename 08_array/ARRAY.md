# Array (অ্যারে) - Go Language

**Array** হলো এমন একটি ডেটা স্ট্রাকচার যেখানে একই ধরনের অনেকগুলো ডেটা একসাথে রাখা যায়। 

অ্যারের একটি প্রধান বৈশিষ্ট্য হলো এর সাইজ ফিক্সড (Fixed)। অর্থাৎ, একবার একটি অ্যারে তৈরি করার সময় যদি বলি এতে ৫টি ডেটা থাকবে, তবে পরে এটি বাড়িয়ে ৬টি করা যাবে সহজভাবে যাবে না। (ডাইনামিক সাইজের জন্য আমরা Slice ব্যবহার করি)।

### Array ডিক্লেয়ার করা:
```go
package main

import "fmt"

func main() {
    // 5 সাইজের একটি Array তৈরি (সব মান প্রথমে 0 থাকবে)
    var marks [5]int

    // Array তে ডাটা রাখা (Index 0 থেকে শুরু হয়)
    marks[0] = 80
    marks[1] = 90
    marks[2] = 85
    
    // Array থেকে ডাটা দেখা
    fmt.Println("প্রথম মার্ক:", marks[0])
    fmt.Println("পুরো Array:", marks)
}
```

### একসাথে তৈরি এবং মান সেট করা:
```go
package main

import "fmt"

func main() {
    // এক লাইনেই Array তৈরি এবং মান দেওয়া
    names := [3]string{"Mamun", "Rahim", "Karim"}
    
    fmt.Println("নামগুলো:", names)
    fmt.Println("Array এর সাইজ (Length):", len(names))
}
```

### সাইজ না বলে Array তৈরি:
আপনি চাইলে `...` ব্যবহার করতে পারেন, এতে যতগুলো ভ্যালু দেবেন, Array ঠিক তত সাইজের হয়ে যাবে।
```go
package main

import "fmt"

func main() {
    colors := [...]string{"Red", "Green", "Blue", "Yellow"}
    fmt.Println("রংগুলো:", colors)
    fmt.Println("মোট রঙ আছে:", len(colors))
}
```
