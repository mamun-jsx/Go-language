# Simple Values (সাধারণ ভ্যালু বা মান) - Go Language

প্রোগ্রামিংয়ের ভাষায়, ডাটা বা তথ্যকে আমরা বিভিন্ন ভাগে ভাগ করি। Go-তে প্রধানত তিন ধরণের সাধারণ ভ্যালু (Simple Values) ব্যবহার করা হয়:

### ১. Numbers (সংখ্যা):
- **Integer (পূর্ণসংখ্যা)**: যেমন `10`, `-5`, `100`। এগুলোকে আমরা `int` বলি।
- **Float (দশমিক সংখ্যা)**: যেমন `10.5`, `3.1416`। এগুলোকে আমরা `float64` বলি।

### ২. String (টেক্সট বা লেখা):
যে কোনো শব্দ বা বাক্যকে String বলা হয়। String-কে সবসময় ডাবল কোটেশন (`" "`) এর ভেতরে রাখতে হয়।
- উদাহরণ: `"Hello Go"`, `"আমার নাম মামুন"`

### ৩. Boolean (সত্য/মিথ্যা):
এটি শুধু দুটি ভ্যালু নিতে পারে: `true` (সত্য) অথবা `false` (মিথ্যা)। সিদ্ধান্ত নেওয়ার কাজে এগুলো খুব দরকারি।

### উদাহরণ:
```go
package main

import "fmt"

func main() {
    // Strings
    fmt.Println("Go" + " " + "Language") // দুটি লেখা জোড়া লাগানো

    // Integers and Floats
    fmt.Println("1 + 1 =", 1+1)
    fmt.Println("7.0 / 3.0 =", 7.0/3.0)

    // Booleans
    fmt.Println("true AND false =", true && false) // AND লজিক
    fmt.Println("true OR false =", true || false)  // OR লজিক
    fmt.Println("NOT true =", !true)               // NOT লজিক
}
```
