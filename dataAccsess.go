package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

/*
*検索キーを使用してデータを取得
*prams countryName, cityName, language
*return country[] or Error
 */
func getCountryByKeyDataAccess(strCountryName string, strCityName string, strLanguage string) (countrys []Country, err error) {

	//stringBuilder
	var strSql strings.Builder

	//SQLを記述
	strSql.WriteString("SELECT ")
	strSql.WriteString(" country.Code As code, ")
	strSql.WriteString(" country.Name As name, ")
	strSql.WriteString(" country.Continent As continent, ")
	strSql.WriteString(" country.Region As region, ")
	strSql.WriteString(" country.SurfaceArea As surfaceArea, ")
	strSql.WriteString(" COALESCE(country.IndepYear, 0) As indepYear, ")
	strSql.WriteString(" country.Population As population, ")
	strSql.WriteString(" COALESCE(country.LifeExpectancy, 0) As lifeExpectancy, ")
	strSql.WriteString(" country.GNP As gnp, ")
	strSql.WriteString(" COALESCE(country.GNPOld, 0) As gnpOld, ")
	strSql.WriteString(" country.LocalName As localName, ")
	strSql.WriteString(" country.GovernmentForm As governmentForm, ")
	strSql.WriteString(" COALESCE(country.HeadOfState, '') As headOfState, ")
	strSql.WriteString(" COALESCE(country.Capital, 0) As capital, ")
	strSql.WriteString(" country.Code2 code2, ")
	strSql.WriteString(" COALESCE(city.ID, 0) As id, ")
	strSql.WriteString(" COALESCE(city.Name, '') As cityName, ")
	strSql.WriteString(" COALESCE(city.District, '') As district, ")
	strSql.WriteString(" COALESCE(city.Population, 0) As population, ")
	strSql.WriteString(" COALESCE(countrylanguage.Language, '') As language, ")
	strSql.WriteString(" COALESCE(countrylanguage.IsOfficial, '') As isOfficial, ")
	strSql.WriteString(" COALESCE(countrylanguage.Percentage, 0) As percentage ")
	strSql.WriteString(" FROM ")
	strSql.WriteString(" world.country ")
	strSql.WriteString(" LEFT OUTER JOIN ")
	strSql.WriteString(" world.city ")
	strSql.WriteString(" ON ")
	strSql.WriteString(" country.Code = city.CountryCode ")
	strSql.WriteString(" LEFT OUTER JOIN ")
	strSql.WriteString(" countrylanguage ")
	strSql.WriteString(" ON ")
	strSql.WriteString(" country.Code = countrylanguage.CountryCode ")

	//検索文字列が存在すれば、WHERE句を追加
	if strCountryName != "" || strCityName != "" || strLanguage != "" {
		strSql.WriteString(" WHERE ")
	}

	//国名が存在すれば検索句を追加
	if strCountryName != "" {
		strSql.WriteString(" country.Name = " + "'" + strCountryName + "' ")
		//町名か、公用語が存在すればAND句を追加
		if strCityName != "" || strLanguage != "" {
			strSql.WriteString(" AND ")
		}
	}

	//町名が存在すれば検索句を追加
	if strCityName != "" {
		strSql.WriteString(" city.Name = " + "'" + strCityName + "' ")
		//公用語が存在すれば、AND句を追加
		if strLanguage != "" {
			strSql.WriteString(" AND ")
		}
	}

	//公用語が存在すれば、検索句を追加
	if strLanguage != "" {
		strSql.WriteString(" countrylanguage.Language = " + "'" + strLanguage + "' ")
	}

	strSql.WriteString(" ORDER BY country.Code ")

	//確認 あとで消す
	//fmt.Println("ストリングビルダーで生成:" + strSql.String())

	//クエリ
	rows, err := Db.Query(strSql.String())

	//エラー処理
	if err != nil {
		log.Fatal(err)
	}

	//Rowを最後に閉じる
	defer rows.Close()

	for rows.Next() {
		country := Country{}

		err := rows.Scan(
			&country.Code,
			&country.Name,
			&country.Continent,
			&country.Region,
			&country.SurfaceArea,
			&country.IndepYear,
			&country.Population,
			&country.LifeExpectancy,
			&country.GNP,
			&country.GNPOld,
			&country.LocalName,
			&country.GovernmentForm,
			&country.HeadOfState,
			&country.Capital,
			&country.Code2,
			&country.Id,
			&country.CityName,
			&country.District,
			&country.CityPopulation,
			&country.Language,
			&country.IsOfficial,
			&country.Percentage)

		if err != nil {
			log.Fatal(err)
			break
		}

		countrys = append(countrys, country)
	}

	return
}

