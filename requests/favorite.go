package requests

type UserFavoriteGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type FavoritePostRequest struct {
	BookID string `form:"book_id" binding:"required"`
}