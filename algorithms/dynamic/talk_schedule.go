package dynamic

import "sort"

// TalkSchedule represents a schedule of a talk.
// DMA described this algorithm on page 211, Section 3.1.5,
// and proved with mathematical induction on page 346.
type TalkSchedule struct {
	StartTime int64
	EndTime   int64
}

type ByEndTime []TalkSchedule

func (e ByEndTime) Len() int {
	return len(e)
}

func (e ByEndTime) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e ByEndTime) Less(i, j int) bool {
	return e[i].EndTime < e[j].EndTime
}

func isScheduleCompatible(talks []TalkSchedule, s TalkSchedule) bool {
	if len(talks) == 0 {
		return true
	}

	last := talks[len(talks)-1]

	return s.StartTime >= last.EndTime
}

func ScheduleTalks(talks []TalkSchedule) []TalkSchedule {
	sort.Sort(ByEndTime(talks))

	var s = make([]TalkSchedule, 0)

	for _, t := range talks {
		ok := isScheduleCompatible(s, t)
		if ok {
			s = append(s, t)
		}
	}

	return s
}
