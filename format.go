package gopandoc

import "errors"

func toPandocInputFormat(format string) (string, error) {
	switch format {
	case "md":
		return "markdown", nil
	case "tex":
		return "latex", nil
	case "txt":
		return "markdown", nil
	case "docx", "epub", "fb2", "html", "odt":
		return format, nil
	default:
		return "", errors.New("unknown input format")
	}
}

func toPandocOutputFormat(format string) (string, error) {
	switch format {
	case "md":
		return "markdown", nil
	case "tex":
		return "latex", nil
	case "txt":
		return "plain", nil
	case "docx", "epub", "fb2", "html", "odt":
		return format, nil
	default:
		return "", errors.New("unknown input format")
	}
}
