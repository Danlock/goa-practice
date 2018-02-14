package controller

import (
	"errors"
	"log"
	"time"

	"github.com/danlock/goa-practice/generated/app"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/goadesign/goa"
	"github.com/streadway/amqp"
)

// AssetController implements the asset resource.
type AssetController struct {
	*goa.Controller
	mgoSesh  *mgo.Session
	amqpConn *amqp.Connection
}

// NewAssetController creates a asset controller.
func NewAssetController(service *goa.Service, mgoSesh *mgo.Session, amqpConn *amqp.Connection) *AssetController {
	return &AssetController{Controller: service.NewController("AssetController"), mgoSesh: mgoSesh, amqpConn: amqpConn}
}

// Delete runs the delete action.
func (c *AssetController) Delete(ctx *app.DeleteAssetContext) error {
	// AssetController_Delete: start_implement

	// Put your logic here

	return ctx.OK(nil)
	// AssetController_Delete: end_implement
}

// Show runs the show action.
func (c *AssetController) Show(ctx *app.ShowAssetContext) error {
	// AssetController_Show: start_implement

	// Put your logic here

	res := &app.Asset{}
	return ctx.OK(res)
	// AssetController_Show: end_implement
}

// ShowAll runs the showAll action.
func (c *AssetController) ShowAll(ctx *app.ShowAllAssetContext) error {
	// AssetController_ShowAll: start_implement
	var assets app.AssetCollection
	err := c.mgoSesh.DB("test").C("asset").Find(nil).All(&assets)
	if err != nil {
		log.Printf("Error finding all assets from mongo!\n%s", err)
		return err
	}

	return ctx.OK(assets)
	// AssetController_ShowAll: end_implement
}

// Create runs the Create action.
func (c *AssetController) Create(ctx *app.CreateAssetContext) error {
	// AssetController_Create: start_implement
	now := time.Now()
	assetID := bson.NewObjectId()
	assetName := ctx.Payload.Name
	assetType := ctx.Payload.Type
	assetData := ctx.Payload.Data
	asset := app.Asset{
		ID:        &assetID,
		Name:      &assetName,
		Type:      &assetType,
		Data:      &assetData,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	err := c.mgoSesh.DB("test").C("asset").Insert(&asset)
	if err != nil {
		log.Printf("Error inserting asset in mongo!\n%s", err)
		return err
	}

	return ctx.OK(&asset)
	// AssetController_Create: end_implement
}

// Update runs the update action.
func (c *AssetController) Update(ctx *app.UpdateAssetContext) error {
	// AssetController_Update: start_implement
	now := time.Now()

	if !bson.IsObjectIdHex(ctx.AssetID) {
		log.Printf("AssetID %s is not an ObjectID!", ctx.AssetID)
		return ctx.NotFound()
	}
	assetID := bson.ObjectIdHex(ctx.AssetID)

	asset := app.Asset{
		Name:      ctx.Payload.Name,
		Type:      ctx.Payload.Type,
		Data:      ctx.Payload.Data,
		UpdatedAt: &now,
	}

	err := c.mgoSesh.DB("test").C("asset").Update(bson.M{"_id": &assetID}, bson.M{"$set": &asset})
	if err != nil {
		log.Panicf("Error updating asset %s!\n%s", ctx.AssetID, err)
		return errors.New("")
	}

	return ctx.OK(nil)
	// AssetController_Update: end_implement
}
