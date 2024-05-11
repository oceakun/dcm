package terminalplot

import (
	"log"
	"encoding/base64"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"

	"fmt"
	"math"
)

func Demo() {
	// Barchart()
	// Canvas()
	// Gauge()
	// Image()
	Plot()
	
}

func Barchart(){
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func Canvas() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	ui.Render(c)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func Gauge() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	g0 := widgets.NewGauge()
	g0.Title = "Slim Gauge"
	g0.SetRect(20, 20, 30, 30)
	g0.Percent = 75
	g0.BarColor = ui.ColorRed
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorCyan

	g2 := widgets.NewGauge()
	g2.Title = "Slim Gauge"
	g2.SetRect(0, 3, 50, 6)
	g2.Percent = 60
	g2.BarColor = ui.ColorYellow
	g2.LabelStyle = ui.NewStyle(ui.ColorBlue)
	g2.BorderStyle.Fg = ui.ColorWhite

	g1 := widgets.NewGauge()
	g1.Title = "Big Gauge"
	g1.SetRect(0, 6, 50, 11)
	g1.Percent = 30
	g1.BarColor = ui.ColorGreen
	g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	g1.TitleStyle.Fg = ui.ColorMagenta
	g1.BorderStyle.Fg = ui.ColorWhite

	g3 := widgets.NewGauge()
	g3.Title = "Gauge with custom label"
	g3.SetRect(0, 11, 50, 14)
	g3.Percent = 50
	g3.Label = fmt.Sprintf("%v%% (100MBs free)", g3.Percent)

	g4 := widgets.NewGauge()
	g4.Title = "Gauge"
	g4.SetRect(0, 14, 50, 17)
	g4.Percent = 50
	g4.Label = "Gauge with custom highlighted label"
	g4.BarColor = ui.ColorGreen
	g4.LabelStyle = ui.NewStyle(ui.ColorYellow)

	ui.Render(g0, g1, g2, g3, g4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}


func Image() {
	var images []image.Image
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			log.Fatalf("failed to fetch image: %v", err)
		}
		image, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatalf("failed to decode fetched image: %v", err)
		}
		images = append(images, image)
	}
	if len(images) == 0 {
		image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(GOPHER_IMAGE)))
		if err != nil {
			log.Fatalf("failed to decode gopher image: %v", err)
		}
		images = append(images, image)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	img := widgets.NewImage(nil)
	img.SetRect(0, 0, 100, 50)
	index := 0
	render := func() {
		img.Image = images[index]
		if !img.Monochrome {
			img.Title = fmt.Sprintf("Color %d/%d", index+1, len(images))
		} else if !img.MonochromeInvert {
			img.Title = fmt.Sprintf("Monochrome(%d) %d/%d", img.MonochromeThreshold, index+1, len(images))
		} else {
			img.Title = fmt.Sprintf("InverseMonochrome(%d) %d/%d", img.MonochromeThreshold, index+1, len(images))
		}
		ui.Render(img)
	}
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Left>", "h":
			index = (index + len(images) - 1) % len(images)
		case "<Right>", "l":
			index = (index + 1) % len(images)
		case "<Up>", "k":
			img.MonochromeThreshold++
		case "<Down>", "j":
			img.MonochromeThreshold--
		case "<Enter>":
			img.Monochrome = !img.Monochrome
		case "<Tab>":
			img.MonochromeInvert = !img.MonochromeInvert
		}
		render()
	}
}

const GOPHER_IMAGE = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

func Plot() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	p0 := widgets.NewPlot()
	p0.Title = "braille-mode Line Chart"
	p0.Data = sinData
	p0.SetRect(0, 0, 50, 15)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen

	p1 := widgets.NewPlot()
	p1.Title = "dot-mode line Chart"
	p1.Marker = widgets.MarkerDot
	p1.Data = [][]float64{{1, 2, 3, 4, 5}}
	p1.SetRect(50, 0, 75, 10)
	p1.DotMarkerRune = '+'
	p1.AxesColor = ui.ColorWhite
	p1.LineColors[0] = ui.ColorYellow
	p1.DrawDirection = widgets.DrawLeft

	p2 := widgets.NewPlot()
	p2.Title = "dot-mode Scatter Plot"
	p2.Marker = widgets.MarkerDot
	p2.Data = make([][]float64, 2)
	p2.Data[0] = []float64{1, 2, 3, 4, 5}
	p2.Data[1] = sinData[1][4:]
	p2.SetRect(0, 15, 50, 30)
	p2.AxesColor = ui.ColorWhite
	p2.LineColors[0] = ui.ColorCyan
	p2.PlotType = widgets.ScatterPlot

	p3 := widgets.NewPlot()
	p3.Title = "braille-mode Scatter Plot"
	p3.Data = make([][]float64, 2)
	p3.Data[0] = []float64{1, 2, 3, 4, 5}
	p3.Data[1] = sinData[1][4:]
	p3.SetRect(45, 15, 80, 30)
	p3.AxesColor = ui.ColorWhite
	p3.LineColors[0] = ui.ColorCyan
	p3.Marker = widgets.MarkerBraille
	p3.PlotType = widgets.ScatterPlot

	ui.Render(p0, p1, p2, p3)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}