package xlsx

import . "gopkg.in/check.v1"

type GoogleDocsExcelSuite struct{}

var _ = Suite(&GoogleDocsExcelSuite{})

// Test that we can successfully read an XLSX file generated by
// Google Docs.
func (g *GoogleDocsExcelSuite) TestGoogleDocsExcel(c *C) {
	xlsxFile, err := OpenFile("./testdocs/googleDocsTest.xlsx")
	c.Assert(err, IsNil)
	c.Assert(xlsxFile, NotNil)
}
