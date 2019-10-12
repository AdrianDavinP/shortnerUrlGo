package Process

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
	"mraladin/shortnerUrl/src/Struct"
	"net/http"
	"time"
)

func DaftarUrl(w http.ResponseWriter, req *http.Request) {
	var url Struct.Url
	_ = json.NewDecoder(req.Body).Decode(&url)
	var data Struct.Data
	var count = 0
	db := ConnectionPostgres()
	defer db.Close()

	if url.URLPanjang == "" {
		w.WriteHeader(401)
		w.Write([]byte("Invalid Value"))
		return
	}

	err := db.Model(&data).Where("url_panjang=?", url.URLPanjang).Count(&count).Error
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}

	if count == 0 {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		now := time.Now()
		data.Id, _ = h.Encode([]int{int(now.Unix())})
		data.Url_pendek = "http://localhost:23456/" + data.Id
		data.Url_panjang = url.URLPanjang

		db.Create(&data)

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
		url.URLPendek = data.Url_pendek
		url.Id = data.Id

	} else {
		w.WriteHeader(401)
		w.Write([]byte("URL Sudah Terdaftar"))
		return
	}
	json.NewEncoder(w).Encode(url)

}

func HitUrlPendek(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var data Struct.Data
	db := ConnectionPostgres()
	defer db.Close()
	db.Where("id= ?", params["id"]).First(&data)
	http.Redirect(w, req, data.Url_panjang, 301)
}

func UbahUrl(w http.ResponseWriter, req *http.Request) {
	var url Struct.Url
	_ = json.NewDecoder(req.Body).Decode(&url)
	var data Struct.Data
	db := ConnectionPostgres()
	defer db.Close()
	err := db.Where("url_panjang=?", url.URLPanjang).Find(&data).Error
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}

	if data != (Struct.Data{}) {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		now := time.Now()
		keyUniq, _ := h.Encode([]int{int(now.Unix())})

		data.Url_pendek = "http://localhost:23456/" + keyUniq
		data.Url_panjang = url.URLPanjang

		//result :=  db.Save(&data)
		err = db.Exec("UPDATE data set id = ?,url_panjang=?, url_pendek=? where id = ?", keyUniq, data.Url_panjang, data.Url_pendek, data.Id).Error
		//fmt.Println(result)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
		url.URLPendek = data.Url_pendek
		url.Id = data.Id

	} else {
		w.WriteHeader(401)
		w.Write([]byte("Data Tidak Ditemukan"))
		return
	}
	json.NewEncoder(w).Encode(url)

}

func Delete(w http.ResponseWriter, req *http.Request) {
	var url Struct.Url
	_ = json.NewDecoder(req.Body).Decode(&url)
	var data Struct.Data
	db := ConnectionPostgres()
	defer db.Close()
	err := db.Where("url_pendek=?", url.URLPendek).Find(&data).Error
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}

	if data != (Struct.Data{}) {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		now := time.Now()
		keyUniq, _ := h.Encode([]int{int(now.Unix())})

		data.Url_pendek = "http://localhost:23456/" + keyUniq
		data.Url_panjang = url.URLPanjang

		err := db.Delete(&data).Error

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

	} else {
		w.WriteHeader(401)
		w.Write([]byte("Data Tidak Ditemukan"))
		return
	}
	json.NewEncoder(w).Encode("Data Berhasil Dihapus")

}
