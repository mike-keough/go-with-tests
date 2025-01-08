package hello

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"

	frenchPrefix  = "Bonjour, "
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishPrefix
	case "French":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
