package main

import (
	"image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Escogiendo variable independiente (parte 2)

	// Abrir archivo de datasets
	f, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear dataframe de archivo CSV
	advertDF := dataframe.ReadCSV(f)

	// Extraer la columna de destino.
	yVals := advertDF.Col("Sales").Float()

	// Cree un diagrama de dispersión para cada una de las características del conjunto de datos.
	for _, colName := range advertDF.Names() {

		// pts mantendrá los valores para trazar
		pts := make(plotter.XYs, advertDF.Nrow())

		// Rellenar pts con datos.
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// Crear la trama.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		// Guardar la trama en un archivo PNG.
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
