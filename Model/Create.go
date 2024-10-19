package Model

type Create struct {
	element *Element

	condition ICondition
}

func NewCreate(delay float64) *Create {
	c := &Create{element: NewElementWithDelay(delay), condition: NewFullRandomCondition(1)}
	c.element.tnext = 0.0
	return c
}

func (c *Create) OutAct() {
	c.element.OutAct()
	c.element.tnext = c.element.tcurr + c.element.GetDelay()

	c.element.nextElements[c.condition.MakeCondition()].InAct()
}

func (c *Create) InAct() {

}

func (c *Create) DoStatistic(time float64) {

}

func (c *Create) GetBaseElement() *Element {
	return c.element
}

func (c *Create) PrintInfo() {
	c.element.PrintInfo()
}

func (c *Create) PrintResult() {
	c.element.PrintResult()
}

func (c *Create) GetId() int {
	return c.element.id
}

func (c *Create) SetNextElement(el IElement) {
	c.element.nextElements = append(c.element.nextElements, el)
}

func (c *Create) SetNextElements(elements []IElement) {
	c.element.nextElements = elements
}

func (c *Create) SetCondition(condition ICondition) {
	c.condition = condition
}

func (c *Create) SetName(name string) {
	c.element.name = name
}

func (c *Create) SetDistribution(distribution string) {
	c.element.distribution = distribution
}

func (c *Create) GetNextElements() []IElement {
	return c.element.nextElements
}

func (c *Create) GetTNext() float64 {
	return c.element.GetTNext()
}

func (c *Create) SetTNext(tNext float64) {
	c.element.SetTNext(tNext)
}

func (c *Create) GetTCurr() float64 {
	return c.element.GetTCurr()
}

func (c *Create) SetTCurr(tCurr float64) {
	c.element.SetTCurr(tCurr)
}
