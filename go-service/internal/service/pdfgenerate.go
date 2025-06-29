package service

import (
	"fmt"
	"strconv"

	"github.com/shamhub/pdfprovider/pkg/config"
	"github.com/shamhub/pdfprovider/types"

	"github.com/shamhub/pdfprovider/internal/dao"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
)

type Client struct {
	creator *creator.Creator
}

type cellStyle struct {
	ColSpan         int
	HAlignment      creator.CellHorizontalAlignment
	BackgroundColor creator.Color
	BorderSide      creator.CellBorderSide
	BorderStyle     creator.CellBorderStyle
	BorderWidth     float64
	BorderColor     creator.Color
	Indent          float64
}

var cellStyles = map[string]cellStyle{
	"heading-left": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"heading-centered": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"centered-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"centered": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"gradingsys-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"gradingsys-row": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"conduct-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-key": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-val": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     3,
	},
}

func GenerateStudentDataPdf(studentData dao.StudentData, fullPathName string) error {
	conf, err := config.NewUniDocCred()
	if err != nil {
		return err
	}

	err = license.SetMeteredKey(conf.Get(types.UNIDOC_LICENSE_API_KEY))
	if err != nil {
		return err
	}

	c := creator.New()
	c.SetPageMargins(40, 40, 0, 0)

	filePath := conf.Get(types.FILE_PATH)
	cr := &Client{creator: c}
	err = cr.generatePdf(studentData, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) generatePdf(studentData dao.StudentData, filePath string) error {
	rect := c.creator.NewRectangle(0, 0, creator.PageSizeLetter[0], 120)
	rect.SetFillColor(creator.ColorRGBFromHex("#dde4e5"))
	rect.SetBorderWidth(0)
	err := c.creator.Draw(rect)
	if err != nil {
		return err
	}

	headerStyle := c.creator.NewTextStyle()
	headerStyle.FontSize = 50

	table := c.creator.NewTable(1)
	table.SetMargins(0, 0, 20, 0)
	err = drawCell(table, c.newPara("Sample Data", headerStyle), cellStyles["centered"])
	if err != nil {
		return err
	}
	err = c.creator.Draw(table)
	if err != nil {
		return err
	}

	err = c.writeData(studentData)
	if err != nil {
		return err
	}

	err = c.creator.WriteToFile(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) newPara(text string, textStyle creator.TextStyle) *creator.StyledParagraph {
	p := c.creator.NewStyledParagraph()
	p.Append(text).Style = textStyle
	p.SetEnableWrap(false)
	return p
}

func drawCell(table *creator.Table, content creator.VectorDrawable, cellStyle cellStyle) error {
	var cell *creator.TableCell
	if cellStyle.ColSpan > 1 {
		cell = table.MultiColCell(cellStyle.ColSpan)
	} else {
		cell = table.NewCell()
	}
	err := cell.SetContent(content)
	if err != nil {
		return err
	}
	cell.SetHorizontalAlignment(cellStyle.HAlignment)
	if cellStyle.BackgroundColor != nil {
		cell.SetBackgroundColor(cellStyle.BackgroundColor)
	}
	cell.SetBorder(cellStyle.BorderSide, cellStyle.BorderStyle, cellStyle.BorderWidth)
	if cellStyle.BorderColor != nil {
		cell.SetBorderColor(cellStyle.BorderColor)
	}
	if cellStyle.Indent > 0 {
		cell.SetIndent(cellStyle.Indent)
	}
	return nil
}

func (c *Client) writeData(studentData dao.StudentData) error {
	headerStyle := c.creator.NewTextStyle()
	table := c.creator.NewTable(2)
	table.SetMargins(0, 0, 50, 0)
	err := drawCell(table, c.newPara("Name: "+studentData.Name, headerStyle), cellStyles["left"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("Address: "+studentData.CurrentAddress, headerStyle), cellStyles["left"])
	if err != nil {
		return err
	}
	err = c.creator.Draw(table)
	if err != nil {
		return err
	}

	table = c.creator.NewTable(20)
	table.SetMargins(0, 0, 20, 0)
	err = table.SetColumnWidths(0.4, 0.2, 0.2, 0.2)
	if err != nil {
		return err
	}
	headingStyle := c.creator.NewTextStyle()
	headingStyle.FontSize = 20
	headingStyle.Color = creator.ColorRGBFromHex("#fdfdfd")
	regularStyle := c.creator.NewTextStyle()

	// Draw table header.
	err = drawCell(table, c.newPara(" Id", headingStyle), cellStyles["heading-left"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("Email", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("SystemAccess", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("Phone", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("Gender", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("DOB", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("ClassName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("SectionName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("Roll", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("FatherName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("FatherPhone", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("MotherName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("MotherPhone", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("GuardianName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("GuardianPhone", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("RelationOfGuardian", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("CurrentAddress", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("PermanentAddress", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("AdmissionDate", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara("ReporterName", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}

	// Draw table values
	err = drawCell(table, c.newPara(" "+strconv.Itoa(studentData.Id), regularStyle), cellStyles["left-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.Email), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.SystemAccess), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.Phone), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.Gender), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.DOB), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.ClassName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.SectionName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.Roll), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.FatherName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.FatherPhone), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.MotherName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.MotherPhone), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.GuardianName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.GuardianPhone), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.RelationOfGuardian), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.CurrentAddress), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.PermanentAddress), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.AdmissionDate), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}
	err = drawCell(table, c.newPara(fmt.Sprintf("%v", studentData.ReporterName), regularStyle), cellStyles["centered-highlighted"])
	if err != nil {
		return err
	}

	err = c.creator.Draw(table)
	if err != nil {
		return err
	}
	return nil
}
