#💡 Use Case: Fetch User Data from Multiple Services Concurrently
Imagine you have a user dashboard that needs to show:

1. Profile info (from Service A)
   
2. Order history (from Service B)

3. Notifications (from Service C)

Instead of calling these one after another (which would be slow), you can run them in parallel using goroutines.



## Findings
WaitGroup is a great way to wait for multiple goroutines to finish.

### Important Functions
- `sync.WaitGroup`
  - `Add(int)`: Add a delta to the WaitGroup counter.
  - `Done()`: Decrements the WaitGroup counter by one.
  - `Wait()`: Blocks until the WaitGroup counter is zero.

## Pitfalls
### Usage of unbuffered channel
In case of Unbuffered channel, the sender and receiver must be ready at the same time. 
If the receiver is not ready, the sender will block until the receiver is ready to receive the value. 
This can lead to deadlocks if not handled properly.