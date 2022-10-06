package freshservice

import "time"

// Assets holds a list of Freshservice asset details
type Assets struct {
	List []AssetDetails `json:"assets"`
}

// Asset holds the details of a specific Freshservice asset
type Asset struct {
	Details AssetDetails `json:"asset"`
}

// AssetDetails are the details related to a specific asset in Freshservice
type AssetDetails struct {
	ID           int       `json:"id"`
	DisplayID    int       `json:"display_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	AssetTypeID  int       `json:"asset_type_id"`
	Impact       string    `json:"impact"`
	AuthorType   string    `json:"author_type"`
	UsageType    string    `json:"usage_type"`
	AssetTag     string    `json:"asset_tag"`
	UserID       int64     `json:"user_id"`
	LocationID   int64     `json:"location_id"`
	DepartmentID int64     `json:"department_id"`
	AgentID      int64     `json:"agent_id"`
	AssignedOn   time.Time `json:"assigned_on"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// AssetListOptions holds the available options that can be
// passed when requesting a list of Freshservice assets
type AssetListOptions struct {
	PageQuery string
	SortBy    *SortOptions
	Embed     *AssetEmbedOptions
}

// AssetEmbedOptions will optonally embed desired metadata in an asset list response
// Each include will consume an additional 2 credits. For example if you embed the stats
// information you will be charged a total of 3 API credits (1 credit for the API call,
// and 2 credits for the additional stats embedding).
type AssetEmbedOptions struct {
	TypeFields bool
	Trashed    bool
}
