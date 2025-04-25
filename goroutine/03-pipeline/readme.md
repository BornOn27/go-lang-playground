# 🏗️ Real-world Scenario: Data Transformation Pipeline
Imagine a system that:

Ingests raw user events (JSON)

Parses the JSON into structs

Filters only login events

Enriches with user metadata

Writes them to a database or logs

Each stage is isolated, independent, and runs in parallel using goroutines — perfect for the pipeline pattern.

### ✅ Why It’s Useful:
Clean, readable flow

Stages are decoupled

Easy to test & extend

Max performance by using concurrency

Pattern is used in ETL systems, log processors, stream analytics, etc.

## Scenario Upgrade:
Let's imagine enriching user data is a slow API call. We want to:

Fan out: spawn multiple workers to enrich concurrently

Fan in: merge the enriched data back into a single output stream

### 🎯 Key Takeaways
✅ Fan-out: create N goroutines to handle a task concurrently

✅ Fan-in: gather results into a single output channel

🚀 Boosts performance for IO-heavy or slow operations

🧼 Keeps each stage clean and testable
