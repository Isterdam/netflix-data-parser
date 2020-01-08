package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotData(infos []Info, name string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = name+" Mbps"
	p.Y.Scale = plot.LogScale{}
	p.Y.Min = 40
	p.X.Label.Text = "Seconds after start"

	bufferLengths := make(plotter.XYs, len(infos))
	playingBitRates := make(plotter.XYs, len(infos))
	bufferBitRates := make(plotter.XYs, len(infos))

	for i, info := range infos {
		bufferLengths[i].Y = info.bufferLength
		bufferLengths[i].X = info.time

		playingBitRates[i].Y = info.playingBitRate
		playingBitRates[i].X = info.time

		bufferBitRates[i].Y = info.bufferBitRate
		bufferBitRates[i].X = info.time
	}

	err = plotutil.AddLinePoints(p,
		"Buffer length (s)", bufferLengths,
		"Playing bit rate (kbps)", playingBitRates,
		"Buffer bit rate (kbps)", bufferBitRates)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, name+".png"); err != nil {
		panic(err)
	}
}
