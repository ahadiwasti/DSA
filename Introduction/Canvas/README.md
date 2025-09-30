# 📝 Algorithm Design Canvas

## 📌 Why Algorithm Design Questions Matter
In coding interviews, **algorithm design questions** are some of the most common — and also the ones many people find the hardest.  

Why do so many struggle with them?  
- Lack of **theoretical knowledge**  
- Lack of **practice**  
- No clear, **systematic approach** to solving problems  

Many people try to prepare by memorizing “cookbook” solutions from the Internet. But memorizing answers rarely helps in a real interview. You need a **method** that works for *any* problem.

---

## 🎯 A Different Approach
Instead of memorizing, we’ll focus on building a **step-by-step process** for tackling problems.  

Some of us have been solving algorithmic problems for **over 20 years**. Through that experience, we developed a **systematic way of thinking** that can help you stay calm and structured during interviews.  

This process is what we call the **Algorithm Design Canvas**.

---

## 🖼️ What is the Algorithm Design Canvas?
The **Algorithm Design Canvas** is a **framework** that captures the way we approach algorithmic problems.  

Think of it as a **worksheet** or **blueprint** that organizes your thought process.  
No matter the problem — big or small, easy or hard — you can fill in the canvas to make sure you don’t miss important steps.

The Canvas has **5 major areas**:
1. Constraints  
2. Ideas  
3. Complexities  
4. Code  
5. Tests  

Taken together, these areas cover everything you need to think about when solving an algorithm design problem.

---

## Problem Name: *[Write problem name here]*

| **Constraints** | **Ideas** | **Complexities** | **Code** | **Tests** |
|-----------------|-----------|------------------|----------|-----------|
| List all problem constraints:<br>• Input size<br>• Data type limits<br>• Special rules | Brainstorm solution ideas:<br>• Idea 1 (brute force)<br>• Idea 2 (optimized) | For each idea:<br>• Time complexity (O(n), O(log n), etc.)<br>• Space complexity | Write the pseudocode or final code here | Add test cases:<br>• Smallest input<br>• Edge cases<br>• Largest input |

---

## Example – Two Sum Problem

| **Constraints** | **Ideas** | **Complexities** | **Code** | **Tests** |
|-----------------|-----------|------------------|----------|-----------|
| • Array length up to 10⁵<br>• Values can be positive/negative<br>• Target value fits in int | • Brute force: check all pairs<br>• HashMap for lookup<br>• Two pointers if array sorted | • Brute force: O(n²) time, O(1) space<br>• HashMap: O(n) time, O(n) space<br>• Two pointers: O(n) time, O(1) space | ```go<br>func twoSum(nums []int, target int) []int {<br>    m := map[int]int{}<br>    for i, num := range nums {<br>        if j, ok := m[target-num]; ok {<br>            return []int{j, i}<br>        }<br>        m[num] = i<br>    }<br>    return nil<br>}<br>``` | • nums = [2,7,11,15], target = 9 → [0,1]<br>• nums = [3,2,4], target = 6 → [1,2]<br>• nums = [3,3], target = 6 → [0,1] |

---

## 🔑 The 5 Areas of the Canvas

### 1️⃣ Constraints
- The **first step** in solving any problem is to **understand the rules and limits**.  
- Examples of constraints:
  - How large can the input array be?  
  - Can the string contain Unicode characters?  
  - Can the graph have negative edge weights?  

💡 *Tip:* Always write down the constraints first. They determine what solutions are possible.  

---

### 2️⃣ Ideas
- Once you know the constraints, start **brainstorming possible solutions**.  
- In an interview, you usually discuss **1–3 ideas**.  
- Start with a simple idea, explain it, then try to improve it.  
- Keep descriptions **short and clear**, so your interviewer can easily follow.

💡 *Tip:* Don’t rush to code the first idea. Explore alternatives first.

---

### 3️⃣ Complexities
- Every idea has a **time complexity** and a **space complexity**.  
- This is where you **estimate performance**:
  - Time = How many operations will it take?  
  - Space = How much memory will it use?  
- Sometimes you trade time for memory, or memory for time.

💡 *Tip:* Always compare your ideas here. Often the interviewer wants to see how you reason about trade-offs.

---

### 4️⃣ Code
- After you and the interviewer agree on a good idea, **then you write the code**.  
- Writing code in an interview is different from writing in an IDE:
  - You need to keep it **clear and bug-free**  
  - Use **good variable names**  
  - Write in a way the interviewer can follow easily  

💡 *Tip:* Talk while you code. Explain what you’re doing.

---

### 5️⃣ Tests
- The last step is to **test your solution** with examples.  
- Many candidates skip this step — big mistake!  
- Good tests show you understand **edge cases** and can **validate your solution**.

Examples of tests:
- Smallest possible input  
- Largest possible input  
- Special cases (like empty arrays, negative numbers, duplicates)

💡 *Tip:* Testing shows professionalism and catches silly mistakes.

---

## ✅ Why This Process Works
By following the Canvas:
- You always know **what to do next**  
- You cover all important areas (constraints → ideas → complexities → code → tests)  
- You show the interviewer that you have a **structured way of thinking**  

Instead of freezing up, you can calmly work through the problem step by step.

---

## 🖨️ Print the Canvas
The best way to practice is to [**print out several copies of the Canvas**](./CanvasDefinition.pdf) and fill them in as you solve problems.  
This builds muscle memory and helps you naturally follow the process in an interview.

---

## 🎬 Example Coming Up
In the next sections, we’ll walk through an actual problem (the classic **Zig-Zag problem from TopCoder**) using the Canvas.  
This way, you’ll see how each area is filled in with real details.

---

## ▶️ What’s Next?
We’ll begin with the [**first area of the Canvas: Constraints**](../Constraints/README.md) — an often overlooked but absolutely critical step in solving problems.
