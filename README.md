# Cuckoo Filter Implementation in Go

The term "Cuckoo" in "Cuckoo Filter" is inspired by the behavior of the common cuckoo bird. Cuckoo birds are known for their behavior of laying their eggs in the nests of other bird species, tricking those birds into raising the cuckoo's young. This phenomenon is known as "brood parasitism."

The name "Cuckoo Filter" draws an analogy between the way cuckoo birds replace the eggs of other birds and the way elements are replaced or evicted from the filter's buckets during insertion. In a Cuckoo Filter, when a new element is inserted and a bucket is already occupied, the existing element is "evicted" or replaced, somewhat resembling the cuckoo bird's behavior of replacing eggs in other birds' nests.

So, the name "Cuckoo Filter" reflects this evict-and-replace behavior during insertion that is similar to how cuckoo birds replace eggs in the nests of other bird species. It's a playful and memorable way to describe the underlying mechanism of the data structure.
## About Cuckoo Filters

A Cuckoo Filter is a space-efficient probabilistic data structure designed for fast set membership queries. It is similar to a Bloom Filter but allows for easy removal of elements and supports efficient dynamic resizing. Cuckoo Filters use cuckoo hashing to store elements and can handle insertions, deletions, and lookups efficiently.

## Hash Functions for Cuckoo Filters

Cuckoo Filters typically use two hash functions to map input elements to two possible locations in the filter's internal buckets. During insertion, if any of the target locations is already occupied, the filter performs a "cuckoo kick" operation to evict the existing element and place the new element in its place.

Choosing hash functions for a Cuckoo Filter involves considering factors such as distribution, performance, and collision resistance. Here's a step-by-step guide to help you choose appropriate hash functions:

`Distribution and Independence`: The hash functions should provide a uniform distribution of hash values for a wide range of input data. Also, they should be as independent as possible, meaning that the hash values for different inputs should not correlate strongly. This helps prevent clustering of elements in the filter buckets.

`Two Hash Functions`: Cuckoo Filters typically use two hash functions to determine the two possible locations for an element in the filter's buckets. Choose two different hash functions that meet the distribution and independence criteria mentioned above.

`Performance`: Hash functions should be efficient to compute. The faster the hash functions are, the better the filter's performance will be. Consider hash functions that can be quickly computed using bit manipulation or simple arithmetic operations.

`Avalanche Effect`: The avalanche effect refers to the property that a small change in the input should produce significant changes in the hash output. This helps ensure that similar inputs do not produce similar hash values, which can lead to collisions. Look for hash functions that exhibit a strong avalanche effect.

`Collision Resistance`: While collisions are expected in a Cuckoo Filter due to the nature of the algorithm, you should still choose hash functions that minimize collisions. Hash functions with strong collision resistance will help reduce the frequency of evictions and improve the filter's performance.

`Cryptographic vs. Non-Cryptographic Hash Functions`: Depending on your application's requirements, you might consider using cryptographic hash functions. However, keep in mind that cryptographic hash functions are often slower than non-cryptographic ones. If speed is a priority and cryptographic strength is not necessary, non-cryptographic hash functions may be preferable.

## Current Implementation

`altHash` Global Variable:
The fingerprint length is one byte (8 bits). The altHash global variable is an array with a size of 256, which corresponds to the range of possible values for an 8-bit byte (0 to 255). This array is used to cache the hash information corresponding to each possible fingerprint value. The purpose of caching this information is to avoid having to recalculate the hash for each fingerprint value every time it's needed later in the code.

`masks` Global Variable:
The masks global variable is an array used to fetch the last n bits of a hash value. The masks array has a size of 65, which means it can store masks for hash values up to 64 bits (since the array is zero-indexed). The masks are used in subsequent operations, where only specific bits of hash values are needed. For example, the code calculates a mask as 1 << i - 1, which generates a mask with i bits set to 1. This mask is then used to extract a specific number of bits from a hash value.

## Pros

- Efficient membership queries: Cuckoo Filters provide fast lookup times for set membership checks.
- Support for deletions: Unlike Bloom Filters, Cuckoo Filters allow for the removal of elements.
- Dynamic resizing: Cuckoo Filters can be resized without needing to rebuild the entire filter.
- Relatively low false positive rate: Cuckoo Filters offer a controllable trade-off between memory usage and false positive rate.

## Cons and Mitigations

- Limited capacity: Cuckoo Filters can become full, leading to eviction loops that can cause performance issues.
  - To mitigate this, implement resizing strategies that increase the filter size when it becomes close to full.
- Higher memory usage: Cuckoo Filters might use more memory compared to other data structures due to the need for two hash values per element.
  - This trade-off is often acceptable given the benefits of deletions and dynamic resizing.

## Usage

To use this Cuckoo Filter implementation in your Go code, follow these steps:

1. Clone the repository.
2. Import the `cuckoofilter` package in your code.
3. Initialize a Cuckoo Filter with the desired parameters.
4. Add elements to the Cuckoo Filter using the `Insert` method.
5. Check for membership using the `Lookup` method.
6. Remove elements using the `Delete` method.
