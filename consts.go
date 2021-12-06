package fdisk

import "fmt"

type SizeUnit int

const (
	SizeUnitHuman = SizeUnit(iota)
	SizeUnitBytes
)

func (s SizeUnit) String() string {
	switch s {
	case SizeUnitHuman:
		return "human"
	case SizeUnitBytes:
		return "bytes"
	default:
		return fmt.Sprintf("unknownUnit(%d)", int(s))
	}
}

type Unit string

const (
	UnitCylinder = Unit("cylinder")
	UnitSector   = Unit("sector")
)

type Direction int

const (
	DirectionUp = iota + 1
	DirectionDown
	DirectionNearest
)

type PartTypeParserFlag int

const (
	PartTypeFlagParseData = PartTypeParserFlag(1 << (iota + 1))
	PartTypeFlagParseDataLast
	PartTypeFlagParseShortcut
	PartTypeFlagParseAlias
	PartTypeFlagParseDeprecated
	PartTypeFlagParseNoUnknown
	PartTypeFlagParseSeqNum
	PartTypeFlagParseName

	PartTypeFlagParseDefault = PartTypeFlagParseData |
		PartTypeFlagParseShortcut |
		PartTypeFlagParseAlias |
		PartTypeFlagParseName |
		PartTypeFlagParseSeqNum
)

type FieldTypes int

const (
	FieldDevice = FieldTypes(iota)
	FieldStart
	FieldEnd
	FieldSectors
	FieldCylinders
	FieldSize
	FieldType
	FieldTypeID
	FieldAttr
	FieldBoot
	FieldBSize
	FieldCPG
	FieldEAddr
	FieldFSize
	FieldName
	FieldSAddr
	FieldUUID
	FieldFSUUID
	FieldFSLabel
	FieldFSType
	NFields
)

type LabelItemGen int

const (
	LabelItemID          = LabelItemGen(0)
	LabelItemNLabelItems = LabelItemGen(8)
)
