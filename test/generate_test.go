package test_test

import (
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
)

var _ = Describe("Generate", func() {
	var (
		configFile     *os.File
		configFileData []byte
		outputFile     string
	)

	BeforeEach(func() {
		outputFile = ""

		var err error
		config := internal.Config{Color: "orange"}
		configFileData, err = json.Marshal(config)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(configFile.Name())).To(Succeed())
		if outputFile != "" {
			Expect(os.Remove(outputFile)).To(Succeed())
		}
	})

	JustBeforeEach(func() {
		var err error
		configFile, err = os.CreateTemp("", "blank-image-config.json")
		Expect(err).ToNot(HaveOccurred())

		configFile.Write(configFileData)
	})

	It("generates a blank image", func() {
		outputFile = "orange.png"
		Run("generate --output " + outputFile + " --config " + configFile.Name())
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
		BeforeEach(func() {
			var err error
			config := internal.Config{Color: "green"}
			configFileData, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("writes the image to stdout", func() {
			Run("generate --height 200 --width 300 --to-stdout --config " + configFile.Name())
			Eventually(CommandSession).Should(Exit(0))

			By("saving the image to a file", func() {
				expectedData, err := os.ReadFile("expected_green.png")
				Expect(err).ToNot(HaveOccurred())
				Expect(CommandSession.Out.Contents()).To(Equal(expectedData))
			})
		})
	})
})
