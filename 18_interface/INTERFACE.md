# Go (Golang) এ Interface (ইন্টারফেস) এর বিস্তারিত ধারণা

Go-তে **Interface** হলো এমন একটি ডেটা টাইপ, যেখানে এক বা একাধিক মেথডের সিগনেচার (Signature) বা ব্লুপ্রিন্ট দেওয়া থাকে, কিন্তু সেগুলোর ভেতরে কোনো কোড লেখা থাকে না। 

কোনো স্ট্রাক্ট যদি ইন্টারফেসে থাকা সবগুলো মেথড তৈরি (implement) করে, তাহলে আমরা বলতে পারি সেই স্ট্রাক্টটি ওই ইন্টারফেসকে ফলো করছে। 

নিচে `interface.go` ফাইলে থাকা কোডটি লাইন-বাই-লাইন ব্যাখ্যা করা হলো:

---

## 💻 সম্পূর্ণ কোড (`interface.go`)

```go
package main

import "fmt"

// ১. Payment ইন্টারফেস ডিক্লেয়ার করা
type Payment interface {
	pay(amount float32)
}

// ২. Bkash স্ট্রাক্ট ডিক্লেয়ার করা
type Bkash struct {
	phoneNumber string
}

// ৩. Bkash এর জন্য Payment ইন্টারফেস ইমপ্লিমেন্ট করা
func (b Bkash) pay(amount float32) {
	fmt.Printf("Making payment of %.2f using bKash number: %s\n", amount, b.phoneNumber)
}

// ৪. Card স্ট্রাক্ট ডিক্লেয়ার করা
type Card struct {
	cardNumber string
	bankName   string
}

// ৫. Card এর জন্য Payment ইন্টারফেস ইমপ্লিমেন্ট করা
func (c Card) pay(amount float32) {
	fmt.Printf("Making payment of %.2f using %s Card: %s\n", amount, c.bankName, c.cardNumber)
}

// ৬. পেমেন্ট প্রসেস করার জন্য একটি ফাংশন যা Payment ইন্টারফেস গ্রহণ করে
func processPayment(p Payment, amount float32) {
	p.pay(amount) // ডাইনামিকালি সঠিক pay() মেথড কল হবে
}

func main() {
	// Bkash এর ইনস্ট্যান্স তৈরি
	myBkash := Bkash{
		phoneNumber: "01711000000",
	}

	// Card এর ইনস্ট্যান্স তৈরি
	myCard := Card{
		cardNumber: "1234-5678-9012",
		bankName:   "DBBL",
	}

	// processPayment ফাংশনে বিভিন্ন পেমেন্ট মেথড পাঠানো হচ্ছে
	processPayment(myBkash, 500.50)
	processPayment(myCard, 1200.00)
}
```

---

## 📝 লাইন-বাই-লাইন ব্যাখ্যা

### ধাপ ১: ইন্টারফেস ডিক্লেয়ার করা
```go
type Payment interface {
	pay(amount float32)
}
```
এখানে আমরা `Payment` নামে একটি ইন্টারফেস তৈরি করেছি যার ভেতরে `pay` নামের একটি মেথড সিগনেচার আছে। এর অর্থ হলো, যে স্ট্রাক্টই নিজেকে `Payment` হিসেবে পরিচয় দিতে চাইবে, তার অবশ্যই একটি `pay` নামের মেথড থাকতে হবে যা একটি `float32` ভ্যালু গ্রহণ করবে।

### ধাপ ২: স্ট্রাক্ট তৈরি করা
```go
type Bkash struct {
	phoneNumber string
}
```
আমরা `Bkash` নামে একটি স্ট্রাক্ট বানালাম।

