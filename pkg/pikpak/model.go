package pikpak

import "time"

type (
	ErrResult struct {
		ErrorCode int    `json:"error_code"`
		Error     string `json:"error"`
	}
	TokenResult struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	Files struct {
		Files         []File `json:"files"`
		NextPageToken string `json:"next_page_token"`
	}
	File struct {
		ID             string    `json:"id"`
		Kind           string    `json:"kind"`
		Name           string    `json:"name"`
		ModifiedTime   time.Time `json:"modified_time"`
		Size           string    `json:"size"`
		ThumbnailLink  string    `json:"thumbnail_link"`
		WebContentLink string    `json:"web_content_link"`
		Medias         []Media   `json:"medias"`
	}
	Media struct {
		MediaId   string `json:"media_id"`
		MediaName string `json:"media_name"`
		Video     struct {
			Height     int    `json:"height"`
			Width      int    `json:"width"`
			Duration   int    `json:"duration"`
			BitRate    int    `json:"bit_rate"`
			FrameRate  int    `json:"frame_rate"`
			VideoCodec string `json:"video_codec"`
			AudioCodec string `json:"audio_codec"`
			VideoType  string `json:"video_type"`
		} `json:"video"`
		Link struct {
			Url    string    `json:"url"`
			Token  string    `json:"token"`
			Expire time.Time `json:"expire"`
		} `json:"link"`
		NeedMoreQuota  bool          `json:"need_more_quota"`
		VipTypes       []interface{} `json:"vip_types"`
		RedirectLink   string        `json:"redirect_link"`
		IconLink       string        `json:"icon_link"`
		IsDefault      bool          `json:"is_default"`
		Priority       int           `json:"priority"`
		IsOrigin       bool          `json:"is_origin"`
		ResolutionName string        `json:"resolution_name"`
		IsVisible      bool          `json:"is_visible"`
		Category       string        `json:"category"`
	}
	DownloadUrl struct {
		Url      string
		Filename string
	}
)
