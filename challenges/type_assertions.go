package type_assertions

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	var dev Developer
	dev.Name = name.(string)
	dev.Age = age.(int)
	return dev
}
