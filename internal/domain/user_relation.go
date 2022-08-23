package domain

type UserRelation struct {
	UserID         string `bson:"userId" json:"userId"`
	UserRelationId string `bson:"userRelationId" json:"userRelationId"`
}
