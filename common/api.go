package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GuoFlight/ghttp"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
	"yuque_export/conf"
)

type API struct{}
type ResGetAllBooks struct {
	Data []struct {
		ID             int       `json:"id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		UserID         int       `json:"user_id"`
		DisplayType    int       `json:"display_type"`
		Name           string    `json:"name"`
		Rank           int       `json:"rank"`
		IsDefault      bool      `json:"is_default"`
		OrganizationID int       `json:"organization_id"`
		Books          []struct {
			ID           int    `json:"id"`
			Type         string `json:"type"`
			Slug         string `json:"slug"`
			Name         string `json:"name"`
			UserID       int    `json:"user_id"`
			Description  string `json:"description"`
			ItemsCount   int    `json:"items_count"`
			LikesCount   int    `json:"likes_count"`
			WatchesCount int    `json:"watches_count"`
			CreatorID    int    `json:"creator_id"`
			Abilities    struct {
				Read          bool `json:"read"`
				Update        bool `json:"update"`
				ModifySetting bool `json:"modify_setting"`
				Destroy       bool `json:"destroy"`
				Share         bool `json:"share"`
				ReadPrivate   bool `json:"read_private"`
				CreateDoc     bool `json:"create_doc"`
			} `json:"abilities"`
			Public           int       `json:"public"`
			ExtendPrivate    int       `json:"extend_private"`
			Scene            any       `json:"scene"`
			Source           any       `json:"source"`
			CreatedAt        time.Time `json:"created_at"`
			UpdatedAt        time.Time `json:"updated_at"`
			ContentUpdatedAt time.Time `json:"content_updated_at"`
			PinnedAt         any       `json:"pinned_at"`
			ArchivedAt       any       `json:"archived_at"`
			Summary          []struct {
				ID                int       `json:"id"`
				Type              string    `json:"type"`
				Title             string    `json:"title"`
				Slug              string    `json:"slug"`
				Cover             string    `json:"cover"`
				Description       string    `json:"description"`
				CustomDescription any       `json:"custom_description"`
				ContentUpdatedAt  time.Time `json:"content_updated_at"`
				CreatedAt         time.Time `json:"created_at"`
				UpdatedAt         time.Time `json:"updated_at"`
				PublishedAt       time.Time `json:"published_at"`
				FirstPublishedAt  time.Time `json:"first_published_at"`
				Meta              struct {
				} `json:"meta"`
				ReadCount       any    `json:"read_count"`
				PrivacyMigrated bool   `json:"privacy_migrated"`
				Book            any    `json:"book"`
				User            any    `json:"user"`
				LastEditor      any    `json:"last_editor"`
				Share           any    `json:"share"`
				Serializer      string `json:"_serializer"`
			} `json:"summary"`
			Status            int    `json:"status"`
			StackID           int    `json:"stack_id"`
			Rank              int    `json:"rank"`
			Layout            string `json:"layout"`
			DocViewport       string `json:"doc_viewport"`
			DocTypography     any    `json:"doc_typography"`
			Cover             any    `json:"cover"`
			CommentCount      any    `json:"comment_count"`
			ReadCount         any    `json:"read_count"`
			Original          any    `json:"original"`
			OrganizationID    int    `json:"organization_id"`
			EnableAutoPublish bool   `json:"enable_auto_publish"`
			PrivacyMigrated   bool   `json:"privacy_migrated"`
			User              struct {
				ID               int       `json:"id"`
				Type             string    `json:"type"`
				Login            string    `json:"login"`
				Name             string    `json:"name"`
				Description      any       `json:"description"`
				Avatar           string    `json:"avatar"`
				AvatarURL        string    `json:"avatar_url"`
				FollowersCount   int       `json:"followers_count"`
				FollowingCount   int       `json:"following_count"`
				Status           int       `json:"status"`
				Public           int       `json:"public"`
				Scene            any       `json:"scene"`
				CreatedAt        time.Time `json:"created_at"`
				UpdatedAt        time.Time `json:"updated_at"`
				ExpiredAt        time.Time `json:"expired_at"`
				IsPaid           bool      `json:"isPaid"`
				MemberLevel      int       `json:"member_level"`
				MemberLevelName  string    `json:"memberLevelName"`
				HasMemberLevel   bool      `json:"hasMemberLevel"`
				IsTopLevel       bool      `json:"isTopLevel"`
				IsNewbie         bool      `json:"isNewbie"`
				Profile          any       `json:"profile"`
				OrganizationUser any       `json:"organizationUser"`
				Serializer       string    `json:"_serializer"`
			} `json:"user"`
			Creator    any    `json:"creator"`
			Serializer string `json:"_serializer"`
		} `json:"books"`
		Serializer string `json:"_serializer"`
	} `json:"data"`
}
type ResGetAllDocs struct {
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
	Data []struct {
		ID                int       `json:"id"`
		SpaceID           int       `json:"space_id"`
		Type              string    `json:"type"`
		SubType           any       `json:"sub_type"`
		Title             string    `json:"title"`
		TitleDraft        any       `json:"title_draft"`
		Tag               any       `json:"tag"`
		Slug              string    `json:"slug"`
		UserID            int       `json:"user_id"`
		BookID            int       `json:"book_id"`
		LastEditorID      int       `json:"last_editor_id"`
		Cover             any       `json:"cover"`
		Description       string    `json:"description"`
		CustomDescription any       `json:"custom_description"`
		Format            string    `json:"format"`
		Status            int       `json:"status"`
		ReadStatus        int       `json:"read_status"`
		ViewStatus        int       `json:"view_status"`
		Public            int       `json:"public"`
		DraftVersion      int       `json:"draft_version"`
		CommentsCount     int       `json:"comments_count"`
		LikesCount        int       `json:"likes_count"`
		ContentUpdatedAt  time.Time `json:"content_updated_at"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		PublishedAt       time.Time `json:"published_at"`
		FirstPublishedAt  time.Time `json:"first_published_at"`
		WordCount         int       `json:"word_count"`
		SelectedAt        any       `json:"selected_at"`
		PinnedAt          any       `json:"pinned_at"`
		ReadCount         any       `json:"read_count"`
		EditorMeta        string    `json:"editor_meta"`
		EditorMetaDraft   string    `json:"editor_meta_draft"`
		Region            string    `json:"region"`
		PrivacyMigrated   bool      `json:"privacy_migrated"`
		Book              any       `json:"book"`
		User              any       `json:"user"`
		LastEditor        any       `json:"last_editor"`
		Share             any       `json:"share"`
		Serializer        string    `json:"_serializer"`
		Meta              struct {
			PrivacyMigrated bool `json:"privacy_migrated"`
		} `json:"meta,omitempty"`
	} `json:"data"`
}
type ResGetDownloadLink struct {
	Data struct {
		State string `json:"state"` // "success"表示成功
		Url   string `json:"url"`   // 下载url
	} `json:"data"`
}

