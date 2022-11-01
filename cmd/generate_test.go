package cmd_test

import (
	"errors"
	"image"

	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/petewall/eink-radiator-image-source-blank/v2/cmd"
	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
	"github.com/petewall/eink-radiator-image-source-blank/v2/internal/internalfakes"
)

var _ = Describe("Generate", func() {
	var (
		img            image.Image
		imageGenerator *internalfakes.FakeImageGenerator
		imageEncoder   *internalfakes.FakeImageEncoder
		imageWriter    *internalfakes.FakeImageWriter
	)

	BeforeEach(func() {
		img = image.NewRGBA(image.Rect(0, 0, 10, 10))
		imageGenerator = &internalfakes.FakeImageGenerator{}
		imageGenerator.GenerateImageReturns(img)
		imageEncoder = &internalfakes.FakeImageEncoder{}
		imageWriter = &internalfakes.FakeImageWriter{}

		cmd.ImageGenerator = imageGenerator
		internal.EncodeImage = imageEncoder.Spy
		internal.WriteImage = imageWriter.Spy

		viper.Set("to-stdout", false)
		viper.Set("output", cmd.DefaultOutputFilename)
		viper.Set("height", 1000)
		viper.Set("width", 2000)
	})

	It("generates a blank image", func() {
		err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
		Expect(err).ToNot(HaveOccurred())

		By("defaulting to writing to blank.png", func() {
			Expect(imageWriter.CallCount()).To(Equal(1))
			filename, image := imageWriter.ArgsForCall(0)
			Expect(filename).To(Equal("blank.png"))
			Expect(image).To(Equal(img))
		})

		By("using the right resolution", func() {
			Expect(imageGenerator.GenerateImageCallCount()).To(Equal(1))
			width, height := imageGenerator.GenerateImageArgsForCall(0)
			Expect(width).To(Equal(2000))
			Expect(height).To(Equal(1000))
		})
	})

	When("using --to-stdout", func() {
		var output *Buffer

		BeforeEach(func() {
			output = NewBuffer()
			cmd.GenerateCmd.SetOut(output)
			viper.Set("to-stdout", true)
		})

		It("outputs the image to stdout", func() {
			err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
			Expect(err).ToNot(HaveOccurred())

			Expect(imageEncoder.CallCount()).To(Equal(1))
			buffer, image := imageEncoder.ArgsForCall(0)
			Expect(buffer).To(Equal(output))
			Expect(image).To(Equal(img))
		})

		When("encoding fails", func() {
			BeforeEach(func() {
				imageEncoder.Returns(errors.New("encode png failed"))
			})

			It("returns an error", func() {
				err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("encode png failed"))
			})
		})
	})

	When("saving the image fails", func() {
		BeforeEach(func() {
			imageWriter.Returns(errors.New("save png failed"))
		})

		It("returns an error", func() {
			err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("save png failed"))
		})
	})
})
