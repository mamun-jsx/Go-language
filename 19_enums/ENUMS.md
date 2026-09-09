# Go-তে Enums (ইনুমস) একদম সহজ ভাষায় 🚀

অন্যান্য প্রোগ্রামিং ভাষা (যেমন- TypeScript, Java, C#) তে `enum` নামক একটি ডেডিকেটেড ফিচার থাকে, যা দিয়ে আমরা নির্দিষ্ট কিছু ভ্যালু ফিক্সড করে দিতে পারি (যেমন- অর্ডারের অবস্থা: Pending, Shipped, Delivered)। 

কিন্তু Go-তে সরাসরি কোনো `enum` কি-ওয়ার্ড নেই! এর বদলে Go-তে আমরা **`type`**, **`const`** এবং **`iota`** ব্যবহার করে Enum-এর মতো কাজ করি। 

---

## 📦 Enums কেন দরকার?
ধরুন, আপনার একটি ই-কমার্স সিস্টেম আছে। সেখানে একটি অর্ডারের স্ট্যাটাস হতে পারে: 
- Recived (০)
- Confirmed (১)
- Shipped (২)
- Delivered (৩)
- Canceled (৪)

এখন আপনি যদি ভ্যালুগুলো স্ট্রিং বা ইন্টিজার হিসেবে সাধারণ ভেরিয়েবলে রাখেন, ভুল হওয়ার সম্ভাবনা থাকে। Enum ব্যবহার করলে আপনার কোড অনেক বেশি নিরাপদ (Type-safe) ও ক্লিন হয়।

---

## 💻 `enums.go` এর সম্পূর্ণ কোড 

নিচে `enums.go` ফাইলের সম্পূর্ণ কোড লাইন-বাই-লাইন ব্যাখ্যাসহ দেওয়া হলো:

```go
package main

import "fmt"

// ==========================================
// ১. Integer Enum তৈরি করা (iota ব্যবহার করে)
// ==========================================

// 'OrderStatus' নামের একটি নতুন টাইপ বানালাম, যা মূলত একটি 'int'
type OrderStatus int

// const ব্লকের ভেতরে iota ব্যবহার করলে ভ্যালুগুলো অটোমেটিক ০, ১, ২, ৩... এভাবে বাড়তে থাকে।
const (
	Recived OrderStatus = iota // ভ্যালু: 0
	Confirmed                  // ভ্যালু: 1
	Shipped                    // ভ্যালু: 2
	Delivered                  // ভ্যালু: 3
	Canceled                   // ভ্যালু: 4
)

// ==========================================
// ২. String Enum তৈরি করা
// ==========================================

// 'DeleteStatus' নামের আরেকটি টাইপ বানালাম, যা মূলত একটি 'string'
type DeleteStatus string

// এখানে আমরা iota ব্যবহার করতে পারবো না কারণ এগুলো স্ট্রিং। তাই ম্যানুয়ালি ভ্যালু সেট করে দিলাম।
const (
	Yes DeleteStatus = "yes"
	No  DeleteStatus = "no"
)

// ==========================================
// ৩. ফাংশনে Enum ব্যবহার করা
// ==========================================

// changeOrderStatus ফাংশনটি শুধু 'OrderStatus' টাইপের ভ্যালু ইনপুট হিসেবে নিবে। 
// আপনি চাইলে সাধারণ int (যেমন 1 বা 2) পাঠাতে পারবেন না, আপনাকে নির্দিষ্ট Enum-ই পাঠাতে হবে।
func changeOrderStatus(status OrderStatus) {
	fmt.Println("change order status to ", status)
}

// deleteOrderStatus ফাংশনটি শুধু 'DeleteStatus' টাইপের ভ্যালু নিবে।
func deleteOrderStatus(status DeleteStatus) {
	fmt.Println("delete status ", status)
}

// ==========================================
// ৪. মেইন ফাংশন
// ==========================================

func main() {
	// OrderStatus টাইপের 'Delivered' (যার মান 3) ফাংশনে পাঠালাম
	changeOrderStatus(Delivered)
	
	// DeleteStatus টাইপের 'Yes' (যার মান "yes") ফাংশনে পাঠালাম
	deleteOrderStatus(Yes)
}
```

### 🖥️ কোডের আউটপুট:
এই কোডটি `go run enums.go` লিখে রান করলে টার্মিনালে নিচের রেজাল্ট আসবে:

```text
change order status to  3
delete status  yes
```

---

## 🧐 Go-তে Enum ব্যবহারের সুবিধা কী?

১. **টাইপ সেফটি (Type Safety):** 
`changeOrderStatus` ফাংশনে আপনি ভুল করে `changeOrderStatus(99)` কল করতে পারবেন না। আপনাকে `OrderStatus` টাইপের ভ্যালুই দিতে হবে। এতে কোডে বাগ আসার সম্ভাবনা কমে যায়।

২. **Iota ম্যাজিক:** 
`iota` ব্যবহার করার ফলে আপনাকে নিজে থেকে `0, 1, 2, 3` লিখতে হয় না। Go নিজেই সেগুলো বসিয়ে নেয়।

৩. **ক্লিন কোড:** 
ভ্যালুর বদলে `Shipped`, `Delivered` নামগুলো পড়লে যে কেউ কোড দেখে সহজেই বুঝতে পারে কী হচ্ছে। 
