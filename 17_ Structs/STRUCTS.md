# Go (Golang) এ Structs এর বিস্তারিত ধারণা

Go-তে **Struct** (স্ট্রাক্ট) হলো অন্যান্য অবজেক্ট-ওরিয়েন্টেড ভাষার `Class` এর মতো কাজ করে। এটি একটি ডেটা স্ট্রাকচারের ব্লুপ্রিন্ট (Blueprint)। Struct ব্যবহার করে আমরা বিভিন্ন ধরনের ডেটা টাইপ (যেমন: string, int, float) একসাথে গ্রুপ করে একটি কমপ্লেক্স ডেটা স্ট্রাকচার তৈরি করতে পারি। 

আপনার `Structs.go` ফাইলে আপনি যা যা প্র্যাকটিস করেছেন, তা নিচে ধাপে ধাপে (Step-by-step) বাংলায় ব্যাখ্যা করা হলো:

---

## ১. Struct ডিক্লেয়ার (Declare) করা
Struct তৈরি করতে `type` এবং `struct` কি-ওয়ার্ড ব্যবহার করতে হয়। 

```go
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // time প্যাকেজ থেকে এসেছে
}
```
এখানে আমরা `order` নামের একটি স্ট্রাক্ট বানিয়েছি যার মধ্যে ৪টি ফিল্ড রয়েছে।

---

## ২. Struct এর Instance (অবজেক্ট) তৈরি করা এবং প্রিন্ট করা
স্ট্রাক্ট ডিক্লেয়ার করার পর, আমাদের এর একটি Instance বা অবজেক্ট বানাতে হয়।

```go
func main() {
	myOrder := order{
		id:     "10",
		amount: 43,
		status: "received",
	}

	fmt.Println("Order struc: ", myOrder)
	// Output: Order struc:  {10 43 received {0 0 <nil>}}
	
	// %+v ব্যবহার করলে ফিল্ডের নামসহ সুন্দরভাবে প্রিন্ট হয়
	fmt.Printf("order %+v\n", myOrder) 
	// Output: order {id:10 amount:43 status:received createdAt:0001-01-01 00:00:00 +0000 UTC}
}
```

---

## ৩. Struct এর ভ্যালু পরিবর্তন (Modify) করা
ডট (`.`) ব্যবহার করে আমরা খুব সহজেই স্ট্রাক্টের ভেতরের যেকোনো ফিল্ড অ্যাক্সেস করতে পারি এবং তার মান পরিবর্তন করতে পারি।

```go
myOrder.createdAt = time.Now()
myOrder.status = "shipped"

fmt.Println("Updated status: ", myOrder.status)
// Output: Updated status:  shipped
```

**নোট:** আপনি যখন `myOrder` এর `status` পরিবর্তন করে "shipped" করবেন, তখন এটি শুধুমাত্র `myOrder` তেই পরিবর্তন হবে। আপনি যদি `myOrder2` নামে অন্য কোনো ইনস্ট্যান্স তৈরি করেন, সেটির উপর এর কোনো প্রভাব পড়বে না।

---

## ৪. Receiver Methods (রিসিভার মেথড)
Go-তে ক্লাস না থাকলেও স্ট্রাক্টের সাথে আমরা মেথড যুক্ত করতে পারি। এটি দুই ধরণের হতে পারে:

### A. Value Receiver (শুধুমাত্র ডেটা পড়ার জন্য)
যখন আমাদের স্ট্রাক্টের কোনো ডেটা শুধুমাত্র পড়তে হবে (পরিবর্তন করার দরকার নেই), তখন আমরা Value Receiver ব্যবহার করি। এখানে পয়েন্টার (Pointer) পাঠাতে হয় না।

```go
func (o order) getAmount() float32 {
	return o.amount
}

// ব্যবহার:
fmt.Println("getAmount: ", myOrder.getAmount())
// Output: getAmount:  43
```

### B. Pointer Receiver (ডেটা পরিবর্তন করার জন্য)
যখন মেথডের ভেতর থেকে আমাদের মূল স্ট্রাক্টের কোনো ভ্যালু পরিবর্তন করতে হয়, তখন আমাদের পয়েন্টার (`*`) ব্যবহার করে মেমোরি লোকেশন পাঠাতে হয়।

