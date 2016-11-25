package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIndex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bin Suite")
}

var _ = Describe("Index", func() {

	createFolders()
	createFiles()

	Describe("Check the results", func() {

		It("should check the multiplication result", func() {
			result := multiplicationOperation()
			Expect(result).To(Equal(240000))
		})

		It("should check the subtraction result", func() {
			result := subtractionOperation()
			Expect(result).To(Equal(0.5990000000000001))
		})

		It("should check the addition result", func() {
			result := additionOperation()
			Expect(result).To(Equal(50))
		})

		It("should check the division result", func() {
			result := divisionOperation()
			Expect(result).To(Equal(0.019047619047619046))
		})
	})

})
