package request_response

import "github.com/jasonstanleyyoman/doremonangis_be/entity"

type AddDorayakiRequest struct {
	Flavor      string `json:"flavor" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImagePath   string `json:"image_path"`
}

func (r AddDorayakiRequest) TransformToDorayaki() entity.Dorayaki {
	return entity.Dorayaki{
		Flavor:      r.Flavor,
		Description: r.Description,
		ImagePath:   r.ImagePath,
	}
}
