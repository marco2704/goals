package functions

type Args struct {
	args map[string]interface{}
}

func (a *Args) Add(name string, value interface{}) {
	if a.args == nil {
		a.args = map[string]interface{}{}
	}

	a.args[name] = value
}

func (a *Args) String(arg string) (string, error) {
	value, ok := a.args[arg]
	if !ok {
		return "", MissingArgErr(arg)
	}

	stringValue, ok := value.(string)
	if !ok {
		return "", WrongArgTypeErr(arg, "string")
	}

	return stringValue, nil
}
