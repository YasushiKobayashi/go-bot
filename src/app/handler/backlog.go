package handler

import (
	"strconv"
	"strings"
	"time"
)

type (
	Backlog struct {
		Created time.Time `json:"created"`
		Project struct {
			Archived          bool   `json:"archived"`
			ProjectKey        string `json:"projectKey"`
			Name              string `json:"name"`
			ChartEnabled      bool   `json:"chartEnabled"`
			ID                int    `json:"id"`
			SubtaskingEnabled bool   `json:"subtaskingEnabled"`
		} `json:"project"`
		ID      int `json:"id"`
		Type    int `json:"type"`
		Content struct {
			Summary     string `json:"summary"`
			KeyID       int    `json:"key_id"`
			Description string `json:"description"`
			Comment     struct {
				ID      int    `json:"id"`
				Content string `json:"content"`
			} `json:"comment"`
			ID int `json:"id"`
		} `json:"content"`
		Notifications []interface{} `json:"notifications"`
		CreatedUser   struct {
			NulabAccount struct {
				NulabID  string `json:"nulabId"`
				Name     string `json:"name"`
				UniqueID string `json:"uniqueId"`
			} `json:"nulabAccount"`
			Name        string      `json:"name"`
			MailAddress interface{} `json:"mailAddress"`
			ID          int         `json:"id"`
			RoleType    int         `json:"roleType"`
			Lang        string      `json:"lang"`
			UserID      interface{} `json:"userId"`
		} `json:"createdUser"`
	}
)

func BacklogContent(params *Backlog, backlogUrl string) (res string) {
	title := "Backlog " + setBackLogType(params.Type)
	key := params.Project.ProjectKey + strconv.Itoa(params.Content.KeyID)
	content := params.Content.Summary
	createdUser := "_by " + params.CreatedUser.Name + "_"
	message := params.Content.Comment.Content

	url := ">>> " + backlogUrl + "view/" + params.Project.ProjectKey +
		"-" + strconv.Itoa(params.Content.KeyID)

	if params.Content.Comment.ID != 0 {
		url = url + "#comment-" + strconv.Itoa(params.Content.Comment.ID)
	}

	array := []string{title, key, content, createdUser, message, url}
	res = strings.Join(array, "\n")
	return res
}

func setBackLogType(backlogType int) (res string) {
	switch backlogType {
	case 1:
		res = "課題の追加"
	case 2, 3:
		res = "課題の更新"
	case 5:
		res = "wikiの追加"
	case 6:
		res = "wikiの更新"
	case 8:
		res = "ファイルの追加"
	case 9:
		res = "ファイルの更新"
	default:
		res = "未定義"
	}
	return res
}
