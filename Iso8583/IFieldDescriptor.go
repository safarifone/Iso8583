package Iso8583
import (
	"github.com/safarifone/Iso8583/Formatter"
	"github.com/safarifone/Iso8583/LengthFormatters"
	"github.com/safarifone/Iso8583/FieldValidator"
)

type IFieldDescriptor interface  {

	Adjuster() Adjuster
	Formatter() formatter.IFormatter
	LengthFormatter() lengthFormatters.ILengthFormatter
	Validator() fieldValidator.IFieldValidator
	IsComposite() bool
	CompositeTemplate() *Template
	Display(string,string,string) string
	GetPackedLength(string) int
	Pack(int, string) ([]byte,error)
	Unpack(int, []byte, int) (string, int, error)
}