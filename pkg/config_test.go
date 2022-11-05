package pkg_test

import (
	"encoding/json"
	"image"
	drawOp "image/draw"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-blank/internal"
	"github.com/petewall/eink-radiator-image-source-blank/internal/internalfakes"
	"github.com/petewall/eink-radiator-image-source-blank/pkg"
)

var _ = Describe("Config", func() {
	Describe("GenerateImage", func() {
		var (
			img      *image.RGBA
			draw     *internalfakes.FakeDrawer
			newImage *internalfakes.FakeImageMaker
		)

		BeforeEach(func() {
			img = image.NewRGBA(image.Rect(0, 0, 10, 10))
			draw = &internalfakes.FakeDrawer{}
			internal.Draw = draw.Spy
			newImage = &internalfakes.FakeImageMaker{}
			newImage.Returns(img)
			internal.NewImage = newImage.Spy
		})

		It("makes a blank image of a certain color", func() {
			config := &pkg.Config{Color: "blanchedalmond"}
			image := config.GenerateImage(100, 200)

			By("creating a new image", func() {
				Expect(newImage.CallCount()).To(Equal(1))
				width, height := newImage.ArgsForCall(0)
				Expect(width).To(Equal(100))
				Expect(height).To(Equal(200))
			})

			By("drawing the color", func() {
				Expect(draw.CallCount()).To(Equal(1))
				dst, rect, color, origin, op := draw.ArgsForCall(0)
				Expect(dst).To(Equal(img))
				Expect(rect.Min.X).To(Equal(0))
				Expect(rect.Min.Y).To(Equal(0))
				Expect(rect.Size().X).To(Equal(10))
				Expect(rect.Size().Y).To(Equal(10))

				Expect(color.At(0, 0)).To(Equal(colornames.Map["blanchedalmond"]))
				Expect(origin.X).To(Equal(0))
				Expect(origin.Y).To(Equal(0))

				Expect(op).To(Equal(drawOp.Src))
			})

			By("returning the image context", func() {
				Expect(image).To(Equal(img))
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
		config := pkg.Config{Color: "chartreuse"}
		var err error
		configFileContents, err = yaml.Marshal(config)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(configFile.Name())).To(Succeed())
	})

	It("parses the image config file", func() {
		config, err := pkg.ParseConfig(configFile.Name())
		Expect(err).ToNot(HaveOccurred())
		Expect(config.Color).To(Equal("chartreuse"))
	})

	Context("config file is json formatted", func() {
		BeforeEach(func() {
			config := pkg.Config{Color: "cyan"}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("parses just fine", func() {
			config, err := pkg.ParseConfig(configFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(config.Color).To(Equal("cyan"))
		})
	})

	When("reading the config file fails", func() {
		It("returns an error", func() {
			_, err := pkg.ParseConfig("this file does not exist")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to read image config file: open this file does not exist: no such file or directory"))
		})
	})

	When("parsing the config file fails", func() {
		BeforeEach(func() {
			configFileContents = []byte("this is invalid yaml!")
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to parse image config file: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `this is...` into pkg.Config"))
		})
	})

	When("the config file has missing data", func() {
		BeforeEach(func() {
			config := pkg.Config{}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: missing color"))
		})
	})

	When("the config file has invalid data", func() {
		BeforeEach(func() {
			config := pkg.Config{
				Color: "golf",
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: unknown color: \"golf\""))
		})
	})
})
