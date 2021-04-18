package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/walpod/bend-it"
	"github.com/walpod/bend-it/cubic"
	"math"
)

type Playground struct {
	//window *widgets.QMainWindow
	//canvas *widgets.QWidget
	spline bendit.Fn2d
	tu     float64
}

func (pg *Playground) build(window *widgets.QMainWindow) {
	//pg.window = window

	/*statusbar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)
	statusbar.ShowMessage("the status bar ...", 0)*/

	central := widgets.NewQWidget(window, 0)
	central.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(central)

	canvas := widgets.NewQWidget(central, 0)
	//canvas.Resize2(800, 500)
	central.Layout().AddWidget(canvas)

	pg.buildSpline()

	canvas.ConnectPaintEvent(func(vqp *gui.QPaintEvent) {
		pg.paint(canvas)
	})
}

func (pg *Playground) buildSpline() {
	/*
		vertsx, vertsy := []float64{10, 100, 150}, []float64{10, 100, 10}
		entryTansx, entryTansy, exitTansx, exitTansy := cubic.NaturalTanf2d{}.Find(vertsx, vertsy, nil)
		pg.spline = cubic.BuildHermiteSpline2d(vertsx, vertsy, entryTansx, entryTansy, exitTansx, exitTansy, nil)
		pg.tu = float64(len(vertsx) - 1)
	*/

	//pg.spline = cubic.NewBezierSpline2d([]float64{200, 400}, []float64{200, 400}, []float64{210, 390}, []float64{200, 400}, nil).Fn()
	pg.spline = cubic.NewBezierSpline2d(
		[]float64{100, 300, 500}, []float64{100, 300, 100},
		[]float64{120, 250, 350, 480}, []float64{150, 300, 300, 150}, nil).Fn()
	pg.tu = 2

}

func (pg *Playground) paint(canvas *widgets.QWidget) {
	qp := gui.NewQPainter2(canvas)
	stepSize := pg.tu / 100
	for t := 0.; t < pg.tu; t += stepSize {
		x, y := pg.spline(t)
		qp.DrawPoint3(int(math.Round(x)), int(math.Round(y)))
	}
	qp.DestroyQPainter()
}
