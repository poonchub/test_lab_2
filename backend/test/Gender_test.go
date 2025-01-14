package test

import (
	"test_2/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestGenderName(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`GenderName is valid`, func(t *testing.T) {
		gender := entity.Gender{
			GenderName: "ชาย",
		}

		ok, error := govalidator.ValidateStruct(gender)

		g.Expect(ok).To(BeTrue())
		g.Expect(error).To(BeNil())
	})

	t.Run(`GenderName is required`, func(t *testing.T){
		gender := entity.Gender{
			GenderName: "",		//ผิดตรงนี้
		}

		ok, error := govalidator.ValidateStruct(gender)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(error).NotTo(BeNil())

		g.Expect(error.Error()).To(Equal("GenderName is required"))
	})
}