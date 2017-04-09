## Channels

Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

### Channel Design

Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

**Language Mechanics:**

* Use channels to orchestrate and coordinate goroutines.
  * Focus on the signaling semantics and not the sharing of data.
  * Signal with data or without data.
  * Question their use for synchronizing access to shared state.
    * _There are cases where channels can be simpler for this but initially question._
* Unbuffered channels:
  * Blocks the sending goroutine from moving forward until a different goroutine has received the data signal.
    * The sending goroutine gains immediate knowledge their data signal has been received.
  * Both goroutines involved must be at the channel at the same time.
    * More important: The Receive happens before the Send.
  * Trade-offs:
    * We take the benefit of knowing the data signal has been received for the cost of higher blocking latency.
* Buffered channels:
  * Does NOT block the sending goroutine from moving forward until a different goroutine has received the data signal.
    * The sending goroutine gains no knowledge that their data signal has been received.
  * Both goroutines involved don't need to be at the channel at the same time.
    * More important: The Send happens before the Receive.
  * Trade-offs:
    * We take the benefit of reducing blocking latency for the cost of not knowing if/when the data signal is received.
* Closing channels:
  * Signaling without data.
  * Perfect for signaling cancellations and deadlines.
* NIL channels:
  * Turn off signaling
  * Perfect for rate limiting or short term stoppages.

**Design Philosophy:**

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

* If any given Send on a channel `CAN` cause the goroutine to block:
  * You have less buffers compared to the number of goroutines that will perform a Send at any given time.
  * An example would be a resource pool of database connections.
  * This requires knowing what happens when the Send blocks because this will create a situation of back pressure inside the application in front of the channel.
  * The things discussed above about [writing concurrent software](https://github.com/ardanlabs/gotraining/blob/master/reading/design_guidelines.md#writing-concurrent-software) must be taken into account for this channel.
  * Not knowing what happens when the Send blocks on the channel is not a license to guess but a directive to STOP, understand and take appropriate measures.
* If any given Send on a channel `WON'T` cause the goroutine to block:
  * You have a buffer for every goroutine that will perform a Send.
  * You will abandon the Send immediately if it can't be performed.
  * An example would be a fan out pattern or pipelining.
  * This requires knowing if the size of the buffer is appropriate and if it is acceptable to abandon the Send.
* Less is more with buffers.
  * Don’t think about performance when thinking about buffers.
  * Buffers can help to reduce blocking latency between signaling.
    * Reducing blocking latency towards zero does not necessarily mean better throughput.
    * If a buffer of one is giving you good enough throughput then keep it.
    * Question buffers that are larger than one and measure for size.
    * Find the smallest buffer possible that provides good enough throughput.
