package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacobintern/meme_coin/controller/command"
	"github.com/jacobintern/meme_coin/pkg/errors"
	"github.com/jacobintern/meme_coin/service"
)

type MemeCoinController struct {
	memeCoinService service.IMemeCoinService
}

func NewMemeCoinController(memeCoinService service.IMemeCoinService) *MemeCoinController {
	return &MemeCoinController{memeCoinService: memeCoinService}
}

// Create Meme Coin godoc
//
//	@Summary		Initialized Meme Coin
//	@Description	Initialized Meme Coin with the provided details.
//	@Tags			MemeCoin
//	@Accept			json
//	@Produce		json
//	@Param			request	body		command.CreateMemeCoinCommand	true	"Meme Coin Creation Command"
//	@Success		200		{object}	controller.Res					"Successfully initialized meme coin information"
//	@Failure		400		{object}	controller.Res					"Invalid request or missing required parameters"
//	@Failure		500		{object}	controller.Res					"Internal server error"
//	@Router			/meme_coin [post]
func (c *MemeCoinController) Create(ctx *gin.Context) {
	command := command.NewCreateMemeCoinCommand()
	if err := ctx.ShouldBindJSON(&command); err != nil {
		Failed(ctx, errors.InvalidArgument.Wrap(err))
		return
	}

	if err := c.memeCoinService.Create(command); err != nil {
		Failed(ctx, err)
		return
	}

	OK(ctx, nil)
}

// Get Meme Coin godoc
//
//	@Summary		fetched Meme Coin
//	@Description	fetched Meme Coin with the provided details.
//	@Tags			MemeCoin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Meme Coin ID"
//	@Success		200	{object}	controller.Res	"Successfully fetched meme coin information"
//	@Failure		400	{object}	controller.Res	"Invalid request or missing required parameters"
//	@Failure		500	{object}	controller.Res	"Internal server error"
//	@Router			/meme_coin/{id} [get]
func (c *MemeCoinController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		Failed(ctx, errors.InvalidParams)
		return
	}

	time.Sleep(7 * time.Second)
	resp, err := c.memeCoinService.Get(id)
	if err != nil {
		Failed(ctx, err)
		return
	}

	OK(ctx, resp)
}

// Patch Meme Coin godoc
//
//	@Summary		Patch Meme Coin
//	@Description	Patch Meme Coin with the provided details.
//	@Tags			MemeCoin
//	@Accept			json
//	@Produce		json
//	@Param			request	body		command.PatchMemeCoinCommand	true	"Meme Coin Patch Command"
//	@Param			id		path		string							true	"Meme Coin ID"
//	@Success		200		{object}	controller.Res					"Successfully patch meme coin description"
//	@Failure		400		{object}	controller.Res					"Invalid request or missing required parameters"
//	@Failure		500		{object}	controller.Res					"Internal server error"
//	@Router			/meme_coin/{id} [patch]
func (c *MemeCoinController) Patch(ctx *gin.Context) {
	command := command.NewPatchMemeCoinCommand()
	if err := ctx.ShouldBindJSON(&command); err != nil {
		Failed(ctx, errors.InvalidArgument.Wrap(err))
		return
	}

	id := ctx.Param("id")
	if id == "" {
		Failed(ctx, errors.InvalidParams)
		return
	}
	command.ID = id

	if err := c.memeCoinService.Patch(command); err != nil {
		Failed(ctx, err)
		return
	}

	OK(ctx, nil)
}

// Delete Meme Coin godoc
//
//	@Summary		Delete Meme Coin
//	@Description	Delete Meme Coin with the provided details.
//	@Tags			MemeCoin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Meme Coin ID"
//	@Success		200	{object}	controller.Res	"Successfully deleted meme coin"
//	@Failure		400	{object}	controller.Res	"Invalid request or missing required parameters"
//	@Failure		500	{object}	controller.Res	"Internal server error"
//	@Router			/meme_coin/{id} [delete]
func (c *MemeCoinController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		Failed(ctx, errors.InvalidParams)
		return
	}

	if err := c.memeCoinService.Delete(id); err != nil {
		Failed(ctx, err)
		return
	}

	OK(ctx, nil)
}

// Meme Coin poke godoc
//
//	@Summary		Meme Coin Poke
//	@Description	Add Meme Coin popularity with the provided details.
//	@Tags			MemeCoin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Meme Coin ID"
//	@Success		200	{object}	controller.Res	"Successfully count add on on meme coin popularity"
//	@Failure		400	{object}	controller.Res	"Invalid request or missing required parameters"
//	@Failure		500	{object}	controller.Res	"Internal server error"
//	@Router			/meme_coin/{id}/poke [post]
func (c *MemeCoinController) Poke(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		Failed(ctx, errors.InvalidParams)
		return
	}

	if err := c.memeCoinService.Poke(id); err != nil {
		Failed(ctx, err)
		return
	}

	OK(ctx, nil)
}
