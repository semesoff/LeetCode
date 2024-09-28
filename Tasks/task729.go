type MyCalendar struct {
    bookings [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{bookings: [][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _, event := range this.bookings {
		if start < event[1] && end > event[0] {
			return false
		}
	}
	this.bookings = append(this.bookings, []int{start, end})
	return true
}



// obj := Constructor();
// param_1 := obj.Book(start,end);
