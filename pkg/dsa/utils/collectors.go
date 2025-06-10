package utils

type Collector struct {
	Elements []string
}

func StringCollector(size int) *Collector {
	arr := make([]string, 0, size)
	return &Collector{Elements: arr}
}
func (c *Collector) Append(s string) {
	c.Elements = append(c.Elements, s)
}
