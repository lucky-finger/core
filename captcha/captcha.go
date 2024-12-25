package captcha

import (
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
)

var slideCapt slide.Captcha

func Init() error {
	builder := slide.NewBuilder(
		slide.WithEnableGraphVerticalRandom(true),
	)

	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		return err
	}

	newGraphs := make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			ShadowImage:  graph.ShadowImage,
			MaskImage:    graph.MaskImage,
		})
	}

	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	slideCapt = builder.Make()
	return nil
}

func Generate() (slide.CaptchaData, error) {
	return slideCapt.Generate()
}
