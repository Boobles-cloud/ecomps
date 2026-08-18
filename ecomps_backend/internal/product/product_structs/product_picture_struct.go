package productstructs

type ProductPictures struct {
	PictureId       uint   `json:"PictureId"`
	PicturePath     string `json:"PicturePath"`
	PicturePosition uint   `json:"PicturePosition"`
	ProductId       uint   `json:"ProductId"`
}
