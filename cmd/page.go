package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var pageCmd = &cobra.Command{
	Use:   "page [nama_page]",
	Short: "Membuat page baru di path: resources/views/pages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pageName := args[0]
		nameParts := strings.Split(pageName, "/")
		pathName := TransformPath(pageName)
		fileHtml := SnakeCase(nameParts[len(nameParts)-1])
		pagePath := fmt.Sprintf(`%s/%s.html`, pathName, fileHtml)
		status, err := createpage(pageName)
		if err != nil {
			fmt.Println("Gagal membuat page:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat page, page", pagePath, "sudah ada!")
			return
		}
		fmt.Println("page", pagePath, "berhasil dibuat!")
	},
}

func createpage(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := TransformPath(name)
	fileHtml := SnakeCase(nameParts[len(nameParts)-1])
	pagePath := fmt.Sprintf(`resources/views/pages/%s/%s.html`, pathName, fileHtml)
	err := os.MkdirAll(fmt.Sprintf("resources/views/pages/%s", pathName), os.ModePerm)
	if err != nil {
		return status, err
	}

	if _, err := os.Stat(pagePath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(pagePath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	code := fmt.Sprintf(`<!-- %s::begin -->
<div class="container">
    <div class="row">
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
        <div class="col-lg-4 col-md-6 col-sm-12 p-2">
            <div class="card">
                <img src="https://source.unsplash.com/1600x900/?news" class="card-img-top" alt="image">
                <div class="card-body">
                    <p class="card-text">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Quis iure maiores quae! Sapiente quos similique in consectetur autem esse saepe repudiandae expedita doloribus provident. Delectus quia natus blanditiis sunt qui!
                    </p>
                    <a href="#" class="btn btn-primary">Read More</a>
                </div>
            </div>
        </div>
    </div>
</div>
<!-- %s::end -->`, fileHtml, fileHtml)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(pageCmd)
}
