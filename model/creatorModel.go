package model

//创建者
type Creator struct {
	CreatorId   int64  `json:"creator_id" gorm:"default:0"` //创建者ID
	CreatorName string `json:"creator_name" gorm:"size:50"` //创建者名称
}

//修改者
type Modifier struct {
	ModifierId   int64  `json:"modifier_id" gorm:"default:0"` //修改者ID
	ModifierName string `json:"modifier_name" gorm:"size:50"` //修改者名称
}

//创建者和修改者
type CreatorAndModifier struct {
	CreatorId  int64 `json:"creator_id" gorm:"default:0"`  //创建者ID
	ModifierId int64 `json:"modifier_id" gorm:"default:0"` //修改者ID
}
