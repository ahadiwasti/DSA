# 💻 Writing the Code

By this point in the Canvas, you’ve:  
- Understood the **constraints**  
- Generated **ideas**  
- Evaluated **complexities**  
- Picked the **best approach**  

👉 Now it’s time to code!

---

## 🚫 Don’t Jump Straight to Code
A common mistake is to **rush into coding immediately**.  
- This is bad in real life.  
- It’s even worse in an interview.  

Always discuss **constraints, ideas, and complexities first**.  
Once both you and your interviewer are aligned, then start coding.

---

## 🖊️ Interview Coding ≠ IDE Coding
Most of us are used to coding with the help of our IDEs.  
In interviews, however, you may be asked to code:  
- On a **whiteboard**  
- On **paper**  
- In a **shared online editor** with limited features  

This can feel strange at first, but with practice, it becomes easier.  
The great news is that coding without your IDE actually makes you a **better engineer overall**.

---

## ✅ Tips for Coding Outside Your IDE
- **Think before you code**: On paper/whiteboard, there’s no easy “undo”.  
- **Follow clean code style**: Use clear variable names, proper indentation, and neat formatting.  
- **Decompose logically**: Break your solution into smaller steps — don’t copy-paste.  
- **Double-check your code**: You can’t rely on a compiler. Read your code carefully before saying it’s done.  

---

## 🌐 Online Interviews
Today, many interviews happen online. Usually, this means:  
- A shared editor (with syntax highlighting but no autocomplete)  
- Sometimes a platform like HackerRank or Codility  
- Rarely a full IDE environment  

💡 **Advice**: Ask in advance how coding will be conducted. Then practice in that environment if possible.  
If unsure, practice on:  
- Paper  
- Whiteboard  
- Simple editors without autocomplete  

This prepares you for any scenario.

---

## 📝 Example – ZigZag Problem
Let’s apply this to the **ZigZag problem**.  
Here’s the C++ implementation from the video (written without relying on IDE helpers):

```cpp
int longestZigZagSequence(int N, std::vector<int> a) {
  std::vector<int> up;
  std::vector<int> down;
  int bestLength = 1;

  up.push_back(1);
  down.push_back(1);

  for (int i = 1; i < N; i++) {
    up.push_back(1);
    down.push_back(1);

    for (int j = 0; j < i; j++) {
      if (a[i] > a[j]) {
        up[i] = max(down[j] + 1, up[i]);
      }
      if (a[i] < a[j]) {
        down[i] = max(up[j] + 1, down[i]);
      }
    }
    bestLength = max(bestLength, max(up[i], down[i]));
  }

  return bestLength;
}


```
### ✅ Summary

In this section you learned:
When to start coding at an interview (after constraints, ideas, and complexities).
Why coding for an interview is different from coding in your IDE.
Practical tips for writing clean, correct code outside your IDE.
An example solution using the Canvas approach.

### ▶️ What’s Next?

The final step of the Canvas is [**Testing**](../Testing/README.md).
You can’t be 100% sure your solution is correct without running it through test cases. Testing not only proves correctness but also demonstrates professionalism. Let’s dive into how to test solutions effectively.