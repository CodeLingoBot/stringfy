package stringfy

// Interfaces

type omissioner interface {
	setOmission(string)
}

type separatorer interface {
	setSeparator(string)
}

type radiuser interface {
	setRadius(int)
}

type lengther interface {
	setLength(int)
}

type lineWidther interface {
	setLineWidth(int)
}

// Types

type omissionOption func(omissioner)
type separatorOption func(separatorer)
type radiusOption func(radiuser)
type lengthOption func(lengther)
type lineWidthOption func(lineWidther)

// AddOmission adds a custom omission
func AddOmission(om string) omissionOption {
	return func(obj omissioner) {
		obj.setOmission(om)
	}
}

// AddSeparator adds a custom separator
func AddSeparator(sep string) separatorOption {
	return func(obj separatorer) {
		obj.setSeparator(sep)
	}
}

// AddRadius adds a custom separator
func AddRadius(rad int) radiusOption {
	return func(obj radiuser) {
		obj.setRadius(rad)
	}
}

// AddLength adds a custom length
func AddLength(l int) lengthOption {
	return func(obj lengther) {
		obj.setLength(l)
	}
}

// AddLineWidth adds a custom line width
func AddLineWidth(lw int) lineWidthOption {
	return func(obj lineWidther) {
		obj.setLineWidth(lw)
	}
}
