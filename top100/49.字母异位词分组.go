/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	m := make(map[string][]string, 0)
	for _, v := range strs {
		t := []byte(v)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		str := string(t)
		m[str] = append(m[str], v)
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

// @lc code=end

