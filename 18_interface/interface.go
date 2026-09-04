package main

import "fmt"

// ১. Learner ইন্টারফেস তৈরি করা (কাজের তালিকা / শর্ত)
// শর্ত হলো: যার ভেতরে Study() নামে মেথড আছে, সে-ই একজন Learner
type Learner interface {
	Study() string
}

// ২. দুটি আলাদা Struct তৈরি করা
type SchoolStudent struct {
	Name string
}

type CollegeStudent struct {
	Name string
}

// ৩. Struct-গুলোর নিজস্ব স্টাইলে Study() মেথড তৈরি করা
func (s SchoolStudent) Study() string {
	return s.Name + " স্কুলের বই মুখস্থ করছে।"
}

func (c CollegeStudent) Study() string {
	return c.Name + " কলেজের প্রজেক্ট নিয়ে গবেষণা করছে।"
}

// ৪. এমন একটি ফাংশন যা শুধু 'Learner' ইন্টারফেস চেনে, কোনো নির্দিষ্ট Struct নয়
func StartLearning(l Learner) {
	// সে জানে l-এর ভেতরে Study() আছেই (যেহেতু সে Learner শর্ত পূরণ করেছে)
	fmt.Println(l.Study())
}

func main() {
	// SchoolStudent এর ইনস্ট্যান্স তৈরি
	schoolBoy := SchoolStudent{Name: "রহিম"}
	
	// CollegeStudent এর ইনস্ট্যান্স তৈরি
	collegeBoy := CollegeStudent{Name: "করিম"}

	// দুজনেই Learner ইন্টারফেসের শর্ত পূরণ করেছে (দুজনেরই Study মেথড আছে)
	// তাই দুজনেই StartLearning-এ ঢুকতে পারবে
	StartLearning(schoolBoy)
	StartLearning(collegeBoy)
}
