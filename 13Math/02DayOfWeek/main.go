package main

//https://leetcode-cn.com/problems/day-of-the-week/comments/

var daysOfMonth = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var daysOfWeek = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func init() {
	for i := 1; i <= 12; i++ {
		daysOfMonth[i] += daysOfMonth[i-1]
	}
}

func dayOfTheWeek(day int, month int, year int) string {
	//查日历得知，1970-12-31是星期四，因此，算出给出日期到基准日期的天数，再加3，模7即可
	ans := 3 //参照基准所在daysOfWeek数组的起始点
	days := 0
	for i := 1971; i < year; i++ {
		if isLeapYear(i) {
			days += 366
		} else {
			days += 365
		}
	}

	//给定年加上月份之前的天数,闰年多加1天
	days += daysOfMonth[month-1]
	if isLeapYear(year) && month > 2 {
		days += 1
	}
	//加上给定年的给定月份的天数
	days += day
	return daysOfWeek[(days+ans)%7] //（距离天数+参照基准起始点） %7
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
