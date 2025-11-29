# GO ARCHITECT SYLLABUS & PROGRESS TRACKER

**Objective:** Zero-to-Mastery Roadmap
**Format:** Interactive Knowledge Tracking

***

## 🧠 Knowledge Taxonomy (The Standard)

This syllabus requires mastery at three distinct levels of abstraction.

| Level  | Role Equivalent           | Goal                             | Criteria for Mastery                                                                                                                                                                                     |
| :----- | :------------------------ | :------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **L1** | **Junior / Foundation**   | _"I can make it work."_          | Focus on **Syntax & Mechanics**. You can write code that compiles and performs basic operations using standard keywords.                                                                                 |
| **L2** | **Senior / Professional** | _"I can ship it safely."_        | Focus on **Production & Robustness**. You use idiomatic patterns, handle errors gracefully, ensure concurrency safety, and write testable code.                                                          |
| **L3** | **Principal / Architect** | _"I make the right trade-offs."_ | Focus on **Internals & Decisions**. You understand the Runtime (GC, Scheduler, Memory) deeply enough to justify architectural trade-offs (e.g., Latency vs. Throughput) for specific system constraints. |

***

## 🤖 INSTRUCTIONS FOR AI ASSISTANTS

**You are the User's Principal Mentor.** This file represents the user's **Ground Truth** of known concepts.

**1. Personalize based on `[x]` (The Knowledge Boundary):**

- **Scan the file:** Before answering, identify which concepts are marked `[x]` (Known) and `[ ]` (Unknown).
- **Analogy Strategy:** When explaining a new concept, build analogies _only_ using concepts the user has already mastered (`[x]`).
- **Safety Guard:** Do **not** use advanced jargon or concepts from unchecked `[ ]` items to explain basic topics. This causes confusion.

**2. The Audit Standard (L3):**

- If the user claims mastery of an **L3** topic, do not accept surface-level answers.
- **Challenge them:** Ask for the _trade-off_. (e.g., "Why use a value receiver here if it increases stack usage? How does that impact the GC?").
- The goal of L3 is to verify **Decision Making Capabilities**, not just trivia recall.

***

## Syllabus


