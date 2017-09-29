package backlogs

import (
	"app/config"
	"app/handler"
	"net/http"

	"github.com/labstack/echo"
)

/*
Post
	バックログのwebhookを受取、整形したものをslackのwebhookに流す
*/
func Post() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		channel := c.Param("channel")
		slackUrl := c.QueryParam("slack_url")
		backlogUrl := c.QueryParam("backlog_url")

		post := new(handler.Backlog)
		if err = c.Bind(post); err != nil {
			return c.JSON(http.StatusConflict, handler.HandleErrorRes("conflict", err))
		}

		sendCont := handler.BacklogContent(post, backlogUrl)
		slack := slackSetting(channel, sendCont)
		err = handler.SendSlack(slack, slackUrl)
		if err != nil {
			return c.JSON(http.StatusConflict, handler.HandleErrorRes("conflict", err))
		}
		return c.JSON(http.StatusOK, slack)
	}
}

func slackSetting(channel string, text string) (res handler.Slack) {
	res = handler.Slack{
		Channel:  channel,
		Username: config.BACKLOG_SLACK_USERNAME,
		IconUrl:  config.BACKLOG_SLACK_ICON,
		Text:     text,
	}
	return res
}
