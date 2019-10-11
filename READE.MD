# Go Library for Concurrent Data Structures

Most of data structures in this library uses a general technique of lock-free programming. Data structures can be transformed from serial to concurrent using compare-and-swap, fetch-and-add instructions, which are supported on most modern CPUs.

The library has no external depenencies.

Install

    go get github.com/morningsend/concurrent