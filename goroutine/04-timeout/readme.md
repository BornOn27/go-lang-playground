## 📦 Real-Life Scenario: External API Call with Timeout
Imagine your Go service needs to call an external payment gateway. You don't want the API to hang forever — so you want to:

Set a timeout (e.g., 2 seconds)

Cancel the request if it takes too long

Ensure any running goroutines exit early

Clean up resources gracefully

## ✅ What You'll Learn:
context.Context propagation

Handling cancellation signals inside goroutines

Avoiding goroutine leaks with proper cleanup

### 🔥 Why This Is Powerful
🚫 Avoid Leaks:
Without cancellation, long-running goroutines might keep running even when the user doesn't care anymore.

✅ Graceful Shutdown:
You can use context.WithCancel in servers to cancel in-flight requests when shutting down.

⛑️ User Experience:
A timeout gives fail-fast behavior, so users aren’t stuck waiting forever.

