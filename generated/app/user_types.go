// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "practice": Application User Types
//
// Command:
// $ goagen
// --design=github.com/danlock/goa-practice/design
// --out=$(GOPATH)/src/github.com/danlock/goa-practice/generated
// --version=v1.3.1

package app

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// assetType user type.
type assetType struct {
	// Object ID attribute
	ID *bson.ObjectId `bson:"_id,omitempty" form:"_id" json:"_id,omitempty"`
	// timestamp of when the asset was created
	CreatedAt *time.Time `bson:"createdAt,omitempty" form:"createdAt" json:"createdAt,omitempty"`
	// Type specific data
	Data *interface{} `bson:"data,omitempty" form:"data" json:"data,omitempty"`
	// Name of asset
	Name *string `bson:"name,omitempty" form:"name" json:"name,omitempty"`
	// Type of asset
	Type *string `bson:"type,omitempty" form:"type" json:"type,omitempty"`
	// timestamp of when the asset was updated
	UpdatedAt *time.Time `bson:"updatedAt,omitempty" form:"updatedAt" json:"updatedAt,omitempty"`
}

// Validate validates the assetType type instance.
func (ut *assetType) Validate() (err error) {
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 3, true))
		}
	}
	if ut.Type != nil {
		if !(*ut.Type == "car") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.type`, *ut.Type, []interface{}{"car"}))
		}
	}
	return
}

// Publicize creates AssetType from assetType
func (ut *assetType) Publicize() *AssetType {
	var pub AssetType
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.Data != nil {
		pub.Data = ut.Data
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	return &pub
}

// AssetType user type.
type AssetType struct {
	// Object ID attribute
	ID *bson.ObjectId `bson:"_id,omitempty" form:"_id" json:"_id,omitempty"`
	// timestamp of when the asset was created
	CreatedAt *time.Time `bson:"createdAt,omitempty" form:"createdAt" json:"createdAt,omitempty"`
	// Type specific data
	Data *interface{} `bson:"data,omitempty" form:"data" json:"data,omitempty"`
	// Name of asset
	Name *string `bson:"name,omitempty" form:"name" json:"name,omitempty"`
	// Type of asset
	Type *string `bson:"type,omitempty" form:"type" json:"type,omitempty"`
	// timestamp of when the asset was updated
	UpdatedAt *time.Time `bson:"updatedAt,omitempty" form:"updatedAt" json:"updatedAt,omitempty"`
}

// Validate validates the AssetType type instance.
func (ut *AssetType) Validate() (err error) {
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 3, true))
		}
	}
	if ut.Type != nil {
		if !(*ut.Type == "car") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.type`, *ut.Type, []interface{}{"car"}))
		}
	}
	return
}
