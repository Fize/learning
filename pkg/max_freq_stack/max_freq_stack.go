/*
给一个数组（例如[1,2,3,2,4,2,1]），将他们依次加入一个特殊的栈，称为频率栈。频率栈的入栈和一般的栈是一样的，但是出栈操作是特殊的：找到还在栈里的数中出现次数最多的数出栈，若有多个，找位置最靠近栈顶的数出栈。
以[1,2,3,2,4,2,1]为例，左为栈底，右为栈顶，此时依次出栈的顺序序列为：
2（因为出现了3次，此时频率栈为[1,2,3,2,4,1]）
1（此时1,2在栈中出现次数都为2，但1更接近栈顶，出栈后该频率栈为[1,2,3,2,4]）
2（此时仅有2出现了2次，出栈后该频率栈为[1,2,3,4]）
4（出栈后该频率栈为[1,2,3]）
3（出栈后该频率栈为[1,2]）
2（出栈后该频率栈为[1]）
1（出栈后该频率栈为[ ]，结束）
给定一个数组来描述该频率栈的当前状态，输出该频率栈的出栈序列。希望能最好能以O(N)的时间复杂度实现。
例如输入数组为[1,2,3,2,4,2,1]，它对应的出栈序列就是[2,1,2,4,3,2,1]
*/

package max_freq_stack

type FreqStack struct {
	Freq int
	Pos  int
}

func MaxFreqStack(list []int) []int {
	var result []int
	var fp map[int]*FreqStack
	for k, v := range list {
		fp[v] = &FreqStack{
			Freq: fp[v].Freq + 1,
		}
		if fp[v].Pos < k {
			fp[v].Pos = k
		}
	}
	return result
}