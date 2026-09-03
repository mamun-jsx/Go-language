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