/*
*国登録
*params Country struct
*return err or result
 */
func insertCountryDataAccess(stCountry Country, tx *sql.Tx) (err error) {

	//ストリングビルダー(SQL成形)
	var strSql strings.Builder

	//SQLを記述
	strSql.WriteString("INSERT INTO ")
	strSql.WriteString(" country(")
	strSql.WriteString(" Code, ")
	strSql.WriteString(" Name, ")
	strSql.WriteString(" Continent, ")
	strSql.WriteString(" Region, ")
	strSql.WriteString(" SurfaceArea, ")
	strSql.WriteString(" IndepYear, ")
	strSql.WriteString(" Population, ")
	strSql.WriteString(" LifeExpectancy, ")
	strSql.WriteString(" GNP, ")
	strSql.WriteString(" GNPOld, ")
	strSql.WriteString(" LocalName, ")
	strSql.WriteString(" GovernmentForm, ")
	strSql.WriteString(" HeadOfState, ")
	strSql.WriteString(" Capital, ")
	strSql.WriteString(" Code2 ")
	strSql.WriteString(" )VALUES( ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ?, ")
	strSql.WriteString(" ? ")
	strSql.WriteString(" )")

	fmt.Println("INSERT文:" + strSql.String())

	stmt, err := tx.Prepare(strSql.String())

	//エラー処理
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		stCountry.Code,
		stCountry.Name,
		stCountry.Continent,
		stCountry.Region,
		stCountry.SurfaceArea,
		stCountry.IndepYear,
		stCountry.Population,
		stCountry.LifeExpectancy,
		stCountry.GNP,
		stCountry.GNPOld,
		stCountry.LocalName,
		stCountry.GovernmentForm,
		stCountry.HeadOfState,
		stCountry.Capital,
		stCountry.Code2)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	insetrtID := result.LastInsertId
	fmt.Println(insetrtID)
	return
}

/**
*町登録
*params City
**/
func insrtCityDataAccess(stCity City, tx *sql.Tx) (err error) {

	//ストリングビルダー(SQL成形)
	var sbSql strings.Builder

	sbSql.WriteString("INSERT INTO ")
	sbSql.WriteString(" city(")
	sbSql.WriteString(" Id, ")
	sbSql.WriteString(" Name, ")
	sbSql.WriteString(" District, ")
	sbSql.WriteString(" Population ")
	sbSql.WriteString(" )VALUES( ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ? ")
	sbSql.WriteString(" )")

	fmt.Println("INSERT文:" + sbSql.String())

	stmt, err := tx.Prepare(sbSql.String())

	//エラー処理
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		stCity.Id,
		stCity.Name,
		stCity.District,
		stCity.Population)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	insetrtID := result.LastInsertId
	fmt.Println(insetrtID)

	return
}

func insrtLanguageDataAccess(stLanguage Language, tx *sql.Tx) (err error) {

	//stringBuilder生成
	var sbSql strings.Builder

	sbSql.WriteString("INSERT INTO ")
	sbSql.WriteString(" language( ")
	sbSql.WriteString(" CountryCode, ")
	sbSql.WriteString(" Language, ")
	sbSql.WriteString(" IsOfficial, ")
	sbSql.WriteString(" Percentage ")
	sbSql.WriteString(" )VALUES( ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ?, ")
	sbSql.WriteString(" ? ")
	sbSql.WriteString(" )")

	fmt.Println("INSERT文:" + sbSql.String())

	stmt, err := tx.Prepare(sbSql.String())

	//エラー処理
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		stLanguage.CountryCode,
		stLanguage.Language,
		stLanguage.IsOfficial,
		stLanguage.Percentage)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.Error())
		return
	}

	insetrtID := result.LastInsertId
	fmt.Println(insetrtID)

	return
}
