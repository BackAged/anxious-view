### Every element appears twice except for one. Find that single one
[136. Single Number](https://leetcode.com/problems/single-number/)<br>
**How**:
<ol>
  <li>Use HashMap and traverse the array twice and find the element with count 1 in map</li>
  <li>Sort and find the element</li>
  <li>Use XOR `A^A=0, A^B^A=B`</li>
</ol>

***

### [0-n] Element, Find the missing number
[268. Missing Number](https://leetcode.com/problems/missing-number/)<br>
**How**:
<ol>
  <li>Use HashMap and traverse 0-n and find which number is missing</li>
  <li>Sort and find the element</li>
  <li>Use math formula to calculate sum `n*(n+1)/2` and traverse the array and find sum and the answer is the substraction</li>
  <li>Modify the array, as the elements are positive, traverse and change arr[arr[i] *= -1, to mark it found </li>
  <li>Swap the value and index</li>
  <li>Binary search</li>
  <li>XOR- how???</li>
</ol>

***

### Find the duplicate element in array without modifying array and constant space
[287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)<br>
**How**:
<ol>
  <li>Sort and traverse</li>
  <li>Floyd's tortoise and hare algorithm</li>
  <li>Bits- how???</li>
  <li>Binary search- how??
</ol>

***
