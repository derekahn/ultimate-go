## Data Races

A data race is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that totally appear random or can be never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

## Notes

* Goroutines need to be coordinated and synchronized.
* When two or more goroutines attempt to access the same resource, we have a data race.
* Atomic functions and mutexes can provide the support we need.

## Cache Coherency and False Sharing
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (30:09-38:30)](https://youtu.be/WDIkqP4JbkE?t=1809)

## Cache Coherency and False Sharing Notes

* Thread memory access matters.
* If your algorithm is not scaling look for false sharing problems.

## Links

[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206)

## False Sharing

Understanding how the hardware works is an critical component to understanding how to write the most performant code you can. Knowing the basics of processor caching can help you make better decisions within the scope of writing idiomatic code.
