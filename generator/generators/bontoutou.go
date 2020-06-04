package generators

import (
    path "path/filepath"
    "image/png"
    "strings"
    "bytes"
    "github.com/bonbot-team/bonapi/utils"
    "github.com/tfriedel6/canvas"
    Backend "github.com/tfriedel6/canvas/backend/softwarebackend"
)

type BonToutou struct {
    // Nothing
}

func(gen *BonToutou) GetName() string {
    return "bontoutou"
}

func(gen *BonToutou) Generate(args map[string][]string)([]byte, *utils.Error) {

    names, ok := args["name"]

    if !ok || len(names[0]) < 1 {
        return nil, utils.NewError(400, "Please specify name parameter")
    }

    name := names[0]

    if len(name) > 12 {
        return nil, utils.NewError(400, "Name parameter must be no more than 12 characters")
    }

    colorss, ok := args["colors"]
    var colors []string

    if !ok || len(colorss[0]) < 1 {
        colors = []string { "#4287f5" }
    } else {
        colorsStr := colorss[0]
        colors = strings.Split(colorsStr, "|")
    }
    
    colors = utils.ConvertColors(colors)
    name = strings.ToLower(name)
    
    var x, y float64
    font := path.Join("assets", "fonts", "bontoutou.ttf")
    
    backend := Backend.New(256, 256)
    backend.MSAA = 1
    
    ctx := canvas.New(backend)
    
    img, err := ctx.LoadImage(path.Join("assets", "imgs", "bontoutou.png"))
    
    if err != nil {
        return nil, utils.NewError(503, err.Error())
    }

    w, h := float64(ctx.Width()), float64(ctx.Height())
    
    ctx.DrawImage(img, 30, 15, w, h)
    
    ctx.SetFont(font, 60)
    ctx.SetFillStyle("#FFFFFF")
    ctx.SetStrokeStyle("#000000")
    ctx.SetLineWidth(2)
    
    x = 10
    y = h - 50
    
    ctx.FillText("BON", x, y)
    ctx.StrokeText("BON", x, y)
    
    ctx.SetLineWidth(4)
    ctx.SetTextBaseline(canvas.Middle)
    
    x = 1
    y = (h / 1.03) - float64(len(name))
    
    fontAspect := 80.0
    ctx.SetFont(font, fontAspect)
    
    for ctx.MeasureText(name).Width > w {
        fontAspect--
        y -= 0.35
        
        ctx.SetFont(font, fontAspect)
    }
    
    if len(colors) <= 2 {
        gradient := ctx.CreateLinearGradient(0.0, 0.0, float64(len(colors)) * 100.0 + 50.0, float64(len(colors)) * 100.0)
        offset := 0.0
        
        for _, color := range colors {
            gradient.AddColorStop(offset, color)
            offset++;
        }
        
        ctx.SetFillStyle(gradient)
        
        ctx.StrokeText(name, x, y)
        ctx.FillText(name, x, y)
    } else {
        letters := strings.Split(name, "")
        
        for len(colors) < len(letters) {
            colors = append(colors, "#4287f5")
        }
        
        for i, letter := range letters {
            if i == 0 {
                x = 1
            } else {
                if string(letter) == string("i") {
                    x += ctx.MeasureText(letter).Width * 2.0
                } else {
                    if string(letters[i - 1]) == string("i") {
                        x += ctx.MeasureText(letter).Width / 1.7
                    } else {
                        x += ctx.MeasureText(letter).Width * 1.2
                    }
                }
            }
            
            ctx.SetFillStyle(colors[i])
            
            ctx.StrokeText(letter, x, y)
            ctx.FillText(letter, x, y)
        }
    }
    
    ctx.Fill()
    ctx.Stroke()
    
    defer img.Delete()

    f := bytes.NewBuffer([]byte {})
    err = png.Encode(f, backend.Image)

    if err != nil {
        return nil, utils.NewError(503, err.Error())
    }

    bytes := f.Bytes()
    
    return bytes, nil
}
