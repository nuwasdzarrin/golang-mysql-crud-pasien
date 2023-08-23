package pasiencontroller

import (
	"fmt"
	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/libraries"
	"github.com/jeypc/go-crud/models"
	"html/template"
	"net/http"
	"strconv"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {
	pasien, _ := pasienModel.GetAll()
	data := map[string]interface{}{
		"pasien": pasien,
	}
	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHP = request.Form.Get("no_hp")

		data := make(map[string]interface{})
		vError := validation.Struct(pasien)
		if vError != nil {
			data["pasien"] = pasien
			data["validation"] = vError
		} else {
			pasienModel.Create(pasien)
			data["pesan"] = "Data pasien berhasil disimpan"
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pasien entities.Pasien
		err := pasienModel.Find(id, &pasien)
		if err != nil {
			return
		}
		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHP = request.Form.Get("no_hp")

		data := make(map[string]interface{})
		vError := validation.Struct(pasien)
		if vError != nil {
			data["pasien"] = pasien
			data["validation"] = vError
		} else {
			err := pasienModel.Update(pasien)
			if err != nil {
				fmt.Println(err)
				return
			}
			data["pesan"] = "Data pasien berhasil diperbarui"
		}

		temp, _ := template.ParseFiles("views/pasien/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
	pasienModel.Delete(id)
	http.Redirect(response, request, "/pasien", http.StatusSeeOther)
}
