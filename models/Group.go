package models

type Group struct {
	Name string	`json:"name" binding:"required"`
	Description string `json:"description"`
	Hidden bool `json:"hidden"`
	OwnerUid int64 `json:"ownerUid"`
}
