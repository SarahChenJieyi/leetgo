# leetgo
## Fast and Slow Double Pointer | 快慢双指针
### General Idea
- 解决环的问题，通常都使用快慢双指针
- 证明有环：慢指针走一步，快指针走两步，若最终重合，证明有环
### Examples

## Monotonic Stack | 单调栈
### General Idea
- 寻找下一个最低点或下一个最高点
    - 维护一个存储下标的单调栈
    - 保持栈中下标对应的数组值是单调递增或递减的
- 维持字典序
    - 维护字典序的单调栈
### Examples
[739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
[42. Trapping Rain Water](https://leetcode-cn.com/problems/trapping-rain-water/)