package controllers

import (
	"officeoa/models"
)

type SeriesController struct {
	FrontController
}

func (c *SeriesController) Get() {
	series, err := models.FindSeries()
	c.checkerr(err, "find series error")
	c.Response(true, "success", series)
}

func (c *SeriesController) Add() {
}
func (c *SeriesController) Edit() {
}
func (c *SeriesController) Delete() {
}
