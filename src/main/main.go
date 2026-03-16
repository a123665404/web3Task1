package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//自我介绍
	selfIntroduction()

	//题目1
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("题目1：", removeDuplicates(nums))

	//题目2
	b := isValid("([{}])")
	fmt.Println("题目2：", b)

	//题目3
	intervals := [][]int{{1, 4}, {0, 4}}
	result := merge(intervals)
	fmt.Println("题目3：", result)

	//题目4
	one := numberAppearsOne([]int{2, 2, 1})
	fmt.Println("题目4：", one)

	//题目5
	fmt.Println("题目5：", isPalindrome1(11011))

	//题目6
	digits := []int{9, 9, 9, 9, 9}
	fmt.Println("题目6：", getResult(digits))

	//题目7
	testCases := []struct {
		input  []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"apple", "apple", "apple"}, "apple"},
		{[]string{"", "abc", "def"}, ""},
		{[]string{"prefix", "preface", "preform"}, "pref"},
		{[]string{"go", "golang", "gopher"}, "go"},
		{[]string{"single"}, "single"},
		{[]string{}, ""},
	}

	for i, tc := range testCases {
		result := longestCommonPrefix(tc.input)
		status := "✓"
		if result != tc.output {
			status = "✗"
		}
		fmt.Printf("题目7测试用例 %d: %-5s 输入: %-30v 输出: %-10s 预期: %-10s\n",
			i+1, status, tc.input, result, tc.output)
	}

	//题目8
	nums1 := []int{2, 7, 11, 15}
	target := 26
	result1 := twoSum(nums1, target)
	fmt.Println("题目8：", result1)
}

// region题目1：给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
func removeDuplicates(nums []int) int {
	// 判断数组是否为空，空数组返回0
	if len(nums) == 0 {
		return 0
	}
	//核心思想是遍历数组，如果当前元素和前一个元素相等，则跳过，否则将当前元素赋给前一个元素
	i := 0 // 记录当前位置
	// 从第二位开始，遍历数组
	for j := 1; j < len(nums); j++ {
		// 如果当前元素和前一个元素不等，则将当前元素赋给前元素
		if nums[i] != nums[j] {
			//需要同时更新i和nums[i]
			i++
			nums[i] = nums[j]
		}
	}
	// 返回新数组的长度，因为i从0开始，所以需要+1
	return i + 1
}

//endregion

//----------------------------------------------------------

//region题目2：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

func isValid(s string) bool {

	// 如果字符串长度为奇数，则不可能所有括号都有匹配的另一半。
	n := len(s)
	if n%2 == 1 || n == 0 {
		return false
	}

	// stack 用作存储遇到的开启括号，以便稍后与闭合括号匹配。
	stack := []byte{}
	// 遍历字符串中的每个字符。
	for i := 0; i < n; i++ {
		//由于已知的只有三种符号，因此这里暂且认为只有三种符号。
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			//左括号压入栈中
			stack = append(stack, s[i])
		} else {
			// 栈不为空，并且栈顶的开启括号与当前闭合括号匹配。
			if len(stack) == 0 {
				return false
			}
			//下面的代码是为了判断栈顶的开启括号是否与当前闭合括号匹配。
			if stack[len(stack)-1] != '(' && s[i] == ')' {
				return false
			}
			if stack[len(stack)-1] != '[' && s[i] == ']' {
				return false
			}
			if stack[len(stack)-1] != '{' && s[i] == '}' {
				return false
			}
			// 移除栈顶的开启括号，因为它已经找到了匹配的闭合括号。
			stack = stack[:len(stack)-1]
		}
	}
	// 如果栈为空，说明所有开启括号都找到了匹配的闭合括号。
	// 如果栈不为空，则说明有开启括号没有找到匹配的闭合括号。
	return len(stack) == 0
}

