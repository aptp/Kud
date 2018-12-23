package entity

import (
	"errors"
	"fmt"
)

type Reply struct {
	UserName      string
	Contributions []int
}

func NewReply(userName string, c []int) *Reply {

	return &Reply{
		UserName:      userName,
		Contributions: c,
	}
}

func (r *Reply) MakeTodaysContributions() (string, error) {

	if len(r.Contributions) != 7 {
		// TODO: make Error define.
		return "", errors.New("Error")
	}

	tc := r.Contributions[len(r.Contributions)-1]
	today := fmt.Sprintf("%s さんの今日のコントリビューションは %d ! ", r.UserName, tc)

	var wc int
	for i := range r.Contributions {
		wc += r.Contributions[i]
	}
	week := fmt.Sprintf("ここ１週間だと %d contributions。", wc)

	switch {
	case wc >= 70:
		week = week + "いい感じですね！"
	case (70 > wc && wc >= 55):
		week = week + "まあまあですね。"
	case (55 > wc && wc >= 30):
		week = week + "微妙ですね... 頑張ってください！"
	case 30 > wc:
		week = week + "ダメダメですね。何やってるんですか。"
	}

	message := fmt.Sprintf("%s\n %s\n", today, week)

	return message, nil
}
