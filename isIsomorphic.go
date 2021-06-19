package main

// 给定两个字符串 s 和 t，判断它们是否是同构的。 输入：s = "egg", t = "add"
//输出：true
//示例 2：

//输入：s = "foo", t = "bar"
//输出：false

// 思路：遍历两个字符串，并统计每个字母出现的顺序。两个字符串必须出现字母数量一致，每个位置的字母出现顺序也要一致
// 击败 88% go提交
func isIsomorphic(s string, t string) bool {

	r1 := []rune(s)
	r2 := []rune(t)
	if len(r1) != len(r2) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for i:=0;i<len(r1);i++{
		seq1,ok1 := map1[r1[i]]
		if !ok1{
			map1[r1[i]] = len(map1)+1
			seq1 = len(map1)+1
		}
		seq2,ok2 := map2[r2[i]]
		if !ok2{
			map2[r2[i]] = len(map2)+1
			seq2 = len(map2)+1
		}

		if seq1 != seq2{
			return false
		}
	}
	return true
}