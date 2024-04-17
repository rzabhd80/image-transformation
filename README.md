# Image Transformation Application

This application is designed for digital image processing using classic linear algebra methods. It provides various transformations such as translation, shearing, scaling, rotation, and identity transformation.

## Getting Started

To get started with the application, follow these instructions:

### Prerequisites

- Go programming language installed on your system.
- Basic understanding of linear algebra concepts.
- Familiarity with Go modules.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your_username/image_transformation.git
    ```

2. Navigate to the project directory:

    ```bash
    cd image_transformation
    ```

### Running the Application

To run the application, execute the following command:

```bash
go run main.go --path <image_path> --format <image_format> --operation <transformation_operation>
```
Replace `<image_path>` with the path to your input image, `<image_format>` with the format of your image (e.g., JPEG, PNG), and `<transformation_operation>` with the desired transformation operation.

## Application Parameters

- `--path`: The file path of the input image.
- `--format`: The file format of the input image.
- `--operation`: The desired transformation operation to be applied to the image.

## Available Transformations

The following transformations are available in the application:

- Identity Transformation
- Translation Transformation
- Shear Vertical Transformation
- Shear Horizontal Transformation
- Scaling Reflection Transformation
- Rotation Transformation

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project utilizes the power of Go programming language and its standard libraries.
- Special thanks to the contributors of the image processing libraries used in this project.

