package option

type ClashAPIOptions struct {
	ExternalController       string `json:"external_controller,omitempty"`
	ExternalUI               string `json:"external_ui,omitempty"`
	ExternalUIDownloadURL    string `json:"external_ui_download_url,omitempty"`
	ExternalUIDownloadDetour string `json:"external_ui_download_detour,omitempty"`
	Secret                   string `json:"secret,omitempty"`
	DefaultMode              string `json:"default_mode,omitempty"`
	StoreMode                bool   `json:"store_mode,omitempty"`
	StoreSelected            bool   `json:"store_selected,omitempty"`
	StoreFakeIP              bool   `json:"store_fakeip,omitempty"`
	CacheFile                string `json:"cache_file,omitempty"`
	CacheID                  string `json:"cache_id,omitempty"`

	ModeList []string `json:"-"`
}

type SelectorOutboundOptions struct {
	Outbounds                 []string `json:"outbounds"`
	Default                   string   `json:"default,omitempty"`
	InterruptExistConnections bool     `json:"interrupt_exist_connections,omitempty"`
}

type URLTestOutboundOptions struct {
	Outbounds                 []string `json:"outbounds"`
	URL                       string   `json:"url,omitempty"`
	Interval                  Duration `json:"interval,omitempty"`
	Tolerance                 uint16   `json:"tolerance,omitempty"`
	InterruptExistConnections bool     `json:"interrupt_exist_connections,omitempty"`
}

type ProviderOutboundOptions struct {
	ProviderType    string           `json:"provider_type"`
	Url             Listable[string] `json:"url,omitempty"`
	Path            Listable[string] `json:"path,omitempty"`
	Default         string           `json:"default,omitempty"`
	Interval        string           `json:"interval"`
	Policy          string           `json:"policy"`
	UrlTest         *UrlTest         `json:"url_test"`
	IncludeKeyWords Listable[string] `json:"include_key_words"`
	ExcludeKeyWords Listable[string] `json:"exclude_key_words"`
}

type UrlTest struct {
	Url       string `json:"url,omitempty"`
	Interval  string `json:"interval"`
	Tolerance int    `json:"tolerance"`
}
