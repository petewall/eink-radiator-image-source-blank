package internal_test

import (
	"encoding/json"
	"image/color"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
	"github.com/petewall/eink-radiator-image-source-blank/v2/internal/internalfakes"
)

var _ = Describe("Config", func() {
	Describe("GenerateImage", func() {
		var (
			newImageContext *internalfakes.FakeImageContextMaker
			imageContext    *internalfakes.FakeImageContext
		)

		BeforeEach(func() {
			imageContext = &internalfakes.FakeImageContext{}
			newImageContext = &internalfakes.FakeImageContextMaker{}
			newImageContext.Returns(imageContext)
			internal.NewImageContext = newImageContext.Spy
		})

		It("makes a blank image of a certain color", func() {
			config := &internal.Config{Color: "blanchedalmond"}
			image := config.GenerateImage(100, 200)

			By("creating a new image context", func() {
				Expect(newImageContext.CallCount()).To(Equal(1))
				width, height := newImageContext.ArgsForCall(0)
				Expect(width).To(Equal(100))
				Expect(height).To(Equal(200))
			})

			By("creating the right image", func() {
				Expect(imageContext.SetColorCallCount()).To(Equal(1))
				Expect(imageContext.SetColorArgsForCall(0)).To(Equal(color.RGBA{0xff, 0xeb, 0xcd, 0xff}))
				Expect(imageContext.DrawRectangleCallCount()).To(Equal(1))
				x, y, w, h := imageContext.DrawRectangleArgsForCall(0)
				Expect(x).To(Equal(0.0))
				Expect(y).To(Equal(0.0))
				Expect(w).To(Equal(100.0))
				Expect(h).To(Equal(200.0))
				Expect(imageContext.FillCallCount()).To(Equal(1))
			})

			By("returning the image context", func() {
				Expect(image).To(Equal(imageContext))
			})
		})
	})
})

var _ = Describe("ParseConfig", func() {
	var (
		configFile         *os.File
		configFileContents []byte
	)

	JustBeforeEach(func() {
		var err error
		configFile, err = os.CreateTemp("", "blank-config.yaml")
		Expect(err).ToNot(HaveOccurred())
		_, err = configFile.Write(configFileContents)
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		config := internal.Config{Color: "chartreuse"}
		var err error
		configFileContents, err = yaml.Marshal(config)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(configFile.Name())).To(Succeed())
	})

	It("parses the image config file", func() {
		config, err := internal.ParseConfig(configFile.Name())
		Expect(err).ToNot(HaveOccurred())
		Expect(config.Color).To(Equal("chartreuse"))
	})

	Context("config file is json formatted", func() {
		BeforeEach(func() {
			config := internal.Config{Color: "cyan"}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("parses just fine", func() {
			config, err := internal.ParseConfig(configFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(config.Color).To(Equal("cyan"))
		})
	})

	When("reading the config file fails", func() {
		It("returns an error", func() {
			_, err := internal.ParseConfig("this file does not exist")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to read image config file: open this file does not exist: no such file or directory"))
		})
	})

	When("parsing the config file fails", func() {
		BeforeEach(func() {
			configFileContents = []byte("this is invalid yaml!")
		})

		It("returns an error", func() {
			_, err := internal.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to parse image config file: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `this is...` into internal.Config"))
		})
	})
})
