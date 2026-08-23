package productstructs

type ProductPictures struct {
	PictureId       uint   `json:"PictureId"`
	PictureName     string `json:"PictureName"`
	PicturePath     string `json:"-"`
	PicturePosition uint   `json:"PicturePosition"`
	ProductId       uint   `json:"ProductId"`
	TenantId        uint   `json:"-"` // The tenant id gets set by the backend
}
