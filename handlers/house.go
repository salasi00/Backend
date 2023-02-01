package handlers

import (
	"encoding/json"
	housedto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerhouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerhouse {
	return &handlerhouse{HouseRepository}
}

var path_file = "http://localhost:5000/uploads/"

func (h *handlerhouse) FindHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	house, err := h.HouseRepository.FindHouses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range house {
		house[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) GetHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house.Image = path_file + house.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) CreateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContext := r.Context().Value("dataFile")
	filename := dataContext.(string)

	cityid, _ := strconv.Atoi(r.FormValue("cityid"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	bedRoom, _ := strconv.Atoi(r.FormValue("bedRoom"))
	bathroom, _ := strconv.Atoi(r.FormValue("bathroom"))
	request := housedto.HouseResponse{
		Name:      r.FormValue("name"),
		CityID:    cityid,
		Address:   r.FormValue("address"),
		Price:     price,
		TypeRent:  r.FormValue("typeRent"),
		Amenities: r.FormValue("amenities"),
		BedRoom:   bedRoom,
		Bathroom:  bathroom,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house := models.House{
		Name:      request.Name,
		CityID:    request.CityID,
		Address:   request.Address,
		Price:     request.Price,
		TypeRent:  request.TypeRent,
		Amenities: request.Amenities,
		BedRoom:   request.BedRoom,
		Bathroom:  request.Bathroom,
		Image:     filename,
		UserID:    userId,
	}

	data, err := h.HouseRepository.CreateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.HouseRepository.GetHouse(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cityid, _ := strconv.Atoi(r.FormValue("cityid"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	bedRoom, _ := strconv.Atoi(r.FormValue("bedRoom"))
	bathroom, _ := strconv.Atoi(r.FormValue("bathroom"))
	request := housedto.HouseResponse{
		Name:      r.FormValue("name"),
		CityID:    cityid,
		Address:   r.FormValue("address"),
		Price:     price,
		TypeRent:  r.FormValue("typeRent"),
		Amenities: r.FormValue("amenities"),
		BedRoom:   bedRoom,
		Bathroom:  bathroom,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		house.Name = request.Name
	}

	if request.Address != "" {
		house.Address = request.Address
	}

	if request.Price != 0 {
		house.Price = request.Price
	}

	if request.TypeRent != "" {
		house.TypeRent = request.TypeRent
	}

	if request.Amenities != "" {
		house.Amenities = request.Amenities
	}

	if request.BedRoom != 0 {
		house.BedRoom = request.BedRoom
	}

	if request.Bathroom != 0 {
		house.Bathroom = request.Bathroom
	}

	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.HouseRepository.GetHouse(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertHouseResponse(u models.House) housedto.HouseResponse {
	return housedto.HouseResponse{
		ID:        u.ID,
		Name:      u.Name,
		CityID:    u.CityID,
		City:      u.City,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Amenities: u.Amenities,
		BedRoom:   u.BedRoom,
		Bathroom:  u.Bathroom,
	}
}
