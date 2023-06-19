package dupan

type (
	ErrResult struct {
		ErrNo   int    `json:"errno"`
		ShowMsg string `json:"show_msg"`
	}
	PResultErr struct {
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}
	ShareRecordInfo struct {
		ShareID         int64   `json:"shareId"`
		FsIds           []int64 `json:"fsIds"`
		ShortLink       string  `json:"shortlink"`
		Status          int     `json:"status"`
		Public          int     `json:"public"`
		TypicalCategory int     `json:"typicalCategory"`
		TypicalPath     string  `json:"typicalPath"`
		ExpireType      int     `json:"expiredType"`
		ExpireTime      int64   `json:"expiredTime"`
		ViewCount       int     `json:"vCnt"`
	}
	ShareURLInfo struct {
		Pwd      string `json:"pwd"`
		ShortURL string `json:"shorturl"`
	}
	Shared struct {
		Link    string `json:"link"`
		Pwd     string `json:"pwd"`
		ShareID int64  `json:"shareid"`
	}
	BlockListJSON struct {
		BlockList []string `json:"block_list"`
	}
	FileDirectory struct {
		FsID     int64  `json:"fs_id"`
		AppID    int64  `json:"app_id"`
		Path     string `json:"path"`
		Filename string `json:"server_filename"`
		Ctime    int64  `json:"ctime"`
		Mtime    int64  `json:"mtime"`
		BlockListJSON
		Size           int64 `json:"size"`
		IsDirInt       int8  `json:"isdir"`
		IfHasSubDirInt int8  `json:"ifhassubdir"`
	}
	FileDirectoryList []*FileDirectory
)
