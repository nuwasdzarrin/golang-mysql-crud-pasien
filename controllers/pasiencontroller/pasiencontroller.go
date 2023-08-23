package pasiencontroller

import (
	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/libraries"
	"github.com/jeypc/go-crud/models"
	"html/template"
	"net/http"
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
		temp, err := template.ParseFiles("views/pasien/edit.html")
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

func Delete(response http.ResponseWriter, request *http.Request) {

}
