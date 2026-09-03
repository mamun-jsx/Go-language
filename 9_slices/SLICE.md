# Slice (স্লাইস) - Go Language

Go-তে **Slice** হলো Array-এর মতো, কিন্তু এর সাইজ নির্দিষ্ট (fixed) থাকে না। আপনি চাইলে এতে নতুন ডাটা যোগ করতে পারবেন বা মুছে ফেলতে পারবেন। সহজ কথায়, Slice হলো ডাইনামিক Array।

### স্লাইস কিভাবে কাজ করে?
Array ডিক্লেয়ার করার সময় আমাদের সাইজ বলে দিতে হয় (যেমন: `[5]int`), কিন্তু Slice-এ কোনো সাইজ বলতে হয় না (যেমন: `[]int`)। 

### উদাহরণ (Example):

```go
package main

import "fmt"

func main() {
    // একটি খালি স্লাইস তৈরি করা
    var mySlice []int

    // স্লাইসে ডাটা যোগ করা (append)
    mySlice = append(mySlice, 10)
    mySlice = append(mySlice, 20)
    mySlice = append(mySlice, 30)

    fmt.Println("আমার স্লাইস:", mySlice)

    // স্লাইস তৈরি করার সময় ভ্যালু দিয়ে দেওয়া
    names := []string{"Mamun", "Hasan", "Rahim"}
    fmt.Println("নামের লিস্ট:", names)
    
    // স্লাইসের কোনো নির্দিষ্ট ডাটা দেখা (Index দিয়ে)
    fmt.Println("প্রথম নাম:", names[0])
}
```

### কেন স্লাইস ব্যবহার করব?
Array-র চাইতে Slice অনেক বেশি ব্যবহার করা হয় কারণ ডাটা কম বা বেশি হলে Slice নিজ থেকেই তার সাইজ ঠিক করে নেয়।
