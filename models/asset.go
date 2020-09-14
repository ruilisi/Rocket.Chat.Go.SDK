package models

type AssetUnSetRequest struct {
	AssetName         string `json:"assetName"`
	RefreshAllClients string `json:"refreshAllClients"`
}
