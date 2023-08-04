package common

type discord struct {
	Enable   bool   `json:"enable"`
	Token    string `json:"token"`
	Upload   string `json:"upload"`
	Download string `json:"download"`
}
type slack struct {
	Enable        bool   `json:"enable"`
	RTMAuthToken  string `json:"rtm_auth_token"`
	AuthToken     string `json:"auth_token"`
	AppToken      string `json:"app_token"`
	Upload        string `json:"upload"`
	Download      string `json:"download"`
	LegacyRTMMode bool   `json:"legacy_rtm_mode"`
}

type googlesheets struct {
	Enable      bool   `json:"enable"`
	Credential  string `json:"credential"`
	SheetId     string `json:"sheet_id"`
	DriveId     string `json:"drive_id"`
	ShowToken   bool   `json:"show_token"`
	EnableHTTP3 bool   `json:"enable_http3"`
}

type o365 struct {
	Enable       bool   `json:"enable"`
	TenantId     string `json:"tenant_id"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	FileName     string `json:"file_name"`
	UserId       string `json:"user_id"`
}
type debugEnable struct {
	Enable bool `json:"enable"`
}

type multiAppEnable struct {
	Enable bool `json:"enable"`
}
type uuidInfo struct {
	HostID   string
	HostName string
}

type Config struct {
	Discord  discord        `json:"discord"`
	Google   googlesheets   `json:"google"`
	O365     o365           `json:"o365"`
	Debug    debugEnable    `json:"debug"`
	MultiApp multiAppEnable `json:"multi_app"`
	Slack    slack          `json:"slack"`
	UUID     uuidInfo
}
