package epub_v2

import (
	"archive/zip"
	"bytes"
	"testing"

	"github.com/mathieu-keller/epub-parser/v2/model"
)

func TestParseOpfHandlesMissingMetadataFields(t *testing.T) {
	book := newBookForOpf(`<package version="2.0" unique-identifier="bookid">
		<metadata />
		<manifest />
		<spine />
	</package>`)

	err := ParseOpf(book)
	if err != nil {
		t.Fatalf("ParseOpf returned error: %v", err)
	}

	assertNonNilSlice(t, "Identifiers", book.Metadata.Identifiers)
	assertNonNilSlice(t, "Titles", book.Metadata.Titles)
	assertNonNilSlice(t, "Languages", book.Metadata.Languages)
	assertNonNilSlice(t, "Creators", book.Metadata.Creators)
	assertNonNilSlice(t, "Contributors", book.Metadata.Contributors)
	assertNonNilSlice(t, "Publishers", book.Metadata.Publishers)
	assertNonNilSlice(t, "Subjects", book.Metadata.Subjects)
	assertNonNilSlice(t, "Descriptions", book.Metadata.Descriptions)
	assertNonNilSlice(t, "Dates", book.Metadata.Dates)
}

func newBookForOpf(opfContent string) *model.Book {
	var buf bytes.Buffer
	writer := zip.NewWriter(&buf)
	file, err := writer.Create("content.opf")
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte(opfContent))
	if err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		panic(err)
	}

	return &model.Book{
		Container: model.Container{Rootfile: model.Rootfile{Path: "content.opf"}},
		ZipReader: zipReader,
	}
}

func assertNonNilSlice(t *testing.T, fieldName string, value interface{}) {
	t.Helper()

	switch typed := value.(type) {
	case *[]model.Identifier:
		if typed == nil {
			t.Fatalf("%s should be initialized", fieldName)
		}
		if len(*typed) != 0 {
			t.Fatalf("%s should be empty", fieldName)
		}
	case *[]model.Title:
		if typed == nil {
			t.Fatalf("%s should be initialized", fieldName)
		}
		if len(*typed) != 0 {
			t.Fatalf("%s should be empty", fieldName)
		}
	case *[]string:
		if typed == nil {
			t.Fatalf("%s should be initialized", fieldName)
		}
		if len(*typed) != 0 {
			t.Fatalf("%s should be empty", fieldName)
		}
	case *[]model.Creator:
		if typed == nil {
			t.Fatalf("%s should be initialized", fieldName)
		}
		if len(*typed) != 0 {
			t.Fatalf("%s should be empty", fieldName)
		}
	case *[]model.DefaultAttributes:
		if typed == nil {
			t.Fatalf("%s should be initialized", fieldName)
		}
		if len(*typed) != 0 {
			t.Fatalf("%s should be empty", fieldName)
		}
	default:
		t.Fatalf("unsupported type %T", value)
	}
}