/*
*
方案2
*/
func isValid1(s string) bool {
	// 如果字符串长度为奇数，则不可能所有括号都有匹配的另一半。
	n := len(s)
	if n%2 == 1 {
		return false
	}
	// pairs 映射定义了每种闭合括号对应的开启括号。
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// stack 用作存储遇到的开启括号，以便稍后与闭合括号匹配。
	stack := []byte{}
	for i := 0; i < n; i++ {
		// 如果当前字符是一个闭合括号，并且在pairs映射中找到了对应的开启括号。
		if pairs[s[i]] > 0 {
			// 检查栈是否为空或者栈顶的开启括号是否与当前闭合括号匹配。
			// 如果不匹配，或者栈为空，则字符串中的括号无效。
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 从栈中移除栈顶的开启括号，因为它已经找到了匹配的闭合括号。
			stack = stack[:len(stack)-1]
		} else {
			// 如果当前字符是一个开启括号，将其添加到栈中。
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

//endregion

//----------------------------------------------------------

//region题目3：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间

/*
*
输入：intervals = [[1,4],[0,4]]，能得到一个结果：[[0,4]]
*/
func merge(intervals [][]int) [][]int {
	// 对intervals进行排序，以便后续处理
	// 这里的排序是基于每个区间起始位置的升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 遍历排序后的intervals，尝试合并重叠的区间
	for i := 0; i < len(intervals)-1; i++ {
		// 如果当前区间的结束位置大于等于下一个区间的起始位置，说明它们有重叠
		if intervals[i][1] >= intervals[i+1][0] {
			// 合并这两个重叠的区间
			// 通过取两个重叠区间中结束位置较大的值作为合并后区间的结束位置
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			// 从intervals中移除被合并的区间
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			// 由于合并了一个区间，为了不跳过对合并后区间的检查，需要将索引i减1
			i--
		}
	}

	// 返回合并重叠区间后的intervals
	return intervals

}

/*
*
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]，能得到一个结果：[[1,6],[8,10],[15,18]]
*/
func merge1(intervals [][]int) [][]int {
	// 如果区间列表为空，则直接返回一个空的区间列表。
	// 这是因为没有区间需要合并，所以结果也是空的。
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化结果列表，将排序后的第一个区间加入其中。
	// 这是因为第一个区间是排序后起始位置最小的，可以作为合并的起点。
	result := [][]int{intervals[0]}

	// 遍历排序后的区间列表，从第二个区间开始。
	// 从第二个区间开始遍历是因为第一个区间已经作为初始值加入到结果列表中。
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的起始位置小于等于结果列表中最后一个区间的结束位置，
		// 则说明这两个区间有重叠，需要进行合并。
		if intervals[i][0] <= result[len(result)-1][1] {
			// 合并区间时，更新结果列表中最后一个区间的结束位置为当前区间结束位置和最后一个区间结束位置的较大值。
			// 这样可以确保合并后的区间覆盖了原来两个区间的范围。
			result[len(result)-1][1] = max(intervals[i][1], result[len(result)-1][1])
		} else {
			// 如果当前区间与结果列表中最后一个区间没有重叠，则将当前区间直接添加到结果列表中。
			// 这是因为当前区间是一个新的、不与之前任何区间重叠的区间。
			result = append(result, intervals[i])
		}
	}

	// 返回合并后的区间列表。
	// 经过上述处理，所有重叠的区间已经被合并，非重叠的区间被保留，因此返回的结果是合并后的区间列表。
	return result

}

//endregion

//----------------------------------------------------------

//region题目4：给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。请找到数组中只出现了一次的元素。

/*
*
使用异或运算
*/
func numberAppearsOne(nums []int) int {
	//任何数与自身异或结果为 0，任何数与 0 异或结果还是它本身。
	result := 0
	// 遍历数组，对所有元素进行异或操作
	for _, num := range nums {
		result ^= num
	}
	return result
}

/*
*
使用map对象存储数字出现的次数，最终选择出次数为1的数字
*/
func numberAppearsOne2(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

//endregion

//----------------------------------------------------------

//region 题目5: 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。 例如，121 是回文，而 123 不是。

/*
*
思路：直接反转后半部分数字，然后与前半部分比较 这个是数学的方法
*/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 100
	}
	return x == revertedNumber || x == revertedNumber/10
}

/*
*
思路：将数字转为字符串，然后使用字符串反转方法
*/
func isPalindrome1(num int) bool {
	//使用转成字符串，然后将字符串反转后进行比较
	str := strconv.Itoa(num)
	// 使用字符串反转方法
	reversed := ReverseString(str)
	return str == reversed
}

func ReverseString(s string) string {
	runes := []rune(s) // 转换为 rune 切片以支持 Unicode
	//声明一个新的切片用于存储反转后的字符
	var runes1 []rune
	for i := len(runes) - 1; i >= 0; i-- {
		runes1 = append(runes1, runes[i])
	}
	//返回反转后的字符串
	return string(runes1)
}

//endregion

//----------------------------------------------------------

//region 题目6：给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。 将大整数加 1，并返回结果的数字数组。

//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//加 1 后得到 123 + 1 = 124。
//因此，结果应该是 [1,2,4]。

func getResult(digits []int) []int {
	//思路：倒序遍历数组，如果当前位是9，则置为0，当下次循环时，依旧判断是否是9，以此类推，当不是9时时，则加1，并返回结果
	//当然还有种情况就是数组全部是9，则需要进位，此时需要创建一个长度为len(digits)+1的数组，并把第一个元素置为1
	for i := len(digits) - 1; i >= 0; i-- {
		//如果当前位是9，则置为0，因为加1后，需要进位
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			//进一位时，直接加1，然后返回
			digits[i]++
			return digits
		}
	}
	//如果所有位都为9，则需要进位，将数组的长度增加1，并把第一个元素置为1
	digits = append([]int{1}, digits...)
	return digits
}