```go
func (o *order) changeStatus(status string) {
	o.status = status 
	// Go নিজে থেকেই pointer কে dereference করে নেয়, তাই *o.status লিখতে হয় না।
}

// ব্যবহার:
myOrder.changeStatus("Confirmed")
fmt.Println("change status receiver: ", myOrder.status)
// Output: change status receiver:  Confirmed
```

---

## ৫. Constructor Function (কনস্ট্রাক্টর)
Go-তে বাই-ডিফল্ট কোনো কনস্ট্রাক্টর নেই, কিন্তু আমরা একটি ফাংশন তৈরি করে কনস্ট্রাক্টরের মতো কাজ করাতে পারি। সাধারণত এই ফাংশনগুলোর নাম `New...` দিয়ে শুরু হয় এবং এরা স্ট্রাক্টের পয়েন্টার রিটার্ন করে।

```go
func NewOrder(id string, amount float32, status string) *order {
	myOrder3 := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder3 // মেমোরি অ্যাড্রেস রিটার্ন করছি
}

// ব্যবহার:
myNewOrder := NewOrder("100", 30, "Not Paid")
fmt.Println("my newOrder: ", myNewOrder)
// Output: my newOrder:  &{100 30 Not Paid {0 0 <nil>}}
```
**আউটপুটে খেয়াল করুন:** `&` চিহ্নটি দেখাচ্ছে কারণ এটি সরাসরি পয়েন্টার বা মেমোরি অ্যাড্রেস প্রিন্ট করছে।

---

## ৬. Inheritance in Go (Struct Embedding / Composition)
Go-তে অন্যান্য অবজেক্ট-ওরিয়েন্টেড ভাষার মতো `extends` বা সরাসরি Inheritance (ইনহেরিট্যান্স) নেই। তবে Go-তে আমরা **Struct Embedding** (বা Composition) ব্যবহার করে ইনহেরিট্যান্সের মতো সুবিধা পেতে পারি। অর্থাৎ, একটি স্ট্রাক্টের ভেতরে অন্য একটি স্ট্রাক্টকে যুক্ত (embed) করা যায়।

```go
// মূল স্ট্রাক্ট (Parent)
type User struct {
	name  string
	email string
}

// আরেকটি স্ট্রাক্ট (Child), যেখানে User কে embed করা হয়েছে
type Admin struct {
	User  // User স্ট্রাক্টটি এখানে inherit/embed করা হলো
	level int
}

// ব্যবহার:
admin1 := Admin{
	User: User{name: "Mamun", email: "mamun@example.com"},
	level: 1,
}

// আমরা সরাসরি প্যারেন্টের ফিল্ড অ্যাক্সেস করতে পারি! (admin1.User.name লেখার দরকার নেই)
fmt.Println("Admin Name: ", admin1.name)
fmt.Println("Admin Email: ", admin1.email)
// Output: Admin Name:  Mamun
// Output: Admin Email:  mamun@example.com
```

---

## ৭. Anonymous Structs (নামবিহীন স্ট্রাক্ট)
মাঝে মাঝে আমাদের এমন কিছু স্ট্রাক্ট দরকার হয় যা একবারই ব্যবহার হবে। তখন আলাদাভাবে `type` দিয়ে ব্লুপ্রিন্ট না বানিয়ে সরাসরি ভ্যারিয়েবল ডিক্লেয়ারেশনের সাথেই আমরা এটি তৈরি করতে পারি।

```go
person := struct {
	name   string
	isGood bool
}{name: "goLang", isGood: true}

fmt.Println("anonymous structs: ", person)
// Output: anonymous structs:  {goLang true}
```

---

### সারসংক্ষেপ (Summary)
১. `Struct` হলো ডেটা স্ট্রাকচারের ব্লুপ্রিন্ট।
২. ভ্যালু মডিফাই করতে হলে মেথডে **Pointer (`*`)** পাঠাতে হবে।
৩. ভ্যালু শুধু রিড (read) করতে হলে Pointer দরকার নেই।
৪. Go-তে Struct এর পয়েন্টার নিয়ে কাজ করার সময় ম্যানুয়ালি dereference (`*`) করার দরকার পড়ে না, Go এটি স্বয়ংক্রিয়ভাবে করে নেয়।
৫. Go-তে ক্লাসিক্যাল Inheritance নেই, এর বদলে **Struct Embedding (Composition)** ব্যবহার করা হয়।
