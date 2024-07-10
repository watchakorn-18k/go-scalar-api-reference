## Overview 🌐

The Scalar package serves as a provider for the [Scalar](https://github.com/scalar/scalar) project. It offers a comprehensive suite of functions designed for generating API references in HTML format, specializing in JSON data handling and web presentation customization. This includes functions to serialize options into JSON, manage HTML escaping, and dynamically handle different types of specification content.

## Features 🚀

### JSON Serialization and HTML Escaping

- **safeJSONConfiguration** 🔒: Serializes configuration options into JSON format and escapes HTML characters to prevent XSS attacks. This ensures that the JSON can be safely embedded within HTML documents.

### Specification Content Handling

- **specContentHandler** 📝: Dynamically handles different types of specification content. It can process content as a function returning a map, a direct map, or a plain string, converting the specification content into JSON format suitable for web use.

### HTML Generation

- **ApiReferenceHTML** 📄: Generates a complete HTML document for API reference. It allows for extensive customization, including themes, layouts, and CDN configuration. It handles both direct specification content and content fetched from a URL, providing error handling for missing specifications.

## Customization Options ⚙️

The package allows extensive customization of the generated API reference through the `Options` struct, supporting:

- **CDN**: URL of the CDN to load additional scripts or styles.
- **Theme**: Customizable themes for styling the API reference.
- **Layout**: Choice between modern and classic layout designs.
- **SpecURL**: URL from which the specification content can be fetched.
- **Dark Mode**: Option to enable a dark theme for the API reference.

### Themes and Styles 🎨

- Provides a default set of CSS styles for both light and dark themes.
- Allows custom CSS injections to tailor the appearance to specific branding or aesthetic requirements.

## Error Handling 🛠️

- Robust error handling for scenarios where necessary parameters like `SpecURL` or `SpecContent` are missing.
- Errors during the fetching of specification content from URLs are properly managed and reported.

## Usage 📚

To use the Scalar package as a provider in your Go project for creating API references, follow the example below:

```go
package main

import (
	"go-fiber-swagger/src/middlewares"
	"log"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	middlewares.Logger(app)
	app.Use("/api/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.yaml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(htmlContent)
	})
	app.Use(recover.New())
	app.Use(cors.New())
	log.Fatal(app.Listen(":3000"))
}
```
## License

This project is forked from [MarceloPetrucio/go-scalar-api-reference](https://github.com/MarceloPetrucio/go-scalar-api-reference) and is licensed under the MIT License.