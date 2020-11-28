package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/**
*Country,City,Languge情報を取得
*
**/
func handleGetCountryInfo(w http.ResponseWriter, r *http.Request) (err error) {

	strCountryName := r.URL.Query().Get("countryName")
	strCityName := r.URL.Query().Get("cityName")
	strLanguage := r.URL.Query().Get("language")

	fmt.Println("countryName:", strCountryName)
	fmt.Println("cityName:", strCityName)
	fmt.Println("language:", strLanguage)

	//検索処理
	countrys, err := getCountryByKeyDataAccess(strCountryName, strCityName, strLanguage)

	//json整形の為にItems構造体に格納
	var items Items = Items{
		Items: countrys,
	}

	//構造体をJsonに変換
	res, err := json.Marshal(items)

	//エラー処理
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Header
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//ステータスコード 200
	w.WriteHeader(http.StatusOK)

	// Response
	w.Write(res)

	return
}

/**
*Country,city,Language情報を登録
*
**/
func handlePostCountryInfo(w http.ResponseWriter, r *http.Request) (err error) {

	// Country
	strCountryCode := r.FormValue("countryCode")
	strCountryName := r.FormValue("countryName")
	strContinent := r.FormValue("continent")
	strRegion := r.FormValue("region")
	strSurfaceArea := r.FormValue("surfaceArea")
	strIndepYear := r.FormValue("indepYear")
	strPopulation := r.FormValue("population")
	strLifeExpectancy := r.FormValue("lifeExpectancy")
	strGNP := r.FormValue("gnp")
	strGNPOld := r.FormValue("gnpOld")
	strLocalName := r.FormValue("localName")
	strGovernmentForm := r.FormValue("governmentForm")
	strHeadOfState := r.FormValue("headOfState")
	strCapital := r.FormValue("capital")
	strCode2 := r.FormValue("code2")

	//City
	strCityId := r.FormValue("cityId")
	strCityName := r.FormValue("cityName")
	strDistrict := r.FormValue("cityDistrict")
	strCityPopulation := r.FormValue("cityPopulation")

	//Language
	strCountryLangugage := r.FormValue("language")
	strIsOfficial := r.FormValue("isOfficial")
	strPercentage := r.FormValue("percentage")

	//City確認
	fmt.Println("cityId:" + strCityId)
	fmt.Println("cityName:" + strCityName)
	fmt.Println("District:" + strDistrict)
	fmt.Println("cityPopulation:" + strCityPopulation)

	//Language確認
	fmt.Println("language:" + strCountryLangugage)
	fmt.Println("isOfficaila:" + strIsOfficial)
	fmt.Println("percentage:" + strPercentage)

	//受け取った値をキャスト 国
	fSurfaceArea, err := strconv.ParseFloat(strSurfaceArea, 64)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intIndepYear, err := strconv.Atoi(strIndepYear)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intPopulation, err := strconv.Atoi(strPopulation)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fLifeExpectancy, err := strconv.ParseFloat(strLifeExpectancy, 64)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fGnp, err := strconv.ParseFloat(strGNP, 64)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fGnpOld, err := strconv.ParseFloat(strGNPOld, 64)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intCapital, err := strconv.Atoi(strCapital)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intCityPopulation, err := strconv.Atoi(strCityPopulation)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fPercentage, err := strconv.ParseFloat(strPercentage, 64)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//構造体にクエリパラメタを格納
	var stRegistCountry Country = Country{
		Code:           strCountryCode,
		Name:           strCountryName,
		Continent:      strContinent,
		Region:         strRegion,
		SurfaceArea:    fSurfaceArea,
		IndepYear:      intIndepYear,
		Population:     intPopulation,
		LifeExpectancy: fLifeExpectancy,
		GNP:            fGnp,
		GNPOld:         fGnpOld,
		LocalName:      strLocalName,
		GovernmentForm: strGovernmentForm,
		HeadOfState:    strHeadOfState,
		Capital:        intCapital,
		Code2:          strCode2,
	}

	var stRegistCity City = City{
		Name:       strCityName,
		District:   strDistrict,
		Population: intCityPopulation,
	}

	var stRegistLanguage Language = Language{
		CountryCode: strCountryCode,
		Language:    strCountryLangugage,
		IsOfficial:  strIsOfficial,
		Percentage:  fPercentage,
	}

	//トランザクション開始
	tx, err := Db.Begin()
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		//ランタイムパニックが起こるとロールバック
		if recover() != nil {
			tx.Rollback()
		}
	}()

	//登録処理 Country
	err = insertCountryDataAccess(stRegistCountry, tx)
	if err != nil {
		fmt.Println(err.Error())
		//DBエラー ロールバック
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//登録処理 City
	err = insrtCityDataAccess(stRegistCity, tx)
	if err != nil {
		fmt.Println(err.Error())
		//DBエラー ロールバック
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//登録処理 Language
	err = insrtLanguageDataAccess(stRegistLanguage, tx)
	if err != nil {
		fmt.Println(err.Error())
		//DBエラー ロールバック
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//トランザクション正常終了
	tx.Commit()
	//ステータスコード 200
	w.WriteHeader(http.StatusOK)

	//正常処理
	return
}
