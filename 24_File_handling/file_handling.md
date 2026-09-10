# Golang-এ File Handling (ফাইল নিয়ে কাজ করা)

যেকোনো প্রোগ্রামিং ল্যাঙ্গুয়েজেই ফাইল নিয়ে কাজ করা খুবই গুরুত্বপূর্ণ একটি বিষয়। Go ল্যাঙ্গুয়েজে `os` এবং `bufio` প্যাকেজ ব্যবহার করে ফাইল এবং ফোল্ডার (ডিরেক্টরি) নিয়ে খুব সহজেই কাজ করা যায়।

আপনার `files.go` কোডের ওপর ভিত্তি করে ফাইল নিয়ে কাজ করার বিভিন্ন ধাপ নিচে ধাপে ধাপে বাংলায় ব্যাখ্যা করা হলো:

---

## ১. ফাইলের তথ্য (File Info) জানা
`os.Open` ব্যবহার করে আপনি একটি ফাইল খুলতে পারেন এবং `f.Stat()` ফাংশন ব্যবহার করে সেই ফাইলের সাইজ, পারমিশন ইত্যাদি তথ্য জানতে পারেন।

**কোড উদাহরণ:**
```go
f, err := os.Open("example.txt")
if err != nil {
    panic(err) // কোনো এরর হলে প্রোগ্রাম বন্ধ হয়ে যাবে
}
defer f.Close() // কাজ শেষে ফাইলটি বন্ধ করা নিশ্চিত করা হচ্ছে

fileInfo, err := f.Stat()
if err != nil {
    panic(err)
}

fmt.Println("Name:", fileInfo.Name())         // ফাইলের নাম
fmt.Println("Size:", fileInfo.Size())         // ফাইলের সাইজ (বাইট এ)
fmt.Println("Modified:", fileInfo.ModTime())  // সর্বশেষ কখন এডিট করা হয়েছে
fmt.Println("Is Directory:", fileInfo.IsDir()) // এটি কি কোনো ফোল্ডার নাকি ফাইল? (ট্রু/ফলস)
fmt.Println("Permissions:", fileInfo.Mode().Perm()) // ফাইলের পারমিশন (যেমন: read, write)
```

---

## ২. বাফার (Buffer) ব্যবহার করে ফাইল পড়া
বড় সাইজের ফাইলের ক্ষেত্রে একসাথে পুরো ফাইল না পড়ে, অল্প অল্প করে (chunk আকারে) বাফারের সাহায্যে পড়া ভালো। এতে র‍্যামের (RAM) ওপর অতিরিক্ত চাপ পড়ে না।

**কোড উদাহরণ:**
```go
f, err := os.Open("example.txt")
if err != nil {
    panic(err)
}
defer f.Close() // কাজ শেষে ফাইল বন্ধ করা

// ১২ বাইটের একটি টেম্পোরারি বাফার (মেমোরি) তৈরি করা হলো
buff := make([]byte, 12)

// ফাইলের ডেটা পড়ে বাফারে রাখা হচ্ছে
d, err := f.Read(buff)
if err != nil {
    panic(err)
}

fmt.Println("Bytes read:", d)       // কত বাইট পড়া হয়েছে
fmt.Println("Content:", string(buff)) // বাফারের ডেটা স্ট্রিং এ কনভার্ট করে প্রিন্ট করা হচ্ছে
```

---

## ৩. ফাইল পড়ার সহজ উপায় (Simple Way)
ফাইলটি যদি ছোট হয়, তবে পুরো ফাইলটি একবারে মেমোরিতে লোড করে পড়া যায় `os.ReadFile` ব্যবহার করে। ছোট টেক্সট বা কনফিগারেশন ফাইল পড়ার জন্য এটি সবচেয়ে সহজ পদ্ধতি।

**কোড উদাহরণ:**
```go
data, err := os.ReadFile("example.txt")
if err != nil {
    panic(err)
}

// পুরো ডেটা স্ট্রিং হিসেবে প্রিন্ট করা হচ্ছে
fmt.Println(string(data))
```
*সতর্কতা: অনেক বড় ফাইলের (যেমন: গিগাবাইট সাইজের ভিডিও) ক্ষেত্রে এই পদ্ধতি ব্যবহার করবেন না, কারণ এটি পুরো র‍্যাম ফুল করে ফেলবে।*

