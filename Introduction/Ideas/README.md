# 💡 Idea Generation

One of the goals of this course is to teach you how to **solve new problems**, not just memorize answers to old ones.  

Why?  
- New interview questions are created all the time — it’s impossible to know them all.  
- Even if you know many solutions, interviewers may **modify the question** and you’ll need to adapt.  
- Employers want to see candidates who can **think on the spot** and come up with ideas.  

This section gives you **strategies for generating ideas** when facing algorithm problems. With practice, these strategies become second nature.

---

## 🔑 Key Strategies for Idea Generation

### 1️⃣ Simplify the Task
When the problem looks too complex, try breaking it into a **simpler version**.  

**Example:**  
> “A map of streets is given as an N x M grid. Each intersection has people. They all want to meet at one intersection. Find the spot that minimizes the total walking distance (Manhattan distance).”

- First, simplify: Imagine only **one street (1D)**.  
- Solution in 1D: The **median position** minimizes total walking distance.  
- Back to 2D: The X and Y positions can be solved separately as two 1D problems.  

👉 By solving the simpler case, we can solve the harder one.

---

### 2️⃣ Try a Few Examples
Don’t rely only on the example given by the interviewer.  
Create your own small examples and test them quickly. This helps you notice **patterns**.  

**Example:**  
> “There are N+1 parking spots (0 to N). N cars are parked, one spot empty. Rearrange so car #i goes to spot #i. The only allowed move is placing a car in the empty spot.”

- At first it looks tricky.  
- Write down 4–5 examples on paper.  
- By experimenting, you’ll notice a pattern in how the empty spot moves.  

👉 Examples often reveal hidden insights.

---

### 3️⃣ Think of Suitable Data Structures
Sometimes a problem almost “hints” that a certain **data structure** is needed.  

**Example:**  
> Design a data structure that supports:  
> • Insert a number in O(log N)  
> • Return the median in O(1)  
> • Delete the median in O(log N)

- Arrays, stacks, or vectors won’t work.  
- A **binary tree** or **heaps** come to mind.  
- Two heaps solution:  
  - Max-heap for the smaller half of numbers  
  - Min-heap for the larger half  
  - Together, they give the median in constant time.  

👉 When you feel a problem needs structure, think about combining **known data structures**.

---

### 4️⃣ Think About Related Problems
If you’re stuck, recall **similar problems** you’ve solved before.  
- Many interview problems are variations of each other.  
- Try to adapt the solution of a related problem.  

👉 This won’t always work, but often it will give you a direction.

---

## 📝 Important Notes
- Practicing is the **only real way** to get good at problem-solving.  
- With enough practice, you’ll start noticing **patterns** that reappear in many problems.  
- The Algorithm Design Canvas helps you write down and organize your **ideas** clearly.  

---

## 📚 Extra Resources
- [TopCoder Tutorial: How to Find a Solution](https://www.topcoder.com/thrive/articles/How%20to%20find%20a%20solution) – useful read on problem-solving strategies.  

---

## 🔍 Example
We’ll apply these strategies to the **ZigZag problem**.  
Fun fact: there’s even a faster O(N) solution available on TopCoder’s website.

---

## ✅ Summary
In this section you learned:  
- It’s better to **learn problem-solving techniques** than to memorize every possible interview question.  
- There are several **strategies** that help generate ideas:
  1. Simplify the task  
  2. Try examples  
  3. Think about data structures  
  4. Look at related problems  
- The more you practice, the better you’ll become at spotting patterns.

---

## ▶️ What’s Next?
After generating ideas, the next step is to think about [**time and space complexity**](../Complexity/README.md).  
This is a critical topic, so let’s dive into it in the next section!
