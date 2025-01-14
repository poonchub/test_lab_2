package test

import (
	"test_2/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStudentID(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`StudentID is required`, func(t *testing.T){
		user := entity.User{
			StudentID: "",
			Name: "Peter",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(error).NotTo(BeNil())
		
		g.Expect(error.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`StudentID is invalid`, func(t *testing.T){
		user := entity.User{
			StudentID: "K1234567",
			Name: "Peter",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(error).NotTo((BeNil()))

		g.Expect(error.Error()).To(Equal("StudentID is invalid format"))
	})

	t.Run(`StudentID is valid`, func(t *testing.T){
		user := entity.User{
			StudentID: "B1234567",
			Name: "Peter",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(error).To((BeNil()))
	})
}

func TestName(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Name is required`, func(t *testing.T){
		user := entity.User{
			StudentID: "B6525163",
			Name: "",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(error).NotTo(BeNil())

		g.Expect(error.Error()).To(Equal("Name is required"))
	})

	t.Run(`Name must be only letters`, func(t *testing.T){
		user := entity.User{
			StudentID: "B6525163",
			Name: "12345",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(error).NotTo(BeNil())

		g.Expect(error.Error()).To(Equal("Name must be only letters"))
	})

	t.Run(`Name is valid`, func(t *testing.T){
		user := entity.User{
			StudentID: "B6525163",
			Name: "Peter",
			Email: "test@gmail.com",
			PhoneNumber: "1234567890",
			GenderID: 1,
		}

		ok, error := govalidator.ValidateStruct(user)
		
		g.Expect(ok).To(BeTrue())
		g.Expect(error).To(BeNil())
	})
}