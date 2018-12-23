package slack

var (
	ErrUnknownCommand error
	ErrFromGitHub     error

	errToSlackMessage = map[error]string{
		ErrUnknownCommand: "？知らないコマンドだよ",
		ErrFromGitHub:     "GitHubからエラーが返ってきたよ...",
	}
	notifyContributionMessage = "今日は %s contributions!"
)