### ধাপ ৩: ইন্টারফেস ইমপ্লিমেন্ট করা
```go
func (b Bkash) pay(amount float32) {
	fmt.Printf("Making payment of %.2f using bKash number: %s\n", amount, b.phoneNumber)
}
```
আমরা `Bkash` স্ট্রাক্টের জন্য `pay` মেথডটি তৈরি করলাম। যেহেতু `Bkash` স্ট্রাক্টে `pay` মেথডটি আছে, তাই Go অটোমেটিকভাবে ধরে নিবে যে `Bkash` স্ট্রাক্টটি `Payment` ইন্টারফেসকে ফলো করছে। (Go-তে Java এর মত `implements Payment` লেখার দরকার হয় না)।

### ধাপ ৪ ও ৫: আরেকটি স্ট্রাক্টে সেম কাজ করা
```go
type Card struct { ... }
func (c Card) pay(amount float32) { ... }
```
একইভাবে আমরা `Card` নামে আরেকটি স্ট্রাক্ট বানালাম এবং সেটির জন্যও `pay` মেথডটি ইমপ্লিমেন্ট করলাম। এখন `Card` স্ট্রাক্টটিও একটি `Payment` ইন্টারফেসের অংশ!

### ধাপ ৬: ইন্টারফেসের আসল ম্যাজিক! (Polymorphism)
```go
func processPayment(p Payment, amount float32) {
	p.pay(amount)
}
```
এই `processPayment` ফাংশনটি কোনো নির্দিষ্ট স্ট্রাক্ট (যেমন Bkash বা Card) রিসিভ করছে না, বরং সে রিসিভ করছে `Payment` ইন্টারফেস। এর মানে হলো, আমরা এই ফাংশনে যেকোনো মেথড পাঠাতে পারবো (Bkash, Card, Rocket বা অন্য কিছু) যদি তাদের `pay` নামের মেথডটি থাকে!

### ধাপ ৭: Main ফাংশনে কল করা
```go
processPayment(myBkash, 500.50)
processPayment(myCard, 1200.00)
```
এখানে আমরা একই ফাংশনে `myBkash` এবং `myCard` পাঠাচ্ছি এবং ফাংশনটি নিজে থেকেই বুঝে নিচ্ছে যে কার জন্য কোন `pay()` মেথডটা কল করতে হবে।

---

## 🖥️ আউটপুট
আপনি কোডটি রান করলে নিচের আউটপুটটি দেখতে পাবেন:

```text
Making payment of 500.50 using bKash number: 01711000000
Making payment of 1200.00 using DBBL Card: 1234-5678-9012
```

### 🤔 এখানে Struct কেন ব্যবহার করা হলো এবং এর কাজ কী?
ইন্টারফেস হলো শুধুমাত্র একটি "চুক্তি" বা "রুলস" (যে কী কী মেথড থাকতে হবে)। কিন্তু এই মেথডগুলো আসলে কোনো একটা ডেটার উপর কাজ করে। 

যেমন, `Bkash` দিয়ে পে করার সময় আমাদের `phoneNumber` দরকার, আবার `Card` দিয়ে পে করার সময় `cardNumber` দরকার। এই আলাদা আলাদা ডেটাগুলো (State) ধরে রাখার জন্যই আমরা **Struct** ব্যবহার করেছি। 

- **Struct এর কাজ:** ডেটা বা প্রপার্টি ধরে রাখা (যেমন: phoneNumber, cardNumber)।
- **Interface এর কাজ:** স্ট্রাক্টগুলো কী কী কাজ (Behavior) করতে পারবে তা নির্দিষ্ট করে দেওয়া (যেমন: pay করা)।

সংক্ষেপে, Struct হলো **"কী ডেটা আছে"**, আর Interface হলো **"সে কী করতে পারে"**।

### 💡 কেন ইন্টারফেস দরকার?
ইন্টারফেস ব্যবহার করার কারণে আমাদের কোড অনেক ফ্লেক্সিবল (Flexible) হয়ে যায়। ভবিষ্যতে যদি আমরা `Rocket` নামে নতুন একটি পেমেন্ট সিস্টেম যোগ করতে চাই, আমাদের `processPayment` ফাংশনে কোনো পরিবর্তন আনতে হবে না। শুধু `Rocket` এর জন্য একটি `pay` মেথড বানালেই কাজ হয়ে যাবে!
