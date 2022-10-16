package cmd_test

import (
	"errors"

	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/petewall/eink-radiator-image-source-blank/v2/cmd"
	"github.com/petewall/eink-radiator-image-source-blank/v2/internal/internalfakes"
)

var _ = Describe("Generate", func() {
	var (
		imageGenerator *internalfakes.FakeImageGenerator
		imageContext   *internalfakes.FakeImageContext
	)

	BeforeEach(func() {
		imageContext = &internalfakes.FakeImageContext{}
		imageGenerator = &internalfakes.FakeImageGenerator{}
		imageGenerator.GenerateImageReturns(imageContext)

		cmd.ImageGenerator = imageGenerator

		viper.Set("to-stdout", false)
		viper.Set("output", cmd.DefaultOutputFilename)
		viper.Set("height", 1000)
		viper.Set("width", 2000)
	})

	It("generates a blank image", func() {
		err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
		Expect(err).ToNot(HaveOccurred())

		By("defaulting to writing to blank.png", func() {
			Expect(imageContext.SavePNGCallCount()).To(Equal(1))
			Expect(imageContext.SavePNGArgsForCall(0)).To(Equal("blank.png"))
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

			Expect(imageContext.EncodePNGCallCount()).To(Equal(1))
			Expect(imageContext.EncodePNGArgsForCall(0)).To(Equal(output))
		})

		When("encoding fails", func() {
			BeforeEach(func() {
				imageContext.EncodePNGReturns(errors.New("encode png failed"))
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
			imageContext.SavePNGReturns(errors.New("save png failed"))
		})

		It("returns an error", func() {
			err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("save png failed"))
		})
	})
})
