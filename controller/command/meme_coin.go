package command

type CreateMemeCoinCommand struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}

func NewCreateMemeCoinCommand() CreateMemeCoinCommand {
	return CreateMemeCoinCommand{}
}

type PatchMemeCoinCommand struct {
	ID          string  `json:"-"`
	Description *string `json:"description"`
}

func NewPatchMemeCoinCommand() PatchMemeCoinCommand {
	return PatchMemeCoinCommand{}
}
