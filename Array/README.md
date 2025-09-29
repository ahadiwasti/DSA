# 📂 Arrays – Introduction

## 📌 What is an Array?
An **array** is a collection of elements stored in **contiguous memory locations**.  
It allows us to store multiple values of the **same type** under a single variable name, and access them using an **index**.

Think of an array like a row of boxes 🗃️, each holding one value, and you can quickly find each box by its position.

### Example in Go
```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(nums[0]) // 1
fmt.Println(nums[3]) // 4
```

## 🎯 Why Arrays Matter in DSA?
Arrays are one of the most fundamental data structures. They are the foundation for solving many coding interview problems, such as:

Searching
Sorting
Prefix sums
Sliding window problems
Two pointers approach

Mastering arrays gives you the skills to tackle more advanced data structures like strings, hash tables, stacks, queues, and dynamic programming.

⚡ Key Properties

Fixed index access → O(1) time to read/write an element by index
Iterating over array → O(n)
Insertion/deletion at the end → O(1) (amortized in dynamic arrays like Go slices)
Insertion/deletion in the middle → O(n) (since elements must be shifted)

🏗️ Common Array Interview Problems

Reverse an array
Find the maximum/minimum element
Move zeros to the end
Two Sum (pair sum problem)
Kadane’s Algorithm (maximum subarray sum)
Merge sorted arrays
Rotate an array

🚀 How We’ll Practice

In this section (arrays/), we will:
Learn array concepts step by step
Solve classic problems using different approaches (brute force → optimized)
Apply techniques like two pointers and sliding window