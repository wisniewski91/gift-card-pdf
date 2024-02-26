package services

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/wisniewski91/gift-card-pdf/types"
)

func GiftCardGenerator(data types.GiftCard) {
	validDate := time.Now().AddDate(0, 3, 7)

	pdf := gofpdf.New("L", "mm", "A5", "static/fonts")
	pdf.AddUTF8Font("Merienda", "L", "Merienda-Light.ttf")
	pdf.SetFont("Merienda", "L", 11)
	pdf.SetTextColor(60, 70, 50)
	pdf.AddPage()
	pdf.ImageOptions(fmt.Sprintf("static/bg/%s", data.Background), 0, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true, AllowNegativePosition: true}, 0, "")
	pdf.AddPage()
	pdf.ImageOptions("static/logo.png", 0, 0, 160, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: false}, 0, "")
	pdf.Text(15, 75, "ALAN MED Gabinet Masażu")
	pdf.Text(15, 83, "ul. Przybyszewskiego 1/2")
	pdf.Text(15, 91, "30-128 Kraków")
	pdf.Text(15, 99, "Tel.: +48 725 41 42 42")
	pdf.Text(15, 117, "https://www.alanmed.eu")
	pdf.Text(15, 125, "gabinetalanmed@gmail.com")

	pdf.SetDrawColor(208, 221, 47)

	pdf.Text(95, 75, "Bon upominkowy dla:")
	pdf.Text(95, 83, data.Name)
	pdf.Line(95, 85, 195, 85)

	pdf.Text(95, 96, "Do wykorzystania w gabinecie masażu na zabieg:")
	pdf.Text(95, 104, data.Service)
	pdf.Line(95, 106, 195, 106)

	pdf.Text(95, 117, "Ważny do:")
	pdf.Text(95, 125, validDate.Format("2006-01-02"))
	pdf.Line(95, 127, 195, 127)

	filename := fmt.Sprintf("%s-%s.pdf", data.Name, validDate.Format("2006-01-02"))

	err := pdf.OutputFileAndClose(fmt.Sprintf("static/giftcards/%s", filename))

	if err != nil {
		panic(err)
	}

	// err = SendEmail(data, filename)
	fmt.Println(data)
	fmt.Println(filename)

	if err != nil {
		panic(err)
	}

}
