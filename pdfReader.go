package main

import (
	"bytes"
	"github.com/ledongthuc/pdf"
	"io/ioutil"
	"strconv"
	"strings"
)

type Info struct {
	bufferLength   float64
	playingBitRate float64
	bufferBitRate  float64
	time           float64
}

var originalTime float64

func main() {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		var infos []Info
		content, err := readPdf(f.Name())
		if err != nil {
			panic(err)
		}
		v := strings.Split(content, "ThisIsADelimiter")
		originalTime = time(v[0])
		for i := 0; i < len(v)-1; i++ {
			infos = append(infos, newInfo(v[i]))
		}
		plotData(infos, strings.Replace(f.Name(), ".pdf", "", 1))
	}
}

func newInfo(content string) Info {
	var i Info
	i.bufferLength = bufferLength(content)
	i.playingBitRate = playingBitRate(content)
	i.bufferBitRate = bufferBitRate(content)
	i.time = time(content)
	return i
}

func bufferLength(content string) float64 {
	v := strings.Split(content, "Buffer size in Seconds (a/v): ")
	return dataToFloat(v)
}

func playingBitRate(content string) float64 {
	v := strings.Split(content, "Playing bitrate (a/v): ")
	return dataToFloat(v)
}

func bufferBitRate(content string) float64 {
	v := strings.Split(content, "Buffering bitrate (a/v): ")
	return dataToFloat(v)
}

func throughput(content string) int {
	v := strings.Split(content, "Throughput:")
	i, err := strconv.Atoi(strings.Split(v[1], " ")[0])
	if err != nil {
		panic(err)
	}
	return i
}

func dataToFloat(v []string) float64 {
	s := strings.Split(strings.Split(v[1], " / ")[1], " ")[0]
	f, err := strconv.ParseFloat(s, 2)
	if err != nil {
		panic(err)
	}
	return f
}

func time(content string) float64 {
	v := strings.Split(content, "Position: ")
	time, err := strconv.ParseFloat(strings.Split(v[1], " ")[0], 2)
	if err != nil {
		panic(err)
	}
	return time - originalTime
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open("./data/"+path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)
	return buf.String(), nil
}
