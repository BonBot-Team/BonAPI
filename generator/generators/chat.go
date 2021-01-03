package generators

import (
	"bytes"
	"github.com/bonbot-team/bonapi/utils"
	"github.com/tfriedel6/canvas"
	Backend "github.com/tfriedel6/canvas/backend/softwarebackend"
	"image/png"
	"net/http"
	path "path/filepath"
)

type Chat struct {
	// Nothing
}

func (gen *Chat) GetName() string {
	return "chat"
}

func (gen *Chat) Generate(args map[string][]string) ([]byte, *utils.Error) {

	images, ok := args["image"]

	if !ok || len(images[0]) < 1 {
		return nil, utils.NewError(400, "Please specify image parameter")
	}

	imageUrl := images[0]
	image, err := utils.DownloadImage(imageUrl)

	if err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err.Error())
	}

	backend := Backend.New(256, 256)
	backend.MSAA = 1

	ctx := canvas.New(backend)

	chatImg, err := ctx.LoadImage(path.Join("assets", "imgs", "chat.png"))

	if err != nil {
		return nil, utils.NewError(503, err.Error())
	}

	img, err := ctx.LoadImage(image)

	if err != nil {
		return nil, utils.NewError(503, err.Error())
	}

	w, h := float64(ctx.Width()), float64(ctx.Height())

	ctx.DrawImage(image, 70, 75, 128, 128)
	ctx.DrawImage(chatImg, 0, 0, w, h)

	defer chatImg.Delete()
	defer img.Delete()

	f := bytes.NewBuffer([]byte{})
	err = png.Encode(f, backend.Image)

	if err != nil {
		return nil, utils.NewError(503, err.Error())
	}

	b := f.Bytes()

	return b, nil
}
