# Map (ম্যাপ) - Go Language

**Map** হলো এমন একটি ডেটা স্ট্রাকচার যেখানে ডাটাগুলো **Key-Value** (কী-ভ্যালু) পেয়ার হিসেবে রাখা হয়। 
সহজ কথায়, Map-এ কোনো ইনডেক্স (0, 1, 2...) থাকে না। এর বদলে আপনি নিজের ইচ্ছেমতো একটি নাম (Key) দিয়ে কোনো ডাটা (Value) সেভ করে রাখতে পারেন। 

যেমন: ডিকশনারিতে আমরা কোনো শব্দের (Key) অর্থ (Value) খুঁজি।

### উদাহরণ (Example):

```go
package main

import "fmt"

func main() {
    // একটি Map তৈরি করা যেখানে Key হবে string এবং Value হবে int
    // উদাহরণ: মানুষের বয়স রাখা
    ages := make(map[string]int)

    // Map-এ ডাটা রাখা
    ages["Mamun"] = 25
    ages["Rahim"] = 30
    ages["Karim"] = 22

    fmt.Println("সবার বয়স:", ages)

    // Map থেকে নির্দিষ্ট কারো বয়স বের করা (Key ব্যবহার করে)
    fmt.Println("Mamun এর বয়স:", ages["Mamun"])

    // Map থেকে কোনো ডাটা মুছে ফেলা (delete)
    delete(ages, "Rahim")
    fmt.Println("Rahim কে বাদ দেওয়ার পর:", ages)
}
```

### কেন Map ব্যবহার করব?
যখন আপনার ডাটাগুলো কোনো নম্বর (Index) দিয়ে খোঁজার বদলে কোনো নির্দিষ্ট নাম বা আইডি (Key) দিয়ে খোঁজার দরকার হয়, তখন Map ব্যবহার করা হয়।
