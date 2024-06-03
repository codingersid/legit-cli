package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var layoutCmd = &cobra.Command{
	Use:   "layout [nama_layout]",
	Short: "Membuat layout baru di path: resources/views/layouts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		layoutName := args[0]
		nameParts := strings.Split(layoutName, "/")
		pathName := TransformPath(layoutName)
		fileHtml := SnakeCase(nameParts[len(nameParts)-1])
		layoutPath := fmt.Sprintf(`%s/%s.html`, pathName, fileHtml)
		status, err := createlayout(layoutName)
		if err != nil {
			fmt.Println("Gagal membuat layout:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat layout, layout", layoutPath, "sudah ada!")
			return
		}
		fmt.Println("layout", layoutPath, "berhasil dibuat!")
	},
}

func createlayout(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := TransformPath(name)
	fileHtml := SnakeCase(nameParts[len(nameParts)-1])
	layoutPath := fmt.Sprintf(`resources/views/layouts/%s/%s.html`, pathName, fileHtml)
	err := os.MkdirAll(fmt.Sprintf("resources/views/layouts/%s", pathName), os.ModePerm)
	if err != nil {
		return status, err
	}

	if _, err := os.Stat(layoutPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(layoutPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	code := `<!doctype html>
<html lang="id">

<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="ROBOTS" content="index, follow" />
	<meta name="author" content="CODINGERS.ID" />
	<meta name="description" content="Legit adalah Go Framework yang dikembangkan oleh CODINGERS.ID sebagai framework yang ditujukan khusus untuk pemula belajar bahasa pemrograman Go." />
	<meta name="keywords" content="Legit Framework, CODINGERS.ID" />
	<meta property="og:title" content="{{ .Title }}" />
	<meta property="og:description" content="Legit adalah Go Framework yang dikembangkan oleh CODINGERS.ID sebagai framework yang ditujukan khusus untuk pemula belajar bahasa pemrograman Go." />
	<meta property="og:image" content="https://raw.githubusercontent.com/codingersid/legit-cli/main/assets/legit-logo/legit%20icon%20color.png" />
	<meta name="twitter:title" content="{{ .Title }}" />
	<meta name="twitter:description" content="Legit adalah Go Framework yang dikembangkan oleh CODINGERS.ID sebagai framework yang ditujukan khusus untuk pemula belajar bahasa pemrograman Go." />
	<meta name="twitter:image" content="https://raw.githubusercontent.com/codingersid/legit-cli/main/assets/legit-logo/legit%20icon%20color.png" />
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" />
	<title>{{ .Title }}</title>
</head>

<body>
	<nav class="navbar sticky-top navbar-expand-lg navbar-light bg-light">
		<div class="container">
			<a class="navbar-brand" href="/">
				<img src="https://raw.githubusercontent.com/codingersid/legit-cli/main/assets/legit-logo/legit%20icon%20color.png" alt="LEGIT FRAMEWORK" height="30" class="d-inline-block align-text-top">
			</a>
			<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarSupportedContent">
				<ul class="navbar-nav me-auto mb-2 mb-lg-0">
					<li class="nav-item">
						<a class="nav-link active" aria-current="page" href="/">Home</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#">Menu 1</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#">Menu 2</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#">Menu 3</a>
					</li>
					<li class="nav-item dropdown">
						<a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
							Menu 4
						</a>
						<ul class="dropdown-menu" aria-labelledby="navbarDropdown">
							<li><a class="dropdown-item" href="#">Sub Menu 1</a></li>
							<li><a class="dropdown-item" href="#">Sub Menu 2</a></li>
							<li>
								<hr class="dropdown-divider">
							</li>
							<li><a class="dropdown-item" href="#">Sub Menu Lainnya</a></li>
						</ul>
					</li>
				</ul>
				<form class="d-flex">
					<a class="btn btn-primary me-2" href="/auth/register">Register</a>
					<a class="btn btn-outline-success" href="/auth/login">Login</a>
				</form>
			</div>
		</div>
	</nav>

	<div class="py-5" style="min-height: 350px;">
		{{ embed }}
	</div>

	<div class="footer-v1 mt-3 bg-light py-5">
		<div class="footer">
			<div class="container">
				<div class="row">
					<div class="col-md-3">
						<div>
							<img src="https://raw.githubusercontent.com/codingersid/legit-cli/main/assets/legit-logo/legit%20landscape%20color.png" width="250" class="img-fluid" alt="Legit Framework Logo" draggable="false">
						</div>
						<p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Iusto ut rerum porro ratione deleniti voluptates quo, libero vero dolorum, molestias accusantium! Error laudantium obcaecati repudiandae asperiores quisquam soluta delectus impedit!
						</p>
					</div>
					<div class="col-md-3 offset-md-3">
						<div class="headline">
							<p>Menus</p>
						</div>
						<ul class="list-unstyled fw-normal pb-1 small">
							<li><a href="#">Menu 1</a></li>
							<li><a href="#">Menu 2</a></li>
							<li><a href="#">Menu 3</a></li>
							<li><a href="#">Menu 4</a></li>
						</ul>
					</div>
					<div class="col-md-3">
						<div class="headline">
							<p>Company</p>
						</div>
						<ul class="list-unstyled fw-normal pb-1 small">
							<li><a href="#">About Us</a></li>
							<li><a href="#">Privacy Policy</a></li>
							<li><a href="#">Terms of Use</a></li>
							<li><a href="#">FAQs</a></li>
							<li><a href="#">Cancellation/Rescheduling policy</a></li>
						</ul>
					</div>
				</div>
			</div>
		</div>
	</div>
	<footer class="text-center bg-dark text-light py-2">
		<small>&copy; yyyy . Legit framework . All Rights Reserved</small>
	</footer>

	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"></script>
</body>

</html>`

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(layoutCmd)
}
