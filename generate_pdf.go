package main

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	m.RegisterHeader(func() {
		m.Row(15, func() {

			m.Col(9, func() {
				m.Text("E-mail Breakdown", props.Text{
					Size:  15,
					Top:   1,
					Align: consts.Center,
					Style: consts.Bold,
				})
			})

			m.Col(3, func() {
				m.Text(time.Now().Format("January 2, 2006"), props.Text{
					Size:   10,
					Top:    7,
					Align:  consts.Center,
					Family: consts.Courier,
				})
			})
		})
	})

	m.Line(1.0,
		props.Line{
			Color: color.Color{0, 0, 0},
			Width: 1.0,
		},
	)

	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("Parsed Headers", props.Text{
				Size:   12,
				Top:    5,
				Align:  consts.Center,
				Family: consts.Helvetica,
			})
		})
	})

	m.TableList(
		[]string{"Header", "Value"},
		[][]string{
			[]string{"From", "anonymous@gmail.com"},
			[]string{"To", "johnwick@gmail.com"},
			[]string{"Sender", "anonymous@gmail.com"},
			[]string{"Message-ID", "141lkjdpqrfe41041734dzaf"},
		},
		props.TableList{
			HeaderProp: props.TableListContent{
				Style: consts.BoldItalic,
				Size:  11,
				Color: color.Color{0, 100, 150},
			},
			ContentProp: props.TableListContent{
				Style: consts.Normal,
				Size:  11,
			},
			//Align: consts.Center,
			Line: true,
		},
	)

	err := m.OutputFileAndClose("./learn.pdf")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
