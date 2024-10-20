package Interfaces

import (
	"Model/Model/Elements"
	"Model/Model/Statistic"
)

type IProcess interface {
	Elements.IProcessElement
	Statistic.IStatistic
}