//endregion

//----------------------------------------------------------

//region题目7：编写一个函数来查找字符串数组中的最长公共前缀。

//如果不存在公共前缀，返回空字符串 ""。
//输入：strs = ["flower","flow","flight"]
//输出："fl"

func longestCommonPrefix(strs []string) string {
	// 如果输入的字符串数组为空，则直接返回空字符串作为结果
	if len(strs) == 0 {
		return ""
	}
	// 初始化前缀为第一个字符串，作为后续比较的基础
	prefix := strs[0]

	// 遍历字符串数组，从第二个字符串开始，以找到共同的前缀
	for i := 1; i < len(strs); i++ {
		// 对当前前缀中的每个字符进行检查
		for j := 0; j < len(prefix); j++ {
			// 如果当前字符串的长度小于前缀长度，或字符不匹配，则当前前缀截止到当前位置
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	// 返回最终确定的共同前缀
	return prefix
}

/*
*
使用二分查找法 （推荐）
思路：对可能的前缀长度进行二分查找，验证中间长度是否有效
时间复杂度：O(S×logM)，其中 M 是最短字符串长度
空间复杂度：O(1)
适用场景：字符串非常长或性能要求高的情况
*/
func binarySearch(strs []string) string {
	//minLength := len(strs[0])
	//for i := 1; i < len(strs); i++ {
	//	if len(strs[i]) < minLength {
	//		minLength = len(strs[i])
	//	}
	//}
	//for i := 0; i < minLength; i++ {
	//	for j := 1; j < len(strs); j++ {
	//		if strs[j][i] != strs[0][i] {
	//			return strs[0][:i]
	//		}
	//	}
	//}
	//return strs[0][:minLength]

	if len(strs) == 0 {
		return ""
	}

	// 找到最短字符串的长度
	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	// 二分查找的左右边界
	low, high := 1, minLen
	for low <= high {
		mid := (low + high) / 2
		if hasCommonPrefix(strs, mid) {
			low = mid + 1 // 可能还有更长的前缀
		} else {
			high = mid - 1 // 前缀太长了
		}
	}

	// 返回最长有效前缀
	return strs[0][:(low+high)/2]
}

// 辅助函数：检查所有字符串是否具有指定长度的公共前缀
func hasCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			return false
		}
	}
	return true
}

//endregion

//----------------------------------------------------------------------

//region题目8：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

//你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。 你可以按任意顺序返回答案。

// selfIntroduction 打印开发者的自我介绍信息。
func selfIntroduction() {
	fmt.Println("========== 自我介绍 ==========")
	fmt.Println("姓名：张军辉（zhangjh）")
	fmt.Println("GitHub：a123665404")
	fmt.Println("邮箱：2244819892@qq.com")
	fmt.Println("简介：热爱编程，专注于 Web3 及区块链技术，熟悉 Go 语言开发，对算法与数据结构有浓厚兴趣。")
	fmt.Println("==============================")
	fmt.Println()
}

// twoSum 函数旨在找出数组中两个数之和等于特定目标值的那两个数的索引。
// 它通过两层循环遍历数组中的每个元素，并检查其他元素是否能与当前元素相加得到目标值。
// 参数 nums 是一个整数数组，target 是需要两数之和达到的目标值。
// 返回值是包含这两个数的索引的切片，如果找不到这样的两个数，则返回空切片。
func twoSum(nums []int, target int) []int {
	// 外层循环遍历数组中的每个元素
	for i := 0; i < len(nums); i++ {
		// 内层循环从当前元素的下一个元素开始，避免重复计算和自身相加
		for j := i + 1; j < len(nums); j++ {
			// 检查当前两个元素的和是否等于目标值
			if nums[i]+nums[j] == target {
				// 如果找到符合条件的两个数，返回它们的索引
				return []int{i, j}
			}
		}
	}
	// 如果遍历完数组后没有找到符合条件的两个数，返回空切片
	return []int{}
}

//endregion