---

## ৪. ডিরেক্টরি বা ফোল্ডার (Folders) পড়া
ফাইলের মতোই একটি ফোল্ডার ওপেন করে `ReadDir` ফাংশনের সাহায্যে সেই ফোল্ডারের ভেতরের সব ফাইল ও ফোল্ডারের তালিকা বের করা যায়।

**কোড উদাহরণ:**
```go
dir, err := os.Open("../") // ফোল্ডারটি ওপেন করা হলো
if err != nil {
    panic(err)
}
defer dir.Close()

// ফোল্ডারের ভেতরের সব ফাইল পড়া হচ্ছে (-1 মানে সব ডেটা)
fileInfos, err := dir.ReadDir(-1)
for _, fi := range fileInfos {
    // প্রতিটা ফাইলের নাম এবং সেটি ফোল্ডার কিনা তা প্রিন্ট করা হচ্ছে
    fmt.Println(fi.Name(), "Is Directory?", fi.IsDir())
}
```

---

## ৫. নতুন ফাইল তৈরি করা এবং তাতে লেখা
`os.Create` ব্যবহার করে নতুন ফাইল তৈরি করা যায়। আর সেটিতে স্ট্রিং বা বাইট ডেটা লেখা যায়।

**কোড উদাহরণ:**
```go
file, err := os.Create("example2.txt") // নতুন ফাইল তৈরি
if err != nil {
    panic(err)
}
defer file.Close()

// সরাসরি স্ট্রিং লেখা
file.WriteString("hi golang")

// বাইট আকারে ডেটা লেখা
byteData := []byte(" hello bro learn go lang")
file.Write(byteData)
```

---

## ৬. এক ফাইল থেকে অন্য ফাইলে ডেটা কপি করা (Copy Data)
বড় ফাইলে ডেটা লস ছাড়া কপি করার জন্য `bufio` প্যাকেজ ব্যবহার করে রিডার (Reader) এবং রাইটার (Writer) দিয়ে বাইট-বাই-বাইট কপি করা সবচেয়ে নিরাপদ উপায়।

**কোড উদাহরণ:**
```go
// ১. সোর্স (মূল) ফাইলটি ওপেন করা হলো
sourceFile, err := os.Open("example.txt")
if err != nil {
    panic(err)
}
defer sourceFile.Close()

// ২. নতুন ফাইল (যেখানে ডেটা কপি হবে) তৈরি করা হলো
destFile, err := os.Create("example3.txt")
if err != nil {
    panic(err)
}
defer destFile.Close()

// ৩. বাফার রিডার এবং রাইটার তৈরি করা হলো
reader := bufio.NewReader(sourceFile)
writer := bufio.NewWriter(destFile)

// ৪. ইনফিনিট লুপের মাধ্যমে এক বাইট করে পড়া এবং লেখা হচ্ছে
for {
    b, err := reader.ReadByte() // এক বাইট পড়া হলো
    if err != nil {
        if err.Error() != "EOF" { // EOF মানে End Of File (ফাইলের শেষ)
            panic(err)
        }
        break // ফাইলের শেষে পৌঁছালে লুপ ব্রেক করা হবে
    }

    err = writer.WriteByte(b) // নতুন ফাইলে এক বাইট লেখা হলো
    if err != nil {
        panic(err)
    }
}

// ৫. Flush করা হচ্ছে যাতে রাইটারে থাকা সব ডেটা নিশ্চিতভাবে হার্ডডিস্কে সেভ হয়
writer.Flush()
fmt.Println("return to new file successfully") // সাকসেস মেসেজ
```

### আউটপুট এবং কমান্ড লাইন (Command Line)
ফাইলটি রান করার জন্য টার্মিনালে নিচের কমান্ডটি লিখুন:
```bash
go run files.go
```
আপনার আনকমেন্ট করা কপি করার কোড ব্লক অনুযায়ী, আউটপুট আসবে:
```text
retuen to new file successfully
```
একই সাথে ফোল্ডারে `example3.txt` নামের নতুন একটি ফাইল তৈরি হবে, যার ভেতরে `example.txt` এর হুবহু ডেটা কপি হয়ে যাবে।
