# 🛡️ Real-Time Use Case: API Rate Limiting (Throttling)

##✅ Scenario:
Suppose you're consuming a third-party API that only allows 1 request per second (to prevent abuse). You want to throttle your calls so that no more than 1 API request is made per second, even if your app tries to send more.

##✅ Learn:
time.Ticker for throttling

select block for managing events

Graceful shutdown with done channel

## ✅ Takeaways:
Use time.Ticker for fixed-rate processing.

Nest select to avoid blocking when no item is ready.

Separate channels for requests and shutdown keep things clean.