// GetAllBooks 获取所有book
func (a API) GetAllBooks() (ResGetAllBooks, error) {
	api := "https://www.yuque.com/api/mine/book_stacks"

	client := ghttp.Client{
		Method: http.MethodGet,
		Url:    api,
		Header: map[string]string{
			"cookie":          conf.GConf.Cookie,
			"accept-language": "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7",
			"accept":          "application/json",
		},
	}
	var ret ResGetAllBooks

	resHttp, err := client.Do()
	if err != nil {
		return ret, err
	}
	// fmt.Println(string(resHttp.Body))
	err = json.Unmarshal(resHttp.Body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// GetAllDocs 获取所有文档
func (a API) GetAllDocs(bookID int) (ResGetAllDocs, error) {
	api := "https://www.yuque.com/api/docs"

	client := ghttp.Client{
		Method: http.MethodGet,
		Url:    api,
		Header: map[string]string{
			"cookie":          conf.GConf.Cookie,
			"accept-language": "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7",
			"accept":          "application/json",
		},
		UrlArgs: map[string]string{
			"book_id": strconv.Itoa(bookID),
		},
	}
	var ret ResGetAllDocs

	resHttp, err := client.Do()
	if err != nil {
		return ret, err
	}
	// fmt.Println(string(resHttp.Body))
	err = json.Unmarshal(resHttp.Body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// GetDownloadLink 获取文档下载链接
func (a API) GetDownloadLink(docID int) (string, error) {
	api := fmt.Sprintf("https://www.yuque.com/api/docs/%d/export", docID)

	client := ghttp.Client{
		Method: http.MethodPost,
		Url:    api,
		Header: map[string]string{
			"cookie":          conf.GConf.Cookie,
			"accept-language": "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7",
			"Content-Type":    "application/json; charset=utf-8",
			"x-csrf-token":    conf.GConf.XCsrfToken,
		},
		ReqBody: `{"type":"markdown","force":0,"options":"{\"latexType\":1,\"enableAnchor\":1,\"enableBreak\":1}"}`,
	}
	var ret ResGetDownloadLink

	resHttp, err := client.Do()
	if err != nil {
		return "", err
	}
	// fmt.Println(string(resHttp.Body))
	err = json.Unmarshal(resHttp.Body, &ret)
	if err != nil {
		return "", err
	}
	if ret.Data.State != "success" {
		return "", errors.New("获取下载链接失败：" + string(resHttp.Body))
	}

	return ret.Data.Url, nil
}

func (a API) DownloadDoc(url string) error {
	client := ghttp.Client{
		Method: http.MethodGet,
		Url:    url,
		Header: map[string]string{
			"cookie": conf.GConf.Cookie,
		},
	}

	resHttp, err := client.Do()
	if err != nil {
		return err
	}
	// fmt.Println(string(resHttp.Body))

	// 从Header中获取文件名
	_, params, err := mime.ParseMediaType(resHttp.Detail.Header.Get("Content-Disposition"))
	fileName := params["filename"] // set to "foo.png"
	if fileName == "" {
		_, fileName = filepath.Split(url) // 如果Content-Disposition中没有文件名，则从URL中获取
	}
	if fileName == "" {
		return errors.New("文件名获取失败：" + url)
	}

	// 创建保存文件
	file, err := os.Create(path.Join(conf.GConf.FileDir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应的内容复制到文件
	_, err = io.WriteString(file, string(resHttp.Body))
	if err != nil {
		return err
	}

	return nil
}
