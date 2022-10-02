package test_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Generate", func() {
	var outputFile string

	BeforeEach(func() {
		outputFile = ""
	})

	AfterEach(func() {
		if outputFile != "" {
			Expect(os.Remove(outputFile)).To(Succeed())
		}
	})

	It("generates a blank image", func() {
		outputFile = "orange.png"
		Run("generate --output orange.png --config orange-config.yaml")
		Eventually(CommandSession).Should(Exit(0))

		By("saving the image to a file", func() {
			actualData, err := os.ReadFile(outputFile)
			Expect(err).ToNot(HaveOccurred())
			expectedData, err := os.ReadFile("expected_orange.png")
			Expect(err).ToNot(HaveOccurred())
			Expect(actualData).To(Equal(expectedData))
		})
	})

	When("using --to-stdout", func() {
		It("writes the image to stdout", func() {
			Run("generate --height 200 --width 300 --to-stdout --config green-config.json")
			Eventually(CommandSession).Should(Exit(0))

			By("saving the image to a file", func() {
				expectedData, err := os.ReadFile("expected_green.png")
				Expect(err).ToNot(HaveOccurred())
				Expect(CommandSession.Out.Contents()).To(Equal(expectedData))
			})
		})
	})
})
