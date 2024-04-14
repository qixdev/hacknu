package helpers

import (
	"fmt"
)

// LinearRegressionModel holds the parameters of the linear regression line
type LinearRegressionModel struct {
	Slope     float64
	Intercept float64
}

// LinearRegression computes the slope and intercept for a set of x and y values
func LinearRegression(x, y []float64) LinearRegressionModel {
	if len(x) != len(y) {
		panic("Mismatched data slices")
	}
	n := float64(len(x))
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXX += x[i] * x[i]
	}
	slope := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return LinearRegressionModel{Slope: slope, Intercept: intercept}
}

// Predict uses the linear regression model to predict a y value for a given x
func PredictHelper(model LinearRegressionModel, x float64) int {
	return int(model.Slope*x + model.Intercept)
}

func Predict(grades []float64) int {
	years := []float64{}
	for i := 0; i < len(grades); i++ {
		years = append(years, float64(i))
	}
	// Perform linear regression
	model := LinearRegression(years, grades)

	// Predict the next grade for the next year/semester
	nextYear := float64(len(grades))
	predictedGrade := PredictHelper(model, nextYear)

	fmt.Printf("Regression Model: y = %.2fx + %.2f\n", model.Slope, model.Intercept)
	fmt.Printf("Prediction: %v", predictedGrade)
	return predictedGrade
}
