package cmd_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/petewall/eink-radiator-image-source-blank/v2/cmd"
)

var _ = XDescribe("Generate", func() {
	It("generates a blank image", func() {

		By("defaulting to writing to blank.png", func() {

		})

		By("defaulting to 640x480", func() {

		})
	})

	When("using --to-stdout", func() {
		var output *Buffer

		BeforeEach(func() {
			output = NewBuffer()
			cmd.GenerateCmd.SetOut(output)
		})

		It("outputs the image to stdout", func() {

		})

		When("encoding fails", func() {
			It("returns an error", func() {
				err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(""))
			})
		})
	})

	When("using --height and --width to change the resolution", func() {
		It("generates an image of the specified resolution", func() {

		})
	})

	When("saving the image fails", func() {
		It("returns an error", func() {
			err := cmd.GenerateCmd.RunE(cmd.GenerateCmd, []string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(""))
		})
	})
})
